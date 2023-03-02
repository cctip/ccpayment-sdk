package ccpayment_go_sdk

import (
	"ccpayment-sdk/golang/sign"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var SignVerifyErr = fmt.Errorf(`data signature verify error`)

// ---------- api create order ----------

func (order *OrderParams) CreateOrder(appId, appSecret string) (data *OrderResultData, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*order, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &OrderResultData{}

	err = sendPost(data, dst, CreateOrderUrl, appId, appSecret, signStr, timeStamp)

	return
}

// ---------- get checkout url ----------

func (pu *GetPaymentUrlParams) GetCheckoutUrl(appId, appSecret string) (data *PaymentUrlResultData, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*pu, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &PaymentUrlResultData{}

	err = sendPost(data, dst, CheckoutUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

// ---------- webhook verify ----------

func (wv *WebhookValidate) WebhookValidate(appId, appSecret string) bool {
	if wv.Signature != "" && signStr_(string(wv.Data), appId, appSecret, wv.Timestamp) == wv.Signature {
		return true
	}
	return false
}

// ---------- webhook data and verify ----------

// result get data and validate
func (result *WebhookParams) GetWebhookDataAndValidate(req *http.Request, appId, appSecret string) error {
	byt, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byt, result)
	if err != nil {
		return err
	}

	// validte
	if result.webhookSignStr(req, appId, appSecret, byt) {
		return nil
	}

	return SignVerifyErr
}

func (result *WebhookParams) webhookSignStr(req *http.Request, appId, appSecret string, byt []byte) bool {
	timestamp := req.Header.Get(TimestampHeaderKey)
	signStr := req.Header.Get(SignHeaderKey)

	if signStr != "" && signStr_(string(byt), appId, appSecret, timestamp) == signStr {
		return true
	}
	return false
}

// ---------- get support token ----------

func (st *SupportTokenParam) GetSupportTokens(appId, appSecret string) (data *SupportTokenResultData, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*st, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &SupportTokenResultData{}

	err = sendPost(data, dst, GetSupportTokenUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

// ---------- get token chain ----------

func (tc *TokenChainParams) GetTokenChain(appId, appSecret string) (data *TokenChainResultData, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tc, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &TokenChainResultData{}

	err = sendPost(data, dst, GetTokenChainUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

// ---------- get token rate ----------

func (tr *GetTokenRateParams) GetTokenRate(appId, appSecret string) (data *TokenRateResultData, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tr, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &TokenRateResultData{}

	err = sendPost(data, dst, GetTokenRateUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func sendPost(data interface{}, dst string, url, appId, appSecret, signStr string, timeStamp int64) (err error) {

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(dst))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(AppIdHeaderKey, appId)
	req.Header.Add(SignHeaderKey, signStr)
	req.Header.Add(TimestampHeaderKey, fmt.Sprintf(`%d`, timeStamp))

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	if strings.Contains(strings.ToLower(url), `https://`) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		byt, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		dt := reflect.TypeOf(data)
		if dt.Kind() != reflect.Ptr {
			return fmt.Errorf(`params: data must be ptr`)
		}

		err = json.Unmarshal(byt, data)
		if err != nil {
			return err
		}

		var code int

		switch data.(type) {
		case *OrderResultData:
			d := data.(*OrderResultData)
			code = d.Code
			goto validate

		case *PaymentUrlResultData:
			d := data.(*PaymentUrlResultData)
			code = d.Code
			goto validate

		case *SupportTokenResultData:
			d := data.(*SupportTokenResultData)
			code = d.Code
			goto validate

		case *TokenChainResultData:
			d := data.(*TokenChainResultData)
			code = d.Code
			goto validate

		case *TokenRateResultData:
			d := data.(*TokenRateResultData)
			code = d.Code
			goto validate

		case *BillTradeResultData:
			d := data.(*BillTradeResultData)
			code = d.Code
			goto validate

		default:
			return fmt.Errorf(`ambiguous receive type`)
		}

	validate:
		if code == 10000 {
			if !getHeadersAndValidate(resp, appId, appSecret, byt) {
				return SignVerifyErr
			}
		}
		return nil
	}

	return fmt.Errorf(`page not found; status_code: %v`, resp.StatusCode)
}

func SignStr(src interface{}, appId, appSecret string, timestamp int64) (dst, signStr string, err error) {
	byt, err := json.Marshal(src)
	if err != nil {
		return "", "", err
	}

	return string(byt), signStr_(string(byt), appId, appSecret, fmt.Sprintf(`%d`, timestamp)), nil
}

func signStr_(dst, appId, appSecret, timestamp string) string {
	return sign.Hash256([]byte(fmt.Sprintf(`%s%s%s%s`, appId, appSecret, timestamp, dst)))
}

func getHeadersAndValidate(resp *http.Response, appId, appSecret string, byt []byte) bool {
	signStr := resp.Header.Get(SignHeaderKey)
	timestamp := resp.Header.Get(TimestampHeaderKey)

	if signStr != "" && signStr_(string(byt), appId, appSecret, timestamp) == signStr {
		return true
	}

	return false
}
