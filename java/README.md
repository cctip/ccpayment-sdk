# CCPayment Java SDK

Official Java SDK for CCPayment API v2.

## Installation

Add the following dependency to your `pom.xml`:

```xml
<dependency>
    <groupId>com.ccpayment</groupId>
    <artifactId>ccpayment-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

## Quick Start

```java
import com.ccpayment.Client;
import com.google.gson.JsonObject;

public class Example {
    public static void main(String[] args) {
        // Create client
        Client client = new Client("your_app_id", "your_app_secret");
        
        try {
            // Get all assets
            JsonObject response = client.merchantAssets().getAppCoinAssetList();
            System.out.println(response);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

## Configuration

### Custom Base URL

```java
client.setBaseUrl("https://custom.ccpayment.com/ccpayment/v2");
```

## API Modules

- **basicInfo** - Token, fiat currency, and chain information queries
- **merchantAssets** - Merchant asset queries
- **merchantDeposit** - Deposit address generation and deposit record queries
- **merchantWithdraw** - Withdrawal requests and withdrawal record queries
- **merchantBatchWithdraw** - Batch withdrawal management
- **userAssets** - User asset queries
- **userDeposit** - User deposit addresses and deposit records
- **userWithdraw** - User withdrawal requests and withdrawal records
- **userTransfer** - User transfers and batch transfers
- **orders** - Order and Invoice orders
- **checkout** - Checkout and Hosted related
- **swap** - Swap related interfaces
- **utilities** - Webhook, address validation, etc.

## Error Handling

```java
import com.ccpayment.APIError;

try {
    JsonObject response = client.merchantAssets().getAppCoinAssetList();
} catch (APIError e) {
    System.out.println("API Error " + e.getCode() + ": " + e.getMessage());
} catch (IOException e) {
    System.out.println("IO Error: " + e.getMessage());
}
```

## API Documentation

For complete API documentation, visit: https://doc.ccpayment.com
