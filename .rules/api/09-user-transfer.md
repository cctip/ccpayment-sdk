# User Transfer Module

## 9.1 User Transfer

**Interface:** `POST /userTransfer`

**Description:** Initiate a transfer between users.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| fromUserId | string | Yes | From user ID | Length 5-64 |
| toUserId | string | Yes | To user ID | Length 5-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| amount | string | Yes | Transfer amount | Length ≥1 |
| remark | string | No | Remark | - |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| recordId | string | Transfer record ID |

## 9.2 Query User Transfer Record

**Interface:** `POST /getUserTransferRecord`

**Description:** Query details of a single user transfer record.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| recordId | string | No | Record ID |
| orderId | string | No | Order ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| record | Object | Transfer record |
| record.recordId | string | Record ID |
| record.coinId | uint64 | Token ID |
| record.coinSymbol | string | Token symbol |
| record.orderId | string | Order ID |
| record.fromUserId | string | From user ID |
| record.toUserId | string | To user ID |
| record.amount | string | Amount |
| record.status | string | Status |
| record.remark | string | Remark (optional) |
| record.coinUSDPrice | string | Token USD price |

## 9.3 Query User Transfer Record List

**Interface:** `POST /getUserTransferRecordList`

**Description:** Query user transfer record list.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| orderIds | Array<string> | No | Order ID list |
| fromUserId | string | No | From user ID |
| toUserId | string | No | To user ID |
| coinId | uint64 | No | Token ID |
| startAt | int64 | No | Start time (default 90 days) |
| endAt | int64 | No | End time |
| nextId | string | No | Next page ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Transfer record list (same structure as getUserTransferRecord) |
| nextId | string | Next page ID (optional) |

## 9.4 User Batch Transfer

**Interface:** `POST /userBatchTransfer`

**Description:** Initiate a batch transfer for users.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| userId | string | Yes | From user ID | Length 5-64 |
| toUsers | Array | Yes | To user list | At least 1 item |
| toUsers[].userId | string | Yes | To user ID | Length 5-64 |
| toUsers[].amount | string | Yes | Transfer amount | Length 5-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| remark | string | No | Remark | - |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| recordId | string | Batch transfer record ID |

## 9.5 Query User Batch Transfer Record

**Interface:** `POST /getUserBatchTransferRecord`

**Description:** Query details of a single user batch transfer record.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| recordId | string | No | Record ID |
| orderId | string | No | Order ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| record | Object | Batch transfer record |
| record.recordId | string | Record ID |
| record.userId | string | From user ID |
| record.coinId | uint64 | Token ID |
| record.coinSymbol | string | Token symbol |
| record.orderId | string | Order ID |
| record.toUsers | Array | To user list |
| record.status | string | Status |
| record.remark | string | Remark (optional) |
| record.coinUSDPrice | string | Token USD price |

## 9.6 Query User Batch Transfer Record List

**Interface:** `POST /getUserBatchTransferRecordList`

**Description:** Query user batch transfer record list.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| orderIds | Array<string> | No | Order ID list |
| userId | string | No | User ID |
| coinId | uint64 | No | Token ID |
| startAt | int64 | No | Start time (default 90 days) |
| endAt | int64 | No | End time |
| nextId | string | No | Next page ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Batch transfer record list (same structure as getUserBatchTransferRecord) |
| nextId | string | Next page ID (optional) |

## 9.7 Query User Batch Transfer Record Details

**Interface:** `POST /getUserBatchTransferRecordDetail`

**Description:** Query the detail list of a user batch transfer record.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| recordId | string | Yes | Batch transfer record ID | Length ≥1 |
| nextId | string | No | Next page ID | - |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| records | Array | Transfer detail list (same structure as getUserTransferRecord) |
| nextId | string | Next page ID (optional) |
