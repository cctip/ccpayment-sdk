---
name: ccpayment.sdk.codegen
version: 2.0.0
description: CCPayment SDK Code Generator - Generate strongly-typed SDK code from API documentation
homepage: https://ccpayment.com
metadata:
  {
    "ccpayment":
      {
        "category": "sdk-generator",
        "api_base": "https://ccpayment.com/ccpayment/v2/",
        "api_docs": ".rules/api/",
        "supported_languages": ["golang", "typescript", "python", "java", "php"]
      },
  }
---
> **AI Agent Instructions** — This is a code generation skill for CCPayment SDK. When invoked with `ccpayment.sdk.codegen <language>`, generate strongly-typed SDK code based on the API documentation in `.rules/api/` directory.

# CCPayment SDK Code Generator

## 概述

这是CCPayment SDK的代码生成工具。它从`.rules/api/`目录中的API文档自动生成强类型的SDK代码。

## 使用方式

```bash
ccpayment.sdk.codegen <language>
```

支持的语言：
- `golang` - 生成Go语言SDK
- `typescript` - 生成TypeScript SDK
- `python` - 生成Python SDK
- `java` - 生成Java SDK
- `php` - 生成PHP SDK

## 代码生成规则

### 1. 输入源

从以下文件读取API定义：
- `.rules/api/README.md` - API概述和认证信息
- `.rules/api/01-basic-info.md` - 基础信息模块
- `.rules/api/02-merchant-assets.md` - 商家资产模块
- `.rules/api/03-merchant-deposit.md` - 商家充值模块
- `.rules/api/04-merchant-withdraw.md` - 商家提现模块
- `.rules/api/05-merchant-batch-withdraw.md` - 商家批量提现模块
- `.rules/api/06-user-assets.md` - 用户资产模块
- `.rules/api/07-user-deposit.md` - 用户充值模块
- `.rules/api/08-user-withdraw.md` - 用户提现模块
- `.rules/api/09-user-transfer.md` - 用户转账模块
- `.rules/api/10-orders.md` - 订单模块
- `.rules/api/11-checkout.md` - 收银台模块
- `.rules/api/12-swap.md` - 换币模块
- `.rules/api/13-utilities.md` - 工具接口
- `.rules/api/appendix.md` - 附录（数据类型、错误码等）

### 2. 输出结构

生成的SDK应包含：

#### 核心组件
1. **Client** - SDK客户端，处理认证和请求
2. **Models** - 所有请求和响应的数据模型
3. **Signature** - 签名生成逻辑
4. **Errors** - 错误处理和错误码定义

#### 模块化API
每个API模块生成独立的服务类：
- BasicInfoService
- MerchantAssetsService
- MerchantDepositService
- MerchantWithdrawService
- MerchantBatchWithdrawService
- UserAssetsService
- UserDepositService
- UserWithdrawService
- UserTransferService
- OrdersService
- CheckoutService
- SwapService
- UtilitiesService

### 3. 强类型要求

#### Golang示例

```go
// 所有字段必须有明确的类型
type GetCoinListResponse struct {
    Coins []CoinEntity `json:"coins"`
}

type CoinEntity struct {
    CoinID       uint64            `json:"coinId"`
    Symbol       string            `json:"symbol"`
    CoinFullName string            `json:"coinFullName"`
    LogoURL      string            `json:"logoUrl"`
    Status       string            `json:"status"`
    Networks     map[string]Network `json:"networks"`
    Price        string            `json:"price"`
}

type Network struct {
    Chain                  string `json:"chain"`
    ChainFullName          string `json:"chainFullName"`
    Contract               string `json:"contract"`
    Precision              uint32 `json:"precision"`
    CanDeposit             bool   `json:"canDeposit"`
    CanWithdraw            bool   `json:"canWithdraw"`
    MinimumDepositAmount   string `json:"minimumDepositAmount"`
    MinimumWithdrawAmount  string `json:"minimumWithdrawAmount"`
    MaximumWithdrawAmount  string `json:"maximumWithdrawAmount"`
    IsSupportMemo          bool   `json:"isSupportMemo"`
    Protocol               string `json:"protocol"`
}
```

#### 验证规则映射

从API文档的验证规则生成代码验证：

| 文档规则 | Golang实现 |
|---------|------------|
| 长度≥1 | `if len(s) < 1 { return ErrInvalidLength }` |
| 长度3-64 | `if len(s) < 3 \|\| len(s) > 64 { return ErrInvalidLength }` |
| ≥1 | `if n < 1 { return ErrInvalidValue }` |
| 邮箱格式 | `if !emailRegex.MatchString(s) { return ErrInvalidEmail }` |
| URI格式 | `if _, err := url.Parse(s); err != nil { return ErrInvalidURI }` |
| 最大50项 | `if len(arr) > 50 { return ErrTooManyItems }` |

### 4. API调用示例

生成的SDK应支持链式调用和清晰的API：

```go
// Golang示例
client := ccpayment.NewClient(appID, appSecret)

// 获取代币列表
resp, err := client.BasicInfo().GetCoinList(ctx)
if err != nil {
    log.Fatal(err)
}

// 创建充值订单
req := &ccpayment.CreateDepositOrderRequest{
    OrderID: "order_001",
    CoinID:  1280,
    Chain:   "POLYGON",
    Price:   "10",
    FiatID:  ptr.Uint64(1033),
}
resp, err := client.MerchantDeposit().CreateOrderDepositAddress(ctx, req)
```

## 生成步骤

当用户执行 `ccpayment.sdk.codegen golang` 时：

1. **读取所有API文档** - 解析`.rules/api/`下的所有markdown文件
2. **提取API定义** - 从表格中提取接口路径、参数、响应结构
3. **生成类型定义** - 为所有请求和响应创建强类型结构
4. **生成验证逻辑** - 根据验证规则生成参数校验代码
5. **生成API方法** - 为每个接口生成对应的方法
6. **生成签名逻辑** - 实现HMAC-SHA256签名算法
7. **生成错误处理** - 创建错误类型和错误码映射
8. **生成文档** - 为生成的代码添加注释和使用示例

## 输出位置

生成的代码应保存到：
```
generated/v2/<language>/
├── client.go/ts/py/java/php
├── models.go/ts/py/java/php
├── signature.go/ts/py/java/php
├── errors.go/ts/py/java/php
├── services/
│   ├── basic_info.go
│   ├── merchant_assets.go
│   ├── merchant_deposit.go
│   └── ...
└── README.md
```

## 参考资源

| 资源 | 链接 |
|------|------|
| API文档 | `.rules/api/README.md` |
| 官方网站 | https://ccpayment.com |
| Base URL | https://ccpayment.com/ccpayment/v2/ |

---

## Platform Overview

### Core Services

| Feature | Description |
| --- | --- |
| Crypto Payments | Accept 900+ tokens on 100+ networks, fees as low as 0.2% |
| Permanent Deposit Addresses | Persistent addresses linked to users — ideal for gaming, streaming |
| Order-based Deposits | Addresses tied to specific orders — ideal for e-commerce |
| Invoice Checkout Pages | Hosted pages where buyers choose currency and network |
| Network Withdrawals | Send crypto to any blockchain address |
| Cwallet Withdrawals | Instant internal transfers to Cwallet accounts (zero network fee) |
| Swap | Exchange between cryptocurrencies via API |
| Wallet System | Per-user wallets with independent balances, deposits, withdrawals, internal transfers, swaps |
| Risk Management | Real-time AML detection with customizable handling |
| Testnet | ETH Sepolia testnet for integration testing |

### Compliance & Licenses

Malta (EU) EMI, Malta VFA Class 4, UK API Payment, Australia Exchange, India Aggregated Payment, South Africa Exchange, USA MSB.

### Performance

- 90+ Countries | $20B+ Volume | 67K+ Partners | 99% Success Rate | ~1.8 min Avg Speed

---

## Authentication

**All API calls** are `POST` requests requiring three custom headers:

