package ccpayment

import "context"

type BasicInfoService struct {
	client *Client
}

// GetCoinList gets all supported token information
func (s *BasicInfoService) GetCoinList(ctx context.Context) (*GetCoinListResponse, error) {
	var result GetCoinListResponse
	err := s.client.post(ctx, "/getCoinList", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCoin gets detailed information for a specified token
func (s *BasicInfoService) GetCoin(ctx context.Context, coinId uint64) (*Coin, error) {
	req := &GetCoinRequest{CoinId: coinId}
	var result Coin
	err := s.client.post(ctx, "/getCoin", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCoinUSDTPrice gets USD prices for tokens in batch
func (s *BasicInfoService) GetCoinUSDTPrice(ctx context.Context, coinIds []uint64) (*GetCoinUSDTPriceResponse, error) {
	req := &GetCoinUSDTPriceRequest{CoinIds: coinIds}
	var result GetCoinUSDTPriceResponse
	err := s.client.post(ctx, "/getCoinUSDTPrice", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFiatList gets all supported fiat currency information
func (s *BasicInfoService) GetFiatList(ctx context.Context) (*GetFiatListResponse, error) {
	var result GetFiatListResponse
	err := s.client.post(ctx, "/getFiatList", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetChainList gets chain information
func (s *BasicInfoService) GetChainList(ctx context.Context, chains []string) (*GetChainListResponse, error) {
	req := &GetChainListRequest{Chains: chains}
	var result GetChainListResponse
	err := s.client.post(ctx, "/getChainList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AllChain gets information for all chains (simplified version)
func (s *BasicInfoService) AllChain(ctx context.Context, chains []string) (*AllChainResponse, error) {
	req := &AllChainRequest{Chains: chains}
	var result AllChainResponse
	err := s.client.post(ctx, "/all/chain", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCwalletUserId verifies and gets CCWallet user information
func (s *BasicInfoService) GetCwalletUserId(ctx context.Context, cwalletUserId string) (*GetCwalletUserIdResponse, error) {
	req := &GetCwalletUserIdRequest{CwalletUserId: cwalletUserId}
	var result GetCwalletUserIdResponse
	err := s.client.post(ctx, "/getCwalletUserId", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWithdrawFee gets the withdrawal fee for a specified token and chain
func (s *BasicInfoService) GetWithdrawFee(ctx context.Context, coinId uint64, chain string) (*GetWithdrawFeeResponse, error) {
	req := &GetWithdrawFeeRequest{CoinId: coinId, Chain: chain}
	var result GetWithdrawFeeResponse
	err := s.client.post(ctx, "/getWithdrawFee", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
