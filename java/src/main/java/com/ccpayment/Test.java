package com.ccpayment;

import com.google.gson.JsonObject;
import com.google.gson.JsonArray;
import java.net.InetSocketAddress;
import java.net.Proxy;

public class Test {
    public static void main(String[] args) {
        // Get credentials from environment variables
        String appId = System.getenv("CCPAYMENT_APP_ID");
        String appSecret = System.getenv("CCPAYMENT_APP_SECRET");
        String proxyUrl = System.getenv("CCPAYMENT_PROXY_URL");

        if (appId == null || appId.isEmpty() || appSecret == null || appSecret.isEmpty()) {
            System.err.println("Error: Please set CCPAYMENT_APP_ID and CCPAYMENT_APP_SECRET environment variables");
            System.exit(1);
        }

        System.out.println("API Key: " + appId);
        System.out.println("Using Proxy: " + (proxyUrl != null && !proxyUrl.isEmpty()));
        System.out.println("=".repeat(50));

        // Create client
        Client client = new Client(appId, appSecret);

        // Configure HTTP proxy if provided
        if (proxyUrl != null && !proxyUrl.isEmpty()) {
            String[] parts = proxyUrl.split(":");
            String host = parts[0];
            int port = Integer.parseInt(parts[1]);
            Proxy proxy = new Proxy(Proxy.Type.HTTP, new InetSocketAddress(host, port));
            client.setProxy(proxy);
            System.out.println("Proxy configured: " + proxyUrl);
        }

        // Test 1: Get coin list
        System.out.println("\n[Test 1] Get Coin List...");
        try {
            JsonObject response = client.post("/getCoinList", null, JsonObject.class);
            JsonArray coins = response.getAsJsonArray("coins");
            int totalCoins = coins != null ? coins.size() : 0;
            System.out.println("✓ Successfully retrieved coin list, total " + totalCoins + " coins");
            if (totalCoins > 0) {
                JsonObject firstCoin = coins.get(0).getAsJsonObject();
                System.out.println("  First coin: " + firstCoin.get("symbol").getAsString() + " (ID: " + firstCoin.get("coinId").getAsInt() + ")");
            }
        } catch (APIError e) {
            System.out.println("✗ API Error: code=" + e.getCode() + ", message=" + e.getMessage());
        } catch (Exception e) {
            System.out.println("✗ Error: " + e.getMessage());
            e.printStackTrace();
        }

        // Test 2: Get merchant assets
        System.out.println("\n[Test 2] Get Merchant Assets...");
        try {
            JsonObject response = client.post("/getAppCoinAssetList", null, JsonObject.class);
            JsonArray assets = response.getAsJsonArray("assets");
            int totalAssets = assets != null ? assets.size() : 0;
            System.out.println("✓ Successfully retrieved merchant assets, total " + totalAssets + " assets");
            if (totalAssets > 0) {
                int limit = Math.min(5, totalAssets);
                for (int i = 0; i < limit; i++) {
                    JsonObject asset = assets.get(i).getAsJsonObject();
                    System.out.println("  " + asset.get("coinSymbol").getAsString() + ": " + asset.get("available").getAsString());
                }
                if (totalAssets > 5) {
                    System.out.println("  ... and " + (totalAssets - 5) + " more");
                }
            }
        } catch (APIError e) {
            System.out.println("✗ API Error: code=" + e.getCode() + ", message=" + e.getMessage());
        } catch (Exception e) {
            System.out.println("✗ Error: " + e.getMessage());
            e.printStackTrace();
        }

        System.out.println("\n" + "=".repeat(50));
        System.out.println("Test completed");
    }
}