| Header | Type | Description |
| --- | --- | --- |
| `Appid` | String | Your App ID from [Dashboard > Developer](https://console.ccpayment.com/developer/config) |
| `Sign` | String | HMAC-SHA256 signature |
| `Timestamp` | String | Current Unix timestamp in **seconds** (10-digit) |

### Credential Setup

1. Register at [https://console.ccpayment.com](https://console.ccpayment.com)
2. Go to **Dashboard > Developer** to get your `App ID` and `App Secret`
3. Optionally configure IP whitelist for security

### Signature Algorithm

```
IF request body is non-empty:
  signText = appId + timestamp + JSON.stringify(requestBody)
ELSE:
  signText = appId + timestamp

sign = HMAC-SHA256(signText, appSecret).hexDigest()
```

### Complete Authentication Example (Node.js)

```javascript
const crypto = require("crypto");
const https = require("https");

const appId = "YOUR_APP_ID";
const appSecret = "YOUR_APP_SECRET";

function callCCPaymentAPI(apiUrl, body = null) {
  const timestamp = Math.floor(Date.now() / 1000);
  const bodyStr = body ? JSON.stringify(body) : "";

  let signText = appId + timestamp;
  if (bodyStr) {
    signText += bodyStr;
  }

  const sign = crypto
    .createHmac("sha256", appSecret)
    .update(signText)
    .digest("hex");

  const headers = {
    "Content-Type": "application/json",
    Appid: appId,
    Sign: sign,
    Timestamp: timestamp.toString(),
  };

  // Use your preferred HTTP client (axios, fetch, etc.)
  // POST to apiUrl with headers and bodyStr
}
```

### Authentication Example (Python)

```python
import hmac, hashlib, json, time, requests

APP_ID = "YOUR_APP_ID"
APP_SECRET = "YOUR_APP_SECRET"

def call_ccpayment_api(url, body=None):
    timestamp = str(int(time.time()))
    body_str = json.dumps(body) if body else ""

    sign_text = APP_ID + timestamp
    if body_str:
        sign_text += body_str

    sign = hmac.new(
        APP_SECRET.encode(), sign_text.encode(), hashlib.sha256
    ).hexdigest()

    headers = {
        "Content-Type": "application/json",
        "Appid": APP_ID,
        "Sign": sign,
        "Timestamp": timestamp,
    }

    resp = requests.post(url, headers=headers, data=body_str)
    return resp.json()
```

### Authentication Example (Go)

```go
package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strconv"
    "time"
)

func generateSign(appId, appSecret, bodyStr string) (string, string) {
    timestamp := strconv.FormatInt(time.Now().Unix(), 10)
    signText := appId + timestamp
    if bodyStr != "" {
        signText += bodyStr
    }
    mac := hmac.New(sha256.New, []byte(appSecret))
    mac.Write([]byte(signText))
    sign := hex.EncodeToString(mac.Sum(nil))
    return sign, timestamp
}
```

### Authentication Example (PHP)

```php
$appId = 'YOUR_APP_ID';
$appSecret = 'YOUR_APP_SECRET';
$timestamp = time();
$bodyStr = json_encode($body);

$signText = $appId . $timestamp;
if ($bodyStr) {
    $signText .= $bodyStr;
}

$sign = hash_hmac('sha256', $signText, $appSecret);

$headers = [
    'Content-Type: application/json',
    'Appid: ' . $appId,
    'Sign: ' . $sign,
    'Timestamp: ' . $timestamp,
];
```

### Authentication Example (Java)

```java
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

String appId = "YOUR_APP_ID";
String appSecret = "YOUR_APP_SECRET";
long timestamp = System.currentTimeMillis() / 1000;
String bodyStr = new ObjectMapper().writeValueAsString(body);

String signText = appId + timestamp;
if (bodyStr != null && !bodyStr.isEmpty()) {
    signText += bodyStr;
}

Mac mac = Mac.getInstance("HmacSHA256");
mac.init(new SecretKeySpec(appSecret.getBytes(), "HmacSHA256"));
byte[] hash = mac.doFinal(signText.getBytes());
String sign = bytesToHex(hash);
```

---

## Response Format

All API responses use this structure:

```json
{
  "code": 10000,
  "msg": "success",
  "data": { ... }
}
```

| Code | Meaning |
| --- | --- |
| `10000` | Success |
| `10001` | Failed — see `msg` for error detail |
| Others | Specific error — see [Error Codes](#error-codes) |

---

## Testnet

Enable ETH Sepolia for testing:

1. Go to [Dashboard > Merchant Settings > ETH test network](https://console.ccpayment.com/merchatsetting/menu/settings) and turn on
2. Get free Sepolia ETH: https://cloud.google.com/application/web3/faucet/ethereum/sepolia
3. Verify transactions: https://sepolia.etherscan.io/
4. **Disable testnet when done** — Sepolia ETH has no real value

---

## Webhooks

### Setup

1. Set URL at [Dashboard > Developer > Webhook URL](https://console.ccpayment.com/developer/config)
2. Whitelist IPs: `54.150.123.157`, `35.72.150.75`, `18.176.186.244`
3. Configure tokens at [Dashboard > Settings > Tokens for your business](https://console.ccpayment.com/merchatsetting/menu/settings)

### Webhook Types

| Type | Trigger |
| --- | --- |
| `DirectDeposit` | Deposit to permanent address |
| `ApiDeposit` | Deposit to order-based address or invoice address |
| `ApiWithdrawal` | Merchant-level network withdrawal status |
| `UserDeposit` | Deposit to user wallet address |
| `UserWithdrawal` | User wallet network withdrawal status |

### Transaction Statuses

| Status | Meaning |
| --- | --- |
| `Processing` | Blockchain is processing |
| `Success` | Transaction confirmed |
| `Failed` | Transaction failed |
| `WaitingApproval` | Withdrawal pending manual approval (if enabled) |
| `Rejected` | Withdrawal rejected by merchant |

### Response Requirement

Your server **must** return HTTP 200 with the string `"Success"` in the response body. Otherwise CCPayment retries.

### Retry Schedule (up to 6 times)

30s → 1m30s → 3m30s → 7m30s → 15m30s → 31m30s

### Idempotency (CRITICAL)

Always check `recordId` or `txId` before processing to prevent duplicate crediting:

```javascript
if (db.hasRecord(recordId)) {
  return res.status(200).send("Success"); // already processed
}
// process transaction, then save recordId to db
res.status(200).send("Success");
```

### Risky Transactions

If `isFlaggedAsRisky` is `true`, funds are NOT auto-credited. Handle at [Dashboard > Transaction > Risky Transaction](https://console.ccpayment.com/transaction/list/risk-list). **Never auto-credit risky transactions to users.**

### Important Webhook Rule

**Webhooks are notifications only.** After receiving a webhook, always call the corresponding Get Record API to confirm the actual payment status. Do not treat webhook status as final — this may result in asset loss.

### Webhook Payload Examples

**DirectDeposit:**

```json
{
  "type": "DirectDeposit",
  "msg": {
    "recordId": "20240313113108...",
    "referenceId": "user_12345",
    "coinId": 1329,
    "coinSymbol": "MATIC",
    "status": "Success",
    "isFlaggedAsRisky": false
  }
}
```

**ApiDeposit:**

```json
{
  "type": "ApiDeposit",
  "msg": {
    "recordId": "20240313121919...",
    "orderId": "order_001",
    "coinId": 1329,
    "coinSymbol": "POL",
    "status": "Success",
    "isFlaggedAsRisky": false
  }
}
```

**ApiWithdrawal:**

```json
{
  "type": "ApiWithdrawal",
  "msg": {
    "recordId": "20240313120722...",
    "orderId": "withdraw_001",
    "coinId": 1891,
    "coinSymbol": "TETH",
    "status": "Success"
  }
}
```

**UserDeposit:**

```json
{
  "type": "UserDeposit",
  "msg": {
    "recordId": "20240313093815...",
    "userId": "user_12345",
    "coinId": 1329,
    "coinSymbol": "MATIC",
    "amount": "0.1",
    "status": "Success",
    "isFlaggedAsRisky": false
  }
}
```

**UserWithdrawal:**

```json
{
  "type": "UserWithdrawal",
  "msg": {
    "recordId": "20240313120733...",
    "orderId": "user_wd_001",
    "userId": "user_12345",
    "coinId": 1891,
    "coinSymbol": "TETH",
    "status": "Success"
  }
}
```

### Resend Webhook API

**Use case:** Resend failed or all webhooks for transactions within a time range.

**POST** `https://ccpayment.com/v2/resendWebhook`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordIds | Array[String] | No | Specific record IDs to resend (max 50) |
| startTimestamp | Integer | Yes | 10-digit start timestamp |
| endTimestamp | Integer | No | 10-digit end timestamp (max 1 hour after start) |
| webhookResult | String | No | `Failed` (default) or `AllResult` |
| transactionType | String | No | `AllType` (default), `ApiDeposit`, `DirectDeposit`, `ApiWithdrawal`, `UserDeposit`, `UserWithdrawal` |

**Response:**

```json
{
  "code": 10000,
  "msg": "Success",
  "data": { "resendCount": 1 }
}
```

---

## Deposit APIs

### Which deposit type should I use?

| Use Case | Deposit Type | API |
| --- | --- | --- |
| E-commerce / fixed-price orders, you specify coin | Order-based (merchant-specified) | Create Deposit Order |
| E-commerce, customer picks coin | Order-based (customer-selected) | Create Invoice Deposit |
| Gaming / streaming / social — user-linked deposits | Permanent address | Get Permanent Deposit Address |
| Building a wallet system with per-user accounts | User deposit | Create/Get User Deposit Address |

---

### 1. Get Permanent Deposit Address

**Use case:** Assign persistent deposit addresses to users. Ideal for gaming, streaming, social platforms where deposits are linked to users, not orders.

**POST** `https://ccpayment.com/v2/getOrCreateDepositAddress`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| referenceId | String | Yes | 3-64 chars. Your unique user ID in your system |
| chain | String | Yes | Chain symbol (e.g. `TRX`, `POLYGON`, `BSC`, `ETH`). Get from Get Token Information API |

**Rules:**
- If the `referenceId + chain` combo already exists, returns the existing address
- If not, creates a new permanent address
- Max 1000 addresses per App ID (contact support for more)
- The address is permanent — user can reuse it indefinitely

**Request Example:**

```javascript
const body = {
  referenceId: "user_12345",
  chain: "TRX"
};
// POST with authentication headers
```

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "address": "TAkmn3f8bgwMAwdVrGfUVSRg4W2XwqgHGc",
    "memo": ""
  }
}
```

| Response Field | Type | Description |
| --- | --- | --- |
| data.address | String | The permanent deposit address |
| data.memo | String | For memo-required chains (XRP, XLM, TON), payer must include this memo |

**Webhook:** `DirectDeposit` — then call Get Deposit Record to confirm.

**Best Practices:**
- Store the `referenceId → address` mapping in your database
- For memo-required chains, always display the memo prominently to the user
- After receiving a `DirectDeposit` webhook, call Get Deposit Record API to confirm status

---

### 2. Create Deposit Order (Merchant-specified Currency)

**Use case:** E-commerce with fixed-price orders where you specify which coin and chain the customer pays with.

**POST** `https://ccpayment.com/v2/createDepositOrder`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | Yes | Coin ID of the payment coin (get from Get Token List) |
| chain | String | Yes | Chain symbol (e.g. `POLYGON`, `BSC`) |
| price | String | Yes | Product price. If `fiatId` is set: fiat price (up to 2 decimal places). If `fiatId` is empty: crypto amount (up to 18 decimal places) |
| orderId | String | Yes | Your order ID, 3-64 characters, must be unique |
| fiatId | Integer | No | Fiat currency ID (get from Get Fiat List). If set, price is in fiat and auto-converted to crypto |
| expiredAt | Integer | No | 10-digit expiration timestamp. Default 10 days max. Rate is locked before expiration |
| buyerEmail | String | No | Customer email for payment notifications |
| generateCheckoutURL | Boolean | No | `true` = generate hosted checkout page URL; `false` (default) |
| product | String | No | Product name on checkout page (max 120 chars). Only when checkout URL enabled |
| returnUrl | String | No | Redirect URL after payment. Only when checkout URL enabled |
| closeUrl | String | No | Redirect URL when buyer clicks "leave page". Only when checkout URL enabled |

**Rules:**
- `expiredAt` tip: For volatile crypto markets, use short expiration to prevent rate loss. After expiration, deposits are still received for 7 days (labeled "Overdue")
- The rate is locked from order creation until expiration

**Request Example:**

```javascript
const body = {
  coinId: 1280,       // USDT
  price: "10",        // 10 USD
  fiatId: 1033,       // USD
  orderId: "order_20240312_001",
  chain: "POLYGON",
  generateCheckoutURL: true,
  product: "Premium Plan",
  returnUrl: "https://yoursite.com/success",
  expiredAt: Math.floor(Date.now() / 1000) + 1800 // 30 min
};
```

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "address": "0x9aBDDCE1EE18D1857C0653EB4a3Fa9d9E0dcd584",
    "amount": "10.008552",
    "memo": "",
    "checkoutUrl": "https://i.ccpayment.com/xxx",
    "confirmsNeeded": 50
  }
}
```

| Response Field | Type | Description |
| --- | --- | --- |
| data.address | String | Deposit address |
| data.amount | String | Crypto amount customer needs to pay |
| data.memo | String | Memo for memo-required chains |
| data.checkoutUrl | String | Hosted checkout page URL (only if `generateCheckoutURL: true`) |
| data.confirmsNeeded | Integer | Block confirmations required |

**Webhook:** `ApiDeposit` — then call Get Order Information to confirm.

---

### 3. Create Invoice Deposit (Customer-selected Currency)

**Use case:** Create a checkout page where the customer selects which cryptocurrency to pay with from your configured token list.

**POST** `https://ccpayment.com/v2/createInvoiceOrder`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| orderId | String | Yes | Your order ID, 3-64 characters |
| price | String | Yes | Product price (fiat or crypto depending on which ID is passed) |
| priceFiatId | String | No | Fiat ID for the denominating currency (get from Get Fiat List) |
| priceCoinId | String | No | Coin ID for the denominating currency (get from Get Token List) |
| product | String | No | Product name on checkout page (max 120 chars) |
| returnUrl | String | No | Redirect URL after payment |
| closeUrl | String | No | Redirect URL when buyer clicks "leave page" |
| expiredAt | Integer | No | 10-digit expiration timestamp. Default 24 hours |
| buyerEmail | String | No | Customer email for notifications |

**Rules:**
- Pass **either** `priceFiatId` **or** `priceCoinId`, **not both**
- Customers can pay with any coin from your "Tokens for your business" list

**Request Example:**

```javascript
const body = {
  orderId: "invoice_001",
  price: "50",
  priceFiatId: "1033",  // USD
  product: "Annual Subscription",
  returnUrl: "https://yoursite.com/thank-you"
};
```

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "invoiceUrl": "https://i.ccpayment.com/xxx"
  }
}
```

| Response Field | Type | Description |
| --- | --- | --- |
| data.invoiceUrl | String | Redirect customer to this URL to pay |

**Webhook:** `ApiDeposit` — then call Get Invoice Order Information to confirm.

---

### 4. Get Order Information (Merchant-specified)

**Use case:** After receiving an `ApiDeposit` webhook for a merchant-specified order, confirm order status and all payment details.

**POST** `https://ccpayment.com/v2/getDepositOrder`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| orderId | String | Yes | Your order ID, 3-64 characters |

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "amountToPay": "0.008552856654122477",
    "coinId": 1329,
    "coinSymbol": "MATIC",
    "chain": "POLYGON",
    "toAddress": "0x9abddce1ee18d1857c0653eb4a3fa9d9e0dcd584",
    "createAt": 1710233933,
    "rate": "1.1692",
    "fiatId": 1033,
    "fiatSymbol": "USD",
    "expiredAt": 1710243931,
    "checkoutUrl": "https://i.ccpayment.com/xxx",
    "paidList": [
      {
        "recordId": "20240312090316119190942031876096",
        "fromAddress": "0x12438f04093ebc87f0ba629bbe93f2451711d967",
        "amount": "0.001",
        "serviceFee": "0.0000005",
        "txid": "0xef4abf7175...",
        "status": "Success",
        "isFlaggedAsRisky": false,
        "arrivedAt": 1710234197,
        "rate": "1.1692"
      }
    ]
  }
}
```

| Response Field | Type | Description |
| --- | --- | --- |
| data.amountToPay | String | Total crypto amount to pay |
| data.coinId / coinSymbol / chain | | Order payment coin info |
| data.toAddress | String | Deposit address |
| data.toMemo | String | Memo for memo-required chains |
| data.rate | String | Locked rate (fiat/crypto). Rate changes if paid after expiration |
| data.fiatId / fiatSymbol | | Fiat info (if fiat-priced order) |
| data.createAt | Integer | Order creation timestamp |
| data.expiredAt | Integer | Order expiration timestamp |
| data.checkoutUrl | String | Checkout URL if created |
| data.paidList | Array | All payments for this order |
| data.paidList[].recordId | String | CCPayment unique transaction ID |
| data.paidList[].fromAddress | String | Sender address (empty for UTXO) |
| data.paidList[].amount | String | Received amount |
| data.paidList[].serviceFee | String | Service fee |
| data.paidList[].txid | String | On-chain transaction hash |
| data.paidList[].status | String | `Success`, `Processing`, `Failed` |
| data.paidList[].isFlaggedAsRisky | Boolean | Risky transaction flag |
| data.paidList[].arrivedAt | Integer | Deposit arrival timestamp |
| data.paidList[].rate | String | Rate at time of receipt |

**Best Practice:** Compare total paid amount vs `amountToPay` to determine: fully paid, partially paid, overpaid, or overdue.

---

### 5. Get Invoice Order Information (Customer-selected)

**Use case:** Confirm payment details for invoice/checkout page orders.

**POST** `https://ccpayment.com/v2/getInvoiceOrder`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| orderId | String | Yes | Your order ID, 3-64 characters |

