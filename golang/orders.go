package ccpayment

import "context"

type OrdersService struct {
	client *Client
}

// GetAppOrderInfo gets Order order details
func (s *OrdersService) GetAppOrderInfo(ctx context.Context, orderId string) (*GetAppOrderInfoResponse, error) {
	req := &GetAppOrderInfoRequest{OrderId: orderId}
	var result GetAppOrderInfoResponse
	err := s.client.post(ctx, "/getAppOrderInfo", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateInvoiceUrl creates an Invoice order and generates a payment URL
func (s *OrdersService) CreateInvoiceUrl(ctx context.Context, req *CreateInvoiceUrlRequest) (*CreateInvoiceUrlResponse, error) {
	var result CreateInvoiceUrlResponse
	err := s.client.post(ctx, "/createInvoiceUrl", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInvoiceOrderInfo gets Invoice order details
func (s *OrdersService) GetInvoiceOrderInfo(ctx context.Context, orderId string) (*GetInvoiceOrderInfoResponse, error) {
	req := &GetInvoiceOrderInfoRequest{OrderId: orderId}
	var result GetInvoiceOrderInfoResponse
	err := s.client.post(ctx, "/getInvoiceOrderInfo", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWebhookInfo gets merchant's Webhook configuration information
func (s *OrdersService) GetWebhookInfo(ctx context.Context) (*GetWebhookInfoResponse, error) {
	var result GetWebhookInfoResponse
	err := s.client.post(ctx, "/getWebhookInfo", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
