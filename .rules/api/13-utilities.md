# 工具接口

## 13.1 验证地址

**接口:** `POST /verifyAddress`

**描述:** 验证区块链地址的有效性。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| chain | string | 是 | 链名称 | 长度≥1 |
| address | string | 是 | 地址 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| valid | bool | 是否有效 |
| message | string | 验证消息 |

## 13.2 废弃地址

**接口:** `POST /abandonAddress`

**描述:** 废弃不再使用的充值地址。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| chain | string | 是 | 链名称 | 长度≥1 |
| address | string | 是 | 地址 | 长度≥1 |

**响应数据:** 空对象

## 13.3 获取Hosted Invoice订单信息

**接口:** `POST /hostedInvoiceOrderInfo`

**描述:** 获取Hosted Invoice订单详情。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| orderId | string | 订单ID |
| product | string | 产品描述 |
| price | string | 价格 |
| priceSymbol | string | 价格符号 |
| invoiceUrl | string | Invoice URL |
| buyerEmail | string | 买家邮箱 |
| expiredAt | int64 | 过期时间 |
| selectedCoinId | uint64 | 选择的代币ID |
| selectedChain | string | 选择的链 |
| toAddress | string | 到账地址 |
| toMemo | string | 备注 |
| amountToPay | string | 应付金额 |
| totalPaidValue | string | 已支付总价值 |
| paidList | Array | 已支付列表 |

## 13.4 获取支付信息

**接口:** `POST /getPayInfo`

**描述:** 获取订单的支付信息（用于前端展示）。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| orderId | string | 订单ID |
| product | string | 产品描述 |
| price | string | 价格 |
| priceSymbol | string | 价格符号 |
| address | string | 支付地址 |
| memo | string | 备注 |
| amount | string | 支付金额 |
| coinSymbol | string | 代币符号 |
| chain | string | 链名称 |
| qrCode | string | 二维码数据 |
| expiredAt | int64 | 过期时间 |

## 13.5 健康检查

**接口:** `POST /health`

**描述:** API健康检查接口。

**请求参数:** 无

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| status | string | 健康状态 |
| timestamp | int64 | 时间戳 |
