# Utilities

## 13.1 Verify Address

**Interface:** `POST /verifyAddress`

**Description:** Verify the validity of a blockchain address.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| chain | string | Yes | Chain name | Length ≥1 |
| address | string | Yes | Address | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| valid | bool | Whether valid |
| message | string | Verification message |

## 13.2 Abandon Address

**Interface:** `POST /abandonAddress`

**Description:** Abandon a deposit address that is no longer in use.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| chain | string | Yes | Chain name | Length ≥1 |
| address | string | Yes | Address | Length ≥1 |

**Response Data:** Empty object

## 13.3 Get Hosted Invoice Order Info

**Interface:** `POST /hostedInvoiceOrderInfo`

**Description:** Get Hosted Invoice order details.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| orderId | string | Order ID |
| product | string | Product description |
| price | string | Price |
| priceSymbol | string | Price symbol |
| invoiceUrl | string | Invoice URL |
| buyerEmail | string | Buyer email |
| expiredAt | int64 | Expiration time |
| selectedCoinId | uint64 | Selected token ID |
| selectedChain | string | Selected chain |
| toAddress | string | Receiving address |
| toMemo | string | Memo |
| amountToPay | string | Amount to pay |
| totalPaidValue | string | Total paid value |
| paidList | Array | Paid list |

## 13.4 Get Pay Info

**Interface:** `POST /getPayInfo`

**Description:** Get the payment information of an order (for frontend display).

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| orderId | string | Order ID |
| product | string | Product description |
| price | string | Price |
| priceSymbol | string | Price symbol |
| address | string | Payment address |
| memo | string | Memo |
| amount | string | Payment amount |
| coinSymbol | string | Token symbol |
| chain | string | Chain name |
| qrCode | string | QR code data |
| expiredAt | int64 | Expiration time |

## 13.5 Health Check

**Interface:** `POST /health`

**Description:** API health check interface.

**Request Parameters:** None

**Response Data:**

| Field | Type | Description |
|------|------|------|
| status | string | Health status |
| timestamp | int64 | Timestamp |
