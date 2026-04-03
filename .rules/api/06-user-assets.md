# User Assets Module

## 6.1 Get User Asset List

**Interface:** `POST /getUserCoinAssetList`

**Description:** Get all token assets of a specified user.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| userId | string | Yes | User ID | Length 5-64 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| userId | string | User ID |
| assets | Array | Asset list |
| assets[].coinId | uint64 | Token ID |
| assets[].coinSymbol | string | Token symbol |
| assets[].available | string | Available balance |

## 6.2 Get User Asset

**Interface:** `POST /getUserCoinAsset`

**Description:** Get a specific token asset of a specified user.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| userId | string | Yes | User ID | Length 5-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| userId | string | User ID |
| asset | Object | Asset information |
| asset.coinId | uint64 | Token ID |
| asset.coinSymbol | string | Token symbol |
| asset.available | string | Available balance |
