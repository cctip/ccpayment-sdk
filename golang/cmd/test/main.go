package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

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

	fmt.Printf("API Key: %s\n", appID)
	fmt.Printf("Using Proxy: %v\n", proxyURL != "")
	fmt.Println("========================================")

	// Create client
	client := ccpayment.NewClient(appID, appSecret)

	// Optional: Configure HTTP proxy
	if proxyURL != "" {
		proxyURLParsed, err := url.Parse("http://" + proxyURL)
		if err != nil {
			log.Printf("Warning: Failed to parse proxy URL: %v", err)
		} else {
			httpClient := &http.Client{
				Transport: &http.Transport{Proxy: http.ProxyURL(proxyURLParsed)},
				Timeout:   30 * time.Second,
			}
			client.SetHTTPClient(httpClient)
			fmt.Printf("Proxy configured: %s\n", proxyURL)
		}
	}

	// Test 1: Get coin list
	fmt.Println("\n[Test 1] Get Coin List...")
	ctx := context.Background()
	coinListResp, err := client.BasicInfo().GetCoinList(ctx)
	if err != nil {
		log.Printf("Error getting coin list: %v", err)
		if apiErr, ok := err.(*ccpayment.APIError); ok {
			log.Printf("API Error Code: %d", apiErr.Code)
		}
	} else {
		fmt.Printf("✓ Successfully retrieved coin list, total %d coins\n", len(coinListResp.Coins))
		if len(coinListResp.Coins) > 0 {
			fmt.Printf("  First coin: %s (ID: %d)\n", coinListResp.Coins[0].Symbol, coinListResp.Coins[0].CoinID)
		}
	}

	// Test 2: Get merchant assets
	fmt.Println("\n[Test 2] Get Merchant Assets...")
	assetsResp, err := client.MerchantAssets().GetAppCoinAssetList(ctx)
	if err != nil {
		log.Printf("Error getting merchant assets: %v", err)
		if apiErr, ok := err.(*ccpayment.APIError); ok {
			log.Printf("API Error Code: %d", apiErr.Code)
		}
	} else {
		fmt.Printf("✓ Successfully retrieved merchant assets, total %d assets\n", len(assetsResp.Assets))
		if len(assetsResp.Assets) > 0 {
			for i, asset := range assetsResp.Assets {
				if i >= 5 {
					fmt.Printf("  ... and %d more\n", len(assetsResp.Assets)-5)
					break
				}
				fmt.Printf("  %s: %s\n", asset.CoinSymbol, asset.Available)
			}
		}
	}

	fmt.Println("\n========================================")
	fmt.Println("Test completed")
}
