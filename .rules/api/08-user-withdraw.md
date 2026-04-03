# User Withdrawal Module

## 8.1 User Withdrawal to Network

**Interface:** `POST /applyUserWithdrawToNetwork`

**Description:** User applies for withdrawal to a blockchain network address.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| userId | string | Yes | User ID | Length 5-64 |
| orderId | string | Yes | Order ID | Length 3-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| chain | string | Yes | Chain name | Length ≥1 |
| address | string | Yes | Withdrawal address | Length ≥1 |
| amount | string | Yes | Withdrawal amount | Length ≥1 |
| memo | string | No | Memo | - |
| networkFeeInquiryID | string | No | Network fee inquiry ID | - |
| notifyUrl | string | No | Notification URL | Max 150 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| recordId | string | Withdrawal record ID |

## 8.2 User Withdrawal to CCWallet

**Interface:** `POST /applyUserWithdrawToCwallet`

**Description:** User applies for withdrawal to a CCWallet account.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| userId | string | Yes | User ID | Length 5-64 |
| orderId | string | Yes | Order ID | Length 3-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| cwalletUser | string | Yes | CCWallet user (email/ID) | Length ≥1 |
| amount | string | Yes | Withdrawal amount | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| recordId | string | Withdrawal record ID |

## 8.3 Query User Withdrawal Record

**Interface:** `POST /getUserWithdrawRecord`

**Description:** Query details of a single user withdrawal record.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| recordId | string | No | Record ID |
| orderId | string | No | Order ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| record | Object | Withdrawal record |
| record.userId | string | User ID |
| record.recordId | string | Record ID |
| record.withdrawType | string | Withdrawal type |
| record.coinId | uint64 | Token ID |
| record.chain | string | Chain name |
| record.orderId | string | Order ID |
| record.txId | string | Transaction hash (optional) |
| record.coinSymbol | string | Token symbol |
| record.fromAddress | string | Source address |
| record.toAddress | string | Target address |
| record.cwalletUser | string | CCWallet user |
| record.toMemo | string | Memo (optional) |
| record.amount | string | Amount |
| record.status | string | Status |
| record.fee | Object | Fee information |
| record.coinUSDPrice | string | Token USD price |

## 8.4 Query User Withdrawal Record List

**Interface:** `POST /getUserWithdrawRecordList`

**Description:** Query user withdrawal record list.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| userId | string | Yes | User ID | Length 5-64 |
| orderIds | Array<string> | No | Order ID list | - |
| chain | string | No | Chain name | - |
| coinId | uint64 | No | Token ID | - |
| startAt | int64 | No | Start time (default 90 days) | - |
| endAt | int64 | No | End time | - |
| toAddress | string | No | Target address | - |
| nextId | string | No | Next page ID | - |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Withdrawal record list (same structure as getUserWithdrawRecord) |
| nextId | string | Next page ID (optional) |
