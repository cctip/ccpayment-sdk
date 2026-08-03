package ccpayment

import "context"

type MerchantDepositService struct {
	client *Client
}

// CreateAppOrderDepositAddress creates a deposit address for a specific order
func (s *MerchantDepositService) CreateAppOrderDepositAddress(ctx context.Context, req *CreateAppOrderDepositAddressRequest) (*CreateAppOrderDepositAddressResponse, error) {
	var result CreateAppOrderDepositAddressResponse
	err := s.client.post(ctx, "/createAppOrderDepositAddress", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrCreateAppDepositAddress gets or creates a direct deposit address
func (s *MerchantDepositService) GetOrCreateAppDepositAddress(ctx context.Context, req *GetOrCreateAppDepositAddressRequest) (*GetOrCreateAppDepositAddressResponse, error) {
	var result GetOrCreateAppDepositAddressResponse
	err := s.client.post(ctx, "/getOrCreateAppDepositAddress", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppDepositRecord queries details of a single deposit record
func (s *MerchantDepositService) GetAppDepositRecord(ctx context.Context, recordId string) (*GetAppDepositRecordResponse, error) {
	req := &GetAppDepositRecordRequest{RecordId: recordId}
	var result GetAppDepositRecordResponse
	err := s.client.post(ctx, "/getAppDepositRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppDepositRecordList queries deposit record list
func (s *MerchantDepositService) GetAppDepositRecordList(ctx context.Context, req *GetAppDepositRecordListRequest) (*GetAppDepositRecordListResponse, error) {
	var result GetAppDepositRecordListResponse
	err := s.client.post(ctx, "/getAppDepositRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
