# CCPayment PHP Integration

This folder contains a lightweight PHP integration client generated from the gateway skill workflow.

## Files

- `src/CCPaymentClient.php`: core client with v1/v2 signing support
- `examples/v1_demo.php`: v1 usage
- `examples/v2_hmac_demo.php`: v2 with HMAC signature
- `examples/v2_rsa_demo.php`: v2 with RSA signature

## Requirements

- PHP 8.0+
- cURL extension
- OpenSSL extension (for RSA mode)

## Quick start

```php
require_once 'src/CCPaymentClient.php';

use CCPayment\SDK\CCPaymentClient;

$client = new CCPaymentClient('https://admin.ccpayment.com', 'app-id', 'app-secret');
$response = $client->getSupportToken();
print_r($response);
```

## Signature rules implemented

- v1 merchant: `sha256(appId + appSecret + timestamp + body)`
- v1 hosted: `sha256(billId + timestamp)`
- v2 hmac: `hmac_sha256(appId + timestamp + body, appSecret)`
- v2 rsa: `RSA-SHA256 sign(appId + timestamp + body)` and Base64 encode

## Notes

- `postV1`, `postV2`, and `postV1Hosted` are generic methods for any endpoint.
- Convenience wrappers are provided for common endpoints used in existing SDKs.
