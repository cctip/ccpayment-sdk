# CCPayment Go SDK

CCPayment API v2 的 Golang SDK，提供强类型的API调用接口。

## 安装

```bash
go get github.com/yourusername/ccpayment-sdk/generated/golang
```

## 快速开始

### 1. 创建客户端

```go
package main

import (
    "context"
    "log"
    
    ccpayment "path/to/ccpayment-sdk/generated/golang"
)

func main() {
    // 使用您的App ID和App Secret创建客户端
    client := ccpayment.NewClient("your_app_id", "your_app_secret")
    
    // 可选: 设置自定义Base URL
    // client.SetBaseURL("https://custom-api.ccpayment.com/v2")
}
```

### 2. 获取商家全部资产

```go
// 获取全部资产
resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
if err != nil {
    log.Fatal(err)
}

// 遍历资产列表
for _, asset := range resp.Assets {
    log.Printf("代币: %s, 可用余额: %s", asset.CoinSymbol, asset.Available)
}
```

### 3. 获取单个币资产

```go
// 获取USDT资产 (coinID=1280)
resp, err := client.MerchantAssets().GetAppCoinAsset(context.Background(), 1280)
if err != nil {
    log.Fatal(err)
}

log.Printf("USDT可用余额: %s", resp.Asset.Available)
```

## API认证

所有API请求都需要在Header中包含以下字段：
- `Appid`: 应用ID
- `Timestamp`: 当前时间戳（秒）
- `Sign`: HMAC-SHA256签名

SDK会自动处理签名生成，您只需提供App ID和App Secret即可。

### 签名算法

```
signText = appID + timestamp + requestBody
sign = HMAC-SHA256(signText, appSecret)
```

## 错误处理

SDK使用标准的Go error处理方式：

```go
resp, err := client.MerchantAssets().GetAppCoinAssetList(ctx)
if err != nil {
    // 检查是否为API错误
    if apiErr, ok := err.(*ccpayment.APIError); ok {
        log.Printf("API错误: code=%d, message=%s", apiErr.Code, apiErr.Message)
    } else {
        log.Printf("其他错误: %v", err)
    }
    return
}
```

## 数据类型

### Asset (资产信息)

| 字段 | 类型 | 说明 |
|------|------|------|
| CoinID | uint64 | 代币ID |
| CoinSymbol | string | 代币符号 |
| Available | string | 可用余额 |

## 完整示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    ccpayment "path/to/ccpayment-sdk/generated/golang"
)

func main() {
    // 1. 创建客户端
    client := ccpayment.NewClient("your_app_id", "your_app_secret")
    
    // 2. 获取全部资产
    assets, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
    if err != nil {
        log.Fatalf("获取资产失败: %v", err)
    }
    
    fmt.Printf("共有 %d 种代币资产\n", len(assets.Assets))
    
    // 3. 打印每种资产
    for _, asset := range assets.Assets {
        fmt.Printf("- %s: %s\n", asset.CoinSymbol, asset.Available)
    }
    
    // 4. 获取特定代币资产
    usdtAsset, err := client.MerchantAssets().GetAppCoinAsset(context.Background(), 1280)
    if err != nil {
        log.Fatalf("获取USDT资产失败: %v", err)
    }
    
    fmt.Printf("\nUSDT可用余额: %s\n", usdtAsset.Asset.Available)
}
```

## API参考

### MerchantAssetsService

#### GetAppCoinAssetList

获取商家的所有代币资产。

```go
func (s *MerchantAssetsService) GetAppCoinAssetList(ctx context.Context) (*GetAppCoinAssetListResponse, error)
```

**参数:**
- `ctx`: 上下文对象

**返回:**
- `*GetAppCoinAssetListResponse`: 包含资产列表的响应
- `error`: 错误信息

#### GetAppCoinAsset

获取商家指定代币的资产。

```go
func (s *MerchantAssetsService) GetAppCoinAsset(ctx context.Context, coinID uint64) (*GetAppCoinAssetResponse, error)
```

**参数:**
- `ctx`: 上下文对象
- `coinID`: 代币ID

**返回:**
- `*GetAppCoinAssetResponse`: 包含单个资产信息的响应
- `error`: 错误信息

## 相关链接

- [CCPayment官网](https://ccpayment.com)
- [API文档](https://ccpayment.com/api/doc)
- [开发者控制台](https://console.ccpayment.com)

## License

MIT
