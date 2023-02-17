package util

type SubmitCreateTradeOrderRequest struct {
	CcpaymentId string       `json:"ccpayment_id" binding:"required"`
	Appid       string       `json:"app_id" binding:"required"`
	Timestamp   int64        `json:"timestamp" binding:"required"`
	JsonContent *JsonContent `json:"json_content" binding:"required"`
	Sign        string       `json:"sign" binding:"required"`
	NotifyUrl   string       `json:"notify_url" binding:"required"`
	Remark      string       `json:"remark"`
	Device      string       `json:"device" binding:"required"`
	Noncestr    string       `json:"noncestr" binding:"required"` // 5 digit random string
}

type JsonContent struct {
	TokenId    string `json:"token_id" binding:"required"`
	Chain      string `json:"chain" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
	Contract   string `json:"contract" binding:"required"`
	OutOrderNo string `json:"out_order_no" binding:"required"`
	FiatName   string `json:"fiat_name" binding:"required"`
}

type EncryptData struct {
	AppId       string              `json:"app_id"`       // Merchant appid
	Timestamp   int64               `json:"timestamp"`    // Current timestamp and it is accurate to the second
	OutOrderNo  string              `json:"out_order_no"` // Merchant order ID
	PayStatus   string              `json:"pay_status"`   // Payment status: success
	Sign        string              `json:"sign"`
	Noncestr    string              `json:"noncestr"` // 5 digit random string
	JsonContent *EncryptJsonContent `json:"json_content"`
}

type EncryptJsonContent struct {
	OriginAmount string `json:"origin_amount"` // Token amount to be paid
	FiatAmount   string `json:"fiat_amount"`   // Amount of fiat currency submitted by the merchant; USD by default
	PaidAmount   string `json:"paid_amount"`   // Token amount paid by the customer
	CurrentRate  string `json:"current_rate"`  // Token price when generating the order
	Chain        string `json:"chain"`         //
	Contract     string `json:"contract"`      // Contract
	OrderNo      string `json:"order_no"`      // ccpayment order ID
	Symbol       string `json:"symbol"`        // Token symbol
}
