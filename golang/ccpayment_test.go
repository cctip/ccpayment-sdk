package ccpayment_go_sdk

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

var (
	appId     = "6cfakPeNnTi0YoHLpbqVLqFw2A#R5EGPH3lJfX75H7x#SO#PXU"
	appSecret = `5SQOObrXo#t82ZsVYfrI02Dn@5MLVPX1Fr1gcpCVS2Ca8ClbAgi@AMp64j0pWO#4jyC#zdS3kis1UaNFsaqW5DGjtxA4gYk1ZEH68#S1Z1McmUd@ph8Gbhn#`
)

// create order
func TestCreateOrder(t *testing.T) {
	order := OrderParams{
		TokenId:         "8e5741cf-6e51-4892-9d04-3d40e1dd0128",
		Amount:          "6",
		MerchantOrderId: "1121241232",
		FiatCurrency:    "USD",
		Remark:          "",
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
	curl := GetPaymentUrlParams{
		Amount:          "12",
		MerchantOrderId: "test_xxxx",
		ProductName:     "name_1",
		ReturnUrl:       "",
		ValidTimestamp:  300, // s
	}

	data, err := curl.GetCheckoutUrl(appId, appSecret)
	if err != nil {
		fmt.Println(`GetCheckoutUrlByRsa error: `, err)
		return
	}

	fmt.Printf(`%+v`, data)
}

// webhook validate
func TestGetWebhookDataAndValidate(t *testing.T) {
	router := gin.Default()

	router.POST("/webhook/verify/simple", PayNotifyBack) // Example of Webhook Verification

	s := &http.Server{
		Addr:           ":8089",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(`ListenAndServe error: `, err)
		return
	}
}

func PayNotifyBack(c *gin.Context) {
	var result WebhookParams

	// style 1: get and verify
	if err := result.GetWebhookDataAndValidate(c.Request, appId, appSecret); err != nil {
		fmt.Printf(`PayNotifyBack - GetWebhookDataAndValidate error: %v`, err)
		c.String(200, "error")
		return
	}

	fmt.Printf(`%+v`, result)

	// style 2: verify
	/*
		byt, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Printf(`PayNotifyBack - ReadAll error: %v`, err)
			return
		}
		wv := &WebhookValidate{
			Data:      byt,
			Signature: c.Request.Header.Get(SignHeaderKey),
			Timestamp: c.Request.Header.Get(TimestampHeaderKey),
		}
		if wv.WebhookValidate(appId, appSecret) {
			fmt.Println(`success`)
		} else {
			fmt.Println(`error`)
			return
		}
	*/

	c.String(200, "success")
}

// get support tokens
func TestGetSupportTokens(t *testing.T) {
	st := &SupportTokenParam{}
	data, err := st.GetSupportTokens(appId, appSecret)
	if err != nil {
		fmt.Println(`GetSupportTokens error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}

// get token chain
func TestGetTokenChain(t *testing.T) {
	tc := &TokenChainParams{
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
	tc := &GetTokenRateParams{
		Amount:  "12",
		TokenId: "8e5741cf-6e51-4892-9d04-3d40e1dd0128",
	}
	data, err := tc.GetTokenRate(appId, appSecret)
	if err != nil {
		fmt.Println(`GetTokenRate error: `, err)
	}
	fmt.Printf(`data: %+v`, data)
}
