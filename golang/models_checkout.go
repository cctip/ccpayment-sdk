package ccpayment

// ==================== Checkout Models ====================

type CheckoutCoin struct {
	CoinId  uint64   `json:"coinId"`
	Symbol  string   `json:"symbol"`
	LogoUrl string   `json:"logoUrl"`
	Chains  []string `json:"chains"`
}

// Request/Response for Checkout

type CreateCheckoutUrlRequest struct {
	OrderId     string `json:"orderId"`
	Product     string `json:"product"`
	PriceFiatId uint64 `json:"priceFiatId,omitempty"`
	PriceCoinId uint64 `json:"priceCoinId,omitempty"`
	Price       string `json:"price"`
	BuyerEmail  string `json:"buyerEmail,omitempty"`
	ReturnUrl   string `json:"returnUrl,omitempty"`
	ExpiredAt   int64  `json:"expiredAt,omitempty"`
	CloseUrl    string `json:"closeUrl,omitempty"`
	NotifyUrl   string `json:"notifyUrl,omitempty"`
}

type CreateCheckoutUrlResponse struct {
	CheckoutUrl string `json:"checkoutUrl"`
}

type HostedOrderInfoRequest struct {
	OrderId string `json:"orderId"`
}

type HostedOrderInfoResponse struct {
	OrderId        string            `json:"orderId"`
	Product        string            `json:"product"`
	Price          string            `json:"price"`
	PriceCoinId    uint64            `json:"priceCoinId"`
	PriceFiatId    uint64            `json:"priceFiatId"`
	PriceSymbol    string            `json:"priceSymbol"`
	CheckoutUrl    string            `json:"checkoutUrl"`
	BuyerEmail     string            `json:"buyerEmail"`
	ExpiredAt      int64             `json:"expiredAt"`
	Step           string            `json:"step"`
	SelectedCoinId uint64            `json:"selectedCoinId"`
	SelectedChain  string            `json:"selectedChain"`
	ToAddress      string            `json:"toAddress"`
	ToMemo         string            `json:"toMemo"`
	AmountToPay    string            `json:"amountToPay"`
	Rate           string            `json:"rate"`
	TotalPaidValue string            `json:"totalPaidValue"`
	PaidList       []InvoicePaidInfo `json:"paidList"`
	RefundList     []interface{}     `json:"refundList"`
}

type HostedOrderInfoFirstRequest struct {
	OrderId string `json:"orderId"`
}

type CreateHostedOrderDepositRequest struct {
	OrderId string `json:"orderId"`
	CoinId  uint64 `json:"coinId"`
	Chain   string `json:"chain"`
}

type CreateHostedOrderDepositResponse struct {
	Address        string `json:"address"`
	Memo           string `json:"memo,omitempty"`
	AmountToPay    string `json:"amountToPay"`
	ConfirmsNeeded uint64 `json:"confirmsNeeded"`
}

type GetHostedCoinUSDTPriceRequest struct {
	CoinIds []uint64 `json:"coinIds"`
}

type GetHostedCoinUSDTPriceResponse struct {
	Prices map[string]string `json:"prices"`
}

type GetMainCoinListResponse struct {
	CoinChainList []CheckoutCoin `json:"coinChainList"`
}

type CreateAppDemoOrderDepositRequest struct {
	OrderId string `json:"orderId"`
	CoinId  uint64 `json:"coinId"`
	Chain   string `json:"chain"`
	Price   string `json:"price"`
}

type CreateAppDemoOrderDepositResponse struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
	Memo    string `json:"memo,omitempty"`
}

type ConfirmTradeRequest struct {
	OrderId string `json:"orderId"`
	TxId    string `json:"txId"`
}

// ==================== Swap Models ====================

type SwapCoin struct {
	CoinId  uint64   `json:"coinId"`
	Symbol  string   `json:"symbol"`
	LogoUrl string   `json:"logoUrl"`
	Chains  []string `json:"chains"`
}

type SwapEstimate struct {
	FromCoinId     uint64 `json:"fromCoinId"`
	FromCoinSymbol string `json:"fromCoinSymbol"`
	FromAmount     string `json:"fromAmount"`
	ToCoinId       uint64 `json:"toCoinId"`
	ToCoinSymbol   string `json:"toCoinSymbol"`
	ToAmount       string `json:"toAmount"`
	Rate           string `json:"rate"`
}

