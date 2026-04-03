# CCPayment API - Shell Scripts

This directory contains shell scripts for testing CCPayment API v2.

## Files

- `env.sh` - Environment configuration with API credentials and signature function
- `basic_info.sh` - Basic info module API examples (8 APIs)

## Usage

1. Configure your API credentials in `env.sh`:
```bash
export APP_ID="your_actual_app_id"
export APP_SECRET="your_actual_app_secret"
```

2. Source the environment and run the script:
```bash
source env.sh && ./basic_info.sh
```

## Requirements

- `curl` - HTTP client
- `openssl` - For HMAC-SHA256 signature generation
- `jq` - JSON processor (optional, for pretty output)

## API Modules

### Basic Info (`basic_info.sh`)
- `getCoinList` - 获取代币列表
- `getCoin` - 获取单个代币
- `getCoinUSDTPrice` - 获取代币USD价格
- `getFiatList` - 获取法币列表
- `getChainList` - 获取链列表
- `all/chain` - 获取所有链列表
- `getCwalletUserId` - 获取CCWallet用户ID
- `getWithdrawFee` - 获取提现手续费

## Signature Generation

The `env.sh` provides a `generate_sign` function that creates HMAC-SHA256 signatures:

```bash
generate_sign '{"coinId":1280}'
```

Signature format: `HMAC-SHA256(AppID + Timestamp + Body, AppSecret)`

## Response Format

All APIs return JSON with the following structure:
```json
{
  "code": 10000,
  "msg": "success",
  "data": {}
}
```

Code 10000 indicates success. Other codes indicate errors.

## Testing All APIs

```bash
# Test basic info module
source env.sh && ./basic_info.sh
```
