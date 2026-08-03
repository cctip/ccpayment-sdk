package ccpayment

import "context"

type UserSwapService struct {
	client *Client
}

// GetUserSwapRecord queries details of a single user swap record
func (s *UserSwapService) GetUserSwapRecord(ctx context.Context, req *GetSwapRecordRequest) (*GetSwapRecordResponse, error) {
	var result GetSwapRecordResponse
	err := s.client.post(ctx, "/getUserSwapRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserSwapRecordList queries user swap record list
func (s *UserSwapService) GetUserSwapRecordList(ctx context.Context, req *GetSwapRecordListRequest) (*GetSwapRecordListResponse, error) {
	var result GetSwapRecordListResponse
	err := s.client.post(ctx, "/getUserSwapRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UserSwap executes a user swap operation
func (s *UserSwapService) UserSwap(ctx context.Context, req *SwapRequest) (*SwapResponse, error) {
	var result SwapResponse
	err := s.client.post(ctx, "/userSwap", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
