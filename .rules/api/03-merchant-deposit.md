# Merchant Deposit Module

## 3.1 Create Order Deposit Address

**Interface:** `POST /createAppOrderDepositAddress`

**Description:** Create a deposit address for a specific order (address mode).

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| fiatId | uint64 | No | Fiat currency ID | - |
| chain | string | Yes | Chain name | Length ≥1 |
| price | string | Yes | Price | Length ≥1 |
| expiredAt | int64 | No | Expiration timestamp | ≥1 |
| buyerEmail | string | No | Buyer email | Email format |
| generateCheckoutURL | bool | No | Whether to generate checkout URL | - |
| product | string | No | Product name | Max 120 characters |
| returnUrl | string | No | Return URL | Max 150 characters, URI format |
| notifyUrl | string | No | Notification URL | Max 150 characters, URI format |
| closeUrl | string | No | Close URL | Max 150 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| address | string | Deposit address |
| amount | string | Deposit amount |
| memo | string | Memo (optional) |
| checkoutUrl | string | Checkout URL (optional) |
| confirmsNeeded | uint64 | Required confirmation count |

## 3.2 Get or Create Deposit Address

**Interface:** `POST /getOrCreateAppDepositAddress`

**Description:** Get or create a direct deposit address (direct address mode).

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| referenceId | string | Yes | Reference ID | Length 3-64 |
| chain | string | Yes | Chain name | Length ≥1 |
| notifyUrl | string | No | Notification URL | Max 150 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| address | string | Deposit address |
| memo | string | Memo (optional) |

## 3.3 Query Deposit Record

**Interface:** `POST /getAppDepositRecord`

**Description:** Query details of a single deposit record.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| recordId | string | Yes | Record ID | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| record | Object | Deposit record |
| record.recordId | string | Record ID |
| record.referenceId | string | Reference ID |
| record.orderId | string | Order ID |
| record.coinId | uint64 | Token ID |
| record.coinSymbol | string | Token symbol |
| record.chain | string | Chain name |
| record.contract | string | Contract address |
| record.coinUSDPrice | string | Token USD price |
| record.fromAddress | string | Source address |
| record.toAddress | string | Receiving address |
| record.toMemo | string | Memo (optional) |
| record.amount | string | Amount |
| record.serviceFee | string | Service fee |
| record.txId | string | Transaction hash |
| record.txIndex | uint64 | Transaction index |
| record.status | string | Status |
| record.arrivedAt | int64 | Arrival time |
| record.isFlaggedAsRisky | bool | Whether marked as risky (optional) |

## 3.4 Query Deposit Record List

**Interface:** `POST /getAppDepositRecordList`

**Description:** Query deposit record list, supporting multiple filter conditions.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| chain | string | No | Chain name | - |
| referenceId | string | No | Reference ID | - |
| orderId | string | No | Order ID | - |
| toAddress | string | No | Receiving address | - |
| coinId | uint64 | No | Token ID | - |
| startAt | int64 | No | Start time (default 90 days) | - |
| endAt | int64 | No | End time | - |
| nextId | string | No | Next page ID | - |
| recordIds | Array<string> | No | Record ID list | Max 50 items |
| referenceIds | Array<string> | No | Reference ID list | Max 50 items |
| orderIds | Array<string> | No | Order ID list | Max 50 items |
| limit | uint64 | No | Items per page | ≤100 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Deposit record list (same structure as getAppDepositRecord) |
| nextId | string | Next page ID (optional) |
