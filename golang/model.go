package golang

// CreateOrderReq api create order params
type CreateOrderReq struct {
	Remark              string `json:"remark"`
	TokenId             string `validate:"required" json:"token_id"`
	ProductPrice        string `validate:"required" json:"product_price"`
	MerchantOrderId     string `validate:"required" json:"merchant_order_id"`
	DenominatedCurrency string `validate:"required" json:"denominated_currency"`
	OrderValidPeriod    int    `json:"order_valid_period"`
	NotifyUrl           string `json:"notify_url"`
	CustomValue         string `json:"custom_value"`
}

// CreateOrderResp api create order response
type CreateOrderResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data OrderInfo `json:"data"`
}

type OrderInfo struct {
	Remark              string `json:"remark"`
	TokenId             string `validate:"required" json:"token_id"`
	ProductPrice        string `validate:"required" json:"product_price"`
	MerchantOrderId     string `validate:"required" json:"merchant_order_id"`
	DenominatedCurrency string `validate:"required" json:"denominated_currency"`
	OrderValidPeriod    int    `json:"order_valid_period"`
	NotifyUrl           string `json:"notify_url"`
	CustomValue         string `json:"custom_value"`
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
	Symbol string        `json:"symbol"`
	Crypto string        `json:"crypto"`
	Name   string        `json:"name"`
	Logo   string        `json:"logo"`
	Min    string        `json:"min"`
	Price  string        `json:"price"`
	CoinId string        `json:"coin_id,omitempty"`
	Status int64         `json:"status"` // 1 normal 2 maintenance 3 To be delisted
	Tokens []*TokenChain `json:"tokens,omitempty"`
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
	TokenId       string `json:"token_id"`
	Crypto        string `json:"crypto"`
	Logo          string `json:"logo"`
	Name          string `json:"name"`
	IsSupportMemo bool   `json:"is_support_memo"`
	Network       string `json:"network"`
	Chain         string `json:"chain"`
	Contract      string `json:"contract"`
	ChainLogo     string `json:"chain_logo"`
	Status        int64  `json:"status"` // 1 normal 2 maintenance 3 To be delisted
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
	MerchantOrderIds []string `json:"merchant_order_ids"`
}

type BillInfoResp struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data []*BillInfoEntity `json:"data"`
}
type BillInfoEntity struct {
	OrderDetail struct {
		ProductPrice        string `json:"product_price"`
		DenominatedCurrency string `json:"denominated_currency"`
		ProductName         string `json:"product_name"`
		MerchantOrderId     string `json:"merchant_order_id"`
		Chain               string `json:"chain"`
		Contract            string `json:"contract"`
		Crypto              string `json:"crypto"`
		OrderAmount         string `json:"order_amount"`
		Status              string `json:"status"`
		TokenRate           string `json:"token_rate"`
		Created             int64  `json:"created"`
	} `json:"order_detail"`
	TradeList []struct {
		Amout      string `json:"amout"`
		Chain      string `json:"chain"`
		Contract   string `json:"contract"`
		Crypto     string `json:"crypto"`
		ServiceFee string `json:"service_fee"`
		NetworkFee string `json:"network_fee"`
		Txid       string `json:"txid"`
		PayTime    int64  `json:"pay_time"`
		TokenRate  string `json:"token_rate"`
	} `json:"trade_list"`
	RefundList []struct {
		RefundAmount         string `json:"refund_amount"`
		NetworkFee           string `json:"network_fee"`
		ActualReceivedAmount string `json:"actual_received_amount"`
		Chain                string `json:"chain"`
		Contract             string `json:"contract"`
		Crypto               string `json:"crypto"`
		Txid                 string `json:"txid"`
		Address              string `json:"address"`
		PayTime              int64  `json:"pay_time"`
		Status               string `json:"status"`
	} `json:"refund_list"`
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

type NetworkChainHeightInfoReq struct {
}

type NetworkChainHeightInfoResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Chain              string  `json:"chain"`
		CurrentChainHeight int64   `json:"current_chain_height"`
		TxConfirmBlockNum  int64   `json:"tx_confirm_block_num"`
		BlockRate          float64 `json:"block_rate"`
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

type AddressReq struct {
	UserId    string `validate:"required" json:"user_id"`
	Chain     string `validate:"required" json:"chain"`
	NotifyUrl string `json:"notify_url"`
}

type AddressResq struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Address string `json:"address,omitempty"`
		Memo    string `json:"memo,omitempty"`
	} `json:"data,omitempty"`
}
