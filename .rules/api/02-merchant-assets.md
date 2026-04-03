# Merchant Assets Module

## 2.1 Get All Assets

**Interface:** `POST /getAppCoinAssetList`

**Description:** Get all token assets of the merchant.

**Request Parameters:** None

**Response Data:**

| Field | Type | Description |
|------|------|------|
| assets | Array | Asset list |
| assets[].coinId | uint64 | Token ID |
| assets[].coinSymbol | string | Token symbol |
| assets[].available | string | Available balance |

## 2.2 Get Single Token Asset

**Interface:** `POST /getAppCoinAsset`

**Description:** Get the asset of a specified token for the merchant.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| coinId | uint64 | Yes | Token ID |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| asset | Object | Asset information |
| asset.coinId | uint64 | Token ID |
| asset.coinSymbol | string | Token symbol |
| asset.available | string | Available balance |
