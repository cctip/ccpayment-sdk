"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const client_1 = require("./src/client");
const errors_1 = require("./src/errors");
// Get credentials from environment variables
const appId = process.env.CCPAYMENT_APP_ID || '';
const appSecret = process.env.CCPAYMENT_APP_SECRET || '';
const proxyUrl = process.env.CCPAYMENT_PROXY_URL || '';
if (!appId || !appSecret) {
    console.error('Error: Please set CCPAYMENT_APP_ID and CCPAYMENT_APP_SECRET environment variables');
    process.exit(1);
}
console.log(`API Key: ${appId}`);
console.log(`Using Proxy: ${!!proxyUrl}`);
console.log('='.repeat(50));
// Create client
const client = new client_1.Client(appId, appSecret);
// Configure HTTP proxy if provided
if (proxyUrl) {
    client.setProxy(`http://${proxyUrl}`);
    console.log(`Proxy configured: ${proxyUrl}`);
}
async function runTests() {
    // Test 1: Get coin list
    console.log('\n[Test 1] Get Coin List...');
    try {
        const response = await client.basicInfo.getCoinList();
        console.log(`✓ Successfully retrieved coin list, total ${response.coins?.length || 0} coins`);
        if (response.coins && response.coins.length > 0) {
            console.log(`  First coin: ${response.coins[0].symbol} (ID: ${response.coins[0].coinId})`);
        }
    }
    catch (error) {
        if (error instanceof errors_1.APIError) {
            console.log(`✗ API Error: code=${error.code}, message=${error.message}`);
        }
        else {
            console.log(`✗ Error: ${error}`);
        }
    }
    // Test 2: Get merchant assets
    console.log('\n[Test 2] Get Merchant Assets...');
    try {
        const response = await client.merchantAssets.getAppCoinAssetList();
        const assets = response.assets || [];
        console.log(`✓ Successfully retrieved merchant assets, total ${assets.length} assets`);
        if (assets.length > 0) {
            for (let i = 0; i < Math.min(5, assets.length); i++) {
                console.log(`  ${assets[i].coinSymbol}: ${assets[i].available}`);
            }
            if (assets.length > 5) {
                console.log(`  ... and ${assets.length - 5} more`);
            }
        }
    }
    catch (error) {
        if (error instanceof errors_1.APIError) {
            console.log(`✗ API Error: code=${error.code}, message=${error.message}`);
        }
        else {
            console.log(`✗ Error: ${error}`);
        }
    }
    // Test 3: Health check
    console.log('\n[Test 3] Health Check...');
    try {
        const response = await client.utilities.health();
        console.log(`✓ Health check passed: ${JSON.stringify(response)}`);
    }
    catch (error) {
        if (error instanceof errors_1.APIError) {
            console.log(`✗ API Error: code=${error.code}, message=${error.message}`);
        }
        else {
            console.log(`✗ Error: ${error}`);
        }
    }
    console.log('\n' + '='.repeat(50));
    console.log('Test completed');
}
runTests();
//# sourceMappingURL=test.js.map