type SwapRecord struct {
	RecordId       string `json:"recordId"`
	OrderId        string `json:"orderId"`
	FromCoinId     uint64 `json:"fromCoinId"`
	FromCoinSymbol string `json:"fromCoinSymbol"`
	FromAmount     string `json:"fromAmount"`
	ToCoinId       uint64 `json:"toCoinId"`
	ToCoinSymbol   string `json:"toCoinSymbol"`
	ToAmount       string `json:"toAmount"`
	Rate           string `json:"rate"`
	Status         string `json:"status"`
	CreatedAt      int64  `json:"createdAt"`
}

// Request/Response for Swap

type GetSwapCoinListResponse struct {
	Coins []SwapCoin `json:"coins"`
}

type SwapEstimateRequest struct {
	FromCoinId uint64 `json:"fromCoinId"`
	ToCoinId   uint64 `json:"toCoinId"`
	FromAmount string `json:"fromAmount"`
}

type SwapEstimateResponse struct {
	SwapEstimate
}

type SwapRequest struct {
	OrderId    string `json:"orderId"`
	FromCoinId uint64 `json:"fromCoinId"`
	ToCoinId   uint64 `json:"toCoinId"`
	FromAmount string `json:"fromAmount"`
}

type SwapResponse struct {
	RecordId string `json:"recordId"`
}

type GetSwapRecordRequest struct {
	RecordId string `json:"recordId,omitempty"`
	OrderId  string `json:"orderId,omitempty"`
}

type GetSwapRecordResponse struct {
	Record SwapRecord `json:"record"`
}

type GetSwapRecordListRequest struct {
	OrderIds   []string `json:"orderIds,omitempty"`
	FromCoinId uint64   `json:"fromCoinId,omitempty"`
	ToCoinId   uint64   `json:"toCoinId,omitempty"`
	StartAt    int64    `json:"startAt,omitempty"`
	EndAt      int64    `json:"endAt,omitempty"`
	NextId     string   `json:"nextId,omitempty"`
}

type GetSwapRecordListResponse struct {
	Records []SwapRecord `json:"records"`
	NextId  string       `json:"nextId,omitempty"`
}

// ==================== Utilities Models ====================

// Request/Response for Utilities

type WebhookResendRequest struct {
	StartTimestamp  int64    `json:"startTimestamp,omitempty"`
	EndTimestamp    int64    `json:"endTimestamp,omitempty"`
	WebhookResult   string   `json:"webhookResult,omitempty"`
	TransactionType string   `json:"transactionType,omitempty"`
	RecordIDs       []string `json:"recordIds,omitempty"`
}

type WebhookResendResponse struct {
	ResendCount *int64 `json:"resendCount,omitempty"`
}

type GetTradeBlockHeightRequest struct {
	RecordID string `json:"record_id"`
}

type GetTradeBlockHeightResponse struct {
	Chain            *string `json:"chain,omitempty"`
	TxBlockHeight    *int64  `json:"txBlockHeight,omitempty"`
	CurrBlockHeight  *int64  `json:"currBlockHeight,omitempty"`
	ReqConfirmations *int64  `json:"reqConfirmations,omitempty"`
}

type CheckWithdrawalAddressValidityRequest struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

type CheckWithdrawalAddressValidityResponse struct {
	AddrIsValid *bool `json:"addrIsValid,omitempty"`
}

type DeprecatedAddressRequest struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

type UnboundAddressDetail struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

type DeprecatedAddressResponse struct {
	Unbound     []UnboundAddressDetail `json:"unbound"`
	UnboundAt   int64                  `json:"unboundAt"`
	UserID      string                 `json:"userID,omitempty"`
	ReferenceID string                 `json:"referenceID,omitempty"`
}

type RescanLostTransactionRequest struct {
	Chain     string `json:"chain"`
	ToAddress string `json:"toAddress"`
	TxID      string `json:"txId"`
	Memo      string `json:"memo,omitempty"`
}

type RescanLostTransactionResponse struct {
	Description string `json:"description"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}
