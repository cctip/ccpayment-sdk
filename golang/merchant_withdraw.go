package ccpayment

import "context"

type MerchantWithdrawService struct {
	client *Client
}

// ApplyAppWithdrawToNetwork applies for withdrawal to a blockchain network address
func (s *MerchantWithdrawService) ApplyAppWithdrawToNetwork(ctx context.Context, req *ApplyAppWithdrawToNetworkRequest) (*ApplyAppWithdrawToNetworkResponse, error) {
	var result ApplyAppWithdrawToNetworkResponse
	err := s.client.post(ctx, "/applyAppWithdrawToNetwork", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ApplyAppWithdrawToCwallet applies for withdrawal to a CCWallet account
func (s *MerchantWithdrawService) ApplyAppWithdrawToCwallet(ctx context.Context, req *ApplyAppWithdrawToCwalletRequest) (*ApplyAppWithdrawToCwalletResponse, error) {
	var result ApplyAppWithdrawToCwalletResponse
	err := s.client.post(ctx, "/applyAppWithdrawToCwallet", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppWithdrawRecord queries details of a single withdrawal record
func (s *MerchantWithdrawService) GetAppWithdrawRecord(ctx context.Context, req *GetAppWithdrawRecordRequest) (*GetAppWithdrawRecordResponse, error) {
	var result GetAppWithdrawRecordResponse
	err := s.client.post(ctx, "/getAppWithdrawRecord", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppWithdrawRecordList queries withdrawal record list
func (s *MerchantWithdrawService) GetAppWithdrawRecordList(ctx context.Context, req *GetAppWithdrawRecordListRequest) (*GetAppWithdrawRecordListResponse, error) {
	var result GetAppWithdrawRecordListResponse
	err := s.client.post(ctx, "/getAppWithdrawRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAutoWithdrawRecordList queries automatic withdrawal record list
func (s *MerchantWithdrawService) GetAutoWithdrawRecordList(ctx context.Context, req *GetAutoWithdrawRecordListRequest) (*GetAutoWithdrawRecordListResponse, error) {
	var result GetAutoWithdrawRecordListResponse
	err := s.client.post(ctx, "/getAutoWithdrawRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRiskyRefundRecordList queries risky refund record list
func (s *MerchantWithdrawService) GetRiskyRefundRecordList(ctx context.Context, req *GetAutoWithdrawRecordListRequest) (*GetRiskyRefundRecordListResponse, error) {
	var result GetRiskyRefundRecordListResponse
	err := s.client.post(ctx, "/getRiskyRefundRecordList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
