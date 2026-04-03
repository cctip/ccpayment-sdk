package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// API凭证
	appID := "adudHubhJ8qQ45dG"
	appSecret := "0831f9c4233da86199cfafd08c0e8b6d"

	// 创建客户端
	client := NewClient(appID, appSecret)

	// 配置HTTP代理
	proxyURL, err := url.Parse("http://127.0.0.1:10808")
	if err != nil {
		log.Fatalf("解析代理URL失败: %v", err)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: 30 * time.Second,
	}
	client.SetHTTPClient(httpClient)

	fmt.Println("=== CCPayment SDK 测试 ===")
	fmt.Printf("App ID: %s\n", appID)
	fmt.Printf("代理: %s\n\n", proxyURL.String())

	// 测试1: 获取全部资产
	fmt.Println("【测试1】获取商家全部资产...")
	ctx := context.Background()
	
	assetsResp, err := client.MerchantAssets().GetAppCoinAssetList(ctx)
	if err != nil {
		log.Fatalf("获取资产列表失败: %v", err)
	}

	fmt.Printf("✓ 成功获取资产列表，共 %d 种代币\n\n", len(assetsResp.Assets))

	// 打印前10个资产
	fmt.Println("资产列表（前10个）:")
	fmt.Println("----------------------------------------")
	count := 0
	for _, asset := range assetsResp.Assets {
		if count >= 10 {
			break
		}
		fmt.Printf("%d. %s (ID: %d)\n", count+1, asset.CoinSymbol, asset.CoinID)
		fmt.Printf("   可用余额: %s\n", asset.Available)
		count++
	}

	if len(assetsResp.Assets) > 10 {
		fmt.Printf("\n... 还有 %d 种代币\n", len(assetsResp.Assets)-10)
	}

	// 测试2: 获取单个币资产（如果列表中有USDT）
	fmt.Println("\n【测试2】获取单个币资产...")
	
	// 查找USDT或第一个有余额的币
	var testCoinID uint64
	var testCoinSymbol string
	
	for _, asset := range assetsResp.Assets {
		if asset.CoinSymbol == "USDT" || asset.Available != "0" {
			testCoinID = asset.CoinID
			testCoinSymbol = asset.CoinSymbol
			break
		}
	}

	if testCoinID == 0 && len(assetsResp.Assets) > 0 {
		testCoinID = assetsResp.Assets[0].CoinID
		testCoinSymbol = assetsResp.Assets[0].CoinSymbol
	}

	if testCoinID > 0 {
		assetResp, err := client.MerchantAssets().GetAppCoinAsset(ctx, testCoinID)
		if err != nil {
			log.Printf("获取单个资产失败: %v", err)
		} else {
			fmt.Printf("✓ 成功获取 %s 资产信息\n", testCoinSymbol)
			fmt.Printf("   代币ID: %d\n", assetResp.Asset.CoinID)
			fmt.Printf("   代币符号: %s\n", assetResp.Asset.CoinSymbol)
			fmt.Printf("   可用余额: %s\n", assetResp.Asset.Available)
		}
	}

	fmt.Println("\n=== 测试完成 ===")
}
