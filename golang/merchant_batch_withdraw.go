package ccpayment

import "context"

type MerchantBatchWithdrawService struct {
	client *Client
}

// CheckWithdrawAddress checks the validity of withdrawal addresses in batch
func (s *MerchantBatchWithdrawService) CheckWithdrawAddress(ctx context.Context, req *CheckWithdrawAddressRequest) (*CheckWithdrawAddressResponse, error) {
	var result CheckWithdrawAddressResponse
	err := s.client.post(ctx, "/checkWithdrawAddress", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ApplyBatchWithdraw creates a batch withdrawal order
func (s *MerchantBatchWithdrawService) ApplyBatchWithdraw(ctx context.Context, req *ApplyBatchWithdrawRequest) (*ApplyBatchWithdrawResponse, error) {
	var result ApplyBatchWithdrawResponse
	err := s.client.post(ctx, "/applyBatchWithdraw", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AppendBatchWithdraw appends tasks to an existing batch withdrawal order
func (s *MerchantBatchWithdrawService) AppendBatchWithdraw(ctx context.Context, req *AppendBatchWithdrawRequest) error {
	return s.client.post(ctx, "/appendBatchWithdraw", req, nil)
}

// ConfirmBatchWithdraw confirms and executes the batch withdrawal order
func (s *MerchantBatchWithdrawService) ConfirmBatchWithdraw(ctx context.Context, req *ConfirmBatchWithdrawRequest) (*ConfirmBatchWithdrawResponse, error) {
	var result ConfirmBatchWithdrawResponse
	err := s.client.post(ctx, "/confirmBatchWithdraw", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AbortBatchWithdraw cancels the batch withdrawal order or part of the tasks
func (s *MerchantBatchWithdrawService) AbortBatchWithdraw(ctx context.Context, req *AbortBatchWithdrawRequest) (*AbortBatchWithdrawResponse, error) {
	var result AbortBatchWithdrawResponse
	err := s.client.post(ctx, "/abortBatchWithdraw", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBatchWithdrawOrder queries batch withdrawal order details
func (s *MerchantBatchWithdrawService) GetBatchWithdrawOrder(ctx context.Context, req *GetBatchWithdrawOrderRequest) (*ConfirmBatchWithdrawResponse, error) {
	var result ConfirmBatchWithdrawResponse
	err := s.client.post(ctx, "/getBatchWithdrawOrder", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBatchWithdrawOrderRecordList queries the task record list of a batch withdrawal order
func (s *MerchantBatchWithdrawService) GetBatchWithdrawOrderRecordList(ctx context.Context, req *GetBatchWithdrawOrderRecordListRequest) (*GetBatchWithdrawOrderRecordListResponse, error) {
	var result GetBatchWithdrawOrderRecordListResponse
	err := s.client.post(ctx, "/getBatchWithdrawOrderRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
