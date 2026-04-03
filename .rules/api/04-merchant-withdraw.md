# 商家提现模块

## 4.1 提现到网络地址

**接口:** `POST /applyAppWithdrawToNetwork`

**描述:** 申请提现到区块链网络地址。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| chain | string | 是 | 链名称 | 长度≥1 |
| address | string | 是 | 提现地址 | 长度≥1 |
| memo | string | 否 | 备注 | - |
| amount | string | 是 | 提现金额 | 长度≥1 |
| merchantPayNetworkFee | bool | 否 | 商家是否支付网络费 | - |
| networkFeeInquiryID | string | 否 | 网络费查询ID | - |
| notifyUrl | string | 否 | 通知URL | 最大150字符，URI格式 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| recordId | string | 提现记录ID |

## 4.2 提现到CCWallet

**接口:** `POST /applyAppWithdrawToCwallet`

**描述:** 申请提现到CCWallet账户。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| cwalletUser | string | 是 | CCWallet用户（邮箱/ID） | 长度≥1 |
| amount | string | 是 | 提现金额 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| recordId | string | 提现记录ID |

## 4.3 查询提现记录

**接口:** `POST /getAppWithdrawRecord`

**描述:** 查询单条提现记录详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| orderId | string | 否 | 订单ID |
| recordId | string | 否 | 记录ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| record | Object | 提现记录 |
| record.recordId | string | 记录ID |
| record.withdrawType | string | 提现类型 |
| record.appId | string | 应用ID |
| record.coinId | uint64 | 代币ID |
| record.coinSymbol | string | 代币符号 |
| record.chain | string | 链名称 |
| record.fromAddress | string | 来源地址 |
| record.toAddress | string | 目标地址 |
| record.cwalletUser | string | CCWallet用户 |
| record.orderId | string | 订单ID |
| record.txId | string | 交易哈希（可选） |
| record.toMemo | string | 备注（可选） |
| record.status | string | 状态 |
| record.amount | string | 金额 |
| record.fee | Object | 手续费信息 |
| record.fee.coinId | uint64 | 代币ID |
| record.fee.coinSymbol | string | 代币符号 |
| record.fee.amount | string | 手续费金额 |
| record.reason | string | 原因（可选） |
| record.coinUSDPrice | string | 代币USD价格 |

## 4.4 查询提现记录列表

**接口:** `POST /getAppWithdrawRecordList`

**描述:** 查询提现记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| chain | string | 否 | 链名称 |
| coinId | uint64 | 否 | 代币ID |
| orderIds | Array<string> | 否 | 订单ID列表 |
| startAt | int64 | 否 | 开始时间（默认90天） |
| endAt | int64 | 否 | 结束时间 |
| toAddress | string | 否 | 目标地址 |
| nextId | string | 否 | 下一页ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 提现记录列表（结构同getAppWithdrawRecord） |
| nextId | string | 下一页ID（可选） |

## 4.5 查询自动提现记录列表

**接口:** `POST /getAutoWithdrawRecordList`

**描述:** 查询自动提现记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| chain | string | 否 | 链名称 |
| coinId | uint64 | 否 | 代币ID |
| recordIds | Array<string> | 否 | 记录ID列表 |
| startAt | int64 | 否 | 开始时间（默认90天） |
| endAt | int64 | 否 | 结束时间 |
| toAddress | string | 否 | 目标地址 |
| nextId | string | 否 | 下一页ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 自动提现记录列表 |
| records[].recordId | string | 记录ID |
| records[].coinId | uint64 | 代币ID |
| records[].coinSymbol | string | 代币符号 |
| records[].chain | string | 链名称 |
| records[].orderId | string | 订单ID |
| records[].toAddress | string | 目标地址 |
| records[].toMemo | string | 备注（可选） |
| records[].amount | string | 金额 |
| records[].txId | string | 交易哈希（可选） |
| records[].status | string | 状态 |
| records[].fee | Object | 手续费信息 |
| records[].serviceFee | string | 服务费 |
| records[].createdAt | int64 | 创建时间 |
| nextId | string | 下一页ID（可选） |

## 4.6 查询风险退款记录列表

**接口:** `POST /getRiskyRefundRecordList`

**描述:** 查询风险退款记录列表。

**请求参数:** 同getAutoWithdrawRecordList

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 风险退款记录列表 |
| records[].recordId | string | 记录ID |
| records[].coinId | uint64 | 代币ID |
| records[].coinSymbol | string | 代币符号 |
| records[].chain | string | 链名称 |
| records[].orderId | string | 订单ID |
| records[].toAddress | string | 目标地址 |
| records[].toMemo | string | 备注（可选） |
| records[].amount | string | 金额 |
| records[].txId | string | 交易哈希（可选） |
| records[].status | string | 状态 |
| records[].fee | Object | 手续费信息 |
| records[].createdAt | int64 | 创建时间 |
| records[].fromDeposit | Array | 来源充值记录 |
| nextId | string | 下一页ID（可选） |
