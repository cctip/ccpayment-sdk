package golang

import "time"

// CreateOrderReq api create order params
type CreateOrderReq struct {
	TokenId             string `json:"token_id"`
	ProductPrice        string `json:"product_price"`
	MerchantOrderId     string `json:"merchant_order_id"`
	DenominatedCurrency string `json:"denominated_currency"` // USD
	Remark              string `json:"remark"`
}

// CreateOrderResp api create order response
type CreateOrderResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data OrderInfo `json:"data"`
}

type OrderInfo struct {
	OrderId            string        `json:"order_id"`
	ProductPrice       string        `json:"product_price"`
	Logo               string        `json:"logo"`
	Network            string        `json:"network"`
	PayAddress         string        `json:"pay_address"`
	Crypto             string        `json:"crypto"`
	TokenId            string        `json:"token_id"`
	Memo               string        `json:"memo"`
	AddressValidPeriod time.Duration `json:"address_valid_period"`
}

// CheckoutUrlReq checkout url params
type CheckoutUrlReq struct {
	OrderValidPeriod int64  `json:"order_valid_period"` // s
	ProductPrice     string `json:"product_price"`
	MerchantOrderId  string `json:"merchant_order_id"`
	ProductName      string `json:"product_name"`
	ReturnUrl        string `json:"return_url"`
}

// CheckoutUrlResp checkout url response
type CheckoutUrlResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PaymentUrl string `json:"payment_url"`
	} `json:"data"`
}

// WebhookValidate verify
type WebhookValidate struct {
	Data      []byte `json:"data"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

// WebhookReq webhook response
type WebhookReq struct {
	PayStatus           string         `json:"pay_status"`
	OrderType           string         `json:"order_type"`
	RecordId            string         `json:"record_id"`
	OrderID             string         `json:"order_id"`
	ProductPrice        string         `json:"product_price"`
	OrderAmount         string         `json:"order_amount"`
	FiatRate            string         `json:"fiat_rate"`
	DenominatedCurrency string         `json:"denominated_currency"`
	PaidAmount          string         `json:"paid_amount"`
	TokenRate           string         `json:"token_rate"`
	Chain               string         `json:"chain"`
	Contract            string         `json:"contract"`
	Crypto              string         `json:"crypto"`
	Extend              *WebhookExtend `json:"extend"`
}

type WebhookExtend struct {
	InvoiceId       string `json:"invoice_id"`
	UserEmail       string `json:"user_email"`
	MerchantOrderId string `json:"merchant_order_id"`
}

// SupportTokenReq get support token list
type SupportTokenReq struct {
}

type SupportTokenResultData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []SupportToken `json:"list"`
	} `json:"data"`
}

type SupportToken struct {
	Crypto  string        `json:"crypto"`
	Name    string        `json:"name"`
	Logo    string        `json:"logo"`
	Min     string        `json:"min"`
	Price   string        `json:"price"`
	TokenId string        `json:"token_id,omitempty"`
	CoinId  string        `json:"coin_id,omitempty"`
	Status  int64         `json:"status"` // 1 normal 2 maintenance 3 To be delisted
	Tokens  []*TokenChain `json:"tokens,omitempty"`
}

// TokenChainReq get token chain
type TokenChainReq struct {
	TokenId string `json:"token_id"`
}

type TokenChainResultData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []TokenChain `json:"list"`
	} `json:"data"`
}

type TokenChain struct {
	TokenId   string `json:"token_id"`
	Crypto    string `json:"crypto"`
	Logo      string `json:"logo"`
	Name      string `json:"name"`
	Network   string `json:"network"`
	Chain     string `json:"chain"`
	Contract  string `json:"contract"`
	ChainLogo string `json:"chain_logo"`
	Status    int64  `json:"status"` // 1 normal 2 maintenance 3 To be delisted
}

// GetTokenRateReq get token rate
type GetTokenRateReq struct {
	Amount  string `json:"amount"`
	TokenId string `json:"token_id"`
}

type TokenRateResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Price string `json:"price"`
		Value string `json:"value"`
	} `json:"data"`
}

// OrderInfoReq get API trade data
type OrderInfoReq struct {
	MerchantOrderId string `json:"merchant_order_id"`
}

type BillInfoResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Detail struct {
			FiatAmount      string `json:"fiat_amount"`
			FiatCurrency    string `json:"fiat_currency"`
			ProductName     string `json:"product_name"`
			MerchantOrderId string `json:"merchant_order_id"`
			Chain           string `json:"chain"`
			Contract        string `json:"contract"`
			Crypto          string `json:"crypto"`
			OriginAmount    string `json:"origin_amount"`
			Status          string `json:"status"`
			Created         int64  `json:"created"`
		} `json:"detail"`
		Received []struct {
			Amount     string `json:"amount"`
			Chain      string `json:"chain"`
			Contract   string `json:"contract"`
			Crypto     string `json:"crypto"`
			ServiceFee string `json:"service_fee"`
			PayTime    int64  `json:"pay_time"`
			TokenRate  string `json:"token_rate"`
		} `json:"received"`
		Refunds []struct {
			Amount               string  `json:"amount"`
			NetworkFee           string  `json:"network_fee"`
			ActualReceivedAmount string  `json:"actual_received_amount"`
			Chain                string  `json:"chain"`
			Contract             string  `json:"contract"`
			Crypto               string  `json:"crypto"`
			Txid                 *string `json:"txid"`
			Address              string  `json:"address"`
			PayTime              int64   `json:"pay_time"`
			Status               string  `json:"status"`
		} `json:"refunds"`
	} `json:"data"`
}

type WithdrawReq struct {
	TokenID         string `json:"token_id"  binding:"required"`
	Address         string `json:"address,omitempty" binding:"required"`
	Memo            string `json:"memo,omitempty"`
	Value           string `json:"value,omitempty" binding:"required"`
	MerchantOrderId string `json:"merchant_order_id"  binding:"required"`
}

type WithdrawResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderId    string `json:"order_id"`
		NetworkFee string `json:"network_fee"`
	} `json:"data"`
}

// ccpay
type CheckUserReq struct {
	CId string `json:"c_id"  binding:"required"`
}

type CheckUserResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CId      string `json:"c_id"` // cwallet id
		Nickname string `json:"nickname"`
	} `json:"data"`
}

// ccpay
type AssetsReq struct {
	TokenId string `json:"token_id"` // 某个币的tokenID
}

type AssetEntity struct {
	TokenId string `json:"token_id"`
	Crypto  string `json:"crypto"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Price   string `json:"price"`
	Logo    string `json:"logo"`
}
type AssetsResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []*AssetEntity `json:"list"`
	} `json:"data"`
}

// ccpay
type NetworkFeeReq struct {
	TokenId string `json:"token_id"  binding:"required"`
	Address string `json:"address"`
	Memo    string `json:"memo"`
}
type NetworkFeeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TokenId string `json:"token_id"`
		Crypto  string `json:"crypto"`
		Fee     string `json:"fee"`
	} `json:"data"`
}

type SupportCoinReq struct {
}

type SupportCoinResultData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []SupportToken `json:"list"`
	} `json:"data"`
}