**Success Response (abbreviated):**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "orderId": "invoice_001",
    "createAt": 1727074590,
    "product": "Premium Plan",
    "price": "50",
    "priceFiatId": 1033,
    "priceSymbol": "USD",
    "invoiceUrl": "https://i.ccpayment.com/xxx",
    "expiredAt": 1727160989,
    "totalPaidValue": "50.022",
    "paidList": [
      {
        "recordId": "20240923065734...",
        "coinId": 1280,
        "coinSymbol": "USDT",
        "chain": "BSC",
        "fromAddress": "0x4766dc...",
        "toAddress": "0xc53be3...",
        "paidAmount": "50",
        "serviceFee": "0.25",
        "rate": "1",
        "paidValue": "50",
        "txid": "0xf227c25...",
        "status": "Success",
        "isFlaggedAsRisky": false,
        "arrivedAt": 1727074654
      }
    ]
  }
}
```

| Key Response Fields | Description |
| --- | --- |
| data.totalPaidValue | Total paid in denominating currency |
| data.paidList[].paidAmount | Crypto amount paid |
| data.paidList[].paidValue | Paid amount in denominating currency (paidAmount × rate) |
| data.paidList[].rate | Rate: payment coin USD value / denominating currency USD value |

---

### 6. Get Deposit Record

**Use case:** Get full details of a specific deposit transaction by its record ID. Essential for confirming transaction status after webhook.

**POST** `https://ccpayment.com/v2/getDepositRecord`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordId | String | Yes | CCPayment unique transaction ID (from webhook) |

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "record": {
      "recordId": "20250116073333231508600365121536",
      "orderId": "1737011983",
      "coinId": 1482,
      "coinSymbol": "TRX",
      "chain": "TRX",
      "contract": "TRX",
      "coinUSDPrice": "0.23717",
      "fromAddress": "TRPKg7eGMy9aZS2RumSPeWoyVkDqTVwLgL",
      "toAddress": "TAkmn3f8bgwMAwdVrGfUVSRg4W2XwqgHGc",
      "toMemo": "",
      "amount": "0.5",
      "serviceFee": "0.0025",
      "txId": "f39abf3275607fe2...",
      "status": "Success",
      "arrivedAt": 1737012813,
      "isFlaggedAsRisky": false,
      "referenceId": "user_12345",
      "orderId": "1737011983"
    }
  }
}
```

| Response Field | Type | Description |
| --- | --- | --- |
| data.record.recordId | String | CCPayment unique transaction ID |
| data.record.coinId / coinSymbol | | Deposit coin info |
| data.record.chain | String | Chain symbol |
| data.record.contract | String | Token contract address |
| data.record.coinUSDPrice | String | Coin USD price at receipt time |
| data.record.fromAddress | String | Sender address (empty for UTXO-type) |
| data.record.toAddress | String | Receiving address |
| data.record.toMemo | String | Memo |
| data.record.amount | String | Received amount |
| data.record.serviceFee | String | Service fee charged |
| data.record.txId | String | On-chain transaction hash |
| data.record.status | String | `Success`, `Processing`, `Failed` |
| data.record.isFlaggedAsRisky | Boolean | Risk flag |
| data.record.arrivedAt | Integer | Arrival timestamp |
| data.record.referenceId | String | Your reference ID (if applicable) |
| data.record.orderId | String | Your order ID (if applicable) |

---

### 7. Get Deposit Record List

**Use case:** Query multiple deposit records with filters.

**POST** `https://ccpayment.com/v2/getDepositRecordList`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | No | Filter by coin ID |
| referenceId | String | No | Filter by reference ID (3-64 chars) |
| orderId | String | No | Filter by order ID. Don't pass both orderId and referenceId |
| chain | String | No | Filter by chain symbol |
| startAt | Integer | No | 10-digit start timestamp |
| endAt | Integer | No | 10-digit end timestamp |
| nextId | String | No | Pagination cursor |

**Pagination:** Results return max 20 records. If more exist, `data.nextId` is returned — pass it in the next request with the same filters.

**Success Response:** Same fields as Get Deposit Record but returns `data.records[]` array and `data.nextId`.

---

### 8. Address Unbinding

**Use case:** Unbind a deposit address from a user/referenceId, typically when the address is flagged as risky.

**POST** `https://ccpayment.com/v2/addressUnbinding`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| chain | String | Yes | Chain symbol of the address |
| address | String | Yes | Address to unbind |

**Rules:**
- If the address is on an EVM-compatible chain (ETH, BSC, Polygon), corresponding addresses on other EVM chains are also auto-unbound
- After unbinding, payments to the old address are still credited to your merchant account
- Call Get Permanent Deposit Address again to get a new address for the same user

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "unbound": [
      { "chain": "POLYGON", "address": "0xa9a363..." }
    ],
    "unboundAt": 1741783734,
    "userID": "",
    "referenceID": "user_12345"
  }
}
```

---

## Withdrawal APIs

### 1. Create Network Withdrawal

**Use case:** Send crypto from your merchant balance to any blockchain address.

**POST** `https://ccpayment.com/v2/applyWithdrawToNetwork`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | Yes | Coin ID |
| chain | String | Yes | Chain symbol |
| address | String | Yes | Destination blockchain address |
| memo | String | No | Memo for memo-required chains (XRP, XLM, TON) |
| orderId | String | Yes | Your withdrawal order ID, 3-64 characters |
| amount | String | Yes | Withdrawal amount |
| merchantPayNetworkFee | Boolean | No | `true`: merchant pays fee (recipient gets full amount). `false` (default): deducted from amount |

**Rules:**
- Check balance first with Get Balance List
- Validate address with Check Withdrawal Address Validity
- Check network fee with Get Withdrawal Network Fee
- Withdrawal may take minutes; if no webhook after 2 hours, call Get Withdrawal Record

**Request Example:**

