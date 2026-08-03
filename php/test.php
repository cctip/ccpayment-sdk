<?php

require_once __DIR__ . '/vendor/autoload.php';

use CCPayment\Client;
use CCPayment\APIError;

// Get credentials from environment variables
$appId = $_ENV['CCPAYMENT_APP_ID'] ?? '';
$appSecret = $_ENV['CCPAYMENT_APP_SECRET'] ?? '';
$proxyUrl = $_ENV['CCPAYMENT_PROXY_URL'] ?? '';

if (empty($appId) || empty($appSecret)) {
    echo "Error: Please set CCPAYMENT_APP_ID and CCPAYMENT_APP_SECRET environment variables\n";
    exit(1);
}

echo "API Key: {$appId}\n";
echo "Using Proxy: " . (!empty($proxyUrl) ? 'Yes' : 'No') . "\n";
echo str_repeat('=', 50) . "\n";

// Create client
$client = new Client($appId, $appSecret);

// Configure HTTP proxy if provided
if (!empty($proxyUrl)) {
    $client->setProxy($proxyUrl);
    echo "Proxy configured: {$proxyUrl}\n";
}

// Test 1: Get coin list
echo "\n[Test 1] Get Coin List...\n";
try {
    $response = $client->basicInfo()->getCoinList();
    $coins = $response['coins'] ?? [];
    $totalCoins = count($coins);
    echo "✓ Successfully retrieved coin list, total {$totalCoins} coins\n";
    if ($totalCoins > 0) {
        echo "  First coin: {$coins[0]['symbol']} (ID: {$coins[0]['coinId']})\n";
    }
} catch (APIError $e) {
    echo "✗ API Error: code={$e->getCode()}, message={$e->getMessage()}\n";
} catch (Exception $e) {
    echo "✗ Error: {$e->getMessage()}\n";
}

// Test 2: Get merchant assets
echo "\n[Test 2] Get Merchant Assets...\n";
try {
    $response = $client->merchantAssets()->getAppCoinAssetList();
    $assets = $response['assets'] ?? [];
    $totalAssets = count($assets);
    echo "✓ Successfully retrieved merchant assets, total {$totalAssets} assets\n";
    if ($totalAssets > 0) {
        $limit = min(5, $totalAssets);
        for ($i = 0; $i < $limit; $i++) {
            echo "  {$assets[$i]['coinSymbol']}: {$assets[$i]['available']}\n";
        }
        if ($totalAssets > 5) {
            echo "  ... and " . ($totalAssets - 5) . " more\n";
        }
    }
} catch (APIError $e) {
    echo "✗ API Error: code={$e->getCode()}, message={$e->getMessage()}\n";
} catch (Exception $e) {
    echo "✗ Error: {$e->getMessage()}\n";
}

echo "\n" . str_repeat('=', 50) . "\n";
echo "Test completed\n";
