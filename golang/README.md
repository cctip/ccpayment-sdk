# CCPayment Go SDK

Official Go SDK for CCPayment API v2.

## Installation

```bash
go get github.com/cctip/ccpayment-sdk/golang
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    
    ccpayment "github.com/cctip/ccpayment-sdk/golang"
)

func main() {
    // Get credentials from environment variables
    appID := os.Getenv("CCPAYMENT_APP_ID")
    appSecret := os.Getenv("CCPAYMENT_APP_SECRET")
    proxyURL := os.Getenv("CCPAYMENT_PROXY_URL")
    
    if appID == "" || appSecret == "" {
        log.Fatal("Please set CCPAYMENT_APP_ID and CCPAYMENT_APP_SECRET environment variables")
    }
    
    // Create client
    client := ccpayment.NewClient(appID, appSecret)
    
    // Optional: Configure HTTP proxy
    if proxyURL != "" {
        proxyURLParsed, _ := url.Parse(proxyURL)
        httpClient := &http.Client{
            Transport: &http.Transport{Proxy: http.ProxyURL(proxyURLParsed)},
        }
        client.SetHTTPClient(httpClient)
    }
    
    // Get all assets
    resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    
    // Print assets
    for _, asset := range resp.Assets {
        fmt.Printf("%s: %s\n", asset.CoinSymbol, asset.Available)
    }
}
```

## Configuration

### HTTP Proxy

```go
proxyURL, _ := url.Parse("http://127.0.0.1:10808")
httpClient := &http.Client{
    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
}
client.SetHTTPClient(httpClient)
```

### Custom Base URL

```go
client.SetBaseURL("https://custom.ccpayment.com/ccpayment/v2")
```

## API Modules

- **BasicInfo** - Token, fiat currency, and chain information queries
- **MerchantAssets** - Merchant asset queries
- **MerchantDeposit** - Deposit address generation and deposit record queries
- **MerchantWithdraw** - Withdrawal requests and withdrawal record queries
- **MerchantBatchWithdraw** - Batch withdrawal management
- **UserAssets** - User asset queries
- **UserDeposit** - User deposit addresses and deposit records
- **UserWithdraw** - User withdrawal requests and withdrawal records
- **UserTransfer** - User transfers and batch transfers
- **Orders** - Order and Invoice orders
- **Swap** - Swap related interfaces
- **User Swap** - User swap management
- **Utilities** - Webhook, address validation, etc.

## Error Handling

```go
resp, err := client.MerchantAssets().GetAppCoinAssetList(ctx)
if err != nil {
    if apiErr, ok := err.(*ccpayment.APIError); ok {
        log.Printf("API Error: code=%d, message=%s", apiErr.Code, apiErr.Message)
    } else {
        log.Printf("Other Error: %v", err)
    }
    return
}
```

## API Documentation

For complete API documentation, visit: https://ccpayment.com/api/doc

## Disclaimer

This SDK is provided by CCPayment on an "as is" and "as available" basis for use at your own risk. CCPayment makes no representations or warranties of any kind, whether express or implied, as to the operation of the SDK, its accuracy, reliability, completeness, or fitness for any particular purpose.

To the fullest extent permitted by law, CCPayment shall not be liable for any losses, damages, or expenses of any kind arising from or in connection with your use of, or inability to use, this SDK, including but not limited to any financial losses resulting from errors, bugs, interruptions, or inaccuracies in the SDK.

Your use of this SDK to access the CCPayment platform is subject to the [CCPayment Terms of Use](https://ccpayment.com/en/privacy-policy#ccpayment-service-agreement), which shall prevail in the event of any conflict with this disclaimer. You are solely responsible for the transactions and API operations executed through your merchant account using this SDK. For a secure integration, please follow the verification practices described in the [API documentation](https://ccpayment.com/api/doc?en).

Plugins or extensions developed by third parties based on this SDK or the CCPayment API are independent works of their authors. Unless expressly designated as official in CCPayment's documentation, they are not maintained by CCPayment, and CCPayment is not responsible for their operation. We welcome community contributions — developers building on CCPayment are invited to reach out at partner@ccpayment.com.

This SDK is not intended to constitute investment advice or a recommendation to buy, sell, or hold any digital asset. You should independently evaluate and verify all information before acting.
