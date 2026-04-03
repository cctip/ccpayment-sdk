package ccpayment

// ==================== Checkout Models ====================

type CheckoutCoin struct {
	CoinId   uint64   `json:"coinId"`
	Symbol   string   `json:"symbol"`
	LogoUrl  string   `json:"logoUrl"`
	Chains   []string `json:"chains"`
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
	Coins []CheckoutCoin `json:"coins"`
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

type SelectHostedInvoiceCoinRequest struct {
	OrderId string `json:"orderId"`
	CoinId  uint64 `json:"coinId"`
	Chain   string `json:"chain"`
}

type SelectHostedInvoiceCoinResponse struct {
	Address     string `json:"address"`
	Memo        string `json:"memo,omitempty"`
	AmountToPay string `json:"amountToPay"`
}

// ==================== Utilities Models ====================

// Request/Response for Utilities

type VerifyAddressRequest struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

type VerifyAddressResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

type AbandonAddressRequest struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

type HostedInvoiceOrderInfoRequest struct {
	OrderId string `json:"orderId"`
}

type HostedInvoiceOrderInfoResponse struct {
	OrderId        string            `json:"orderId"`
	Product        string            `json:"product"`
	Price          string            `json:"price"`
	PriceSymbol    string            `json:"priceSymbol"`
	InvoiceUrl     string            `json:"invoiceUrl"`
	BuyerEmail     string            `json:"buyerEmail"`
	ExpiredAt      int64             `json:"expiredAt"`
	SelectedCoinId uint64            `json:"selectedCoinId"`
	SelectedChain  string            `json:"selectedChain"`
	ToAddress      string            `json:"toAddress"`
	ToMemo         string            `json:"toMemo"`
	AmountToPay    string            `json:"amountToPay"`
	TotalPaidValue string            `json:"totalPaidValue"`
	PaidList       []InvoicePaidInfo `json:"paidList"`
}

type GetPayInfoRequest struct {
	OrderId string `json:"orderId"`
}

type GetPayInfoResponse struct {
	OrderId    string `json:"orderId"`
	Product    string `json:"product"`
	Price      string `json:"price"`
	PriceSymbol string `json:"priceSymbol"`
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	Amount     string `json:"amount"`
	CoinSymbol string `json:"coinSymbol"`
	Chain      string `json:"chain"`
	QrCode     string `json:"qrCode"`
	ExpiredAt  int64  `json:"expiredAt"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}
