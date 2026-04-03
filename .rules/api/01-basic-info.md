# 基础信息模块

## 1.1 获取代币列表

**接口:** `POST /getCoinList`

**描述:** 获取平台支持的所有代币信息，包括网络配置和价格。

**请求参数:** 无

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| coins | Array | 代币列表 |
| coins[].coinId | uint64 | 代币ID |
| coins[].symbol | string | 代币符号 |
| coins[].coinFullName | string | 代币全称 |
| coins[].logoUrl | string | Logo URL |
| coins[].status | string | 状态（Normal/Maintain） |
| coins[].networks | Object | 网络信息映射（key为链名称） |
| coins[].networks[].chain | string | 链名称 |
| coins[].networks[].chainFullName | string | 链全称 |
| coins[].networks[].contract | string | 合约地址 |
| coins[].networks[].precision | uint32 | 精度 |
| coins[].networks[].canDeposit | bool | 是否支持充值 |
| coins[].networks[].canWithdraw | bool | 是否支持提现 |
| coins[].networks[].minimumDepositAmount | string | 最小充值金额 |
| coins[].networks[].minimumWithdrawAmount | string | 最小提现金额 |
| coins[].networks[].maximumWithdrawAmount | string | 最大提现金额 |
| coins[].networks[].isSupportMemo | bool | 是否支持Memo |
| coins[].networks[].protocol | string | 协议类型 |
| coins[].price | string | USD价格 |

## 1.2 获取单个代币

**接口:** `POST /getCoin`

**描述:** 获取指定代币的详细信息。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| coinId | uint64 | 是 | 代币ID | ≥1 |

**响应数据:** 同getCoinList中的单个coin对象

## 1.3 获取代币USD价格

**接口:** `POST /getCoinUSDTPrice`

**描述:** 批量获取代币的USD价格。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| coinIds | Array<uint64> | 是 | 代币ID列表 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| prices | Object | 价格映射（key为coinId，value为价格字符串） |

## 1.4 获取法币列表

**接口:** `POST /getFiatList`

**描述:** 获取平台支持的所有法币信息。

**请求参数:** 无

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| fiats | Array | 法币列表 |
| fiats[].fiatId | uint64 | 法币ID |
| fiats[].symbol | string | 法币符号 |
| fiats[].logoUrl | string | Logo URL |
| fiats[].mark | string | 标记 |
| fiats[].usdRate | string | 对USD汇率 |

## 1.5 获取链列表

**接口:** `POST /getChainList`

**描述:** 查询指定链的状态信息，不传chains参数则查询所有链。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| chains | Array<string> | 否 | 链名称列表 | 每项长度1-16 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| chains | Array | 链信息列表 |
| chains[].chain | string | 链名称 |
| chains[].chainFullName | string | 链全称 |
| chains[].explorer | string | 区块浏览器URL |
| chains[].logoUrl | string | Logo URL |
| chains[].status | string | 状态（Normal/Maintain） |
| chains[].confirmations | int32 | 确认数 |
| chains[].baseUrl | string | 基础URL |
| chains[].isEVM | bool | 是否为EVM链 |
| chains[].supportMemo | bool | 是否支持Memo |

## 1.6 获取所有链列表

**接口:** `POST /all/chain`

**描述:** 获取所有链的信息（简化版）。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| chains | Array<string> | 否 | 链名称列表 | 每项长度1-16 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| chains | Array | 链信息列表 |
| chains[].chain | string | 链名称 |
| chains[].chainFullName | string | 链全称 |
| chains[].explorer | string | 区块浏览器URL |
| chains[].logoUrl | string | Logo URL |
| chains[].status | string | 状态 |
| chains[].confirmNum | int32 | 确认数 |
| chains[].isEVM | bool | 是否为EVM链 |

## 1.7 获取CCWallet用户ID

**接口:** `POST /getCwalletUserId`

**描述:** 验证并获取CCWallet用户信息。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| cwalletUserId | string | 是 | CCWallet用户ID | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| cwalletUserId | string | CCWallet用户ID |
| cwalletUserName | string | CCWallet用户名 |

## 1.8 获取提现手续费

**接口:** `POST /getWithdrawFee`

**描述:** 获取指定代币和链的提现手续费。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| coinId | uint64 | 是 | 代币ID | ≥1 |
| chain | string | 是 | 链名称 | 长度≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| fee | Object | 手续费信息 |
| fee.coinId | uint64 | 代币ID |
| fee.coinSymbol | string | 代币符号 |
| fee.amount | string | 手续费金额 |
| networkFeeInquiryID | string | 网络手续费查询ID |
