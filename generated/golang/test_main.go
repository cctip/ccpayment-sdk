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
	proxyURL, _ := url.Parse("http://127.0.0.1:10808")
	httpClient := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		Timeout:   30 * time.Second,
	}
	client.SetHTTPClient(httpClient)

	fmt.Println("=== 商家资产模块测试 ===")
	fmt.Printf("App ID: %s\n\n", appID)

	// 测试1: 获取全部资产
	fmt.Println("【测试1】获取全部资产...")
	resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✓ 成功获取资产列表，共 %d 种代币\n\n", len(resp.Assets))

	// 显示前10个
	fmt.Println("资产列表（前10个）:")
	fmt.Println("--------------------------------------------------")
	for i, asset := range resp.Assets {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %s (ID: %d)\n", i+1, asset.CoinSymbol, asset.CoinID)
		fmt.Printf("   可用余额: %s\n", asset.Available)
	}

	if len(resp.Assets) > 10 {
		fmt.Printf("\n... 还有 %d 种代币\n", len(resp.Assets)-10)
	}

	// 测试2: 获取单个币资产
	if len(resp.Assets) > 0 {
		fmt.Println("\n【测试2】获取单个币资产...")

		// 查找USDT或第一个资产
		testAsset := resp.Assets[0]
		for _, asset := range resp.Assets {
			if asset.CoinSymbol == "USDT" {
				testAsset = asset
				break
			}
		}

		assetResp, err := client.MerchantAssets().GetAppCoinAsset(context.Background(), testAsset.CoinID)
		if err != nil {
			log.Printf("获取失败: %v", err)
		} else {
			fmt.Printf("✓ 成功获取 %s 资产信息\n", assetResp.Asset.CoinSymbol)
			fmt.Printf("   代币ID: %d\n", assetResp.Asset.CoinID)
			fmt.Printf("   代币符号: %s\n", assetResp.Asset.CoinSymbol)
			fmt.Printf("   可用余额: %s\n", assetResp.Asset.Available)
		}
	}

	fmt.Println("\n=== 测试完成 ===")
}
