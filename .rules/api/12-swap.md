# 换币模块

## 12.1 获取换币代币列表

**接口:** `POST /getSwapCoinList`

**描述:** 获取支持换币的代币列表。

**请求参数:** 无

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| coins | Array | 代币列表 |
| coins[].coinId | uint64 | 代币ID |
| coins[].symbol | string | 代币符号 |
| coins[].logoUrl | string | Logo URL |
| coins[].chains | Array<string> | 支持的链列表 |

## 12.2 换币估算

**接口:** `POST /swapEstimate`

**描述:** 估算换币金额。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| fromCoinId | uint64 | 是 | 换出代币ID | ≥1 |
| toCoinId | uint64 | 是 | 换入代币ID | ≥1 |
| fromAmount | string | 是 | 换出金额 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| fromCoinId | uint64 | 换出代币ID |
| fromCoinSymbol | string | 换出代币符号 |
| fromAmount | string | 换出金额 |
| toCoinId | uint64 | 换入代币ID |
| toCoinSymbol | string | 换入代币符号 |
| toAmount | string | 换入金额 |
| rate | string | 汇率 |

## 12.3 执行换币

**接口:** `POST /swap`

**描述:** 执行换币操作。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| fromCoinId | uint64 | 是 | 换出代币ID | ≥1 |
| toCoinId | uint64 | 是 | 换入代币ID | ≥1 |
| fromAmount | string | 是 | 换出金额 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| recordId | string | 换币记录ID |

## 12.4 查询换币记录

**接口:** `POST /getSwapRecord`

**描述:** 查询单条换币记录详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| recordId | string | 否 | 记录ID | - |
| orderId | string | 否 | 订单ID | - |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| record | Object | 换币记录 |
| record.recordId | string | 记录ID |
| record.orderId | string | 订单ID |
| record.fromCoinId | uint64 | 换出代币ID |
| record.fromCoinSymbol | string | 换出代币符号 |
| record.fromAmount | string | 换出金额 |
| record.toCoinId | uint64 | 换入代币ID |
| record.toCoinSymbol | string | 换入代币符号 |
| record.toAmount | string | 换入金额 |
| record.rate | string | 汇率 |
| record.status | string | 状态 |
| record.createdAt | int64 | 创建时间 |

## 12.5 查询换币记录列表

**接口:** `POST /getSwapRecordList`

**描述:** 查询换币记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| orderIds | Array<string> | 否 | 订单ID列表 |
| fromCoinId | uint64 | 否 | 换出代币ID |
| toCoinId | uint64 | 否 | 换入代币ID |
| startAt | int64 | 否 | 开始时间（默认90天） |
| endAt | int64 | 否 | 结束时间 |
| nextId | string | 否 | 下一页ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 换币记录列表（结构同getSwapRecord） |
| nextId | string | 下一页ID（可选） |

## 12.6 选择Hosted Invoice代币

**接口:** `POST /selectHostedInvoiceCoin`

**描述:** 为Hosted Invoice订单选择支付代币。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度≥1 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| chain | string | 是 | 链名称 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| address | string | 充值地址 |
| memo | string | 备注（可选） |
| amountToPay | string | 应付金额 |
