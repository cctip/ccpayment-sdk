package ccpayment

import "context"

type UserDepositService struct {
	client *Client
}

// GetOrCreateUserDepositAddress gets or creates a user's deposit address
func (s *UserDepositService) GetOrCreateUserDepositAddress(ctx context.Context, req *GetOrCreateUserDepositAddressRequest) (*GetOrCreateUserDepositAddressResponse, error) {
	var result GetOrCreateUserDepositAddressResponse
	err := s.client.post(ctx, "/getOrCreateUserDepositAddress", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserDepositRecord queries details of a single user deposit record
func (s *UserDepositService) GetUserDepositRecord(ctx context.Context, recordId string) (*GetUserDepositRecordResponse, error) {
	req := &GetUserDepositRecordRequest{RecordId: recordId}
	var result GetUserDepositRecordResponse
	err := s.client.post(ctx, "/getUserDepositRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserDepositRecordList queries user deposit record list
func (s *UserDepositService) GetUserDepositRecordList(ctx context.Context, req *GetUserDepositRecordListRequest) (*GetUserDepositRecordListResponse, error) {
	var result GetUserDepositRecordListResponse
	err := s.client.post(ctx, "/getUserDepositRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