```javascript
const body = {
  coinId: 1280,
  chain: "POLYGON",
  address: "0xBb9C4e7F3687aca1AE2828c18f9E3ae067F569FE",
  orderId: "wd_" + Date.now(),
  amount: "10.5",
  merchantPayNetworkFee: false
};
```

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "recordId": "202403120909501767478092588126208"
  }
}
```

**Webhook:** `ApiWithdrawal` with statuses: `Processing` → `Success` (or `Failed`). If approval enabled: `WaitingApproval` → `Processing` → `Success`.

**Important:** In rare cases of network jitter, you may not receive a response. Call Get Withdrawal Record to confirm status.

---

### 2. Withdrawal to Cwallet Account

**Use case:** Instant transfer to a Cwallet user. Zero network fee.

**POST** `https://ccpayment.com/v2/applyWithdrawToCwallet`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | Yes | Coin ID |
| cwalletUser | String | Yes | Cwallet ID or email address |
| amount | String | Yes | Amount (min 0.001 USD equivalent) |
| orderId | String | Yes | Your order ID, 3-64 characters |

**Rules:**
- No webhook for Cwallet transfers — call Get Withdrawal Record to confirm
- Validate Cwallet user with Get Cwallet User Information first

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "recordId": "202403120913091767478929213362176"
  }
}
```

---

### 3. Get Withdrawal Record

**Use case:** Confirm withdrawal status and details.

**POST** `https://ccpayment.com/v2/getWithdrawRecord`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordId | String | No | CCPayment transaction ID |
| orderId | String | No | Your order ID |

**Rules:** Pass either `recordId` or `orderId`, not both empty.

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "record": {
      "recordId": "202403120909501767478092588126208",
      "withdrawType": "Network",
      "coinId": 1891,
      "coinSymbol": "TETH",
      "chain": "POLYGON",
      "fromAddress": "0x1AE2828c...",
      "toAddress": "0xBb9C4e7F...",
      "orderId": "1710234589577",
      "txId": "0xb55bb282...",
      "toMemo": "",
      "status": "Success",
      "amount": "0.004",
      "fee": {
        "coinId": 1891,
        "coinSymbol": "TETH",
        "amount": "0.001"
      }
    }
  }
}
```

| Response Field | Type | Description |
| --- | --- | --- |
| data.record.withdrawType | String | `Network` or `Cwallet` |
| data.record.cwalletUser | String | Cwallet user ID (only for Cwallet withdrawals) |
| data.record.txId | String | On-chain hash (only for network withdrawals) |
| data.record.fee | Object | Network fee details (only for network withdrawals) |
| data.record.status | String | `Processing`, `Success`, `Failed`, `WaitingApproval`, `Rejected` |

---

### 4. Get Withdrawal Record List

**POST** `https://ccpayment.com/v2/getWithdrawRecordList`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | No | Filter by coin |
| orderIds | Array[String] | No | Filter by order IDs (max 20) |
| chain | String | No | Filter by chain |
| startAt | Integer | No | 10-digit start timestamp |
| endAt | Integer | No | 10-digit end timestamp |
| nextId | String | No | Pagination cursor |

**Pagination:** Max 20 records per page. Use `nextId` for next page.

---

## Swap APIs

### 1. Get Swap Quote

**Use case:** Preview swap output before executing.

**POST** `https://ccpayment.com/v2/getSwapQuote`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinIdIn | Integer | Yes | Input coin ID |
| amountIn | String | Yes | Input amount |
| coinIdOut | Integer | Yes | Output coin ID |

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "coinIdIn": 1280,
    "coinIdOut": 1329,
    "amountOut": "194.09158706508187",
    "amountIn": "100",
    "swapRate": "1.9409158706508187",
    "feeRate": "0.005",
    "fee": "0.97045793532540935",
    "netAmountOut": "193.12112912975646065"
  }
}
```

| Response Field | Description |
| --- | --- |
| amountOut | Gross output amount |
| netAmountOut | Net output (amountOut - fee) |
| swapRate | amountIn / amountOut |
| feeRate | CCPayment fee rate |
| fee | Fee amount. Minimum 0.1 USDT equivalent |

---

### 2. Create and Fulfill Swap Order

**Use case:** Execute a swap between two cryptocurrencies from merchant balance.

**POST** `https://ccpayment.com/v2/createSwapOrder`

**Parameters:**

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| orderId | String | Yes | Your swap order ID, 3-64 chars |
| coinIdIn | Integer | Yes | Input coin ID |
| amountIn | String | Yes | Input amount |
| coinIdOut | Integer | Yes | Output coin ID |
| amountOutMinimum | String | No | Slippage protection — swap fails if output < this |

**Success Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "recordId": "20240719085639165937311982694400",
    "orderId": "swap_001",
    "coinIdIn": 1280,
    "coinIdOut": 1329,
    "amountOut": "1.9428459962937677",
    "amountIn": "1",
    "swapRate": "1.9428459962937677",
    "fee": "0.1950648590656393",
    "feeRate": "0.1004016064257028",
    "netAmountOut": "1.7477811372281284"
  }
}
```

---

### 3. Get Swap Record

**POST** `https://ccpayment.com/v2/getSwapRecord`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordId | String | No | CCPayment transaction ID |
| orderId | String | No | Your order ID |

**Rules:** Pass either `recordId` or `orderId`, not both.

---

### 4. Get Swap Record List

**POST** `https://ccpayment.com/v2/getSwapRecordList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordIds | Array[String] | No | Filter by record IDs |
| orderIds | Array[String] | No | Filter by order IDs |
| coinIdIn | Integer | No | Filter by input coin |
| coinIdOut | Integer | No | Filter by output coin |
| startAt | Integer | No | 10-digit start timestamp |
| endAt | Integer | No | 10-digit end timestamp (max 3 months from startAt) |
| nextId | String | No | Pagination cursor |

---

## Wallet System APIs

Build a wallet system where each user has an independent account with deposits, withdrawals, swaps, and internal transfers.

**User ID Rules:**
- String, 5-64 characters
- Cannot start with `"sys"`
- The merchant's own App ID is reserved as the merchant account user ID — do not assign it to users

---

### User Balance

#### Get User Balance List

**POST** `https://ccpayment.com/v2/getUserBalanceList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| userId | String | Yes | User ID (5-64 chars, not starting with "sys") |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "userId": "user_12345",
    "assets": [
      { "coinId": 1280, "coinSymbol": "USDT", "available": "100.5" },
      { "coinId": 1482, "coinSymbol": "TRX", "available": "250.0" }
    ]
  }
}
```

#### Get Coin Balance of User

**POST** `https://ccpayment.com/v2/getUserCoinBalance`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| coinId | Integer | Yes | Coin ID |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "userId": "user_12345",
    "asset": { "coinId": 1280, "coinSymbol": "USDT", "available": "100.5" }
  }
}
```

---

### User Deposit

#### Create or Get User Deposit Address

**POST** `https://ccpayment.com/v2/getOrCreateUserDepositAddress`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| userId | String | Yes | User ID (5-64 chars) |
| chain | String | Yes | Chain symbol |

**Rules:** Max 1000 user addresses per App ID.

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": { "address": "0x...", "memo": "" }
}
```

**Webhook:** `UserDeposit` — then call Get User Deposit Record to confirm.

#### Get User Deposit Record

**POST** `https://ccpayment.com/v2/getUserDepositRecord`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordId | String | Yes | CCPayment transaction ID |

**Response:** Same structure as merchant Get Deposit Record, plus `data.record.userId`.

#### Get User Deposit Record List

**POST** `https://ccpayment.com/v2/getUserDepositRecordList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| coinId | Integer | No | Filter by coin |
| chain | String | No | Filter by chain |
| startAt | Integer | No | 10-digit timestamp |
| endAt | Integer | No | 10-digit timestamp |
| nextId | String | No | Pagination cursor |

---

### User Withdrawal

#### Withdrawal to Blockchain Address

**POST** `https://ccpayment.com/v2/applyUserWithdrawToNetwork`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| coinId | Integer | Yes | Coin ID |
| chain | String | Yes | Chain symbol |
| address | String | Yes | Destination address |
| memo | String | No | Memo for memo-required chains |
| orderId | String | Yes | Your order ID, 3-64 chars |
| amount | String | Yes | Withdrawal amount |

**Webhook:** `UserWithdrawal` — then call Get User Withdrawal Record.

**Tip:** If no `Success` webhook after 2 hours, proactively call Get User Withdrawal Record.

#### User Withdrawal to Cwallet

**POST** `https://ccpayment.com/v2/applyUserWithdrawToCwallet`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| coinId | Integer | Yes | Coin ID |
| cwalletUser | String | Yes | Cwallet ID or email |
| amount | String | Yes | Amount (min 0.001 USD) |
| orderId | String | Yes | Your order ID, 3-64 chars |

**No webhook** — call Get User Withdrawal Record to confirm.

#### Get User Withdrawal Record

**POST** `https://ccpayment.com/v2/getUserWithdrawRecord`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordId | String | No | CCPayment transaction ID |
| orderId | String | No | Your order ID |

**Rules:** Pass either `recordId` or `orderId`, not both empty.

#### Get User Withdrawal Record List

**POST** `https://ccpayment.com/v2/getUserWithdrawRecordList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| coinId | Integer | No | Filter by coin |
| toAddress | String | No | Filter by destination address |
| chain | String | No | Filter by chain |
| startAt / endAt | Integer | No | 10-digit timestamps |
| nextId | String | No | Pagination cursor |

---

### User Internal Transaction

**Use case:** Transfer funds between users within your merchant wallet system (e.g., tipping, marketplace payouts). Also used to sweep user funds to the merchant account.

#### Create Internal Transaction

**POST** `https://ccpayment.com/v2/createUserInternalTransaction`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| fromUserId | String | Yes | Source user ID (5-64 chars) |
| toUserId | String | Yes | Destination user ID. Use your App ID to transfer to merchant account |
| coinId | Integer | Yes | Coin ID |
| amount | String | Yes | Amount (min 0.001 USD) |
| orderId | String | Yes | Your order ID, 3-64 chars |
| remark | String | No | Transaction note (max 255 chars) |

**Rules:**
- **No webhook** for internal transactions
- If response `code` is `10000`, the transfer is immediately successful
- Instant and zero-fee

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": { "recordId": "202403220335..." }
}
```

#### Get User Internal Transaction Record

**POST** `https://ccpayment.com/v2/getUserInternalTransactionRecord`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| recordId | String | No | CCPayment transaction ID |
| orderId | String | No | Your order ID |

**Rules:** Pass either, not both empty.

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "record": {
      "recordId": "20240322033752...",
      "coinId": 1329,
      "coinSymbol": "MATIC",
      "orderId": "internal_001",
      "fromUserId": "user_12345",
      "toUserId": "user_67890",
      "amount": "0.002",
      "status": "Success",
      "remark": "tip"
    }
  }
}
```

#### Get User Internal Transaction Record List

**POST** `https://ccpayment.com/v2/getUserInternalTransactionRecordList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| fromUserId | String | Yes | Source user ID |
| toUserId | String | Yes | Destination user ID |
| coinId | Integer | No | Filter by coin |
| startAt / endAt | Integer | No | 10-digit timestamps |
| nextId | String | No | Pagination cursor |

