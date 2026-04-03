# Checkout Module

## 11.1 Create Checkout URL

**Interface:** `POST /createCheckoutUrl`

**Description:** Create a Hosted Checkout payment URL.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| product | string | Yes | Product description | Length ≥1 |
| priceFiatId | uint64 | No | Fiat currency ID | - |
| priceCoinId | uint64 | No | Token ID | - |
| price | string | Yes | Price | Length ≥1 |
| buyerEmail | string | No | Buyer email | Email format |
| returnUrl | string | No | Return URL | Max 150 characters, URI format |
| expiredAt | int64 | No | Expiration timestamp | ≥1 |
| closeUrl | string | No | Close URL | Max 150 characters, URI format |
| notifyUrl | string | No | Notification URL | Max 150 characters, URI format |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| checkoutUrl | string | Checkout payment URL |

## 11.2 Get Hosted Order Info

**Interface:** `POST /hostedOrderInfo`

**Description:** Get Hosted order details.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| orderId | string | Order ID |
| product | string | Product description |
| price | string | Price |
| priceCoinId | uint64 | Token ID |
| priceFiatId | uint64 | Fiat currency ID |
| priceSymbol | string | Price symbol |
| checkoutUrl | string | Checkout URL |
| buyerEmail | string | Buyer email |
| expiredAt | int64 | Expiration time |
| step | string | Current step |
| selectedCoinId | uint64 | Selected token ID |
| selectedChain | string | Selected chain |
| toAddress | string | Receiving address |
| toMemo | string | Memo |
| amountToPay | string | Amount to pay |
| rate | string | Exchange rate |
| totalPaidValue | string | Total paid value |
| paidList | Array | Paid list |
| refundList | Array | Refund list |

## 11.3 Get Hosted Order First Info

**Interface:** `POST /hostedOrderInfoFirst`

**Description:** Get the first visit information of a Hosted order (for frontend display).

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |

**Response Data:** Same as hostedOrderInfo, but includes more configuration information needed for frontend display

## 11.4 Create Hosted Order Deposit

**Interface:** `POST /createHostedOrderDeposit`

**Description:** Create a deposit address after selecting token and chain for a Hosted order.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| chain | string | Yes | Chain name | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| address | string | Deposit address |
| memo | string | Memo (optional) |
| amountToPay | string | Amount to pay |
| confirmsNeeded | uint64 | Required confirmation count |

## 11.5 Get Hosted Coin USDT Price

**Interface:** `POST /getHostedCoinUSDTPrice`

**Description:** Get the USDT price of a token in Hosted mode.

**Request Parameters:**

| Field | Type | Required | Description |
|------|------|------|------|
| coinIds | Array<uint64> | Yes | Token ID list |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| prices | Object | Price map |

## 11.6 Get Main Coin List

**Interface:** `POST /getMainCoinList`

**Description:** Get the list of main supported tokens in Hosted mode.

**Request Parameters:** None

**Response Data:**

| Field | Type | Description |
|------|------|------|
| coins | Array | Token list |
| coins[].coinId | uint64 | Token ID |
| coins[].symbol | string | Token symbol |
| coins[].logoUrl | string | Logo URL |
| coins[].chains | Array | Supported chain list |

## 11.7 Create Demo Order Deposit

**Interface:** `POST /createAppDemoOrderDeposit`

**Description:** Create a Demo order deposit address (for testing).

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length 3-64 |
| coinId | uint64 | Yes | Token ID | ≥1 |
| chain | string | Yes | Chain name | Length ≥1 |
| price | string | Yes | Price | Length ≥1 |

**Response Data:**

| Field | Type | Description |
|------|------|------|
| address | string | Deposit address |
| amount | string | Deposit amount |
| memo | string | Memo (optional) |

## 11.8 Confirm Trade

**Interface:** `POST /confirmTrade`

**Description:** Confirm the trade of a Hosted order.

**Request Parameters:**

| Field | Type | Required | Description | Validation Rules |
|------|------|------|------|----------|
| orderId | string | Yes | Order ID | Length ≥1 |
| txId | string | Yes | Transaction hash | Length ≥1 |

**Response Data:** Empty object
