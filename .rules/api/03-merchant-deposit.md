# 商家充值模块

## 3.1 创建订单充值地址

**接口:** `POST /createAppOrderDepositAddress`

**描述:** 为特定订单创建充值地址（地址模式）。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| fiatId | uint64 | 否 | 法币ID | - |
| chain | string | 是 | 链名称 | 长度≥1 |
| price | string | 是 | 价格 | 长度≥1 |
| expiredAt | int64 | 否 | 过期时间戳 | ≥1 |
| buyerEmail | string | 否 | 买家邮箱 | 邮箱格式 |
| generateCheckoutURL | bool | 否 | 是否生成收银台URL | - |
| product | string | 否 | 产品名称 | 最大120字符 |
| returnUrl | string | 否 | 返回URL | 最大150字符，URI格式 |
| notifyUrl | string | 否 | 通知URL | 最大150字符，URI格式 |
| closeUrl | string | 否 | 关闭URL | 最大150字符，URI格式 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| address | string | 充值地址 |
| amount | string | 充值金额 |
| memo | string | 备注（可选） |
| checkoutUrl | string | 收银台URL（可选） |
| confirmsNeeded | uint64 | 需要的确认数 |

## 3.2 获取或创建充值地址

**接口:** `POST /getOrCreateAppDepositAddress`

**描述:** 获取或创建直接充值地址（直接地址模式）。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| referenceId | string | 是 | 引用ID | 长度3-64 |
| chain | string | 是 | 链名称 | 长度≥1 |
| notifyUrl | string | 否 | 通知URL | 最大150字符，URI格式 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| address | string | 充值地址 |
| memo | string | 备注（可选） |

## 3.3 查询充值记录

**接口:** `POST /getAppDepositRecord`

**描述:** 查询单条充值记录详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| recordId | string | 是 | 记录ID | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| record | Object | 充值记录 |
| record.recordId | string | 记录ID |
| record.referenceId | string | 引用ID |
| record.orderId | string | 订单ID |
| record.coinId | uint64 | 代币ID |
| record.coinSymbol | string | 代币符号 |
| record.chain | string | 链名称 |
| record.contract | string | 合约地址 |
| record.coinUSDPrice | string | 代币USD价格 |
| record.fromAddress | string | 来源地址 |
| record.toAddress | string | 到账地址 |
| record.toMemo | string | 备注（可选） |
| record.amount | string | 金额 |
| record.serviceFee | string | 服务费 |
| record.txId | string | 交易哈希 |
| record.txIndex | uint64 | 交易索引 |
| record.status | string | 状态 |
| record.arrivedAt | int64 | 到账时间 |
| record.isFlaggedAsRisky | bool | 是否标记为风险（可选） |

## 3.4 查询充值记录列表

**接口:** `POST /getAppDepositRecordList`

**描述:** 查询充值记录列表，支持多种筛选条件。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| chain | string | 否 | 链名称 | - |
| referenceId | string | 否 | 引用ID | - |
| orderId | string | 否 | 订单ID | - |
| toAddress | string | 否 | 到账地址 | - |
| coinId | uint64 | 否 | 代币ID | - |
| startAt | int64 | 否 | 开始时间（默认90天） | - |
| endAt | int64 | 否 | 结束时间 | - |
| nextId | string | 否 | 下一页ID | - |
| recordIds | Array<string> | 否 | 记录ID列表 | 最大50项 |
| referenceIds | Array<string> | 否 | 引用ID列表 | 最大50项 |
| orderIds | Array<string> | 否 | 订单ID列表 | 最大50项 |
| limit | uint64 | 否 | 每页数量 | ≤100 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 充值记录列表（结构同getAppDepositRecord） |
| nextId | string | 下一页ID（可选） |
