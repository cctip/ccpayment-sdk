package main

type SubmitCreateTradeOrderRequest struct {
	CcpaymentId string       `json:"ccpayment_id" binding:"required"`
	Appid       string       `json:"app_id" binding:"required"`
	Timestamp   int64        `json:"timestamp" binding:"required"`
	JsonContent *JsonContent `json:"json_content" binding:"required"`
	Sign        string       `json:"sign" binding:"required"`
	NotifyUrl   string       `json:"notify_url" binding:"required"`
	Remark      string       `json:"remark"`
	Device      string       `json:"device" binding:"required"`
	Noncestr    string       `json:"noncestr" binding:"required"`
}

type Request struct {
	CcpaymentId   string       `json:"ccpayment_id"`
	TransactionId string       `json:"transaction_id"`
	Msg           string       `json:"msg"`
	Appid         string       `json:"appid"`
	Timestamp     string       `json:"timestamp"`
	JsonContent   *JsonContent `json:"json_content"`
}

type JsonContent struct {
	Chain      string `json:"chain" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
	Contract   string `json:"contract" binding:"required"`
	OutOrderNo string `json:"out_order_no" binding:"required"`
	//Symbol     string `json:"symbol" binding:"required"`
	//PayAddress   string `json:"pay_address" binding:"required"`
	TokenId  string `json:"token_id" binding:"required"`
	FiatName string `json:"fiat_name" binding:"required"`
}