---

### User Swap

#### User Get Swap Quote

**POST** `https://ccpayment.com/v2/getUserSwapQuote`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinIdIn | Integer | Yes | Input coin ID |
| amountIn | String | Yes | Input amount |
| coinIdOut | Integer | Yes | Output coin ID |
| extraFeeRate | String | No | Your markup fee rate (e.g., `"0.01"` = 1%). Charged from user's output |

**Response includes:** `amountOut`, `netAmountOut`, `fee`, `extraFee`, `swapRate`

`netAmountOut = amountOut - fee - extraFee`

#### Create User Swap Order

**POST** `https://ccpayment.com/v2/createUserSwapOrder`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| orderId | String | Yes | Your order ID, 3-64 chars |
| userId | String | Yes | User ID |
| coinIdIn | Integer | Yes | Input coin ID |
| amountIn | String | Yes | Input amount |
| coinIdOut | Integer | Yes | Output coin ID |
| extraFeeRate | String | No | Your markup fee rate |
| amountOutMinimum | String | No | Slippage protection |

#### Get User Swap Record / List

Same structure as merchant swap record APIs, with additional `userId`, `extraFee`, `extraFeeRate` fields.

**POST** `https://ccpayment.com/v2/getUserSwapRecord` (pass `recordId` or `orderId`)

**POST** `https://ccpayment.com/v2/getUserSwapRecordList` (supports `userId`, `recordIds`, `orderIds`, `coinIdIn`, `coinIdOut`, `startAt`, `endAt`, `nextId`)

---

## Common APIs

### Get Token List

**Use case:** Retrieve all tokens enabled in your merchant "Tokens for your business" config with full network details.

**POST** `https://ccpayment.com/v2/getTokenList`

**Parameters:** None required.

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "coins": [
      {
        "coinId": 1207,
        "symbol": "LINK",
        "coinFullName": "ChainLink Token",
        "logoUrl": "https://resource.ccpayment.com/token/icon/link.png",
        "status": "Normal",
        "networks": {
          "BSC": {
            "chain": "BSC",
            "chainFullName": "Binance Smart Chain",
            "contract": "0xf8a0bf9cf54bb92f17374d9e9a321e6a111a51bd",
            "precision": 18,
            "canDeposit": true,
            "canWithdraw": true,
            "minimumDepositAmount": "0",
            "minimumWithdrawAmount": "0.025",
            "maximumWithdrawAmount": "0",
            "isSupportMemo": false
          }
        }
      }
    ]
  }
}
```

| Field | Description |
| --- | --- |
| coinId | Use this ID in deposit/withdrawal/swap APIs |
| status | `Normal` (active), `Maintain` (withdrawal paused), `Pre-delisting` (deposit paused), `Delisted` (all paused) |
| networks[chain].chain | Use this in `chain` parameter of APIs |
| canDeposit / canWithdraw | Check before creating deposits/withdrawals |
| minimumDepositAmount / minimumWithdrawAmount | Enforce in your UI |
| maximumWithdrawAmount | 0 means no maximum |
| isSupportMemo | If true, memo is required for this chain |

---

### Get Token Information

**Use case:** Get detailed info for a specific token.

**POST** `https://ccpayment.com/v2/getTokenInfo`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | Yes | Coin ID |

**Response:** Same structure as a single coin in Get Token List.

---

### Get Token Price

**Use case:** Get current USD-equivalent prices for tokens.

**POST** `https://ccpayment.com/v2/getTokenPrice`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinIds | Array[Integer] | Yes | Array of coin IDs |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "prices": {
      "1329": "1.1683",
      "1280": "1.0001"
    }
  }
}
```

Prices are in USDT equivalent.

---

### Get Balance List (Merchant)

**Use case:** Get all balances in your merchant account.

**POST** `https://ccpayment.com/v2/getBalanceList`

**Parameters:** None required.

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "assets": [
      { "coinId": 1280, "coinSymbol": "USDT", "available": "1523.45" }
    ]
  }
}
```

### Get Coin Balance (Merchant)

**POST** `https://ccpayment.com/v2/getCoinBalance`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | Yes | Coin ID |

---

### Rescan Lost Transaction

**Use case:** User confirms on-chain success but funds not credited. Trigger a manual rescan.

**POST** `https://ccpayment.com/v2/rescanLostTransaction`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| chain | String | Yes | Chain symbol |
| toAddress | String | Yes | Receiving deposit address |
| memo | String | No | Memo (for XRP, XLM, TON) |
| txId | String | Yes | On-chain transaction hash |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "description": "The deposit has been rescanned. Please wait for the transaction to be credited."
  }
}
```

---

### Get Cwallet User Information

**Use case:** Validate a Cwallet user before sending a Cwallet withdrawal.

**POST** `https://ccpayment.com/v2/getCwalletUserInfo`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| cwalletUserId | String | Yes | Cwallet user ID |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "cwalletUserId": "35255142",
    "cwalletUserName": "j***@proton.me"
  }
}
```

---

### Check Withdrawal Address Validity

**Use case:** Validate an address before creating a withdrawal.

**POST** `https://ccpayment.com/v2/checkWithdrawAddressValidity`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| chain | String | Yes | Chain symbol |
| address | String | Yes | Address to validate |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": { "addrIsValid": true }
}
```

---

### Get Withdrawal Network Fee

**Use case:** Get real-time network fee before creating a withdrawal.

**POST** `https://ccpayment.com/v2/getWithdrawNetworkFee`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | Yes | Coin ID |
| chain | String | Yes | Chain symbol |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "fee": {
      "coinId": 1329,
      "coinSymbol": "MATIC",
      "amount": "0.0213814"
    }
  }
}
```

---

### Get Fiat List

**Use case:** Get supported fiat currencies for pricing orders.

**POST** `https://ccpayment.com/v2/getFiatList`

**Parameters:** None.

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "fiats": [
      {
        "fiatId": 1033,
        "symbol": "USD",
        "logoUrl": "https://resource.ccpayment.com/fiat/USD.png",
        "mark": "US",
        "usdRate": "1"
      }
    ]
  }
}
```

Use `fiatId` in deposit order APIs.

---

### Get Swap Coin List

**Use case:** Get coins available for swap.

**POST** `https://ccpayment.com/v2/getSwapCoinList`

**Parameters:** None.

**Note:** Only returns coins from your "Tokens for your business" list that are swappable.

---

### Get Chain List

**Use case:** Get all supported chains with status and configuration details.

**POST** `https://ccpayment.com/v2/getChainList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| chains | Array[String] | No | Filter by specific chains (e.g., `["ETH", "POLYGON"]`) |

**Response:**

```json
{
  "code": 10000,
  "msg": "success",
  "data": {
    "chains": [
      {
        "chain": "ETH",
        "chainFullName": "Ethereum",
        "explorer": "https://etherscan.io/",
        "baseUrl": "https://etherscan.io/tx/%s",
        "confirmations": 2,
        "isEVM": true,
        "supportMemo": false,
        "logoUrl": "https://resource.ccpayment.com/token/icon/ETH.png",
        "status": "Normal"
      }
    ]
  }
}
```

| Field | Description |
| --- | --- |
| baseUrl | Append txId to build explorer link: `baseUrl.replace('%s', txId)` |
| confirmations | Block confirmations needed |
| isEVM | EVM-compatible chain |
| supportMemo | Whether memo/tag is required |
| status | `Normal` or `Maintenance` |

---

### Get Auto-withdrawal Record List

**POST** `https://ccpayment.com/v2/getAutoWithdrawRecordList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | No | Filter by coin |
| recordIds | Array[String] | No | Filter by record IDs |
| chain | String | No | Filter by chain |
| startAt / endAt | Integer | No | 10-digit timestamps |
| nextId | String | No | Pagination cursor |

---

### Get Risk Refund Record List

**POST** `https://ccpayment.com/v2/getRiskyRefundRecordList`

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| coinId | Integer | No | Filter by coin |
| chain | String | No | Filter by chain |
| toAddress | String | No | Filter by refund address |
| startAt / endAt | Integer | No | 10-digit timestamps |
| nextId | String | No | Pagination cursor |

**Note:** When a risk refund succeeds, the system deducts the risk refund fee and the corresponding service fee from the merchant's account.

---

## Proto-Verified HTTP Schema Reference

This supplement is verified against the current HTTP interface definitions in `src/merchant-common/pb/gateway/api/interface/v2/*.proto`.

Use this section as the authoritative request/response schema reference when integrating against the current codebase. The earlier examples in this file remain useful for workflow understanding, but several endpoint names and field lists are simplified there.

### Conventions

- All paths below are the current HTTP paths exposed by `interface.proto`
- All requests are `POST`
- All responses are wrapped as:

```json
{
  "code": 10000,
  "msg": "success",
  "data": { "...endpoint specific payload..." }
}
```

- `recordId` is CCPayment's transaction or record ID
- `orderId` is the merchant-defined business order ID
- `nextId` is the cursor for pagination
- Fields marked with `?` are optional in the proto definition

### Shared Response Entities

#### `WithdrawFee`

| Field | Type | Description |
| --- | --- | --- |
| coinId | UInt64 | Fee coin ID |
| coinSymbol | String | Fee coin symbol |
| amount | String | Fee amount |

#### `AssetEntity`

| Field | Type | Description |
| --- | --- | --- |
| coinId | UInt64 | Coin ID |
| coinSymbol | String | Coin symbol |
| available? | String | Available balance |

#### `AppDepositRecordEntity`

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Deposit record ID |
| referenceId | String | Merchant-side permanent address binding ID |
| orderId | String | Merchant order ID for order-mode deposits |
| coinId | UInt64 | Coin ID |
| coinSymbol | String | Coin symbol |
| chain | String | Chain symbol |
| contract | String | Token contract or native token marker |
| coinUSDPrice | String | Coin USD price when credited |
| fromAddress | String | Sender address |
| toAddress | String | Receiving address |
| toMemo? | String | Receiving memo/tag |
| amount | String | Gross credited amount |
| serviceFee | String | Service fee charged |
| txId | String | On-chain transaction hash |
| txIndex | UInt64 | Transaction index on chain, when applicable |
| status | String | Deposit status |
| arrivedAt | Int64 | Credit timestamp |
| isFlaggedAsRisky? | Bool | Whether the deposit is marked risky |

