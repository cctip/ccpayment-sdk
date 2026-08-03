# CCPayment Java SDK

Official Java SDK for CCPayment API v2.

## Installation

### Maven

Add to your `pom.xml`:

```xml
<dependency>
    <groupId>com.ccpayment</groupId>
    <artifactId>ccpayment-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

### Gradle

```groovy
implementation 'com.ccpayment:ccpayment-sdk:1.0.0'
```

## Quick Start

```java
import com.ccpayment.Client;
import com.ccpayment.APIError;
import com.google.gson.JsonObject;

public class Example {
    public static void main(String[] args) {
        // Create client
        Client client = new Client("your_app_id", "your_app_secret");
        
        // Get all assets
        try {
            JsonObject response = client.post("/getAppCoinAssetList", null, JsonObject.class);
            System.out.println(response);
        } catch (APIError e) {
            System.out.println("API Error: " + e.getMessage());
        }
    }
}
```

## Configuration

### HTTP Proxy

```java
import java.net.InetSocketAddress;
import java.net.Proxy;

Proxy proxy = new Proxy(Proxy.Type.HTTP, new InetSocketAddress("127.0.0.1", 10808));
client.setProxy(proxy);
```

### Custom Base URL

```java
client.setBaseUrl("https://custom.ccpayment.com/ccpayment/v2");
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

```java
import com.ccpayment.Client;
import com.ccpayment.APIError;

try {
    JsonObject response = client.post("/getAppCoinAssetList", null, JsonObject.class);
} catch (APIError e) {
    System.out.println("API Error " + e.getCode() + ": " + e.getMessage());
} catch (Exception e) {
    System.out.println("Error: " + e.getMessage());
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
