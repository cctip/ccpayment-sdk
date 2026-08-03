package ccpayment

// ==================== User Transfer Models ====================

type ToUser struct {
	UserId string `json:"userId"`
	Amount string `json:"amount"`
}

type TransferRecord struct {
	RecordId     string `json:"recordId"`
	CoinId       uint64 `json:"coinId"`
	CoinSymbol   string `json:"coinSymbol"`
	OrderId      string `json:"orderId"`
	FromUserId   string `json:"fromUserId"`
	ToUserId     string `json:"toUserId"`
	Amount       string `json:"amount"`
	Status       string `json:"status"`
	Remark       string `json:"remark,omitempty"`
	CoinUSDPrice string `json:"coinUSDPrice"`
}

type BatchTransferRecord struct {
	RecordId     string   `json:"recordId"`
	UserId       string   `json:"userId"`
	CoinId       uint64   `json:"coinId"`
	CoinSymbol   string   `json:"coinSymbol"`
	OrderId      string   `json:"orderId"`
	ToUsers      []ToUser `json:"toUsers"`
	Status       string   `json:"status"`
	Remark       string   `json:"remark,omitempty"`
	CoinUSDPrice string   `json:"coinUSDPrice"`
}

// Request/Response for User Transfer

type UserTransferRequest struct {
	OrderId    string `json:"orderId"`
	FromUserId string `json:"fromUserId"`
	ToUserId   string `json:"toUserId"`
	CoinId     uint64 `json:"coinId"`
	Amount     string `json:"amount"`
	Remark     string `json:"remark,omitempty"`
}

type UserTransferResponse struct {
	RecordId string `json:"recordId"`
}

type GetUserTransferRecordRequest struct {
	RecordId string `json:"recordId,omitempty"`
	OrderId  string `json:"orderId,omitempty"`
}

type GetUserTransferRecordResponse struct {
	Record TransferRecord `json:"record"`
}

type GetUserTransferRecordListRequest struct {
	OrderIds   []string `json:"orderIds,omitempty"`
	FromUserId string   `json:"fromUserId,omitempty"`
	ToUserId   string   `json:"toUserId,omitempty"`
	CoinId     uint64   `json:"coinId,omitempty"`
	StartAt    int64    `json:"startAt,omitempty"`
	EndAt      int64    `json:"endAt,omitempty"`
	NextId     string   `json:"nextId,omitempty"`
}

type GetUserTransferRecordListResponse struct {
	Records []TransferRecord `json:"records"`
	NextId  string           `json:"nextId,omitempty"`
}

type UserBatchTransferRequest struct {
	OrderId string   `json:"orderId"`
	UserId  string   `json:"userId"`
	ToUsers []ToUser `json:"toUsers"`
	CoinId  uint64   `json:"coinId"`
	Remark  string   `json:"remark,omitempty"`
}

type UserBatchTransferResponse struct {
	RecordId string `json:"recordId"`
}

type GetUserBatchTransferRecordRequest struct {
	RecordId string `json:"recordId,omitempty"`
	OrderId  string `json:"orderId,omitempty"`
}

type GetUserBatchTransferRecordResponse struct {
	Record BatchTransferRecord `json:"record"`
}

type GetUserBatchTransferRecordListRequest struct {
	OrderIds []string `json:"orderIds,omitempty"`
	UserId   string   `json:"userId,omitempty"`
	CoinId   uint64   `json:"coinId,omitempty"`
	StartAt  int64    `json:"startAt,omitempty"`
	EndAt    int64    `json:"endAt,omitempty"`
	NextId   string   `json:"nextId,omitempty"`
}

type GetUserBatchTransferRecordListResponse struct {
	Records []BatchTransferRecord `json:"records"`
	NextId  string                `json:"nextId,omitempty"`
}

type GetUserBatchTransferRecordDetailRequest struct {
	RecordId string `json:"recordId"`
	NextId   string `json:"nextId,omitempty"`
}

