package ccpayment

// ==================== Basic Info Models ====================

// Coin represents a token with its network information
type Coin struct {
	CoinID     uint64                  `json:"coinId"`
	Symbol     string                  `json:"symbol"`
	CoinFullName string                `json:"coinFullName"`
	LogoUrl    string                  `json:"logoUrl"`
	Status     string                  `json:"status"`
	Networks   map[string]NetworkInfo  `json:"networks"`
	Price      string                  `json:"price"`
}

type NetworkInfo struct {
	Chain                 string `json:"chain"`
	ChainFullName         string `json:"chainFullName"`
	Contract              string `json:"contract"`
	Precision             uint32 `json:"precision"`
	CanDeposit            bool   `json:"canDeposit"`
	CanWithdraw           bool   `json:"canWithdraw"`
	MinimumDepositAmount  string `json:"minimumDepositAmount"`
	MinimumWithdrawAmount string `json:"minimumWithdrawAmount"`
	MaximumWithdrawAmount string `json:"maximumWithdrawAmount"`
	IsSupportMemo         bool   `json:"isSupportMemo"`
	Protocol              string `json:"protocol"`
}

type Fiat struct {
	FiatId  uint64 `json:"fiatId"`
	Symbol  string `json:"symbol"`
	LogoUrl string `json:"logoUrl"`
	Mark    string `json:"mark"`
	UsdRate string `json:"usdRate"`
}

type Chain struct {
	Chain          string `json:"chain"`
	ChainFullName  string `json:"chainFullName"`
	Explorer       string `json:"explorer"`
	LogoUrl        string `json:"logoUrl"`
	Status         string `json:"status"`
	Confirmations  int32  `json:"confirmations"`
	BaseUrl        string `json:"baseUrl"`
	IsEVM          bool   `json:"isEVM"`
	SupportMemo    bool   `json:"supportMemo"`
}

type ChainSimple struct {
	Chain          string `json:"chain"`
	ChainFullName  string `json:"chainFullName"`
	Explorer       string `json:"explorer"`
	LogoUrl        string `json:"logoUrl"`
	Status         string `json:"status"`
	ConfirmNum     int32  `json:"confirmNum"`
	IsEVM          bool   `json:"isEVM"`
}

type Fee struct {
	CoinId       uint64 `json:"coinId"`
	CoinSymbol   string `json:"coinSymbol"`
	Amount       string `json:"amount"`
}

// Request/Response for Basic Info

type GetCoinRequest struct {
	CoinId uint64 `json:"coinId"`
}

type GetCoinListResponse struct {
	Coins []Coin `json:"coins"`
}

type GetCoinUSDTPriceRequest struct {
	CoinIds []uint64 `json:"coinIds"`
}

type GetCoinUSDTPriceResponse struct {
	Prices map[string]string `json:"prices"`
}

type GetFiatListResponse struct {
	Fiats []Fiat `json:"fiats"`
}

type GetChainListRequest struct {
	Chains []string `json:"chains,omitempty"`
}

type GetChainListResponse struct {
	Chains []Chain `json:"chains"`
}

type AllChainRequest struct {
	Chains []string `json:"chains,omitempty"`
}

type AllChainResponse struct {
	Chains []ChainSimple `json:"chains"`
}

type GetCwalletUserIdRequest struct {
	CwalletUserId string `json:"cwalletUserId"`
}

type GetCwalletUserIdResponse struct {
	CwalletUserId   string `json:"cwalletUserId"`
	CwalletUserName string `json:"cwalletUserName"`
}

type GetWithdrawFeeRequest struct {
	CoinId uint64 `json:"coinId"`
	Chain  string `json:"chain"`
}

type GetWithdrawFeeResponse struct {
	Fee                 Fee    `json:"fee"`
	NetworkFeeInquiryID string `json:"networkFeeInquiryID"`
}

// ==================== Merchant Assets Models ====================

type Asset struct {
	CoinId     uint64 `json:"coinId"`
	CoinSymbol string `json:"coinSymbol"`
	Available  string `json:"available"`
}

// Request/Response for Merchant Assets

type GetAppCoinAssetListResponse struct {
	Assets []Asset `json:"assets"`
}

type GetAppCoinAssetRequest struct {
	CoinId uint64 `json:"coinId"`
}

type GetAppCoinAssetResponse struct {
	Asset Asset `json:"asset"`
}

// ==================== Merchant Deposit Models ====================

