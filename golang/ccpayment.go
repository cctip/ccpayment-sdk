package golang

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cctip/ccpayment-sdk/golang/sign"
	"io/ioutil"
	"net/http"
	"reflect"
	"resty.dev/v3"
	"sync"
	"time"
)

var SignVerifyErr = fmt.Errorf(`data signature verify error`)

// ---------- api create order ----------

func (order *CreateOrderReq) CreateOrder(appId, appSecret string) (data *CreateOrderResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*order, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &CreateOrderResp{}

	err = sendPost(data, dst, CreateOrderUrl, appId, appSecret, signStr, timeStamp)

	return
}

// ---------- get checkout url ----------

func (pu *CheckoutUrlReq) CheckoutUrl(appId, appSecret string) (data *CheckoutUrlResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*pu, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &CheckoutUrlResp{}

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
func (result *WebhookReq) GetWebhookDataAndValidate(req *http.Request, appId, appSecret string) error {
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

func (result *WebhookReq) webhookSignStr(req *http.Request, appId, appSecret string, byt []byte) bool {
	timestamp := req.Header.Get(TimestampHeaderKey)
	signStr := req.Header.Get(SignHeaderKey)

	if signStr != "" && signStr_(string(byt), appId, appSecret, timestamp) == signStr {
		return true
	}
	return false
}

// ---------- get support token ----------

func (st *SupportCoinReq) GetSupportCoin(appId, appSecret string) (data *SupportCoinResultData, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*st, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &SupportCoinResultData{}

	err = sendPost(data, dst, GetSupportCoinUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

// ---------- get support token ----------

func (st *SupportTokenReq) GetSupportToken(appId, appSecret string) (data *SupportTokenResultData, err error) {
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

func (tc *TokenChainReq) GetTokenChain(appId, appSecret string) (data *TokenChainResultData, err error) {
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

func (tr *GetTokenRateReq) GetTokenRate(appId, appSecret string) (data *TokenRateResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tr, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &TokenRateResp{}

	err = sendPost(data, dst, GetTokenRateUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func (tr *WithdrawReq) Withdraw(appId, appSecret string) (data *WithdrawResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tr, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &WithdrawResp{}

	err = sendPost(data, dst, WithdrawApiUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func (tr *CheckUserReq) CheckUser(appId, appSecret string) (data *CheckUserResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tr, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &CheckUserResp{}

	err = sendPost(data, dst, CheckUserUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func (tr *AssetsReq) Assets(appId, appSecret string) (data *AssetsResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tr, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &AssetsResp{}

	err = sendPost(data, dst, AssetsUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func (tr *NetworkFeeReq) NetworkFee(appId, appSecret string) (data *NetworkFeeResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tr, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &NetworkFeeResp{}

	err = sendPost(data, dst, NetworkFeeUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func (tr *NetworkChainHeightInfoReq) GetChainHeightInfo(appId, appSecret string) (data *NetworkChainHeightInfoResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*tr, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &NetworkChainHeightInfoResp{}

	err = sendPost(data, dst, NetworkChainHeightInfoUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func (oi *OrderInfoReq) GetAPIOrderInfo(appId, appSecret string) (data *BillInfoResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*oi, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &BillInfoResp{}

	err = sendPost(data, dst, ApiOrderInfoUrl, appId, appSecret, signStr, timeStamp)

	return data, err
}

func sendPost(data interface{}, dst string, uri, appId, appSecret, signStr string, timeStamp int64) (err error) {

	r := NewClient().R().SetContext(context.Background())

	// headers
	r.SetHeaders(map[string]string{
		"Content-Type":     "application/json",
		AppIdHeaderKey:     appId,
		TimestampHeaderKey: fmt.Sprintf(`%d`, timeStamp),
		SignHeaderKey:      signStr,
	})
	r.SetBody(dst)

	resp, err := r.Post(uri)
	if err != nil {
		return err
	}
	// defer res

	if resp.StatusCode() == http.StatusOK {
		byt := resp.String()

		dt := reflect.TypeOf(data)
		if dt.Kind() != reflect.Ptr {
			return fmt.Errorf(`params: data must be ptr`)
		}

		err = json.Unmarshal([]byte(byt), data)
		if err != nil {
			return err
		}

		/*var code int

			switch data.(type) {
			case *CreateOrderResp:
				d := data.(*CreateOrderResp)
				code = d.Code
				goto validate

			case *CheckoutUrlResp:
				d := data.(*CheckoutUrlResp)
				code = d.Code
				goto validate

			case *SupportTokenResultData:
				d := data.(*SupportTokenResultData)
				code = d.Code
				goto validate
			case *SupportCoinResultData:
				d := data.(*SupportCoinResultData)
				code = d.Code
				goto validate

			case *TokenChainResultData:
				d := data.(*TokenChainResultData)
				code = d.Code
				goto validate

			case *TokenRateResp:
				d := data.(*TokenRateResp)
				code = d.Code
				goto validate

			case *BillInfoResp:
				d := data.(*BillInfoResp)
				code = d.Code
				goto validate
			case *WithdrawResp:
				d := data.(*WithdrawResp)
				code = d.Code
				goto validate
			case *CheckUserResp:
				d := data.(*CheckUserResp)
				code = d.Code
				goto validate
			case *NetworkFeeResp:
				d := data.(*NetworkFeeResp)
				code = d.Code
				goto validate
			case *AssetsResp:
				d := data.(*AssetsResp)
				code = d.Code
				goto validate
			case *AddressResp:
				d := data.(*AddressResp)
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
			}*/

		value := reflect.ValueOf(data).Elem().FieldByName("Code")
		if !value.CanInt() {
			return fmt.Errorf(`ambiguous receive type`)
		}

		if value.Int() == 10000 {
			if !getHeadersAndValidate(resp.RawResponse, appId, appSecret, []byte(byt)) {
				return SignVerifyErr
			}
		}

		return nil
	}

	return fmt.Errorf(`page not found; status_code: %v`, resp.StatusCode())
}

var (
	clientOnce sync.Once
	client     *resty.Client
)

func NewClient() *resty.Client {
	clientOnce.Do(func() {
		client = resty.New()
		client.SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
	})

	return client
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

func (ar *AddressReq) GetOtherPaymentAddress(appId, appSecret string) (data *AddressResp, err error) {
	timeStamp := time.Now().Unix()

	dst, signStr, err := SignStr(*ar, appId, appSecret, timeStamp)
	if err != nil {
		return nil, err
	}

	data = &AddressResp{}

	err = sendPost(data, dst, GetOtherPaymentAddress, appId, appSecret, signStr, timeStamp)

	return data, err
}
