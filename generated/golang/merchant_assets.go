package main

import "context"

// MerchantAssetsService 商家资产服务
type MerchantAssetsService struct {
	client *Client
}

// GetAppCoinAssetList 获取商家的所有代币资产
//
// 接口: POST /getAppCoinAssetList
// 描述: 获取商家的所有代币资产
// 请求参数: 无
func (s *MerchantAssetsService) GetAppCoinAssetList(ctx context.Context) (*GetAppCoinAssetListResponse, error) {
	var result GetAppCoinAssetListResponse
	err := s.client.post(ctx, "/getAppCoinAssetList", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAppCoinAsset 获取商家指定代币的资产
//
// 接口: POST /getAppCoinAsset
// 描述: 获取商家指定代币的资产
// 请求参数:
//   - coinId: 代币ID (必填)
func (s *MerchantAssetsService) GetAppCoinAsset(ctx context.Context, coinID uint64) (*GetAppCoinAssetResponse, error) {
	req := &GetAppCoinAssetRequest{
		CoinID: coinID,
	}
	
	var result GetAppCoinAssetResponse
	err := s.client.post(ctx, "/getAppCoinAsset", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
