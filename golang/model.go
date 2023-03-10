package golang

// CreateOrderReq api create order params
type CreateOrderReq struct {
	TokenId         string `json:"token_id"`
	Amount          string `json:"amount"`
	MerchantOrderId string `json:"merchant_order_id"`
	FiatCurrency    string `json:"fiat_currency"` // USD
	Remark          string `json:"remark"`
}

// CreateOrderResp api create order response
type CreateOrderResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data OrderInfo `json:"data"`
}

type OrderInfo struct {
	OrderId    string `json:"order_id"`
	Amount     string `json:"amount"`
	Logo       string `json:"logo"`
	Network    string `json:"network"`
	PayAddress string `json:"pay_address"`
	Crypto     string `json:"crypto"`
}

// CheckoutUrlReq checkout url params
type CheckoutUrlReq struct {
	ValidTimestamp  int64  `json:"valid_timestamp"` // s
	Amount          string `json:"amount"`
	MerchantOrderId string `json:"merchant_order_id"`
	ProductName     string `json:"product_name"`
	ReturnUrl       string `json:"return_url"`
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
	PayStatus    string         `json:"pay_status"`
	OrderType    string         `json:"order_type"`
	RecordId     string         `json:"record_id"`
	OrderID      string         `json:"order_id"`
	OriginPrice  string         `json:"origin_price"`
	OriginAmount string         `json:"origin_amount"`
	FiatRate     string         `json:"fiat_rate"`
	PaidAmount   string         `json:"paid_amount"`
	TokenRate    string         `json:"token_rate"`
	Chain        string         `json:"chain"`
	Contract     string         `json:"contract"`
	Crypto       string         `json:"crypto"`
	Extend       *WebhookExtend `json:"extend"`
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
	Crypto  string `json:"crypto"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	Min     string `json:"min"`
	Price   string `json:"price"`
	TokenId string `json:"token_id"`
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

// OrderTradeReq get API trade data
type OrderTradeReq struct {
	MerchantOrderId string `json:"merchant_order_id"`
}

type BillTradeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		FiatAmount      string `json:"fiat_amount"`
		FiatCurrency    string `json:"fiat_currency"`
		ProductName     string `json:"product_name"`
		MerchantOrderId string `json:"merchant_order_id"`
		Chain           string `json:"chain"`
		Contract        string `json:"contract"`
		Crypto          string `json:"crypto"`
		List            []struct {
			Amount     string `json:"amount"`
			Chain      string `json:"chain"`
			Contract   string `json:"contract"`
			Crypto     string `json:"crypto"`
			NetworkFee string `json:"network_fee"`
			Network    string `json:"network"`
			PayTime    int64  `json:"pay_time"`
			Txid       string `json:"txid"`
			TokenRate  string `json:"token_rate"`
		} `json:"list"`
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
		Type       string `json:"type"`
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
