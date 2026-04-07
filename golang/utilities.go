package ccpayment

import "context"

type UtilitiesService struct {
	client *Client
}

// WebhookResend resends webhook notifications
func (s *UtilitiesService) WebhookResend(ctx context.Context, req *WebhookResendRequest) (*WebhookResendResponse, error) {
	var result WebhookResendResponse
	err := s.client.post(ctx, "/webhook/resend", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTradeBlockHeight gets trade confirmation height information
func (s *UtilitiesService) GetTradeBlockHeight(ctx context.Context, req *GetTradeBlockHeightRequest) (*GetTradeBlockHeightResponse, error) {
	var result GetTradeBlockHeightResponse
	err := s.client.post(ctx, "/getTradeBlockHeight", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckWithdrawalAddressValidity checks whether a withdrawal address is valid
func (s *UtilitiesService) CheckWithdrawalAddressValidity(ctx context.Context, req *CheckWithdrawalAddressValidityRequest) (*CheckWithdrawalAddressValidityResponse, error) {
	var result CheckWithdrawalAddressValidityResponse
	err := s.client.post(ctx, "/checkWithdrawalAddressValidity", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeprecatedAddress deprecates an unused address
func (s *UtilitiesService) DeprecatedAddress(ctx context.Context, req *DeprecatedAddressRequest) (*DeprecatedAddressResponse, error) {
	var result DeprecatedAddressResponse
	err := s.client.post(ctx, "/addressUnbinding", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RescanLostTransaction triggers a rescan for a lost transaction
func (s *UtilitiesService) RescanLostTransaction(ctx context.Context, req *RescanLostTransactionRequest) (*RescanLostTransactionResponse, error) {
	var result RescanLostTransactionResponse
	err := s.client.post(ctx, "/rescanLostTransaction", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
