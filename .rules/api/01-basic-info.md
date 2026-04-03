# Basic Info Module

## 1.1 Get Token List

**Interface:** `POST /getCoinList`

**Description:** Get all supported token information on the platform, including network configurations and prices.

**Request Parameters:** None

**Response Data:**

| Field | Type | Description |
|------|------|------|
| coins | Array | Token list |
| coins[].coinId | uint64 | Token ID |
| coins[].symbol | string | Token symbol |
| coins[].coinFullName | string | Token full name |
| coins[].logoUrl | string | Logo URL |
| coins[].status | string | Status (Normal/Maintain) |
| coins[].networks | Object | Network information map (key is chain name) |
| coins[].networks[].chain | string | Chain name |
| coins[].networks[].chainFullName | string | Chain full name |
| coins[].networks[].contract | string | Contract address |
| coins[].networks[].precision | uint32 | Precision |
| coins[].networks[].canDeposit | bool | Whether deposit is supported |
| coins[].networks[].canWithdraw | bool | Whether withdrawal is supported |
| coins[].networks[].minimumDepositAmount | string | Minimum deposit amount |
| coins[].networks[].minimumWithdrawAmount | string | Minimum withdrawal amount |
| coins[].networks[].maximumWithdrawAmount | string | Maximum withdrawal amount |
| coins[].networks[].isSupportMemo | bool | Whether Memo is supported |
| coins[].networks[].protocol | string | Protocol type |
| coins[].price | string | USD price |

## 1.2 Get Single Token

**Interface:** `POST /getCoin`

**Description:** Get detailed information for a specified token.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| coinId | uint64 | Yes | Token ID | ≥1 |

**Response Data:** Same as individual coin object in getCoinList

## 1.3 Get Token USD Price

**Interface:** `POST /getCoinUSDTPrice`

**Description:** Get USD prices for tokens in batch.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| coinIds | Array<uint64> | Yes | Token ID list |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| prices | Object | Price map (key is coinId, value is price string) |

## 1.4 Get Fiat Currency List

**Interface:** `POST /getFiatList`

**Description:** Get all supported fiat currency information on the platform.

**Request Parameters:** None

**Response Data:**

| Field | Type | Description |
|------|------|------|
| fiats | Array | Fiat currency list |
| fiats[].fiatId | uint64 | Fiat currency ID |
| fiats[].symbol | string | Fiat currency symbol |
| fiats[].logoUrl | string | Logo URL |
| fiats[].mark | string | Mark |
| fiats[].usdRate | string | USD exchange rate |

## 1.5 Get Chain List

**Interface:** `POST /getChainList`

**Description:** Query status information for specified chains. If chains parameter is not provided, query all chains.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| chains | Array<string> | No | Chain name list | Each item length 1-16 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| chains | Array | Chain information list |
| chains[].chain | string | Chain name |
| chains[].chainFullName | string | Chain full name |
| chains[].explorer | string | Block explorer URL |
| chains[].logoUrl | string | Logo URL |
| chains[].status | string | Status (Normal/Maintain) |
| chains[].confirmations | int32 | Confirmation count |
| chains[].baseUrl | string | Base URL |
| chains[].isEVM | bool | Whether it's an EVM chain |
| chains[].supportMemo | bool | Whether Memo is supported |

## 1.6 Get All Chain List

**Interface:** `POST /all/chain`

**Description:** Get information for all chains (simplified version).

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| chains | Array<string> | No | Chain name list | Each item length 1-16 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| chains | Array | Chain information list |
| chains[].chain | string | Chain name |
| chains[].chainFullName | string | Chain full name |
| chains[].explorer | string | Block explorer URL |
| chains[].logoUrl | string | Logo URL |
| chains[].status | string | Status |
| chains[].confirmNum | int32 | Confirmation count |
| chains[].isEVM | bool | Whether it's an EVM chain |

## 1.7 Get CCWallet User ID

**Interface:** `POST /getCwalletUserId`

**Description:** Verify and get CCWallet user information.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| cwalletUserId | string | Yes | CCWallet user ID | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| cwalletUserId | string | CCWallet user ID |
| cwalletUserName | string | CCWallet username |

## 1.8 Get Withdrawal Fee

**Interface:** `POST /getWithdrawFee`

**Description:** Get the withdrawal fee for a specified token and chain.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| coinId | uint64 | Yes | Token ID | ≥1 |
| chain | string | Yes | Chain name | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| fee | Object | Fee information |
| fee.coinId | uint64 | Token ID |
| fee.coinSymbol | string | Token symbol |
| fee.amount | string | Fee amount |
| networkFeeInquiryID | string | Network fee inquiry ID |
