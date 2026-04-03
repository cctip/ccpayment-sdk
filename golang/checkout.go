package ccpayment

import "context"

type CheckoutService struct {
	client *Client
}

// CreateCheckoutUrl creates a Hosted Checkout payment URL
func (s *CheckoutService) CreateCheckoutUrl(ctx context.Context, req *CreateCheckoutUrlRequest) (*CreateCheckoutUrlResponse, error) {
	var result CreateCheckoutUrlResponse
	err := s.client.post(ctx, "/createCheckoutUrl", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// HostedOrderInfo gets Hosted order details
func (s *CheckoutService) HostedOrderInfo(ctx context.Context, orderId string) (*HostedOrderInfoResponse, error) {
	req := &HostedOrderInfoRequest{OrderId: orderId}
	var result HostedOrderInfoResponse
	err := s.client.post(ctx, "/hostedOrderInfo", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// HostedOrderInfoFirst gets the first visit information of a Hosted order
func (s *CheckoutService) HostedOrderInfoFirst(ctx context.Context, orderId string) (*HostedOrderInfoResponse, error) {
	req := &HostedOrderInfoFirstRequest{OrderId: orderId}
	var result HostedOrderInfoResponse
	err := s.client.post(ctx, "/hostedOrderInfoFirst", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateHostedOrderDeposit creates a deposit address after selecting token and chain for a Hosted order
func (s *CheckoutService) CreateHostedOrderDeposit(ctx context.Context, req *CreateHostedOrderDepositRequest) (*CreateHostedOrderDepositResponse, error) {
	var result CreateHostedOrderDepositResponse
	err := s.client.post(ctx, "/createHostedOrderDeposit", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHostedCoinUSDTPrice gets the USDT price of a token in Hosted mode
func (s *CheckoutService) GetHostedCoinUSDTPrice(ctx context.Context, coinIds []uint64) (*GetHostedCoinUSDTPriceResponse, error) {
	req := &GetHostedCoinUSDTPriceRequest{CoinIds: coinIds}
	var result GetHostedCoinUSDTPriceResponse
	err := s.client.post(ctx, "/getHostedCoinUSDTPrice", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMainCoinList gets the list of main supported tokens in Hosted mode
func (s *CheckoutService) GetMainCoinList(ctx context.Context) (*GetMainCoinListResponse, error) {
	var result GetMainCoinListResponse
	err := s.client.post(ctx, "/getMainCoinList", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateAppDemoOrderDeposit creates a Demo order deposit address
func (s *CheckoutService) CreateAppDemoOrderDeposit(ctx context.Context, req *CreateAppDemoOrderDepositRequest) (*CreateAppDemoOrderDepositResponse, error) {
	var result CreateAppDemoOrderDepositResponse
	err := s.client.post(ctx, "/createAppDemoOrderDeposit", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ConfirmTrade confirms the trade of a Hosted order
func (s *CheckoutService) ConfirmTrade(ctx context.Context, req *ConfirmTradeRequest) error {
	return s.client.post(ctx, "/confirmTrade", req, nil)
}
