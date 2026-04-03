# Merchant Batch Withdrawal Module

## 5.1 Check Withdrawal Address

**Interface:** `POST /checkWithdrawAddress`

**Description:** Check the validity of withdrawal addresses in batch.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| chain | string | Yes | Chain name | Length ≥1 |
| addressInfoList | Array | Yes | Address information list | 1-500 items |
| addressInfoList[].address | string | Yes | Address | Length ≥1 |
| addressInfoList[].memo | string | No | Memo | - |
| addressInfoList[].seq | uint32 | Yes | Sequence number | ≥1 |
| addressInfoList[].codes | Array<int32> | No | Error codes | - |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| addressInfoResults | Array | Address check results |
| addressInfoResults[].seq | uint32 | Sequence number |
| addressInfoResults[].codes | Array<int32> | Error code list (empty means valid) |

## 5.2 Apply Batch Withdrawal

**Interface:** `POST /applyBatchWithdraw`

**Description:** Create a batch withdrawal order.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| batchOrderId | string | Yes | Batch order ID | Length 3-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| chain | string | Yes | Chain name | Length ≥1 |
| productName | string | No | Product name | - |
| orders | Array | No | Order list | 0-500 items |
| orders[].seq | uint32 | Yes | Sequence number | >0 |
| orders[].address | string | Yes | Address | Length ≥1 |
| orders[].memo | string | No | Memo | Max 16 characters |
| orders[].amount | string | Yes | Amount | Length ≥1 |
| orders[].remark | string | No | Remark | Max 64 characters |
| mode | string | Yes | Execution mode (Single/Batch) | Length ≥1 |
| notifyUrl | string | No | Notification URL | Max 120 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| batchOrderId | string | Batch order ID |
| billId | string | Bill ID |

## 5.3 Append Batch Withdrawal

**Interface:** `POST /appendBatchWithdraw`

**Description:** Append tasks to an existing batch withdrawal order.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| batchOrderId | string | Yes | Batch order ID | Length 3-64 |
| orders | Array | Yes | Order list | 1-500 items |

**Response Data:** Empty object

## 5.4 Confirm Batch Withdrawal

**Interface:** `POST /confirmBatchWithdraw`

**Description:** Confirm and execute the batch withdrawal order.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| batchOrderId | string | Yes | Batch order ID | Length 3-64 |
| delaySeconds | int64 | No | Delay execution time (seconds) | 0-3600 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| batchOrderId | string | Batch order ID |
| productName | string | Product name |
| billId | string | Bill ID |
| coin | Object | Token information |
| coin.coin_id | uint64 | Token ID |
| coin.coin_symbol | string | Token symbol |
| coin.coin_price | string | Token price |
| amount | string | Total amount |
| networkFee | string | Network fee |
| networkFeeCoin | Object | Fee token information |
| status | string | Status (Init/Reviewing/Rejected/Pending/Processing/Completed) |
| reason | string | Reason |
| mode | string | Execution mode |
| stats | Object | Statistics |
| stats.total | int32 | Total count |
| stats.succeeded | int32 | Succeeded count |
| stats.failed | int32 | Failed count |
| stats.canceled | int32 | Canceled count |
| stats.processing | int32 | Processing count |
| stats.execSeq | uint32 | Executed sequence |
| createdAt | int64 | Creation time |
| updatedAt | int64 | Update time |

## 5.5 Cancel Batch Withdrawal

**Interface:** `POST /abortBatchWithdraw`

**Description:** Cancel the batch withdrawal order or part of the tasks.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| batchOrderId | string | Yes | Batch order ID | Length 3-64 |
| seqs | Array<uint32> | No | List of sequence numbers to cancel | 0-500 items |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| batchOrderId | string | Batch order ID |
| canceledSeqs | Array<uint32> | Canceled sequence numbers |
| ignoredSeqs | Array<uint32> | Ignored sequence numbers |

## 5.6 Get Batch Withdrawal Order

**Interface:** `POST /getBatchWithdrawOrder`

**Description:** Query batch withdrawal order details.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| batchOrderId | string | Yes | Batch order ID | Length 3-64 |
| verbose | uint32 | No | Verbosity level (0-3) | 0-3 |

**Response Data:** Same as confirmBatchWithdraw

## 5.7 Get Batch Withdrawal Record List

**Interface:** `POST /getBatchWithdrawOrderRecordList`

**Description:** Query the task record list of a batch withdrawal order.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| batchOrderId | string | Yes | Batch order ID | Length 3-64 |
| nextId | string | No | Next page ID | - |
| limit | uint64 | No | Items per page | ≤100 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| nextId | string | Next page ID |
| records | Array | Task record list |
| records[].seq | uint32 | Sequence number |
| records[].address | string | Address |
| records[].memo | string | Memo |
| records[].amount | string | Amount |
| records[].remark | string | Remark |
| records[].recordId | string | Record ID |
| records[].orderId | string | Order ID |
| records[].status | string | Status |
| records[].networkFee | string | Network fee |
| records[].txId | string | Transaction hash |
| records[].reason | string | Reason |
| records[].createdAt | int64 | Creation time |
| records[].updatedAt | int64 | Update time |
