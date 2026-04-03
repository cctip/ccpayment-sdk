package ccpayment

// ==================== Merchant Batch Withdraw Models ====================

type AddressInfo struct {
	Address string  `json:"address"`
	Memo    string  `json:"memo,omitempty"`
	Seq     uint32  `json:"seq"`
	Codes   []int32 `json:"codes,omitempty"`
}

type AddressInfoResult struct {
	Seq   uint32  `json:"seq"`
	Codes []int32 `json:"codes"`
}

type BatchWithdrawOrder struct {
	Seq       uint32 `json:"seq"`
	Address   string `json:"address"`
	Memo      string `json:"memo,omitempty"`
	Amount    string `json:"amount"`
	Remark    string `json:"remark,omitempty"`
	RecordId  string `json:"recordId,omitempty"`
	OrderId   string `json:"orderId,omitempty"`
	Status    string `json:"status,omitempty"`
	NetworkFee string `json:"networkFee,omitempty"`
	TxId      string `json:"txId,omitempty"`
	Reason    string `json:"reason,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

type BatchWithdrawStats struct {
	Total      int32  `json:"total"`
	Succeeded  int32  `json:"succeeded"`
	Failed     int32  `json:"failed"`
	Canceled   int32  `json:"canceled"`
	Processing int32  `json:"processing"`
	ExecSeq    uint32 `json:"execSeq"`
}

type BatchWithdrawCoin struct {
	CoinId     uint64 `json:"coin_id"`
	CoinSymbol string `json:"coin_symbol"`
	CoinPrice  string `json:"coin_price"`
}

// Request/Response for Merchant Batch Withdraw

type CheckWithdrawAddressRequest struct {
	Chain           string        `json:"chain"`
	AddressInfoList []AddressInfo `json:"addressInfoList"`
}

type CheckWithdrawAddressResponse struct {
	AddressInfoResults []AddressInfoResult `json:"addressInfoResults"`
}

type ApplyBatchWithdrawRequest struct {
	BatchOrderId string             `json:"batchOrderId"`
	CoinId       uint64             `json:"coinId"`
	Chain        string             `json:"chain"`
	ProductName  string             `json:"productName,omitempty"`
	Orders       []BatchWithdrawOrder `json:"orders,omitempty"`
	Mode         string             `json:"mode"`
	NotifyUrl    string             `json:"notifyUrl,omitempty"`
}

type ApplyBatchWithdrawResponse struct {
	BatchOrderId string `json:"batchOrderId"`
	BillId       string `json:"billId"`
}

type AppendBatchWithdrawRequest struct {
	BatchOrderId string             `json:"batchOrderId"`
	Orders       []BatchWithdrawOrder `json:"orders"`
}

type ConfirmBatchWithdrawRequest struct {
	BatchOrderId  string `json:"batchOrderId"`
	DelaySeconds  int64  `json:"delaySeconds,omitempty"`
}

type ConfirmBatchWithdrawResponse struct {
	BatchOrderId   string             `json:"batchOrderId"`
	ProductName    string             `json:"productName"`
	BillId         string             `json:"billId"`
	Coin           BatchWithdrawCoin  `json:"coin"`
	Amount         string             `json:"amount"`
	NetworkFee     string             `json:"networkFee"`
	NetworkFeeCoin BatchWithdrawCoin  `json:"networkFeeCoin"`
	Status         string             `json:"status"`
	Reason         string             `json:"reason,omitempty"`
	Mode           string             `json:"mode"`
	Stats          BatchWithdrawStats `json:"stats"`
	CreatedAt      int64              `json:"createdAt"`
	UpdatedAt      int64              `json:"updatedAt"`
}

type AbortBatchWithdrawRequest struct {
	BatchOrderId string   `json:"batchOrderId"`
	Seqs         []uint32 `json:"seqs,omitempty"`
}

type AbortBatchWithdrawResponse struct {
	BatchOrderId   string   `json:"batchOrderId"`
	CanceledSeqs   []uint32 `json:"canceledSeqs"`
	IgnoredSeqs    []uint32 `json:"ignoredSeqs"`
}

type GetBatchWithdrawOrderRequest struct {
	BatchOrderId string `json:"batchOrderId"`
	Verbose      uint32 `json:"verbose,omitempty"`
}

type GetBatchWithdrawOrderRecordListRequest struct {
	BatchOrderId string `json:"batchOrderId"`
	NextId       string `json:"nextId,omitempty"`
	Limit        uint64 `json:"limit,omitempty"`
}

type GetBatchWithdrawOrderRecordListResponse struct {
	NextId  []BatchWithdrawOrder `json:"records"`
	Records []BatchWithdrawOrder `json:"records"`
}

// ==================== User Assets Models ====================

// Request/Response for User Assets

type GetUserCoinAssetListRequest struct {
	UserId string `json:"userId"`
}

type GetUserCoinAssetListResponse struct {
	UserId string  `json:"userId"`
	Assets []Asset `json:"assets"`
}

type GetUserCoinAssetRequest struct {
	UserId string `json:"userId"`
	CoinId uint64 `json:"coinId"`
}

type GetUserCoinAssetResponse struct {
	UserId string `json:"userId"`
	Asset  Asset  `json:"asset"`
}

// ==================== User Deposit Models ====================

// Request/Response for User Deposit

type GetOrCreateUserDepositAddressRequest struct {
	UserId    string `json:"userId"`
	Chain     string `json:"chain"`
	NotifyUrl string `json:"notifyUrl,omitempty"`
}

type GetOrCreateUserDepositAddressResponse struct {
	Address string `json:"address"`
	Memo    string `json:"memo,omitempty"`
}

type UserDepositRecord struct {
	UserId           string `json:"userId"`
	RecordId         string `json:"recordId"`
	CoinId           uint64 `json:"coinId"`
	Chain            string `json:"chain"`
	Contract         string `json:"contract"`
	CoinSymbol       string `json:"coinSymbol"`
	TxId             string `json:"txId"`
	CoinUSDPrice     string `json:"coinUSDPrice"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	ToMemo           string `json:"toMemo,omitempty"`
	Amount           string `json:"amount"`
	ServiceFee       string `json:"serviceFee"`
	Status           string `json:"status"`
	ArrivedAt        int64  `json:"arrivedAt"`
	IsFlaggedAsRisky bool   `json:"isFlaggedAsRisky,omitempty"`
}

