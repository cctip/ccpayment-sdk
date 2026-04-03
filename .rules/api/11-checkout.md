# 收银台模块

## 11.1 创建Checkout URL

**接口:** `POST /createCheckoutUrl`

**描述:** 创建Hosted Checkout支付URL。

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
| checkoutUrl | string | Checkout支付URL |

## 11.2 获取Hosted订单信息

**接口:** `POST /hostedOrderInfo`

**描述:** 获取Hosted订单详情。

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
| priceCoinId | uint64 | 代币ID |
| priceFiatId | uint64 | 法币ID |
| priceSymbol | string | 价格符号 |
| checkoutUrl | string | Checkout URL |
| buyerEmail | string | 买家邮箱 |
| expiredAt | int64 | 过期时间 |
| step | string | 当前步骤 |
| selectedCoinId | uint64 | 选择的代币ID |
| selectedChain | string | 选择的链 |
| toAddress | string | 到账地址 |
| toMemo | string | 备注 |
| amountToPay | string | 应付金额 |
| rate | string | 汇率 |
| totalPaidValue | string | 已支付总价值 |
| paidList | Array | 已支付列表 |
| refundList | Array | 退款列表 |

## 11.3 获取Hosted订单首次信息

**接口:** `POST /hostedOrderInfoFirst`

**描述:** 获取Hosted订单的首次访问信息（用于前端展示）。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度≥1 |

**响应数据:** 同hostedOrderInfo，但包含更多前端展示所需的配置信息

## 11.4 创建Hosted订单充值

**接口:** `POST /createHostedOrderDeposit`

**描述:** 为Hosted订单选择代币和链后创建充值地址。

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
| confirmsNeeded | uint64 | 需要的确认数 |

## 11.5 获取Hosted代币USDT价格

**接口:** `POST /getHostedCoinUSDTPrice`

**描述:** 获取Hosted模式下代币的USDT价格。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| coinIds | Array<uint64> | 是 | 代币ID列表 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| prices | Object | 价格映射 |

## 11.6 获取主要代币列表

**接口:** `POST /getMainCoinList`

**描述:** 获取Hosted模式下支持的主要代币列表。

**请求参数:** 无

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| coins | Array | 代币列表 |
| coins[].coinId | uint64 | 代币ID |
| coins[].symbol | string | 代币符号 |
| coins[].logoUrl | string | Logo URL |
| coins[].chains | Array | 支持的链列表 |

## 11.7 创建Demo订单充值

**接口:** `POST /createAppDemoOrderDeposit`

**描述:** 创建Demo订单充值地址（用于测试）。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度3-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |
| chain | string | 是 | 链名称 | 长度≥1 |
| price | string | 是 | 价格 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| address | string | 充值地址 |
| amount | string | 充值金额 |
| memo | string | 备注（可选） |

## 11.8 确认交易

**接口:** `POST /confirmTrade`

**描述:** 确认Hosted订单的交易。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| orderId | string | 是 | 订单ID | 长度≥1 |
| txId | string | 是 | 交易哈希 | 长度≥1 |

**响应数据:** 空对象
