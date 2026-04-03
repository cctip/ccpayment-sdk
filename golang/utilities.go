package ccpayment

import "context"

type UtilitiesService struct {
	client *Client
}

// VerifyAddress verifies the validity of a blockchain address
func (s *UtilitiesService) VerifyAddress(ctx context.Context, chain, address string) (*VerifyAddressResponse, error) {
	req := &VerifyAddressRequest{Chain: chain, Address: address}
	var result VerifyAddressResponse
	err := s.client.post(ctx, "/verifyAddress", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AbandonAddress abandons a deposit address that is no longer in use
func (s *UtilitiesService) AbandonAddress(ctx context.Context, chain, address string) error {
	req := &AbandonAddressRequest{Chain: chain, Address: address}
	return s.client.post(ctx, "/abandonAddress", req, nil)
}

// HostedInvoiceOrderInfo gets Hosted Invoice order details
func (s *UtilitiesService) HostedInvoiceOrderInfo(ctx context.Context, orderId string) (*HostedInvoiceOrderInfoResponse, error) {
	req := &HostedInvoiceOrderInfoRequest{OrderId: orderId}
	var result HostedInvoiceOrderInfoResponse
	err := s.client.post(ctx, "/hostedInvoiceOrderInfo", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPayInfo gets the payment information of an order
func (s *UtilitiesService) GetPayInfo(ctx context.Context, orderId string) (*GetPayInfoResponse, error) {
	req := &GetPayInfoRequest{OrderId: orderId}
	var result GetPayInfoResponse
	err := s.client.post(ctx, "/getPayInfo", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Health API health check interface
func (s *UtilitiesService) Health(ctx context.Context) (*HealthResponse, error) {
	// Health endpoint may return plain text, use custom handling
	var result HealthResponse
	err := s.client.postWithStringResponse(ctx, "/health", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
