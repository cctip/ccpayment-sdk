# CCPayment API v2 Complete Documentation

Last Updated: 2026-04-03

## Overview

This document is compiled based on the CCPayment API v2 proto definitions, containing detailed request parameters and response data descriptions for all public interfaces.

**Base URL:** `https://ccpayment.com/ccpayment/v2/`

**Authentication:** All API requests must include the following headers:
- `Appid`: Application ID  
- `Timestamp`: Current timestamp (milliseconds)
- `Sign`: Request signature

**Common Response Format:**
```json
{
  "code": 10000,
  "msg": "success",
  "data": {}
}
```

## Error Code Reference

For a complete list of error codes, please refer to the Code enum definition in the proto file. Common error codes:
- 10000: Success
- 11000: Parameter error
- 11005: Signature verification failed
- 14000: Insufficient balance
- 15002: Order expired

---

## API Module List

This document is organized by functional modules, containing a total of 13 modules:

### Merchant APIs

1. [Basic Info Module](./01-basic-info.md) - Token, fiat currency, and chain information queries
2. [Merchant Assets Module](./02-merchant-assets.md) - Merchant asset queries
3. [Merchant Deposit Module](./03-merchant-deposit.md) - Deposit address generation and deposit record queries
4. [Merchant Withdrawal Module](./04-merchant-withdraw.md) - Withdrawal requests and withdrawal record queries
5. [Merchant Batch Withdrawal Module](./05-merchant-batch-withdraw.md) - Batch withdrawal management

### User APIs

6. [User Assets Module](./06-user-assets.md) - User asset queries
7. [User Deposit Module](./07-user-deposit.md) - User deposit addresses and deposit records
8. [User Withdrawal Module](./08-user-withdraw.md) - User withdrawal requests and withdrawal records
9. [User Transfer Module](./09-user-transfer.md) - User transfers and batch transfers

### Orders & Payments

10. [Orders Module](./10-orders.md) - Order and Invoice orders
11. [Checkout Module](./11-checkout.md) - Checkout and Hosted related

### Other Features

12. [Swap Module](./12-swap.md) - Swap related interfaces
13. [Utilities](./13-utilities.md) - Webhook, address validation, etc.

---

## Appendix

- [Data Type Reference](./appendix.md#data-type-reference)
- [Status Enums](./appendix.md#status-enums)
- [Validation Rules](./appendix.md#validation-rules)
- [Important Notes](./appendix.md#important-notes)

---

**Document Version:** v2.0  
**Last Updated:** 2026-04-03  
**Based on:** Proto file definitions
