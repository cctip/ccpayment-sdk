package golang

// OrderParams api create order params
type OrderParams struct {
	TokenId         string `json:"token_id"`
	Amount          string `json:"amount"`
	MerchantOrderId string `json:"merchant_order_id"`
	FiatCurrency    string `json:"fiat_currency"` // USD
	Remark          string `json:"remark"`
}

// OrderResultData api create order response
type OrderResultData struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data OrderInfo `json:"data"`
}

type OrderInfo struct {
	BillId     string `json:"bill_id"`
	Amount     string `json:"amount"`
	Logo       string `json:"logo"`
	Network    string `json:"network"`
	PayAddress string `json:"pay_address"`
	Crypto     string `json:"crypto"`
}

// GetPaymentUrlParams checkout url params
type GetPaymentUrlParams struct {
	ValidTimestamp  int64  `json:"valid_timestamp"` // s
	Amount          string `json:"amount"`
	MerchantOrderId string `json:"merchant_order_id"`
	ProductName     string `json:"product_name"`
	ReturnUrl       string `json:"return_url"`
}

// PaymentUrlResultData checkout url response
type PaymentUrlResultData struct {
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

// WebhookParams webhook response
type WebhookParams struct {
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

// SupportTokenParam get support token list
type SupportTokenParam struct {
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

// TokenChainParams get token chain
type TokenChainParams struct {
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

// GetTokenRateParams get token rate
type GetTokenRateParams struct {
	Amount  string `json:"amount"`
	TokenId string `json:"token_id"`
}

type TokenRateResultData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Price string `json:"price"`
		Value string `json:"value"`
	} `json:"data"`
}

// OrderTradeParams get API trade data
type OrderTradeParams struct {
	MerchantOrderId string `json:"merchant_order_id"`
}

type BillTradeResultData struct {
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
