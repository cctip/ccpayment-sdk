# 用户资产模块

## 6.1 获取用户资产列表

**接口:** `POST /getUserCoinAssetList`

**描述:** 获取指定用户的所有代币资产。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| userId | string | 是 | 用户ID | 长度5-64 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| userId | string | 用户ID |
| assets | Array | 资产列表 |
| assets[].coinId | uint64 | 代币ID |
| assets[].coinSymbol | string | 代币符号 |
| assets[].available | string | 可用余额 |

## 6.2 获取用户资产

**接口:** `POST /getUserCoinAsset`

**描述:** 获取指定用户的特定代币资产。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 | 验证规则 |
|------|------|------|------|----------|
| userId | string | 是 | 用户ID | 长度5-64 |
| coinId | uint64 | 是 | 代币ID | ≥1 |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| userId | string | 用户ID |
| asset | Object | 资产信息 |
| asset.coinId | uint64 | 代币ID |
| asset.coinSymbol | string | 代币符号 |
| asset.available | string | 可用余额 |
