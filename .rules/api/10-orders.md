# Orders Module

## 10.1 Get Order Info

**Interface:** `POST /getAppOrderInfo`

**Description:** Get Order order details.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| amountToPay | string | Amount to pay |
| coinId | uint64 | Token ID |
| coinSymbol | string | Token symbol |
| chain | string | Chain name |
| toAddress | string | Receiving address |
| toMemo | string | Memo |
| createAt | int64 | Creation time |
| rate | string | Exchange rate |
| fiatId | uint64 | Fiat currency ID |
| fiatSymbol | string | Fiat currency symbol |
| expiredAt | int64 | Expiration time |
| checkoutUrl | string | Checkout URL |
| buyerEmail | string | Buyer email |
| paidList | Array | Paid list |
| paidList[].recordId | string | Record ID |
| paidList[].fromAddress | string | Source address |
| paidList[].amount | string | Amount |
| paidList[].serviceFee | string | Service fee |
| paidList[].txid | string | Transaction hash |
| paidList[].status | string | Status |
| paidList[].arrivedAt | int64 | Arrival time |
| paidList[].rate | string | Exchange rate |
| paidList[].isFlaggedAsRisky | bool | Whether marked as risky |

## 10.2 Create Invoice URL

**Interface:** `POST /createInvoiceUrl`

**Description:** Create an Invoice order and generate a payment URL.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| product | string | Yes | Product description | Length ≥1 |
| priceFiatId | uint64 | No | Fiat currency ID | - |
| priceCoinId | uint64 | No | Token ID | - |
| price | string | Yes | Price | Length ≥1 |
| buyerEmail | string | No | Buyer email | Email format |
| returnUrl | string | No | Return URL | Max 150 characters, URI format |
| expiredAt | int64 | No | Expiration timestamp | ≥1 |
| closeUrl | string | No | Close URL | Max 150 characters, URI format |
| notifyUrl | string | No | Notification URL | Max 150 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| invoiceUrl | string | Invoice payment URL |

## 10.3 Get Invoice Order Info

**Interface:** `POST /getInvoiceOrderInfo`

**Description:** Get Invoice order details.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| orderId | string | Order ID |
| createAt | int64 | Creation time |
| product | string | Product description |
| price | string | Price |
| priceCoinId | uint64 | Token ID |
| priceFiatId | uint64 | Fiat currency ID |
| priceSymbol | string | Price symbol |
| invoiceUrl | string | Invoice URL |
| buyerEmail | string | Buyer email |
| expiredAt | int64 | Expiration time |
| totalPaidValue | string | Total paid value |
| paidList | Array | Paid list |
| paidList[].recordId | string | Record ID |
| paidList[].coinId | uint64 | Token ID |
| paidList[].coinSymbol | string | Token symbol |
| paidList[].chain | string | Chain name |
| paidList[].fromAddress | string | Source address |
| paidList[].toAddress | string | Receiving address |
| paidList[].toMemo | string | Memo |
| paidList[].paidAmount | string | Paid amount |
| paidList[].serviceFee | string | Service fee |
| paidList[].txid | string | Transaction hash |
| paidList[].status | string | Status |
| paidList[].arrivedAt | int64 | Arrival time |
| paidList[].rate | string | Exchange rate |
| paidList[].paidValue | string | Paid value |
| paidList[].isFlaggedAsRisky | bool | Whether marked as risky |

## 10.4 Get Webhook Info

**Interface:** `POST /getWebhookInfo`

**Description:** Get merchant's Webhook configuration information.

**Request Parameters:** None

**Response Data:**

| Field | Type | Description |
|------|------|------|
| webhookUrl | string | Webhook URL |
| webhookSecret | string | Webhook secret |
