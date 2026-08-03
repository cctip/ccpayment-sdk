# CCPayment Python SDK

Official Python SDK for CCPayment API v2.

## Installation

```bash
pip install ccpayment-sdk
```

Or install from source:

```bash
git clone https://github.com/ccpayment/ccpayment-sdk-python.git
cd ccpayment-sdk-python
pip install -r requirements.txt
```

## Quick Start

```python
from ccpayment import Client

# Create client
client = Client("your_app_id", "your_app_secret")

# Get all assets
response = client.merchant_assets().get_app_coin_asset_list()
for asset in response["assets"]:
    print(f"{asset['coinSymbol']}: {asset['available']}")
```

## Configuration

### HTTP Proxy

```python
client.set_proxy("http://127.0.0.1:10808")
```

### Custom Base URL

```python
client.set_base_url("https://custom.ccpayment.com/ccpayment/v2")
```

## API Modules

- **basic_info** - Token, fiat currency, and chain information queries
- **merchant_assets** - Merchant asset queries
- **merchant_deposit** - Deposit address generation and deposit record queries
- **merchant_withdraw** - Withdrawal requests and withdrawal record queries
- **merchant_batch_withdraw** - Batch withdrawal management
- **user_assets** - User asset queries
- **user_deposit** - User deposit addresses and deposit records
- **user_withdraw** - User withdrawal requests and withdrawal records
- **user_transfer** - User transfers and batch transfers
- **orders** - Order and Invoice orders
- **swap** - Swap related interfaces
- **user_swap** - User swap management
- **utilities** - Webhook, address validation, etc.

## Error Handling

```python
from ccpayment import Client, APIError

client = Client("your_app_id", "your_app_secret")

try:
    response = client.merchant_assets().get_app_coin_asset_list()
except APIError as e:
    print(f"API Error {e.code}: {e.message}")
```

## API Documentation

For complete API documentation, visit: https://ccpayment.com/api/doc

## Disclaimer

This SDK is provided by CCPayment on an "as is" and "as available" basis for use at your own risk. CCPayment makes no representations or warranties of any kind, whether express or implied, as to the operation of the SDK, its accuracy, reliability, completeness, or fitness for any particular purpose.

To the fullest extent permitted by law, CCPayment shall not be liable for any losses, damages, or expenses of any kind arising from or in connection with your use of, or inability to use, this SDK, including but not limited to any financial losses resulting from errors, bugs, interruptions, or inaccuracies in the SDK.

Your use of this SDK to access the CCPayment platform is subject to the [CCPayment Terms of Use](https://ccpayment.com/en/privacy-policy#ccpayment-service-agreement), which shall prevail in the event of any conflict with this disclaimer. You are solely responsible for the transactions and API operations executed through your merchant account using this SDK. For a secure integration, please follow the verification practices described in the [API documentation](https://ccpayment.com/api/doc?en).

Plugins or extensions developed by third parties based on this SDK or the CCPayment API are independent works of their authors. Unless expressly designated as official in CCPayment's documentation, they are not maintained by CCPayment, and CCPayment is not responsible for their operation. We welcome community contributions — developers building on CCPayment are invited to reach out at partner@ccpayment.com.

This SDK is not intended to constitute investment advice or a recommendation to buy, sell, or hold any digital asset. You should independently evaluate and verify all information before acting.
