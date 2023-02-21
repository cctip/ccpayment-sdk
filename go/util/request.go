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
	AppId        string         `json:"app_id"`
	Timestamp    int64          `json:"timestamp"`     //当前时间戳
	Noncestr     string         `json:"noncestr"`      // 随机字符串
	PayStatus    string         `json:"pay_status"`    // 支付状态
	Sign         string         `json:"sign"`          // 使用 sha-256 的签名字段
	BillType     string         `json:"bill_type"`     // Api、Invoice
	RecordId     string         `json:"record_id"`     // ccpyament的交易记录，唯一
	OrderNo      string         `json:"order_no"`      // ccpyament的订单号码
	OriginPrice  string         `json:"origin_price"`  // 原始定价 法币或者代币，商家自己提供的 100人民币
	OriginAmount string         `json:"origin_amount"` // 本身该支付的代币，
	FiatRate     string         `json:"fiat_rate"`     // 法币价格 单位USD
	PaidAmount   string         `json:"paid_amount"`   // 支付了的代币
	TokenRate    string         `json:"token_rate"`    // 支付token价格，单位USD
	Chain        string         `json:"chain"`
	Contract     string         `json:"contract"`
	Symbol       string         `json:"symbol"`
	Extend       *WebhookExtend `json:"extend"` // 扩展信息
}

type WebhookExtend struct {
	InvoiceId  string `json:"invoice_id,omitempty"`   // invoice 使用
	UserEmail  string `json:"user_email,omitempty"`   // invoice 使用
	OutOrderNo string `json:"out_order_no,omitempty"` // api 使用
}
