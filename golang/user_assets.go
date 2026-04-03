package ccpayment

import "context"

type UserAssetsService struct {
	client *Client
}

// GetUserCoinAssetList gets all token assets of a specified user
func (s *UserAssetsService) GetUserCoinAssetList(ctx context.Context, userId string) (*GetUserCoinAssetListResponse, error) {
	req := &GetUserCoinAssetListRequest{UserId: userId}
	var result GetUserCoinAssetListResponse
	err := s.client.post(ctx, "/getUserCoinAssetList", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserCoinAsset gets a specific token asset of a specified user
func (s *UserAssetsService) GetUserCoinAsset(ctx context.Context, userId string, coinId uint64) (*GetUserCoinAssetResponse, error) {
	req := &GetUserCoinAssetRequest{UserId: userId, CoinId: coinId}
	var result GetUserCoinAssetResponse
	err := s.client.post(ctx, "/getUserCoinAsset", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
