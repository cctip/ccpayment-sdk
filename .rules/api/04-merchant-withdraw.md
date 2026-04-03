# Merchant Withdrawal Module

## 4.1 Withdraw to Network Address

**Interface:** `POST /applyAppWithdrawToNetwork`

**Description:** Apply for withdrawal to a blockchain network address.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| chain | string | Yes | Chain name | Length ≥1 |
| address | string | Yes | Withdrawal address | Length ≥1 |
| memo | string | No | Memo | - |
| amount | string | Yes | Withdrawal amount | Length ≥1 |
| merchantPayNetworkFee | bool | No | Whether merchant pays network fee | - |
| networkFeeInquiryID | string | No | Network fee inquiry ID | - |
| notifyUrl | string | No | Notification URL | Max 150 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| recordId | string | Withdrawal record ID |

## 4.2 Withdraw to CCWallet

**Interface:** `POST /applyAppWithdrawToCwallet`

**Description:** Apply for withdrawal to a CCWallet account.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| cwalletUser | string | Yes | CCWallet user (email/ID) | Length ≥1 |
| amount | string | Yes | Withdrawal amount | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| recordId | string | Withdrawal record ID |

## 4.3 Query Withdrawal Record

**Interface:** `POST /getAppWithdrawRecord`

**Description:** Query details of a single withdrawal record.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| orderId | string | No | Order ID |
| recordId | string | No | Record ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| record | Object | Withdrawal record |
| record.recordId | string | Record ID |
| record.withdrawType | string | Withdrawal type |
| record.appId | string | Application ID |
| record.coinId | uint64 | Token ID |
| record.coinSymbol | string | Token symbol |
| record.chain | string | Chain name |
| record.fromAddress | string | Source address |
| record.toAddress | string | Target address |
| record.cwalletUser | string | CCWallet user |
| record.orderId | string | Order ID |
| record.txId | string | Transaction hash (optional) |
| record.toMemo | string | Memo (optional) |
| record.status | string | Status |
| record.amount | string | Amount |
| record.fee | Object | Fee information |
| record.fee.coinId | uint64 | Token ID |
| record.fee.coinSymbol | string | Token symbol |
| record.fee.amount | string | Fee amount |
| record.reason | string | Reason (optional) |
| record.coinUSDPrice | string | Token USD price |

## 4.4 Query Withdrawal Record List

**Interface:** `POST /getAppWithdrawRecordList`

**Description:** Query withdrawal record list.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| chain | string | No | Chain name |
| coinId | uint64 | No | Token ID |
| orderIds | Array<string> | No | Order ID list |
| startAt | int64 | No | Start time (default 90 days) |
| endAt | int64 | No | End time |
| toAddress | string | No | Target address |
| nextId | string | No | Next page ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Withdrawal record list (same structure as getAppWithdrawRecord) |
| nextId | string | Next page ID (optional) |

## 4.5 Query Auto Withdrawal Record List

**Interface:** `POST /getAutoWithdrawRecordList`

**Description:** Query automatic withdrawal record list.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| chain | string | No | Chain name |
| coinId | uint64 | No | Token ID |
| recordIds | Array<string> | No | Record ID list |
| startAt | int64 | No | Start time (default 90 days) |
| endAt | int64 | No | End time |
| toAddress | string | No | Target address |
| nextId | string | No | Next page ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Automatic withdrawal record list |
| records[].recordId | string | Record ID |
| records[].coinId | uint64 | Token ID |
| records[].coinSymbol | string | Token symbol |
| records[].chain | string | Chain name |
| records[].orderId | string | Order ID |
| records[].toAddress | string | Target address |
| records[].toMemo | string | Memo (optional) |
| records[].amount | string | Amount |
| records[].txId | string | Transaction hash (optional) |
| records[].status | string | Status |
| records[].fee | Object | Fee information |
| records[].serviceFee | string | Service fee |
| records[].createdAt | int64 | Creation time |
| nextId | string | Next page ID (optional) |

## 4.6 Query Risky Refund Record List

**Interface:** `POST /getRiskyRefundRecordList`

**Description:** Query risky refund record list.

**Request Parameters:** Same as getAutoWithdrawRecordList

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Risky refund record list |
| records[].recordId | string | Record ID |
| records[].coinId | uint64 | Token ID |
| records[].coinSymbol | string | Token symbol |
| records[].chain | string | Chain name |
| records[].orderId | string | Order ID |
| records[].toAddress | string | Target address |
| records[].toMemo | string | Memo (optional) |
| records[].amount | string | Amount |
| records[].txId | string | Transaction hash (optional) |
| records[].status | string | Status |
| records[].fee | Object | Fee information |
| records[].createdAt | int64 | Creation time |
| records[].fromDeposit | Array | Source deposit records |
| nextId | string | Next page ID (optional) |
