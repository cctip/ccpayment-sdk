package ccpayment

import "context"

// SwapPairV2 is an enabled swap pair returned by GetSwapPairList.
type SwapPairV2 struct {
	PairSymbol  string `json:"pairSymbol"`
	CoinAId     uint64 `json:"coinAId"`
	CoinASymbol string `json:"coinASymbol"`
	CoinBId     uint64 `json:"coinBId"`
	CoinBSymbol string `json:"coinBSymbol"`
	FeeCoinId   uint64 `json:"feeCoinId"`
}

type GetSwapPairListResponse struct {
	Pairs []SwapPairV2 `json:"pairs"`
}

// SwapV2Request follows merchant-common SwapReqV2; amounts and rates are decimal strings.
type SwapV2Request struct {
	OrderId       string `json:"orderId,omitempty"`
	UserId        string `json:"userId,omitempty"`
	FromCoinId    uint64 `json:"fromCoinId"`
	FromAmount    string `json:"fromAmount"`
	ToCoinId      uint64 `json:"toCoinId"`
	MarkupFeeRate string `json:"markupFeeRate,omitempty"`
	AppId         string `json:"appId,omitempty"`
	SlippageRate  string `json:"slippageRate,omitempty"`
}

type SwapEstimateV2Response struct {
	CoinIdIn      uint64 `json:"coinIdIn"`
	CoinIdOut     uint64 `json:"coinIdOut"`
	AmountOut     string `json:"amountOut"`
	AmountIn      string `json:"amountIn"`
	SwapRate      string `json:"swapRate"`
	FeeRate       string `json:"feeRate"`
	Fee           string `json:"fee"`
	ExtraFee      string `json:"extraFee"`
	NetAmountOut  string `json:"netAmountOut"`
	FromCoinId    uint64 `json:"fromCoinId"`
	ToCoinId      uint64 `json:"toCoinId"`
	FromAmount    string `json:"fromAmount"`
	NetFromAmount string `json:"netFromAmount"`
	ToAmount      string `json:"toAmount"`
	NetToAmount   string `json:"netToAmount"`
	FeeCoinId     uint64 `json:"feeCoinId"`
	FeeDeductSide string `json:"feeDeductSide"`
	MarkupFeeRate string `json:"markupFeeRate"`
	MarkupFee     string `json:"markupFee"`
	SwapType      string `json:"swapType"`
}

type SwapV2Response struct {
	RecordId          string `json:"recordId"`
	OrderId           string `json:"orderId"`
	CoinIdIn          uint64 `json:"coinIdIn"`
	CoinIdOut         uint64 `json:"coinIdOut"`
	AmountOut         string `json:"amountOut"`
	AmountIn          string `json:"amountIn"`
	AmountOutMinimum  string `json:"amountOutMinimum"`
	SwapRate          string `json:"swapRate"`
	Fee               string `json:"fee"`
	FeeRate           string `json:"feeRate"`
	ExtraFee          string `json:"extraFee"`
	NetAmountOut      string `json:"netAmountOut"`
	ToCoinPrice       string `json:"to_coin_price"`
	FromCoinPrice     string `json:"from_coin_price"`
	FromCoinId        uint64 `json:"fromCoinId"`
	ToCoinId          uint64 `json:"toCoinId"`
	FromAmount        string `json:"fromAmount"`
	NetFromAmount     string `json:"netFromAmount"`
	ToAmount          string `json:"toAmount"`
	NetToAmount       string `json:"netToAmount"`
	FeeCoinId         uint64 `json:"feeCoinId"`
	FeeDeductSide     string `json:"feeDeductSide"`
	SwapType          string `json:"swapType"`
	MarkupFeeRate     string `json:"markupFeeRate"`
	MarkupFee         string `json:"markupFee"`
	SuccessFromAmount string `json:"successFromAmount"`
	RefundAmount      string `json:"refundAmount"`
	UserId            string `json:"userId"`
}

type GetSwapRecordV2Request struct {
	RecordId string `json:"recordId,omitempty"`
	OrderId  string `json:"orderId,omitempty"`
	Ty       int32  `json:"ty,omitempty"`
}

type GetSwapRecordListV2Request struct {
	RecordIds []string `json:"recordIds,omitempty"`
	OrderIds  []string `json:"orderIds,omitempty"`
	UserId    string   `json:"userId,omitempty"`
	CoinIdIn  uint64   `json:"coinIdIn,omitempty"`
	CoinIdOut uint64   `json:"coinIdOut,omitempty"`
	StartAt   int64    `json:"startAt,omitempty"`
	EndAt     int64    `json:"endAt,omitempty"`
	NextId    string   `json:"nextId,omitempty"`
}

