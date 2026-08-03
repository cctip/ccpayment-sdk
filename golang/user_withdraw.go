package ccpayment

import "context"

type UserWithdrawService struct {
	client *Client
}

// ApplyUserWithdrawToNetwork user applies for withdrawal to a blockchain network address
func (s *UserWithdrawService) ApplyUserWithdrawToNetwork(ctx context.Context, req *ApplyUserWithdrawToNetworkRequest) (*ApplyUserWithdrawToNetworkResponse, error) {
	var result ApplyUserWithdrawToNetworkResponse
	err := s.client.post(ctx, "/applyUserWithdrawToNetwork", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ApplyUserWithdrawToCwallet user applies for withdrawal to a CCWallet account
func (s *UserWithdrawService) ApplyUserWithdrawToCwallet(ctx context.Context, req *ApplyUserWithdrawToCwalletRequest) (*ApplyUserWithdrawToCwalletResponse, error) {
	var result ApplyUserWithdrawToCwalletResponse
	err := s.client.post(ctx, "/applyUserWithdrawToCwallet", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserWithdrawRecord queries details of a single user withdrawal record
func (s *UserWithdrawService) GetUserWithdrawRecord(ctx context.Context, req *GetUserWithdrawRecordRequest) (*GetUserWithdrawRecordResponse, error) {
	var result GetUserWithdrawRecordResponse
	err := s.client.post(ctx, "/getUserWithdrawRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserWithdrawRecordList queries user withdrawal record list
func (s *UserWithdrawService) GetUserWithdrawRecordList(ctx context.Context, req *GetUserWithdrawRecordListRequest) (*GetUserWithdrawRecordListResponse, error) {
	var result GetUserWithdrawRecordListResponse
	err := s.client.post(ctx, "/getUserWithdrawRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
