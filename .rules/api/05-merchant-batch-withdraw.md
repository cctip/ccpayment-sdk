# 商家批量提现模块

## 5.1 检查提现地址

**接口:** `POST /checkWithdrawAddress`

**描述:** 批量检查提现地址的有效性。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| chain | string | 是 | 链名称 | 长度≥1 |
| addressInfoList | Array | 是 | 地址信息列表 | 1-500项 |
| addressInfoList[].address | string | 是 | 地址 | 长度≥1 |
| addressInfoList[].memo | string | 否 | 备注 | - |
| addressInfoList[].seq | uint32 | 是 | 序号 | ≥1 |
| addressInfoList[].codes | Array<int32> | 否 | 错误码 | - |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| addressInfoResults | Array | 地址检查结果 |
| addressInfoResults[].seq | uint32 | 序号 |
| addressInfoResults[].codes | Array<int32> | 错误码列表（空表示有效） |

## 5.2 申请批量提现

**接口:** `POST /applyBatchWithdraw`

**描述:** 创建批量提现订单。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| batchOrderId | string | 是 | 批量订单ID | 长度3-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| chain | string | 是 | 链名称 | 长度≥1 |
| productName | string | 否 | 产品名称 | - |
| orders | Array | 否 | 订单列表 | 0-500项 |
| orders[].seq | uint32 | 是 | 序号 | >0 |
| orders[].address | string | 是 | 地址 | 长度≥1 |
| orders[].memo | string | 否 | 备注 | 最大16字符 |
| orders[].amount | string | 是 | 金额 | 长度≥1 |
| orders[].remark | string | 否 | 备注说明 | 最大64字符 |
| mode | string | 是 | 执行模式（Single/Batch） | 长度≥1 |
| notifyUrl | string | 否 | 通知URL | 最大120字符，URI格式 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| batchOrderId | string | 批量订单ID |
| billId | string | 账单ID |

## 5.3 追加批量提现

**接口:** `POST /appendBatchWithdraw`

**描述:** 向已存在的批量提现订单追加任务。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| batchOrderId | string | 是 | 批量订单ID | 长度3-64 |
| orders | Array | 是 | 订单列表 | 1-500项 |

**响应数据:** 空对象

## 5.4 确认批量提现

**接口:** `POST /confirmBatchWithdraw`

**描述:** 确认并执行批量提现订单。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| batchOrderId | string | 是 | 批量订单ID | 长度3-64 |
| delaySeconds | int64 | 否 | 延迟执行时间（秒） | 0-3600 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| batchOrderId | string | 批量订单ID |
| productName | string | 产品名称 |
| billId | string | 账单ID |
| coin | Object | 代币信息 |
| coin.coin_id | uint64 | 代币ID |
| coin.coin_symbol | string | 代币符号 |
| coin.coin_price | string | 代币价格 |
| amount | string | 总金额 |
| networkFee | string | 网络手续费 |
| networkFeeCoin | Object | 手续费代币信息 |
| status | string | 状态（Init/Reviewing/Rejected/Pending/Processing/Completed） |
| reason | string | 原因 |
| mode | string | 执行模式 |
| stats | Object | 统计信息 |
| stats.total | int32 | 总笔数 |
| stats.succeeded | int32 | 成功笔数 |
| stats.failed | int32 | 失败笔数 |
| stats.canceled | int32 | 取消笔数 |
| stats.processing | int32 | 处理中笔数 |
| stats.execSeq | uint32 | 已执行序号 |
| createdAt | int64 | 创建时间 |
| updatedAt | int64 | 更新时间 |

## 5.5 取消批量提现

**接口:** `POST /abortBatchWithdraw`

**描述:** 取消批量提现订单或部分任务。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| batchOrderId | string | 是 | 批量订单ID | 长度3-64 |
| seqs | Array<uint32> | 否 | 要取消的序号列表 | 0-500项 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| batchOrderId | string | 批量订单ID |
| canceledSeqs | Array<uint32> | 已取消的序号 |
| ignoredSeqs | Array<uint32> | 被忽略的序号 |

## 5.6 获取批量提现订单

**接口:** `POST /getBatchWithdrawOrder`

**描述:** 查询批量提现订单详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| batchOrderId | string | 是 | 批量订单ID | 长度3-64 |
| verbose | uint32 | 否 | 详细程度（0-3） | 0-3 |

**响应数据:** 同confirmBatchWithdraw

## 5.7 获取批量提现记录列表

**接口:** `POST /getBatchWithdrawOrderRecordList`

**描述:** 查询批量提现订单的任务记录列表。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| batchOrderId | string | 是 | 批量订单ID | 长度3-64 |
| nextId | string | 否 | 下一页ID | - |
| limit | uint64 | 否 | 每页数量 | ≤100 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| nextId | string | 下一页ID |
| records | Array | 任务记录列表 |
| records[].seq | uint32 | 序号 |
| records[].address | string | 地址 |
| records[].memo | string | 备注 |
| records[].amount | string | 金额 |
| records[].remark | string | 备注说明 |
| records[].recordId | string | 记录ID |
| records[].orderId | string | 订单ID |
| records[].status | string | 状态 |
| records[].networkFee | string | 网络手续费 |
| records[].txId | string | 交易哈希 |
| records[].reason | string | 原因 |
| records[].createdAt | int64 | 创建时间 |
| records[].updatedAt | int64 | 更新时间 |