#### `AppWithdrawRecordEntity`

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Withdrawal record ID |
| withdrawType | String | `Network` or `Cwallet` |
| appId | String | Merchant app ID |
| coinId | UInt64 | Coin ID |
| coinSymbol | String | Coin symbol |
| chain | String | Chain symbol |
| fromAddress | String | Source address |
| toAddress | String | Destination address |
| cwalletUser | String | Cwallet user identifier for Cwallet withdrawals |
| orderId | String | Merchant order ID |
| txId? | String | On-chain hash |
| toMemo? | String | Destination memo/tag |
| status | String | Withdrawal status |
| amount | String | Withdrawal amount |
| fee | Object | See `WithdrawFee` |
| reason? | String | Failure or rejection reason |
| coinUSDPrice | String | Coin USD price at processing time |

#### `UserDepositRecord`

| Field | Type | Description |
| --- | --- | --- |
| userId | String | Wallet-system user ID |
| recordId | String | Deposit record ID |
| coinId | UInt64 | Coin ID |
| chain | String | Chain symbol |
| contract | String | Token contract |
| coinSymbol | String | Coin symbol |
| txId | String | On-chain transaction hash |
| coinUSDPrice | String | Coin USD price when credited |
| fromAddress | String | Sender address |
| toAddress | String | Receiving address |
| toMemo? | String | Receiving memo/tag |
| amount | String | Gross credited amount |
| serviceFee | String | Service fee |
| status | String | Deposit status |
| arrivedAt | Int64 | Credit timestamp |
| isFlaggedAsRisky? | Bool | Risk flag |

#### `UserWithdrawRecord`

| Field | Type | Description |
| --- | --- | --- |
| userId | String | Wallet-system user ID |
| recordId | String | Withdrawal record ID |
| withdrawType | String | `Network` or `Cwallet` |
| coinId | UInt64 | Coin ID |
| chain | String | Chain symbol |
| orderId | String | Merchant order ID |
| txId? | String | On-chain hash |
| coinSymbol | String | Coin symbol |
| fromAddress | String | Source address |
| toAddress | String | Destination address |
| cwalletUser | String | Cwallet user identifier |
| toMemo? | String | Destination memo/tag |
| amount | String | Withdrawal amount |
| status | String | Withdrawal status |
| fee | Object | See `WithdrawFee` |
| coinUSDPrice | String | Coin USD price at processing time |

#### `SwapRecordEntity`

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Swap record ID |
| orderId | String | Merchant order ID |
| coinInSymbol | String | Input coin symbol |
| coinIdIn | UInt64 | Input coin ID |
| coinOutSymbol | String | Output coin symbol |
| coinIdOut | UInt64 | Output coin ID |
| amountOut | String | Gross output amount |
| amountIn | String | Input amount |
| amountOutMinimum | String | Slippage-protection minimum output |
| netAmountOut | String | Net output after fees |
| userId | String | Wallet-system user ID for user swap |
| extraFee | String | Merchant markup fee from user swap |
| extraFeeRate | String | Merchant markup fee rate |
| swapRate | String | Swap rate |
| feeRate | String | CCPayment fee rate |
| fee | String | CCPayment fee amount |
| status | String | Swap status |
| createdAt | Int64 | Creation timestamp |
| arrivedAt | Int64 | Completion timestamp |

### Merchant Deposit Endpoints

#### `/getOrCreateAppDepositAddress`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| referenceId | String | Yes | 3-64 chars |
| notifyUrl | String | No | Optional webhook override URL, max 150 chars, must be URI |
| chain | String | Yes | Chain symbol |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| address | String | Permanent deposit address |
| memo? | String | Memo/tag for memo-required chains |

#### `/createAppOrderDepositAddress`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | 3-64 chars |
| coinId | UInt64 | Yes | Payment coin ID |
| fiatId | UInt64 | No | Pricing fiat ID |
| chain | String | Yes | Payment chain |
| price | String | Yes | Fiat price or crypto amount |
| expiredAt | Int64 | No | Expiration timestamp |
| buyerEmail | String | No | Buyer email |
| generateCheckoutURL | Bool | No | Whether to return checkout URL |
| product | String | No | Max 120 chars |
| returnUrl | String | No | URI, max 150 chars |
| notifyUrl | String | No | URI, max 150 chars |
| closeUrl | String | No | URI, max 150 chars |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| address | String | Deposit address |
| amount | String | Amount the payer should send |
| memo? | String | Memo/tag |
| checkoutUrl? | String | Hosted payment URL when enabled |
| confirmsNeeded | UInt64 | Required confirmation count |

#### `/createInvoiceUrl`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | 3-64 chars |
| product | String | No | 1-128 chars |
| priceFiatId | UInt64 | No | Denominated fiat ID |
| priceCoinId | UInt64 | No | Denominated coin ID |
| price | String | Yes | Order price |
| buyerEmail | String | No | Buyer email |
| returnUrl | String | No | URI, max 150 chars |
| expiredAt | Int64 | No | Expiration timestamp |
| closeUrl | String | No | URI, max 150 chars |
| notifyUrl | String | No | URI, max 150 chars |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| invoiceUrl | String | Hosted invoice URL |

#### `/getAppOrderInfo`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | Merchant order ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| amountToPay | String | Required payment amount |
| coinId | UInt64 | Payment coin ID |
| coinSymbol | String | Payment coin symbol |
| chain | String | Payment chain |
| toAddress | String | Deposit address |
| toMemo | String | Deposit memo/tag |
| createAt | Int64 | Created timestamp |
| rate | String | Locked exchange rate |
| fiatId | UInt64 | Pricing fiat ID |
| fiatSymbol | String | Pricing fiat symbol |
| expiredAt | Int64 | Expiration timestamp |
| checkoutUrl | String | Hosted checkout URL |
| buyerEmail | String | Buyer email |
| paidList | Array | Paid transaction list |

