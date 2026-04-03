# CCPayment Go SDK - Merchant Assets Module

CCPayment API v2 的 Golang SDK，专注于商家资产模块。

## 快速开始

### 1. 创建客户端

```go
package main

import (
    "context"
    "log"
)

func main() {
    // 使用您的App ID和App Secret创建客户端
    client := NewClient("your_app_id", "your_app_secret")
    
    // 可选: 设置HTTP代理
    // proxyURL, _ := url.Parse("http://127.0.0.1:10808")
    // httpClient := &http.Client{
    //     Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
    // }
    // client.SetHTTPClient(httpClient)
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

## 配置HTTP代理

```go
import (
    "net/http"
    "net/url"
    "time"
)

proxyURL, _ := url.Parse("http://127.0.0.1:10808")
httpClient := &http.Client{
    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
    Timeout: 30 * time.Second,
}
client.SetHTTPClient(httpClient)
```

## API认证

所有API请求需要在Header中包含：
- `Appid`: 应用ID
- `Timestamp`: 当前时间戳（秒）
- `Sign`: HMAC-SHA256签名

SDK会自动处理签名生成。

## 错误处理

```go
resp, err := client.MerchantAssets().GetAppCoinAssetList(ctx)
if err != nil {
    if apiErr, ok := err.(*APIError); ok {
        log.Printf("API错误: code=%d, message=%s", apiErr.Code, apiErr.Message)
    } else {
        log.Printf("其他错误: %v", err)
    }
    return
}
```

## 数据模型

### Asset (资产信息)

| 字段 | 类型 | 说明 |
|------|------|------|
| CoinID | uint64 | 代币ID |
| CoinSymbol | string | 代币符号 |
| Available | string | 可用余额 |

### GetAppCoinAssetListResponse

| 字段 | 类型 | 说明 |
|------|------|------|
| Assets | []Asset | 资产列表 |

### GetAppCoinAssetResponse

| 字段 | 类型 | 说明 |
|------|------|------|
| Asset | Asset | 资产信息 |

## 运行测试

```bash
cd generated/golang
go run .
```

## 注意事项

⚠️ **IP白名单**: 首次测试前必须在CCPayment开发者后台配置IP白名单。

1. 登录 https://console.ccpayment.com/developer/config
2. 找到"IP白名单"设置
3. 添加您的公网IP地址
4. 保存配置

## 相关链接

- [CCPayment官网](https://ccpayment.com)
- [开发者控制台](https://console.ccpayment.com)
