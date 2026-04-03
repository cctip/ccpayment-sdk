# 用户提现模块

## 8.1 用户提现到网络

**接口:** `POST /applyUserWithdrawToNetwork`

**描述:** 用户申请提现到区块链网络地址。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| userId | string | 是 | 用户ID | 长度5-64 |
| orderId | string | 是 | 订单ID | 长度3-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| chain | string | 是 | 链名称 | 长度≥1 |
| address | string | 是 | 提现地址 | 长度≥1 |
| amount | string | 是 | 提现金额 | 长度≥1 |
| memo | string | 否 | 备注 | - |
| networkFeeInquiryID | string | 否 | 网络费查询ID | - |
| notifyUrl | string | 否 | 通知URL | 最大150字符，URI格式 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| recordId | string | 提现记录ID |

## 8.2 用户提现到CCWallet

**接口:** `POST /applyUserWithdrawToCwallet`

**描述:** 用户申请提现到CCWallet账户。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| userId | string | 是 | 用户ID | 长度5-64 |
| orderId | string | 是 | 订单ID | 长度3-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| cwalletUser | string | 是 | CCWallet用户（邮箱/ID） | 长度≥1 |
| amount | string | 是 | 提现金额 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| recordId | string | 提现记录ID |

## 8.3 查询用户提现记录

**接口:** `POST /getUserWithdrawRecord`

**描述:** 查询单条用户提现记录详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| recordId | string | 否 | 记录ID |
| orderId | string | 否 | 订单ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| record | Object | 提现记录 |
| record.userId | string | 用户ID |
| record.recordId | string | 记录ID |
| record.withdrawType | string | 提现类型 |
| record.coinId | uint64 | 代币ID |
| record.chain | string | 链名称 |
| record.orderId | string | 订单ID |
| record.txId | string | 交易哈希（可选） |
| record.coinSymbol | string | 代币符号 |
| record.fromAddress | string | 来源地址 |
| record.toAddress | string | 目标地址 |
| record.cwalletUser | string | CCWallet用户 |
| record.toMemo | string | 备注（可选） |
| record.amount | string | 金额 |
| record.status | string | 状态 |
| record.fee | Object | 手续费信息 |
| record.coinUSDPrice | string | 代币USD价格 |

## 8.4 查询用户提现记录列表

**接口:** `POST /getUserWithdrawRecordList`

**描述:** 查询用户提现记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| userId | string | 是 | 用户ID | 长度5-64 |
| orderIds | Array<string> | 否 | 订单ID列表 | - |
| chain | string | 否 | 链名称 | - |
| coinId | uint64 | 否 | 代币ID | - |
| startAt | int64 | 否 | 开始时间（默认90天） | - |
| endAt | int64 | 否 | 结束时间 | - |
| toAddress | string | 否 | 目标地址 | - |
| nextId | string | 否 | 下一页ID | - |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 提现记录列表（结构同getUserWithdrawRecord） |
| nextId | string | 下一页ID（可选） |