`paidList[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Deposit record ID |
| fromAddress | String | Sender address |
| amount | String | Paid amount |
| serviceFee | String | Service fee |
| txid | String | Transaction hash |
| status | String | Payment status |
| arrivedAt | Int64 | Credit timestamp |
| rate | String | Effective rate for this payment |
| isFlaggedAsRisky? | Bool | Risk flag |

#### `/getInvoiceOrderInfo`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | Merchant order ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| orderId | String | Merchant order ID |
| createAt | Int64 | Created timestamp |
| product | String | Product name |
| price | String | Denominated price |
| priceCoinId | UInt64 | Denominated coin ID when coin-priced |
| priceFiatId | UInt64 | Denominated fiat ID when fiat-priced |
| priceSymbol | String | Denomination symbol |
| invoiceUrl | String | Hosted invoice URL |
| buyerEmail | String | Buyer email |
| expiredAt | Int64 | Expiration timestamp |
| totalPaidValue | String | Total paid value in the denominated currency |
| paidList | Array | Successful or processed payment records |

`paidList[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Deposit record ID |
| coinId | UInt64 | Paid coin ID |
| coinSymbol | String | Paid coin symbol |
| chain | String | Paid chain |
| fromAddress | String | Sender address |
| toAddress | String | Receiving address |
| toMemo? | String | Receiving memo/tag |
| paidAmount | String | Paid coin amount |
| serviceFee | String | Service fee |
| txid | String | Transaction hash |
| status | String | Payment status |
| arrivedAt | Int64 | Credit timestamp |
| rate | String | Coin USD price divided by denominated asset USD price |
| paidValue | String | Paid value in the denominated currency |
| isFlaggedAsRisky? | Bool | Risk flag |

#### `/getAppDepositRecord`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| recordId | String | Yes | Deposit record ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| record | Object | See `AppDepositRecordEntity` |

#### `/getAppDepositRecordList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| chain | String | No | Chain filter |
| referenceId | String | No | Single reference filter |
| orderId | String | No | Single order filter |
| toAddress | String | No | Receiving address filter |
| coinId | UInt64 | No | Coin filter |
| startAt | Int64 | No | Start timestamp |
| endAt | Int64 | No | End timestamp |
| nextId | String | No | Pagination cursor |
| recordIds | Array[String] | No | Up to 50 record IDs |
| referenceIds | Array[String] | No | Up to 50 reference IDs |
| orderIds | Array[String] | No | Up to 50 order IDs |
| limit | UInt64 | No | Max 100 |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| records | Array | Array of `AppDepositRecordEntity` |
| nextId? | String | Cursor for the next page |

#### `/addressUnbinding`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| chain | String | Yes | Chain symbol |
| address | String | Yes | Bound address to unbind |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| unbound | Array | Unbound address list |
| unboundAt | Int64 | Unbind timestamp |
| userID? | String | User wallet ID when user address is unbound |
| referenceID? | String | Merchant reference ID when app address is unbound |

### Merchant Withdrawal Endpoints

#### `/applyAppWithdrawToNetwork`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | 3-64 chars |
| coinId | UInt64 | Yes | Coin ID |
| chain | String | Yes | Chain symbol |
| address | String | Yes | Destination address |
| memo | String | No | Destination memo/tag |
| amount | String | Yes | Withdrawal amount |
| merchantPayNetworkFee | Bool | No | Merchant covers network fee when `true` |
| networkFeeInquiryID? | String | No | Fee quote ID from `/getWithdrawFee` |
| notifyUrl | String | No | Optional webhook override URL |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Withdrawal record ID |

#### `/applyAppWithdrawToCwallet`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | 3-64 chars |
| coinId | UInt64 | Yes | Coin ID |
| cwalletUser | String | Yes | Cwallet email or user ID |
| amount | String | Yes | Withdrawal amount |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Withdrawal record ID |

#### `/getAppWithdrawRecord`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | No | Merchant order ID |
| recordId | String | No | Withdrawal record ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| record | Object | See `AppWithdrawRecordEntity` |

#### `/getAppWithdrawRecordList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| chain | String | No | Chain filter |
| coinId | UInt64 | No | Coin filter |
| orderIds | Array[String] | No | Order IDs |
| startAt | Int64 | No | Start timestamp |
| endAt | Int64 | No | End timestamp |
| toAddress | String | No | Destination address |
| nextId | String | No | Pagination cursor |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| records | Array | Array of `AppWithdrawRecordEntity` |
| nextId? | String | Cursor for the next page |

### Wallet System Endpoints

#### `/getUserCoinAssetList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| userId | String | Yes | 5-64 chars |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| userId | String | User ID |
| assets | Array | User asset list |

`assets[]` fields: `coinId`, `coinSymbol`, `available`

#### `/getUserCoinAsset`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| userId | String | Yes | 5-64 chars |
| coinId | UInt64 | Yes | Coin ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| userId | String | User ID |
| asset | Object | Single user asset |

#### `/getOrCreateUserDepositAddress`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| userId | String | Yes | 5-64 chars, should not start with `sys` |
| chain | String | Yes | Chain symbol |
| notifyUrl | String | No | Optional webhook override URL |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| address | String | User deposit address |
| memo? | String | Memo/tag |

#### `/getUserDepositRecord`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| recordId | String | Yes | Deposit record ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| record | Object | See `UserDepositRecord` |

#### `/getUserDepositRecordList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| chain | String | No | Chain filter |
| coinId | UInt64 | No | Coin filter |
| toAddress | String | No | Receiving address filter |
| startAt | Int64 | No | Start timestamp |
| endAt | Int64 | No | End timestamp |
| nextId | String | No | Pagination cursor |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| records | Array | Array of `UserDepositRecord` |
| nextId? | String | Cursor for the next page |

#### `/applyUserWithdrawToNetwork`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| orderId | String | Yes | 3-64 chars |
| coinId | UInt64 | Yes | Coin ID |
| chain | String | Yes | Chain symbol |
| address | String | Yes | Destination address |
| amount | String | Yes | Withdrawal amount |
| memo | String | No | Destination memo/tag |
| networkFeeInquiryID? | String | No | Fee quote ID from `/getWithdrawFee` |
| notifyUrl | String | No | Optional webhook override URL |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Withdrawal record ID |

#### `/applyUserWithdrawToCwallet`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| orderId | String | Yes | 3-64 chars |
| coinId | UInt64 | Yes | Coin ID |
| cwalletUser | String | Yes | Cwallet email or user ID |
| amount | String | Yes | Withdrawal amount |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Withdrawal record ID |

#### `/getUserWithdrawRecord`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| recordId | String | No | Withdrawal record ID |
| orderId | String | No | Merchant order ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| record | Object | See `UserWithdrawRecord` |

#### `/getUserWithdrawRecordList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| userId | String | Yes | User ID |
| orderIds | Array[String] | No | Order IDs |
| chain | String | No | Chain filter |
| coinId | UInt64 | No | Coin filter |
| startAt | Int64 | No | Start timestamp |
| endAt | Int64 | No | End timestamp |
| toAddress | String | No | Destination address |
| nextId | String | No | Pagination cursor |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| records | Array | Array of `UserWithdrawRecord` |
| nextId? | String | Cursor for the next page |

#### `/userTransfer`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | 3-64 chars |
| fromUserId | String | Yes | Source user ID |
| toUserId | String | Yes | Destination user ID |
| coinId | UInt64 | Yes | Coin ID |
| amount | String | Yes | Transfer amount |
| remark | String | No | Remark text |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Internal transfer record ID |

#### `/getUserTransferRecord`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| recordId | String | No | Transfer record ID |
| orderId | String | No | Merchant order ID |

Response `data.record` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Transfer record ID |
| coinId | UInt64 | Coin ID |
| coinSymbol | String | Coin symbol |
| orderId | String | Merchant order ID |
| fromUserId | String | Source user ID |
| toUserId | String | Destination user ID |
| amount | String | Transfer amount |
| status | String | Transfer status |
| remark? | String | Remark text |
| coinUSDPrice | String | Coin USD price |

#### `/getUserTransferRecordList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderIds | Array[String] | No | Order IDs |
| fromUserId | String | No | Source user filter |
| toUserId | String | No | Destination user filter |
| coinId | UInt64 | No | Coin filter |
| startAt | Int64 | No | Start timestamp |
| endAt | Int64 | No | End timestamp |
| nextId | String | No | Pagination cursor |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| records | Array | Array of user transfer records |
| nextId? | String | Cursor for the next page |

### Swap Endpoints

#### `/estimate`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | No | Ignored for quote-only use |
| userId | String | No | Present for user quote scenarios |
| coinIdIn | UInt64 | Yes | Input coin ID |
| amountIn | String | Yes | Input amount |
| coinIdOut | UInt64 | Yes | Output coin ID |
| amountOutMinimum | String | No | Slippage protection hint |
| extraFeeRate | String | No | Merchant markup rate for user quote |
| appId | String | No | Internal business field |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| coinIdIn | UInt64 | Input coin ID |
| coinIdOut | UInt64 | Output coin ID |
| amountOut | String | Gross output amount |
| amountIn | String | Input amount |
| swapRate | String | Swap rate |
| feeRate | String | CCPayment fee rate |
| fee | String | CCPayment fee |
| extraFee | String | Merchant markup fee for user quote |
| netAmountOut | String | Net output after fees |

#### `/swap`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | Merchant order ID |
| coinIdIn | UInt64 | Yes | Input coin ID |
| amountIn | String | Yes | Input amount |
| coinIdOut | UInt64 | Yes | Output coin ID |
| amountOutMinimum | String | No | Minimum acceptable output |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Swap record ID |
| orderId | String | Merchant order ID |
| coinIdIn | UInt64 | Input coin ID |
| coinIdOut | UInt64 | Output coin ID |
| amountOut | String | Gross output amount |
| amountIn | String | Input amount |
| amountOutMinimum | String | Minimum output |
| swapRate | String | Swap rate |
| fee | String | CCPayment fee |
| feeRate | String | CCPayment fee rate |
| extraFee | String | Merchant markup fee, usually empty for merchant swap |
| netAmountOut | String | Net output amount |
| to_coin_price | String | Output coin reference price |
| from_coin_price | String | Input coin reference price |

#### `/userSwap`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| orderId | String | Yes | Merchant order ID |
| userId | String | Yes | User ID |
| coinIdIn | UInt64 | Yes | Input coin ID |
| amountIn | String | Yes | Input amount |
| coinIdOut | UInt64 | Yes | Output coin ID |
| amountOutMinimum | String | No | Minimum acceptable output |
| extraFeeRate | String | No | Merchant markup fee rate |

Response `data` fields:

The response shape matches `/swap`, but `extraFee` is meaningful for user swaps.

#### `/getSwapRecord`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| recordId | String | No | Swap record ID |
| orderId | String | No | Merchant order ID |
| ty | Int32 | No | Internal type selector |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| record | Object | See `SwapRecordEntity` |

#### `/getUserSwapRecord`

Request fields are the same as `/getSwapRecord`.

Response `data.record` uses the same `SwapRecordEntity` shape, with `userId`, `extraFee`, and `extraFeeRate` populated for user swaps.

#### `/getSwapRecordList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| recordIds | Array[String] | No | Max 20 |
| orderIds | Array[String] | No | Max 20 |
| userId | String | No | Usually empty for merchant swap |
| coinIdIn | UInt64 | No | Input coin filter |
| coinIdOut | UInt64 | No | Output coin filter |
| startAt | Int64 | No | Start timestamp |
| endAt | Int64 | No | End timestamp |
| nextId | String | No | Pagination cursor |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| records | Array | Array of `SwapRecordEntity` |
| nextId? | String | Cursor for the next page |

#### `/getUserSwapRecordList`

Request and response fields are the same as `/getSwapRecordList`, but `userId` is commonly used and the returned records contain user-swap-specific fields.

### Common Endpoints

#### `/getCoinList`

Request fields: none.

Response `data.coins[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| coinId | UInt64 | Coin ID |
| symbol | String | Coin symbol |
| coinFullName | String | Full coin name |
| logoUrl | String | Coin logo URL |
| status | String | Coin status |
| networks | Map | Keyed by chain |
| price | String | Current reference price |

`networks[chain]` fields:

| Field | Type | Description |
| --- | --- | --- |
| chain | String | Chain symbol |
| chainFullName | String | Chain name |
| contract | String | Token contract |
| precision | UInt32 | Decimal precision |
| canDeposit? | Bool | Deposit availability |
| canWithdraw? | Bool | Withdrawal availability |
| minimumDepositAmount | String | Minimum deposit amount |
| minimumWithdrawAmount | String | Minimum withdrawal amount |
| maximumWithdrawAmount | String | Maximum withdrawal amount, `0` means unlimited |
| isSupportMemo? | Bool | Whether memo/tag is used |
| protocol | String | Network protocol label |

#### `/getCoin`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| coinId | UInt64 | Yes | Coin ID |

Response `data.coin` uses the same shape as `/getCoinList` items.

#### `/getCoinUSDTPrice`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| coinIds | Array[UInt64] | Yes | Coin IDs |

Response `data.prices` is a map keyed by `coinId`, with USDT price strings as values.

#### `/getAppCoinAssetList`

Request fields: none.

Response `data.assets[]` uses the `AssetEntity` shape.

#### `/getAppCoinAsset`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| coinId | UInt64 | Yes | Coin ID |

Response `data.asset` uses the `AssetEntity` shape.

#### `/getCwalletUserId`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| cwalletUserId | String | Yes | Cwallet email or user ID |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| cwalletUserId | String | Resolved Cwallet user ID |
| cwalletUserName | String | Masked Cwallet user name |

#### `/checkWithdrawalAddressValidity`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| chain | String | Yes | Chain symbol |
| address | String | Yes | Destination address |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| addrIsValid? | Bool | Whether the address is valid |

#### `/getWithdrawFee`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| coinId | UInt64 | Yes | Coin ID |
| chain | String | Yes | Chain symbol |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| fee | Object | See `WithdrawFee` |
| networkFeeInquiryID | String | Fee quote ID, can be reused in withdrawal requests |

#### `/getFiatList`

Request fields: none.

Response `data.fiats[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| fiatId | UInt64 | Fiat ID |
| symbol | String | Fiat symbol |
| logoUrl | String | Fiat logo URL |
| mark | String | Region mark |
| usdRate | String | USD conversion rate |

#### `/getSwapCoinList`

Request fields: none.

Response `data.coins[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| coinId | UInt64 | Swappable coin ID |
| symbol | String | Coin symbol |
| logoUrl | String | Coin logo URL |

#### `/getChainList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| chains | Array[String] | No | Specific chains to query |

Response `data.chains[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| chain | String | Chain symbol |
| chainFullName | String | Full chain name |
| explorer | String | Explorer home URL |
| logoUrl | String | Chain logo URL |
| status | String | Chain status |
| confirmations? | Int32 | Required confirmations |
| baseUrl | String | Explorer transaction URL template |
| isEVM? | Bool | Whether the chain is EVM-compatible |
| supportMemo? | Bool | Whether memo/tag is used |

#### `/rescanLostTransaction`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| chain | String | Yes | Chain symbol |
| toAddress | String | Yes | Receiving address |
| txId | String | Yes | Transaction hash |
| memo | String | No | Receiving memo/tag |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| description | String | Rescan result description |

#### `/webhook/resend`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| startTimestamp | Int64 | No | Start of resend window |
| endTimestamp | Int64 | No | End of resend window |
| webhookResult | String | No | Result filter |
| transactionType | String | No | Transaction type filter |
| recordIds | Array[String] | No | Explicit record IDs |

Response `data` fields:

| Field | Type | Description |
| --- | --- | --- |
| resendCount? | Int64 | Number of webhook resend jobs created |

#### `/getAutoWithdrawRecordList`

Request fields:

| Field | Type | Required | Notes |
| --- | --- | --- | --- |
| chain | String | No | Chain filter |
| coinId | UInt64 | No | Coin filter |
| recordIds | Array[String] | No | Record IDs |
| startAt | Int64 | No | Start timestamp |
| endAt | Int64 | No | End timestamp |
| toAddress | String | No | Destination address |
| nextId | String | No | Pagination cursor |

Response `data.records[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Withdrawal record ID |
| coinId | UInt64 | Coin ID |
| coinSymbol | String | Coin symbol |
| chain | String | Chain symbol |
| orderId | String | Merchant order ID |
| toAddress | String | Destination address |
| toMemo? | String | Destination memo/tag |
| amount | String | Withdrawal amount |
| txId? | String | On-chain hash |
| status | String | Withdrawal status |
| fee | Object | See `WithdrawFee` |
| serviceFee | String | Service fee |
| createdAt | Int64 | Creation timestamp |

The response also includes `nextId?`.

#### `/getRiskyRefundRecordList`

Request fields are identical to `/getAutoWithdrawRecordList`.

Response `data.records[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Refund record ID |
| coinId | UInt64 | Coin ID |
| coinSymbol | String | Coin symbol |
| chain | String | Chain symbol |
| orderId | String | Related merchant order ID |
| toAddress | String | Refund destination address |
| toMemo? | String | Destination memo/tag |
| amount | String | Refund amount |
| txId? | String | On-chain hash |
| status | String | Refund status |
| fee | Object | See `WithdrawFee` |
| nextId? | String | Cursor-like field present on each record in proto |
| createdAt | Int64 | Creation timestamp |
| fromDeposit | Array | Source deposit records used to build the refund |

`fromDeposit[]` fields:

| Field | Type | Description |
| --- | --- | --- |
| recordId | String | Source deposit record ID |

---

## Error Codes

### Status Codes

| Code | Meaning |
| --- | --- |
| 10000 | Success |
| 10001 | Failed — check msg for details |

### Common Errors

| Code | Tip | Description |
| --- | --- | --- |
| 11000 | InvalidArgument | Invalid parameter |
| 11001 | HeaderInvalidArgument | Invalid header parameter |
| 11002 | Internal | Server error — retry later or contact support |
| 11003 | NotFound | Data does not exist |
| 11004 | RateLimit | Rate limit reached — reduce request frequency |
| 11005 | VerifySignFailed | HMAC signature verification failed — check sign algorithm |
| 11006 | ReqExpired | Request timestamp too old — check server clock |
| 11007 | RepeatedSubmit | Duplicate request |
| 11008 | QueryDurationTooMax | Query time range too large |
| 11009 | ReqDailyLimit | Daily request limit exceeded |
| 11010 | QueryNumMax | Query result too large (max 100) |
| 11011 | OrderDuplicate | Order ID already exists |
| 11012 | ExpiredAtTooMax | Expiration exceeds 10-day maximum |
| 11013 | NoSupportVersion | Account restricted to V1 API only |
| 11014 | MaliciousReq | Malicious request — IP banned |
| 11015 | UserIdNotFound | User ID does not exist |

### Account Errors

| Code | Tip | Description |
| --- | --- | --- |
| 12000 | MerchantDisabled | Merchant account disabled |
| 12001 | MerchantNotFound | Merchant does not exist |
| 12002 | IpNotInWhitelist | IP not in whitelist — update at Dashboard > Developer |
| 12003 | MerchantApiDisabled | Complete website verification on dashboard |

### Token Errors

| Code | Tip | Description |
| --- | --- | --- |
| 13000 | InvalidCoin | Unsupported coin |
| 13001 | InvalidChain | Unsupported network for this token |
| 13002 | AbnormalCoinPrice | Abnormal coin price — retry later |
| 13003 | AbnormalCoinPriceNotSupportMode | Price error, only merchant-pay-fee mode available |
| 13004 | UnstableBlockchain | Blockchain unstable — withdrawal unavailable |

### Withdrawal Errors

| Code | Tip | Description |
| --- | --- | --- |
| 14000 | BalanceInsufficient | Not enough balance |
| 14001 | WithdrawFeeTooLow | Network fee too low — check real-time fee |
| 14002 | AddressNotActive | Invalid/inactive address |
| 14003 | AddressEmptyMemo | Memo required for this chain |
| 14004 | ChainStopWithdraw | Withdrawals suspended on this chain |
| 14005 | WithdrawValueLessThanLimit | Below minimum withdrawal amount |
| 14006 | WithdrawValueMoreThanLimit | Exceeds maximum withdrawal amount |
| 14007 | WithdrawAddrFormat | Incorrect address format |
| 14008 | WithdrawCannotSelf | Cannot withdraw to own CCPayment address |
| 14009 | CoinNoSupportMemo | This coin does not support memo |
| 14010 | NoSupportCoin | Coin not in merchant's supported tokens |
| 14011 | WithdrawFeeBalanceNotEnough | Insufficient native token for network fee |
| 14012 | NotSupportMerchantTransfer | Merchant transfer only between main/sub accounts |
| 14013 | CoinPrecisionLimit | Exceeded coin precision limit |
| 14014 | NotFinishCollection | Unpaid asset aggregation fees — check dashboard |
| 14015 | WithdrawToContractAddress | Cannot withdraw to contract addresses |
| 14016 | WithdrawInvalidAddressType | Wrong address type (only data accounts supported) |

### Deposit Errors

| Code | Tip | Description |
| --- | --- | --- |
| 15000 | GenerateAddressFailed | Failed to generate address — retry |
| 15001 | PaymentAddressNumTooMuch | Address limit reached (1000 per App ID) |
| 15003 | ChainStopDeposit | Deposits suspended on this chain |

---

## Integration Flows

### Flow 1: E-commerce Payment (Fixed Price)

```
1. Call Get Token List → find coinId and chain for accepted currencies
2. Call Create Deposit Order with coinId, chain, price, orderId
   → Get address + amount (+ optional checkoutUrl)
3. Display address/amount to customer (or redirect to checkoutUrl)
4. Receive ApiDeposit webhook
5. Call Get Order Information with orderId → confirm actual payment
6. Compare paid amount vs expected → fulfill order
```

### Flow 2: E-commerce Payment (Customer Chooses Crypto)

```
1. Call Get Fiat List → find fiatId for your currency
2. Call Create Invoice Deposit with orderId, price, priceFiatId
   → Get invoiceUrl
3. Redirect customer to invoiceUrl
4. Receive ApiDeposit webhook
5. Call Get Invoice Order Information with orderId → confirm payment
6. Check totalPaidValue vs price → fulfill order
```

### Flow 3: Permanent Address Deposits (Gaming/Social)

```
1. Call Get Permanent Deposit Address with referenceId + chain
   → Get address (+ memo for XRP/XLM)
2. Display address to user (save in your DB)
3. Receive DirectDeposit webhook
4. Call Get Deposit Record with recordId → confirm amount & status
5. Credit user in your system
```

### Flow 4: Send Withdrawal

```
1. Call Check Withdrawal Address Validity → verify address
2. Call Get Withdrawal Network Fee → show fee to user
3. Call Get Coin Balance → verify sufficient balance
4. Call Create Network Withdrawal
5. Receive ApiWithdrawal webhook (Processing → Success)
6. Call Get Withdrawal Record → confirm final status
```

### Flow 5: User Wallet System

```
1. Call Create/Get User Deposit Address for each user
2. Receive UserDeposit webhook → call Get User Deposit Record
3. User balances: Get User Balance List
4. User-to-user transfers: Create Internal Transaction
5. User withdrawals: User Withdrawal to Blockchain
6. User swaps: Create User Swap Order
```

---

## Best Practices

### Security
- **Never expose** `appSecret` in client-side code or public repos
- Use **IP whitelisting** on Dashboard > Developer for production
- Always **verify webhook source** by checking the sender IP against CCPayment's whitelist IPs
- Always **call Get Record APIs** to confirm transaction status — never trust webhooks alone

### Reliability
- Implement **idempotency** using `recordId` to prevent double-crediting
- Handle **webhook retries** gracefully — same event may arrive multiple times
- For withdrawals, if no response due to network issues, call **Get Withdrawal Record** to check status
- Set reasonable **timeouts** with retry logic for API calls

### Performance
- **Cache** token lists and fiat lists — they don't change frequently
- Use **pagination** (nextId) for record list queries
- Respect **rate limits** — implement exponential backoff on 11004 errors

### Money Safety
- Never auto-credit **risky transactions** (`isFlaggedAsRisky: true`)
- For order-based deposits, compare **total paid vs expected** amount to handle partial/overpayments
- Use **amountOutMinimum** in swap orders for slippage protection
- Check **canDeposit / canWithdraw** flags before showing options to users

### Testnet
- **Always test on Sepolia first** before going live
- Remember to **disable testnet** when switching to production

---

## Links

| Resource | URL |
| --- | --- |
| CCPayment Website | [https://ccpayment.com](https://ccpayment.com) |
| API Documentation (V2) | [https://ccpayment.com/api/doc/?en](https://ccpayment.com/api/doc/?en#introduction) |
| Merchant Dashboard | [https://console.ccpayment.com](https://console.ccpayment.com) |
| Developer Config | [https://console.ccpayment.com/developer/config](https://console.ccpayment.com/developer/config) |
| Merchant Settings | [https://console.ccpayment.com/merchatsetting/menu/settings](https://console.ccpayment.com/merchatsetting/menu/settings) |
| Webhook Logs | [https://console.ccpayment.com/webhook/index](https://console.ccpayment.com/webhook/index) |
| Risky Transactions | [https://console.ccpayment.com/transaction/list/risk-list](https://console.ccpayment.com/transaction/list/risk-list) |
| Referral Program | [https://ccpayment.com/en/referral](https://ccpayment.com/en/referral) |
| X / Twitter | [https://x.com/CCPaymentX](https://x.com/CCPaymentX) |
| Telegram Community | [https://t.me/ccpayment_com](https://t.me/ccpayment_com) |
| Telegram Support | [https://t.me/CCPaymentSupportBot](https://t.me/CCPaymentSupportBot) |
| LinkedIn | [https://www.linkedin.com/company/ccpayment](https://www.linkedin.com/company/ccpayment) |
| Medium Blog | [https://ccpayment.medium.com](https://ccpayment.medium.com/) |
| Support Email | [support@ccpayment.com](mailto:support@ccpayment.com) |
| Partnership Email | [partner@ccpayment.com](mailto:partner@ccpayment.com) |
| Support Center | [https://ccpayment.com/support-center](https://ccpayment.com/support-center) |
