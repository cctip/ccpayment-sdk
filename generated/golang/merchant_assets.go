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
//
// 示例:
//
//	resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, asset := range resp.Assets {
//	    fmt.Printf("%s: %s\n", asset.CoinSymbol, asset.Available)
//	}
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
//
// 参数:
//   - coinID: 代币ID (必填)
//
// 示例:
//
//	resp, err := client.MerchantAssets().GetAppCoinAsset(context.Background(), 1280)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("USDT可用余额: %s\n", resp.Asset.Available)
func (s *MerchantAssetsService) GetAppCoinAsset(ctx context.Context, coinID uint64) (*GetAppCoinAssetResponse, error) {
	req := &GetAppCoinAssetRequest{CoinID: coinID}
	var result GetAppCoinAssetResponse
	err := s.client.post(ctx, "/getAppCoinAsset", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
