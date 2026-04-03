# User Deposit Module

## 7.1 Get or Create User Deposit Address

**Interface:** `POST /getOrCreateUserDepositAddress`

**Description:** Get or create a user's deposit address.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| userId | string | Yes | User ID (cannot start with sys) | Length 5-64 |
| chain | string | Yes | Chain name | Length ≥1 |
| notifyUrl | string | No | Notification URL | Max 150 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| address | string | Deposit address |
| memo | string | Memo (optional) |

## 7.2 Query User Deposit Record

**Interface:** `POST /getUserDepositRecord`

**Description:** Query details of a single user deposit record.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| recordId | string | Yes | Record ID | Length 5-64 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| record | Object | Deposit record |
| record.userId | string | User ID |
| record.recordId | string | Record ID |
| record.coinId | uint64 | Token ID |
| record.chain | string | Chain name |
| record.contract | string | Contract address |
| record.coinSymbol | string | Token symbol |
| record.txId | string | Transaction hash |
| record.coinUSDPrice | string | Token USD price |
| record.fromAddress | string | Source address |
| record.toAddress | string | Receiving address |
| record.toMemo | string | Memo (optional) |
| record.amount | string | Amount |
| record.serviceFee | string | Service fee |
| record.status | string | Status |
| record.arrivedAt | int64 | Arrival time |
| record.isFlaggedAsRisky | bool | Whether marked as risky (optional) |

## 7.3 Query User Deposit Record List

**Interface:** `POST /getUserDepositRecordList`

**Description:** Query user deposit record list.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| userId | string | Yes | User ID | Length 5-64 |
| chain | string | No | Chain name | - |
| coinId | uint64 | No | Token ID | - |
| toAddress | string | No | Receiving address | - |
| startAt | int64 | No | Start time (default 90 days) | - |
| endAt | int64 | No | End time | - |
| nextId | string | No | Next page ID | - |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Deposit record list (same structure as getUserDepositRecord) |
| nextId | string | Next page ID (optional) |
