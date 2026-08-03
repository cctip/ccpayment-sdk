package ccpayment

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func ExampleClient() {
	// Create client
	client := NewClient("your_app_id", "your_app_secret")

	// Optional: Configure HTTP proxy
	proxyURL, _ := url.Parse("http://127.0.0.1:10808")
	httpClient := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		Timeout:   30 * time.Second,
	}
	client.SetHTTPClient(httpClient)

	// Get all assets
	resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, asset := range resp.Assets {
		fmt.Printf("%s: %s\n", asset.CoinSymbol, asset.Available)
	}
}
