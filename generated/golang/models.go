package main

// Response API通用响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// APIError API错误
type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}

// Asset 资产信息
type Asset struct {
	CoinID      uint64 `json:"coinId"`
	CoinSymbol  string `json:"coinSymbol"`
	Available   string `json:"available"`
}

// GetAppCoinAssetListResponse 获取全部资产响应
type GetAppCoinAssetListResponse struct {
	Assets []Asset `json:"assets"`
}

// GetAppCoinAssetRequest 获取单个币资产请求
type GetAppCoinAssetRequest struct {
	CoinID uint64 `json:"coinId"`
}

// GetAppCoinAssetResponse 获取单个币资产响应
type GetAppCoinAssetResponse struct {
	Asset Asset `json:"asset"`
}
