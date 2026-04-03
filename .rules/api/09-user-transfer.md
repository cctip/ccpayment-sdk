# 用户转账模块

## 9.1 用户转账

**接口:** `POST /userTransfer`

**描述:** 发起用户间转账。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| fromUserId | string | 是 | 转出用户ID | 长度5-64 |
| toUserId | string | 是 | 转入用户ID | 长度5-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| amount | string | 是 | 转账金额 | 长度≥1 |
| remark | string | 否 | 备注 | - |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| recordId | string | 转账记录ID |

## 9.2 查询用户转账记录

**接口:** `POST /getUserTransferRecord`

**描述:** 查询单条用户转账记录详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| recordId | string | 否 | 记录ID |
| orderId | string | 否 | 订单ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| record | Object | 转账记录 |
| record.recordId | string | 记录ID |
| record.coinId | uint64 | 代币ID |
| record.coinSymbol | string | 代币符号 |
| record.orderId | string | 订单ID |
| record.fromUserId | string | 转出用户ID |
| record.toUserId | string | 转入用户ID |
| record.amount | string | 金额 |
| record.status | string | 状态 |
| record.remark | string | 备注（可选） |
| record.coinUSDPrice | string | 代币USD价格 |

## 9.3 查询用户转账记录列表

**接口:** `POST /getUserTransferRecordList`

**描述:** 查询用户转账记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| orderIds | Array<string> | 否 | 订单ID列表 |
| fromUserId | string | 否 | 转出用户ID |
| toUserId | string | 否 | 转入用户ID |
| coinId | uint64 | 否 | 代币ID |
| startAt | int64 | 否 | 开始时间（默认90天） |
| endAt | int64 | 否 | 结束时间 |
| nextId | string | 否 | 下一页ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 转账记录列表（结构同getUserTransferRecord） |
| nextId | string | 下一页ID（可选） |

## 9.4 用户批量转账

**接口:** `POST /userBatchTransfer`

**描述:** 发起用户批量转账。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| userId | string | 是 | 转出用户ID | 长度5-64 |
| toUsers | Array | 是 | 转入用户列表 | 至少1项 |
| toUsers[].userId | string | 是 | 转入用户ID | 长度5-64 |
| toUsers[].amount | string | 是 | 转账金额 | 长度5-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| remark | string | 否 | 备注 | - |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| recordId | string | 批量转账记录ID |

## 9.5 查询用户批量转账记录

**接口:** `POST /getUserBatchTransferRecord`

**描述:** 查询单条用户批量转账记录详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| recordId | string | 否 | 记录ID |
| orderId | string | 否 | 订单ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| record | Object | 批量转账记录 |
| record.recordId | string | 记录ID |
| record.userId | string | 转出用户ID |
| record.coinId | uint64 | 代币ID |
| record.coinSymbol | string | 代币符号 |
| record.orderId | string | 订单ID |
| record.toUsers | Array | 转入用户列表 |
| record.status | string | 状态 |
| record.remark | string | 备注（可选） |
| record.coinUSDPrice | string | 代币USD价格 |

## 9.6 查询用户批量转账记录列表

**接口:** `POST /getUserBatchTransferRecordList`

**描述:** 查询用户批量转账记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| orderIds | Array<string> | 否 | 订单ID列表 |
| userId | string | 否 | 用户ID |
| coinId | uint64 | 否 | 代币ID |
| startAt | int64 | 否 | 开始时间（默认90天） |
| endAt | int64 | 否 | 结束时间 |
| nextId | string | 否 | 下一页ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 批量转账记录列表（结构同getUserBatchTransferRecord） |
| nextId | string | 下一页ID（可选） |

## 9.7 查询用户批量转账记录明细

**接口:** `POST /getUserBatchTransferRecordDetail`

**描述:** 查询用户批量转账记录的明细列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| recordId | string | 是 | 批量转账记录ID | 长度≥1 |
| nextId | string | 否 | 下一页ID | - |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 转账明细列表（结构同getUserTransferRecord） |
| nextId | string | 下一页ID（可选） |
