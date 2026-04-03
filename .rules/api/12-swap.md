# Swap Module

## 12.1 Get Swap Coin List

**Interface:** `POST /getSwapCoinList`

**Description:** Get the list of tokens supported for swapping.

**Request Parameters:** None

**Response Data:**

| Field | Type | Description |
|------|------|------|
| coins | Array | Token list |
| coins[].coinId | uint64 | Token ID |
| coins[].symbol | string | Token symbol |
| coins[].logoUrl | string | Logo URL |
| coins[].chains | Array<string> | Supported chain list |

## 12.2 Swap Estimate

**Interface:** `POST /swapEstimate`

**Description:** Estimate the swap amount.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| fromCoinId | uint64 | Yes | From token ID | ≥1 |
| toCoinId | uint64 | Yes | To token ID | ≥1 |
| fromAmount | string | Yes | From amount | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| fromCoinId | uint64 | From token ID |
| fromCoinSymbol | string | From token symbol |
| fromAmount | string | From amount |
| toCoinId | uint64 | To token ID |
| toCoinSymbol | string | To token symbol |
| toAmount | string | To amount |
| rate | string | Exchange rate |

## 12.3 Execute Swap

**Interface:** `POST /swap`

**Description:** Execute a swap operation.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| fromCoinId | uint64 | Yes | From token ID | ≥1 |
| toCoinId | uint64 | Yes | To token ID | ≥1 |
| fromAmount | string | Yes | From amount | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| recordId | string | Swap record ID |

## 12.4 Query Swap Record

**Interface:** `POST /getSwapRecord`

**Description:** Query details of a single swap record.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| recordId | string | No | Record ID | - |
| orderId | string | No | Order ID | - |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| record | Object | Swap record |
| record.recordId | string | Record ID |
| record.orderId | string | Order ID |
| record.fromCoinId | uint64 | From token ID |
| record.fromCoinSymbol | string | From token symbol |
| record.fromAmount | string | From amount |
| record.toCoinId | uint64 | To token ID |
| record.toCoinSymbol | string | To token symbol |
| record.toAmount | string | To amount |
| record.rate | string | Exchange rate |
| record.status | string | Status |
| record.createdAt | int64 | Creation time |

## 12.5 Query Swap Record List

**Interface:** `POST /getSwapRecordList`

**Description:** Query swap record list.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| orderIds | Array<string> | No | Order ID list |
| fromCoinId | uint64 | No | From token ID |
| toCoinId | uint64 | No | To token ID |
| startAt | int64 | No | Start time (default 90 days) |
| endAt | int64 | No | End time |
| nextId | string | No | Next page ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Swap record list (same structure as getSwapRecord) |
| nextId | string | Next page ID (optional) |

## 12.6 Select Hosted Invoice Coin

**Interface:** `POST /selectHostedInvoiceCoin`

**Description:** Select the payment token for a Hosted Invoice order.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| chain | string | Yes | Chain name | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| address | string | Deposit address |
| memo | string | Memo (optional) |
| amountToPay | string | Amount to pay |
