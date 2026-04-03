# 用户充值模块

## 7.1 获取或创建用户充值地址

**接口:** `POST /getOrCreateUserDepositAddress`

**描述:** 获取或创建用户的充值地址。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| userId | string | 是 | 用户ID（不能以sys开头） | 长度5-64 |
| chain | string | 是 | 链名称 | 长度≥1 |
| notifyUrl | string | 否 | 通知URL | 最大150字符，URI格式 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| address | string | 充值地址 |
| memo | string | 备注（可选） |

## 7.2 查询用户充值记录

**接口:** `POST /getUserDepositRecord`

**描述:** 查询单条用户充值记录详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| recordId | string | 是 | 记录ID | 长度5-64 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| record | Object | 充值记录 |
| record.userId | string | 用户ID |
| record.recordId | string | 记录ID |
| record.coinId | uint64 | 代币ID |
| record.chain | string | 链名称 |
| record.contract | string | 合约地址 |
| record.coinSymbol | string | 代币符号 |
| record.txId | string | 交易哈希 |
| record.coinUSDPrice | string | 代币USD价格 |
| record.fromAddress | string | 来源地址 |
| record.toAddress | string | 到账地址 |
| record.toMemo | string | 备注（可选） |
| record.amount | string | 金额 |
| record.serviceFee | string | 服务费 |
| record.status | string | 状态 |
| record.arrivedAt | int64 | 到账时间 |
| record.isFlaggedAsRisky | bool | 是否标记为风险（可选） |

## 7.3 查询用户充值记录列表

**接口:** `POST /getUserDepositRecordList`

**描述:** 查询用户充值记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| userId | string | 是 | 用户ID | 长度5-64 |
| chain | string | 否 | 链名称 | - |
| coinId | uint64 | 否 | 代币ID | - |
| toAddress | string | 否 | 到账地址 | - |
| startAt | int64 | 否 | 开始时间（默认90天） | - |
| endAt | int64 | 否 | 结束时间 | - |
| nextId | string | 否 | 下一页ID | - |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| records | Array | 充值记录列表（结构同getUserDepositRecord） |
| nextId | string | 下一页ID（可选） |