type GetUserBatchTransferRecordDetailResponse struct {
	Records []TransferRecord `json:"records"`
	NextId  string           `json:"nextId,omitempty"`
}

// ==================== Orders Models ====================

type PaidInfo struct {
	RecordId         string `json:"recordId"`
	FromAddress      string `json:"fromAddress"`
	Amount           string `json:"amount"`
	ServiceFee       string `json:"serviceFee"`
	Txid             string `json:"txid"`
	Status           string `json:"status"`
	ArrivedAt        int64  `json:"arrivedAt"`
	Rate             string `json:"rate"`
	IsFlaggedAsRisky bool   `json:"isFlaggedAsRisky"`
}

type InvoicePaidInfo struct {
	RecordId         string `json:"recordId"`
	CoinId           uint64 `json:"coinId"`
	CoinSymbol       string `json:"coinSymbol"`
	Chain            string `json:"chain"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	ToMemo           string `json:"toMemo"`
	PaidAmount       string `json:"paidAmount"`
	ServiceFee       string `json:"serviceFee"`
	Txid             string `json:"txid"`
	Status           string `json:"status"`
	ArrivedAt        int64  `json:"arrivedAt"`
	Rate             string `json:"rate"`
	PaidValue        string `json:"paidValue"`
	IsFlaggedAsRisky bool   `json:"isFlaggedAsRisky"`
}

// Request/Response for Orders

type GetAppOrderInfoRequest struct {
	OrderId string `json:"orderId"`
}

type GetAppOrderInfoResponse struct {
	AmountToPay   string     `json:"amountToPay"`
	CoinId        uint64     `json:"coinId"`
	CoinSymbol    string     `json:"coinSymbol"`
	Chain         string     `json:"chain"`
	ToAddress     string     `json:"toAddress"`
	ToMemo        string     `json:"toMemo"`
	CreateAt      int64      `json:"createAt"`
	Rate          string     `json:"rate"`
	FiatId        uint64     `json:"fiatId"`
	FiatSymbol    string     `json:"fiatSymbol"`
	ExpiredAt     int64      `json:"expiredAt"`
	CheckoutUrl   string     `json:"checkoutUrl"`
	BuyerEmail    string     `json:"buyerEmail"`
	PaidList      []PaidInfo `json:"paidList"`
}

type CreateInvoiceUrlRequest struct {
	OrderId    string `json:"orderId"`
	Product    string `json:"product"`
	PriceFiatId uint64 `json:"priceFiatId,omitempty"`
	PriceCoinId uint64 `json:"priceCoinId,omitempty"`
	Price      string `json:"price"`
	BuyerEmail string `json:"buyerEmail,omitempty"`
	ReturnUrl  string `json:"returnUrl,omitempty"`
	ExpiredAt  int64  `json:"expiredAt,omitempty"`
	CloseUrl   string `json:"closeUrl,omitempty"`
	NotifyUrl  string `json:"notifyUrl,omitempty"`
}

type CreateInvoiceUrlResponse struct {
	InvoiceUrl string `json:"invoiceUrl"`
}

type GetInvoiceOrderInfoRequest struct {
	OrderId string `json:"orderId"`
}

type GetInvoiceOrderInfoResponse struct {
	OrderId        string            `json:"orderId"`
	CreateAt       int64             `json:"createAt"`
	Product        string            `json:"product"`
	Price          string            `json:"price"`
	PriceCoinId    uint64            `json:"priceCoinId"`
	PriceFiatId    uint64            `json:"priceFiatId"`
	PriceSymbol    string            `json:"priceSymbol"`
	InvoiceUrl     string            `json:"invoiceUrl"`
	BuyerEmail     string            `json:"buyerEmail"`
	ExpiredAt      int64             `json:"expiredAt"`
	TotalPaidValue string            `json:"totalPaidValue"`
	PaidList       []InvoicePaidInfo `json:"paidList"`
}

type GetWebhookInfoResponse struct {
	WebhookUrl    string `json:"webhookUrl"`
	WebhookSecret string `json:"webhookSecret"`
}
