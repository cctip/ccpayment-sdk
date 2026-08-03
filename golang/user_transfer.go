package ccpayment

import "context"

type UserTransferService struct {
	client *Client
}

// UserTransfer initiates a transfer between users
func (s *UserTransferService) UserTransfer(ctx context.Context, req *UserTransferRequest) (*UserTransferResponse, error) {
	var result UserTransferResponse
	err := s.client.post(ctx, "/userTransfer", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserTransferRecord queries details of a single user transfer record
func (s *UserTransferService) GetUserTransferRecord(ctx context.Context, req *GetUserTransferRecordRequest) (*GetUserTransferRecordResponse, error) {
	var result GetUserTransferRecordResponse
	err := s.client.post(ctx, "/getUserTransferRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserTransferRecordList queries user transfer record list
func (s *UserTransferService) GetUserTransferRecordList(ctx context.Context, req *GetUserTransferRecordListRequest) (*GetUserTransferRecordListResponse, error) {
	var result GetUserTransferRecordListResponse
	err := s.client.post(ctx, "/getUserTransferRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UserBatchTransfer initiates a batch transfer for users
func (s *UserTransferService) UserBatchTransfer(ctx context.Context, req *UserBatchTransferRequest) (*UserBatchTransferResponse, error) {
	var result UserBatchTransferResponse
	err := s.client.post(ctx, "/userBatchTransfer", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserBatchTransferRecord queries details of a single user batch transfer record
func (s *UserTransferService) GetUserBatchTransferRecord(ctx context.Context, req *GetUserBatchTransferRecordRequest) (*GetUserBatchTransferRecordResponse, error) {
	var result GetUserBatchTransferRecordResponse
	err := s.client.post(ctx, "/getUserBatchTransferRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserBatchTransferRecordList queries user batch transfer record list
func (s *UserTransferService) GetUserBatchTransferRecordList(ctx context.Context, req *GetUserBatchTransferRecordListRequest) (*GetUserBatchTransferRecordListResponse, error) {
	var result GetUserBatchTransferRecordListResponse
	err := s.client.post(ctx, "/getUserBatchTransferRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserBatchTransferRecordDetail queries the detail list of a user batch transfer record
func (s *UserTransferService) GetUserBatchTransferRecordDetail(ctx context.Context, req *GetUserBatchTransferRecordDetailRequest) (*GetUserBatchTransferRecordDetailResponse, error) {
	var result GetUserBatchTransferRecordDetailResponse
	err := s.client.post(ctx, "/getUserBatchTransferRecordDetail", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
