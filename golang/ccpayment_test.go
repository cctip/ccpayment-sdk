package golang

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var (
	appId     = "202307200931401681960050962821120"
	appSecret = `3c9c804fb62e39efe6e5b999a7d91e06`
	//appId     = "2023051616364210105509935845378"
	//appSecret = `9ae5247d0f8dca9f3300e328d0ad9a96`
)

// create order
func TestCreateOrder(t *testing.T) {
	order := CreateOrderReq{
		Remark:              "",
		TokenId:             "0912e09a-d8e2-41d7-a0bc-a25530892988",
		ProductPrice:        "0.5",
		MerchantOrderId:     "3735077979050379",
		DenominatedCurrency: "USD",
		OrderValidPeriod:    823456,
	}
	data, err := order.CreateOrder(appId, appSecret)
	if err != nil {
		fmt.Println(`TestCreateOrder error: `, err)
		return
	}
	fmt.Printf(`%+v`, data)
}

// get checkout url
func TestCheckoutUrl(t *testing.T) {
	curl := CheckoutUrlReq{
		ProductPrice:     "12",
		MerchantOrderId:  fmt.Sprintf("test_xxxx%d", time.Now().UnixMicro()),
		ProductName:      "name_1",
		ReturnUrl:        "",
		OrderValidPeriod: 300, // s
	}

	data, err := curl.CheckoutUrl(appId, appSecret)
	if err != nil {
		fmt.Println(`GetCheckoutUrlByRsa error: `, err)
		return
	}

	fmt.Printf(`%+v`, data)
}

// webhook validate
//func TestGetWebhookDataAndValidate(t *testing.T) {
//	router := gin.Default()
//
//	router.POST("/webhook/verify/simple", PayNotifyBack) // Example of Webhook Verification
//
//	s := &http.Server{
//		Addr:           ":8089",
//		Handler:        router,
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//	err := s.ListenAndServe()
//	if err != nil {
//		fmt.Println(`ListenAndServe error: `, err)
//		return
//	}
//}

// webhook
//func PayNotifyBack(c *gin.Context) {
//	var result WebhookReq
//
//	// style 1: get and verify
//	if err := result.GetWebhookDataAndValidate(c.Request, appId, appSecret); err != nil {
//		fmt.Printf(`PayNotifyBack - GetWebhookDataAndValidate error: %v`, err)
//		c.String(200, "error")
//		return
//	}
//
//	fmt.Printf(`%+v`, result)
//
//	// style 2: verify
//
//	byt, err := ioutil.ReadAll(c.Request.Body)
//	if err != nil {
//		fmt.Printf(`PayNotifyBack - ReadAll error: %v`, err)
//		return
//	}
//	wv := &WebhookValidate{
//		Data:      byt,
//		Signature: c.Request.Header.Get(SignHeaderKey),
//		Timestamp: c.Request.Header.Get(TimestampHeaderKey),
//	}
//	if wv.WebhookValidate(appId, appSecret) {
//		fmt.Println(`success`)
//	} else {
//		fmt.Println(`error`)
//		return
//	}
//
//	c.String(200, "success")
//}

func TestGetSupportCoin(t *testing.T) {
	st := &SupportCoinReq{}
	data, err := st.GetSupportCoin(appId, appSecret)
	if err != nil {
		fmt.Println(`GetSupportCoin error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// get support tokens
func TestGetSupportToken(t *testing.T) {
	st := &SupportTokenReq{}
	data, err := st.GetSupportToken(appId, appSecret)
	if err != nil {
		fmt.Println(`GetSupportToken error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// get token chain
func TestGetTokenChain(t *testing.T) {
	tc := &TokenChainReq{
		TokenId: "8e5741cf-6e51-4892-9d04-3d40e1dd0128",
	}
	data, err := tc.GetTokenChain(appId, appSecret)
	if err != nil {
		fmt.Println(`GetTokenChain error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// get token rate
func TestGetTokenRate(t *testing.T) {
	tc := &GetTokenRateReq{
		Amount:  "12",
		TokenId: "8e5741cf-6e51-4892-9d04-3d40e1dd0128",
	}
	data, err := tc.GetTokenRate(appId, appSecret)
	if err != nil {
		fmt.Println(`GetTokenRate error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// create api withdrawal
func TestWithdraw(t *testing.T) {
	tc := &WithdrawReq{
		TokenID:         "8e5741cf-6e51-4892-9d04-3d40e1dd0128",
		Address:         "9454818",
		Value:           "12",
		MerchantOrderId: fmt.Sprintf("%d", time.Now().UnixMicro()),
		MerchantPaysFee: false,
	}
	data, err := tc.Withdraw(appId, appSecret)
	if err != nil {
		fmt.Println(`Withdraw error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// c user check
func TestCheckUser(t *testing.T) {
	tc := &CheckUserReq{
		CId: "9454818",
	}
	data, err := tc.CheckUser(appId, appSecret)
	if err != nil {
		fmt.Println(`CheckUser error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// get token assets
func TestAssets(t *testing.T) {
	tc := &AssetsReq{
		TokenId: "8e5741cf-6e51-4892-9d04-3d40e1dd0128",
	}
	data, err := tc.Assets(appId, appSecret)
	if err != nil {
		fmt.Println(`Assets error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// get network fee
func TestNetworkFee(t *testing.T) {
	tc := &NetworkFeeReq{
		TokenId: "f137d42c-f3a6-4f23-9402-76f0395d0cfe",
	}
	data, err := tc.NetworkFee(appId, appSecret)
	if err != nil {
		fmt.Println(`GetTokenRate error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

func TestGetChainHeightInfo(t *testing.T) {
	tc := &NetworkChainHeightInfoReq{}
	data, err := tc.GetChainHeightInfo(appId, appSecret)
	if err != nil {
		fmt.Println(`GetTokenRate error: `, err)
	}
	
	dataJson, _ := json.Marshal(data)
	fmt.Printf(`data: %s`, dataJson)
}

// get bill trade

// get bill trade
func TestGetAPIOrderInfo(t *testing.T) {
	bt := &OrderInfoReq{
		MerchantOrderIds: []string{"2376655808480575", "4445821684092051"}, // 2376655808480575  4445821684092051
	}
	data, err := bt.GetAPIOrderInfo(appId, appSecret)
	if err != nil {
		fmt.Println(`GetAPIOrderInfo error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

func TestGetOtherPaymentAddress(t *testing.T) {
	ar := &AddressReq{
		UserId:    "gmm123456789",
		Chain:     "BSC",
		NotifyUrl: "",
	}
	data, err := ar.GetOtherPaymentAddress(appId, appSecret)
	if err != nil {
		fmt.Println(`GetOtherPaymentAddress error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}