type GetUserDepositRecordRequest struct {
	RecordId string `json:"recordId"`
}

type GetUserDepositRecordResponse struct {
	Record UserDepositRecord `json:"record"`
}

type GetUserDepositRecordListRequest struct {
	UserId    string `json:"userId"`
	Chain     string `json:"chain,omitempty"`
	CoinId    uint64 `json:"coinId,omitempty"`
	ToAddress string `json:"toAddress,omitempty"`
	StartAt   int64  `json:"startAt,omitempty"`
	EndAt     int64  `json:"endAt,omitempty"`
	NextId    string `json:"nextId,omitempty"`
}

type GetUserDepositRecordListResponse struct {
	Records []UserDepositRecord `json:"records"`
	NextId  string              `json:"nextId,omitempty"`
}

// ==================== User Withdraw Models ====================

// Request/Response for User Withdraw

type ApplyUserWithdrawToNetworkRequest struct {
	UserId              string `json:"userId"`
	OrderId             string `json:"orderId"`
	CoinId              uint64 `json:"coinId"`
	Chain               string `json:"chain"`
	Address             string `json:"address"`
	Amount              string `json:"amount"`
	Memo                string `json:"memo,omitempty"`
	NetworkFeeInquiryID string `json:"networkFeeInquiryID,omitempty"`
	NotifyUrl           string `json:"notifyUrl,omitempty"`
}

type ApplyUserWithdrawToNetworkResponse struct {
	RecordId string `json:"recordId"`
}

type ApplyUserWithdrawToCwalletRequest struct {
	UserId      string `json:"userId"`
	OrderId     string `json:"orderId"`
	CoinId      uint64 `json:"coinId"`
	CwalletUser string `json:"cwalletUser"`
	Amount      string `json:"amount"`
}

type ApplyUserWithdrawToCwalletResponse struct {
	RecordId string `json:"recordId"`
}

type UserWithdrawRecord struct {
	UserId       string      `json:"userId"`
	RecordId     string      `json:"recordId"`
	WithdrawType string      `json:"withdrawType"`
	CoinId       uint64      `json:"coinId"`
	Chain        string      `json:"chain"`
	OrderId      string      `json:"orderId"`
	TxId         string      `json:"txId,omitempty"`
	CoinSymbol   string      `json:"coinSymbol"`
	FromAddress  string      `json:"fromAddress"`
	ToAddress    string      `json:"toAddress"`
	CwalletUser  string      `json:"cwalletUser"`
	ToMemo       string      `json:"toMemo,omitempty"`
	Amount       string      `json:"amount"`
	Status       string      `json:"status"`
	Fee          WithdrawFee `json:"fee"`
	CoinUSDPrice string      `json:"coinUSDPrice"`
}

type GetUserWithdrawRecordRequest struct {
	RecordId string `json:"recordId,omitempty"`
	OrderId  string `json:"orderId,omitempty"`
}

type GetUserWithdrawRecordResponse struct {
	Record UserWithdrawRecord `json:"record"`
}

type GetUserWithdrawRecordListRequest struct {
	UserId   string   `json:"userId"`
	OrderIds []string `json:"orderIds,omitempty"`
	Chain    string   `json:"chain,omitempty"`
	CoinId   uint64   `json:"coinId,omitempty"`
	StartAt  int64    `json:"startAt,omitempty"`
	EndAt    int64    `json:"endAt,omitempty"`
	ToAddress string  `json:"toAddress,omitempty"`
	NextId   string   `json:"nextId,omitempty"`
}

type GetUserWithdrawRecordListResponse struct {
	Records []UserWithdrawRecord `json:"records"`
	NextId  string                 `json:"nextId,omitempty"`
}
