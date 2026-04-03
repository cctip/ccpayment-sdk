package ccpayment

import "context"

type MerchantAssetsService struct {
	client *Client
}

// GetAppCoinAssetList gets all token assets of the merchant
func (s *MerchantAssetsService) GetAppCoinAssetList(ctx context.Context) (*GetAppCoinAssetListResponse, error) {
	var result GetAppCoinAssetListResponse
	err := s.client.post(ctx, "/getAppCoinAssetList", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppCoinAsset gets the asset of a specified token for the merchant
func (s *MerchantAssetsService) GetAppCoinAsset(ctx context.Context, coinId uint64) (*GetAppCoinAssetResponse, error) {
	req := &GetAppCoinAssetRequest{CoinId: coinId}
	var result GetAppCoinAssetResponse
	err := s.client.post(ctx, "/getAppCoinAsset", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
