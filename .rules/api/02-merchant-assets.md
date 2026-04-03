# 商家资产模块

## 2.1 获取全部资产

**接口:** `POST /getAppCoinAssetList`

**描述:** 获取商家的所有代币资产。

**请求参数:** 无

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| assets | Array | 资产列表 |
| assets[].coinId | uint64 | 代币ID |
| assets[].coinSymbol | string | 代币符号 |
| assets[].available | string | 可用余额 |

## 2.2 获取单个币资产

**接口:** `POST /getAppCoinAsset`

**描述:** 获取商家指定代币的资产。

**请求参数:**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| coinId | uint64 | 是 | 代币ID |

**响应数据:**

| 字段 | 类型 | 说明 |
|------|------|------|
| asset | Object | 资产信息 |
| asset.coinId | uint64 | 代币ID |
| asset.coinSymbol | string | 代币符号 |
| asset.available | string | 可用余额 |