// Request/Response for Merchant Deposit

type CreateAppOrderDepositAddressRequest struct {
	OrderId               string `json:"orderId"`
	CoinId                uint64 `json:"coinId"`
	FiatId                uint64 `json:"fiatId,omitempty"`
	Chain                 string `json:"chain"`
	Price                 string `json:"price"`
	ExpiredAt             int64  `json:"expiredAt,omitempty"`
	BuyerEmail            string `json:"buyerEmail,omitempty"`
	GenerateCheckoutURL   bool   `json:"generateCheckoutURL,omitempty"`
	Product               string `json:"product,omitempty"`
	ReturnUrl             string `json:"returnUrl,omitempty"`
	NotifyUrl             string `json:"notifyUrl,omitempty"`
	CloseUrl              string `json:"closeUrl,omitempty"`
}

type CreateAppOrderDepositAddressResponse struct {
	Address        string `json:"address"`
	Amount         string `json:"amount"`
	Memo           string `json:"memo,omitempty"`
	CheckoutUrl    string `json:"checkoutUrl,omitempty"`
	ConfirmsNeeded uint64 `json:"confirmsNeeded"`
}

type GetOrCreateAppDepositAddressRequest struct {
	ReferenceId string `json:"referenceId"`
	Chain       string `json:"chain"`
	NotifyUrl   string `json:"notifyUrl,omitempty"`
}

type GetOrCreateAppDepositAddressResponse struct {
	Address string `json:"address"`
	Memo    string `json:"memo,omitempty"`
}