type SwapRecordV2 struct {
	RecordId          string `json:"recordId"`
	OrderId           string `json:"orderId"`
	CoinInSymbol      string `json:"coinInSymbol"`
	CoinIdIn          uint64 `json:"coinIdIn"`
	CoinOutSymbol     string `json:"coinOutSymbol"`
	CoinIdOut         uint64 `json:"coinIdOut"`
	AmountOut         string `json:"amountOut"`
	AmountIn          string `json:"amountIn"`
	AmountOutMinimum  string `json:"amountOutMinimum"`
	NetAmountOut      string `json:"netAmountOut"`
	UserId            string `json:"userId"`
	ExtraFee          string `json:"extraFee"`
	ExtraFeeRate      string `json:"extraFeeRate"`
	SwapRate          string `json:"swapRate"`
	FeeRate           string `json:"feeRate"`
	Fee               string `json:"fee"`
	Status            string `json:"status"`
	CreatedAt         int64  `json:"createdAt"`
	ArrivedAt         int64  `json:"arrivedAt"`
	FromCoinSymbol    string `json:"fromCoinSymbol"`
	FromCoinId        uint64 `json:"fromCoinId"`
	FromAmount        string `json:"fromAmount"`
	NetFromAmount     string `json:"netFromAmount"`
	ToCoinSymbol      string `json:"toCoinSymbol"`
	ToCoinId          uint64 `json:"toCoinId"`
	ToAmount          string `json:"toAmount"`
	NetToAmount       string `json:"netToAmount"`
	SwapType          string `json:"swapType"`
	MarkupFeeRate     string `json:"markupFeeRate"`
	MarkupFee         string `json:"markupFee"`
	FeeCoinId         uint64 `json:"feeCoinId"`
	FeeDeductSide     string `json:"feeDeductSide"`
	SuccessFromAmount string `json:"successFromAmount"`
	RefundAmount      string `json:"refundAmount"`
	SlippageRate      string `json:"slippageRate"`
}

type GetSwapRecordV2Response struct {
	Record SwapRecordV2 `json:"record"`
}

type GetSwapRecordListV2Response struct {
	Records []SwapRecordV2 `json:"records"`
	NextId  string         `json:"nextId,omitempty"`
}

func (s *SwapService) GetSwapPairList(ctx context.Context) (*GetSwapPairListResponse, error) {
	var result GetSwapPairListResponse
	if err := s.client.post(ctx, "/swapv2/getSwapPairList", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SwapService) SwapEstimateV2(ctx context.Context, req *SwapV2Request) (*SwapEstimateV2Response, error) {
	var result SwapEstimateV2Response
	if err := s.client.post(ctx, "/swapv2/estimate", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SwapService) SwapV2(ctx context.Context, req *SwapV2Request) (*SwapV2Response, error) {
	var result SwapV2Response
	if err := s.client.post(ctx, "/swapv2/swap", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SwapService) GetSwapRecordV2(ctx context.Context, req *GetSwapRecordV2Request) (*GetSwapRecordV2Response, error) {
	var result GetSwapRecordV2Response
	if err := s.client.post(ctx, "/swapv2/getSwapRecord", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SwapService) GetSwapRecordListV2(ctx context.Context, req *GetSwapRecordListV2Request) (*GetSwapRecordListV2Response, error) {
	var result GetSwapRecordListV2Response
	if err := s.client.post(ctx, "/swapv2/getSwapRecordList", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserSwapService) UserSwapV2(ctx context.Context, req *SwapV2Request) (*SwapV2Response, error) {
	var result SwapV2Response
	if err := s.client.post(ctx, "/swapv2/userSwap", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserSwapService) GetUserSwapRecordV2(ctx context.Context, req *GetSwapRecordV2Request) (*GetSwapRecordV2Response, error) {
	var result GetSwapRecordV2Response
	if err := s.client.post(ctx, "/swapv2/getUserSwapRecord", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserSwapService) GetUserSwapRecordListV2(ctx context.Context, req *GetSwapRecordListV2Request) (*GetSwapRecordListV2Response, error) {
	var result GetSwapRecordListV2Response
	if err := s.client.post(ctx, "/swapv2/getUserSwapRecordList", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
