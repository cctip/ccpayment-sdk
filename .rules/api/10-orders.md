# 订单模块

## 10.1 获取订单信息

**接口:** `POST /getAppOrderInfo`

**描述:** 获取Order订单详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| amountToPay | string | 应付金额 |
| coinId | uint64 | 代币ID |
| coinSymbol | string | 代币符号 |
| chain | string | 链名称 |
| toAddress | string | 到账地址 |
| toMemo | string | 备注 |
| createAt | int64 | 创建时间 |
| rate | string | 汇率 |
| fiatId | uint64 | 法币ID |
| fiatSymbol | string | 法币符号 |
| expiredAt | int64 | 过期时间 |
| checkoutUrl | string | 收银台URL |
| buyerEmail | string | 买家邮箱 |
| paidList | Array | 已支付列表 |
| paidList[].recordId | string | 记录ID |
| paidList[].fromAddress | string | 来源地址 |
| paidList[].amount | string | 金额 |
| paidList[].serviceFee | string | 服务费 |
| paidList[].txid | string | 交易哈希 |
| paidList[].status | string | 状态 |
| paidList[].arrivedAt | int64 | 到账时间 |
| paidList[].rate | string | 汇率 |
| paidList[].isFlaggedAsRisky | bool | 是否标记为风险 |

## 10.2 创建发票URL

**接口:** `POST /createInvoiceUrl`

**描述:** 创建Invoice订单并生成支付URL。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| product | string | 是 | 产品描述 | 长度≥1 |
| priceFiatId | uint64 | 否 | 法币ID | - |
| priceCoinId | uint64 | 否 | 代币ID | - |
| price | string | 是 | 价格 | 长度≥1 |
| buyerEmail | string | 否 | 买家邮箱 | 邮箱格式 |
| returnUrl | string | 否 | 返回URL | 最大150字符，URI格式 |
| expiredAt | int64 | 否 | 过期时间戳 | ≥1 |
| closeUrl | string | 否 | 关闭URL | 最大150字符，URI格式 |
| notifyUrl | string | 否 | 通知URL | 最大150字符，URI格式 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| invoiceUrl | string | 发票支付URL |

## 10.3 获取发票订单信息

**接口:** `POST /getInvoiceOrderInfo`

**描述:** 获取Invoice订单详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| orderId | string | 订单ID |
| createAt | int64 | 创建时间 |
| product | string | 产品描述 |
| price | string | 价格 |
| priceCoinId | uint64 | 代币ID |
| priceFiatId | uint64 | 法币ID |
| priceSymbol | string | 价格符号 |
| invoiceUrl | string | 发票URL |
| buyerEmail | string | 买家邮箱 |
| expiredAt | int64 | 过期时间 |
| totalPaidValue | string | 已支付总价值 |
| paidList | Array | 已支付列表 |
| paidList[].recordId | string | 记录ID |
| paidList[].coinId | uint64 | 代币ID |
| paidList[].coinSymbol | string | 代币符号 |
| paidList[].chain | string | 链名称 |
| paidList[].fromAddress | string | 来源地址 |
| paidList[].toAddress | string | 到账地址 |
| paidList[].toMemo | string | 备注 |
| paidList[].paidAmount | string | 支付金额 |
| paidList[].serviceFee | string | 服务费 |
| paidList[].txid | string | 交易哈希 |
| paidList[].status | string | 状态 |
| paidList[].arrivedAt | int64 | 到账时间 |
| paidList[].rate | string | 汇率 |
| paidList[].paidValue | string | 支付价值 |
| paidList[].isFlaggedAsRisky | bool | 是否标记为风险 |

## 10.4 获取Webhook信息

**接口:** `POST /getWebhookInfo`

**描述:** 获取商家的Webhook配置信息。

**请求参数:** 无

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| webhookUrl | string | Webhook URL |
| webhookSecret | string | Webhook密钥 |