type DepositRecord struct {
	RecordId         string `json:"recordId"`
	ReferenceId      string `json:"referenceId"`
	OrderId          string `json:"orderId"`
	CoinId           uint64 `json:"coinId"`
	CoinSymbol       string `json:"coinSymbol"`
	Chain            string `json:"chain"`
	Contract         string `json:"contract"`
	CoinUSDPrice     string `json:"coinUSDPrice"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	ToMemo           string `json:"toMemo,omitempty"`
	Amount           string `json:"amount"`
	ServiceFee       string `json:"serviceFee"`
	TxId             string `json:"txId"`
	TxIndex          uint64 `json:"txIndex"`
	Status           string `json:"status"`
	ArrivedAt        int64  `json:"arrivedAt"`
	IsFlaggedAsRisky bool   `json:"isFlaggedAsRisky,omitempty"`
}

type GetAppDepositRecordRequest struct {
	RecordId string `json:"recordId"`
}

type GetAppDepositRecordResponse struct {
	Record DepositRecord `json:"record"`
}

type GetAppDepositRecordListRequest struct {
	Chain         string   `json:"chain,omitempty"`
	ReferenceId   string   `json:"referenceId,omitempty"`
	OrderId       string   `json:"orderId,omitempty"`
	ToAddress     string   `json:"toAddress,omitempty"`
	CoinId        uint64   `json:"coinId,omitempty"`
	StartAt       int64    `json:"startAt,omitempty"`
	EndAt         int64    `json:"endAt,omitempty"`
	NextId        string   `json:"nextId,omitempty"`
	RecordIds     []string `json:"recordIds,omitempty"`
	ReferenceIds  []string `json:"referenceIds,omitempty"`
	OrderIds      []string `json:"orderIds,omitempty"`
	Limit         uint64   `json:"limit,omitempty"`
}

type GetAppDepositRecordListResponse struct {
	Records []DepositRecord `json:"records"`
	NextId  string          `json:"nextId,omitempty"`
}

// ==================== Merchant Withdraw Models ====================

// Request/Response for Merchant Withdraw

type ApplyAppWithdrawToNetworkRequest struct {
	OrderId               string `json:"orderId"`
	CoinId                uint64 `json:"coinId"`
	Chain                 string `json:"chain"`
	Address               string `json:"address"`
	Memo                  string `json:"memo,omitempty"`
	Amount                string `json:"amount"`
	MerchantPayNetworkFee bool   `json:"merchantPayNetworkFee,omitempty"`
	NetworkFeeInquiryID   string `json:"networkFeeInquiryID,omitempty"`
	NotifyUrl             string `json:"notifyUrl,omitempty"`
}

type ApplyAppWithdrawToNetworkResponse struct {
	RecordId string `json:"recordId"`
}

type ApplyAppWithdrawToCwalletRequest struct {
	OrderId      string `json:"orderId"`
	CoinId       uint64 `json:"coinId"`
	CwalletUser  string `json:"cwalletUser"`
	Amount       string `json:"amount"`
}

type ApplyAppWithdrawToCwalletResponse struct {
	RecordId string `json:"recordId"`
}

type WithdrawFee struct {
	CoinId     uint64 `json:"coinId"`
	CoinSymbol string `json:"coinSymbol"`
	Amount     string `json:"amount"`
}

type WithdrawRecord struct {
	RecordId     string      `json:"recordId"`
	WithdrawType string      `json:"withdrawType"`
	AppId        string      `json:"appId"`
	CoinId       uint64      `json:"coinId"`
	CoinSymbol   string      `json:"coinSymbol"`
	Chain        string      `json:"chain"`
	FromAddress  string      `json:"fromAddress"`
	ToAddress    string      `json:"toAddress"`
	CwalletUser  string      `json:"cwalletUser"`
	OrderId      string      `json:"orderId"`
	TxId         string      `json:"txId,omitempty"`
	ToMemo       string      `json:"toMemo,omitempty"`
	Status       string      `json:"status"`
	Amount       string      `json:"amount"`
	Fee          WithdrawFee `json:"fee"`
	Reason       string      `json:"reason,omitempty"`
	CoinUSDPrice string      `json:"coinUSDPrice"`
}

type GetAppWithdrawRecordRequest struct {
	OrderId  string `json:"orderId,omitempty"`
	RecordId string `json:"recordId,omitempty"`
}

type GetAppWithdrawRecordResponse struct {
	Record WithdrawRecord `json:"record"`
}

type GetAppWithdrawRecordListRequest struct {
	Chain   string   `json:"chain,omitempty"`
	CoinId  uint64   `json:"coinId,omitempty"`
	OrderIds []string `json:"orderIds,omitempty"`
	StartAt int64    `json:"startAt,omitempty"`
	EndAt   int64    `json:"endAt,omitempty"`
	ToAddress string `json:"toAddress,omitempty"`
	NextId  string   `json:"nextId,omitempty"`
}

type GetAppWithdrawRecordListResponse struct {
	Records []WithdrawRecord `json:"records"`
	NextId  string           `json:"nextId,omitempty"`
}

type AutoWithdrawRecord struct {
	RecordId    string      `json:"recordId"`
	CoinId      uint64      `json:"coinId"`
	CoinSymbol  string      `json:"coinSymbol"`
	Chain       string      `json:"chain"`
	OrderId     string      `json:"orderId"`
	ToAddress   string      `json:"toAddress"`
	ToMemo      string      `json:"toMemo,omitempty"`
	Amount      string      `json:"amount"`
	TxId        string      `json:"txId,omitempty"`
	Status      string      `json:"status"`
	Fee         WithdrawFee `json:"fee"`
	ServiceFee  string      `json:"serviceFee"`
	CreatedAt   int64       `json:"createdAt"`
}

type GetAutoWithdrawRecordListRequest struct {
	Chain     string   `json:"chain,omitempty"`
	CoinId    uint64   `json:"coinId,omitempty"`
	RecordIds []string `json:"recordIds,omitempty"`
	StartAt   int64    `json:"startAt,omitempty"`
	EndAt     int64    `json:"endAt,omitempty"`
	ToAddress string   `json:"toAddress,omitempty"`
	NextId    string   `json:"nextId,omitempty"`
}

type GetAutoWithdrawRecordListResponse struct {
	Records []AutoWithdrawRecord `json:"records"`
	NextId  string               `json:"nextId,omitempty"`
}

type RiskyRefundRecord struct {
	RecordId    string      `json:"recordId"`
	CoinId      uint64      `json:"coinId"`
	CoinSymbol  string      `json:"coinSymbol"`
	Chain       string      `json:"chain"`
	OrderId     string      `json:"orderId"`
	ToAddress   string      `json:"toAddress"`
	ToMemo      string      `json:"toMemo,omitempty"`
	Amount      string      `json:"amount"`
	TxId        string      `json:"txId,omitempty"`
	Status      string      `json:"status"`
	Fee         WithdrawFee `json:"fee"`
	CreatedAt   int64       `json:"createdAt"`
	FromDeposit []string    `json:"fromDeposit"`
}

type GetRiskyRefundRecordListResponse struct {
	Records []RiskyRefundRecord `json:"records"`
	NextId  string              `json:"nextId,omitempty"`
}
