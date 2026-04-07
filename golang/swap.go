package ccpayment

import "context"

type SwapService struct {
	client *Client
}

// GetSwapCoinList gets the list of tokens supported for swapping
func (s *SwapService) GetSwapCoinList(ctx context.Context) (*GetSwapCoinListResponse, error) {
	var result GetSwapCoinListResponse
	err := s.client.post(ctx, "/getSwapCoinList", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SwapEstimate estimates the swap amount
func (s *SwapService) SwapEstimate(ctx context.Context, req *SwapEstimateRequest) (*SwapEstimateResponse, error) {
	var result SwapEstimateResponse
	err := s.client.post(ctx, "/estimate", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Swap executes a swap operation
func (s *SwapService) Swap(ctx context.Context, req *SwapRequest) (*SwapResponse, error) {
	var result SwapResponse
	err := s.client.post(ctx, "/swap", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSwapRecord queries details of a single swap record
func (s *SwapService) GetSwapRecord(ctx context.Context, req *GetSwapRecordRequest) (*GetSwapRecordResponse, error) {
	var result GetSwapRecordResponse
	err := s.client.post(ctx, "/getSwapRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSwapRecordList queries swap record list
func (s *SwapService) GetSwapRecordList(ctx context.Context, req *GetSwapRecordListRequest) (*GetSwapRecordListResponse, error) {
	var result GetSwapRecordListResponse
	err := s.client.post(ctx, "/getSwapRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
