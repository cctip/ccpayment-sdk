# CCPayment SDK

Official multi-language SDKs for [CCPayment](https://ccpayment.com) API v2.

## Languages

| Language | Path | Documentation |
|----------|------|----------------|
| Go | [`golang/`](golang/) | [README](golang/README.md) |
| Python | [`python/`](python/) | [README](python/README.md) |
| TypeScript | [`typescript/`](typescript/) | [README](typescript/README.md) |
| JavaScript | [`javascript/`](javascript/) | [README](javascript/README.md) |
| Java | [`java/`](java/) | [README](java/README.md) |
| PHP | [`php/`](php/) | [README](php/README.md) |
| Ruby | [`ruby/`](ruby/) | [README](ruby/README.md) |

## Documentation

- [API Documentation](https://ccpayment.com/api/doc)
- [Developer Console](https://console.ccpayment.com)

## Swap V2

All seven language SDKs expose the eight `/swapv2/*` routes through the existing swap and user-swap services. The Swap V1 methods remain callable but are deprecated. Use the V2 methods for new integrations.

| Swap service | User-swap service |
| --- | --- |
| `getSwapPairList` | `userSwapV2` |
| `swapEstimateV2` | `getUserSwapRecordV2` |
| `swapV2` | `getUserSwapRecordListV2` |
| `getSwapRecordV2` | |
| `getSwapRecordListV2` | |

Go uses exported names such as `SwapV2`; Python and Ruby use snake_case names such as `swap_v2`. V2 order and estimate requests use `fromCoinId`, `fromAmount`, `toCoinId`, optional `markupFeeRate`, and optional `slippageRate`. These fields have different meanings from V1's `coinIdIn`, `amountIn`, and `coinIdOut`; migrate request bodies explicitly.

## Disclaimer

This SDK is provided by CCPayment on an "as is" and "as available" basis for use at your own risk. CCPayment makes no representations or warranties of any kind, whether express or implied, as to the operation of the SDK, its accuracy, reliability, completeness, or fitness for any particular purpose.

To the fullest extent permitted by law, CCPayment shall not be liable for any losses, damages, or expenses of any kind arising from or in connection with your use of, or inability to use, this SDK, including but not limited to any financial losses resulting from errors, bugs, interruptions, or inaccuracies in the SDK.

Your use of this SDK to access the CCPayment platform is subject to the [CCPayment Terms of Use](https://ccpayment.com/en/privacy-policy#ccpayment-service-agreement), which shall prevail in the event of any conflict with this disclaimer. You are solely responsible for the transactions and API operations executed through your merchant account using this SDK. For a secure integration, please follow the verification practices described in the [API documentation](https://ccpayment.com/api/doc?en).

Plugins or extensions developed by third parties based on this SDK or the CCPayment API are independent works of their authors. Unless expressly designated as official in CCPayment's documentation, they are not maintained by CCPayment, and CCPayment is not responsible for their operation. We welcome community contributions — developers building on CCPayment are invited to reach out at partner@ccpayment.com.

This SDK is not intended to constitute investment advice or a recommendation to buy, sell, or hold any digital asset. You should independently evaluate and verify all information before acting.
