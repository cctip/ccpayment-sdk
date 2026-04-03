# CCPayment API v2 完整文档

最后更新时间: 2026-04-03

## 概述

本文档基于CCPayment API v2的proto定义整理，包含所有公开接口的详细请求参数和响应数据说明。

**Base URL:** `https://ccpayment.com/ccpayment/v2/`

**认证方式:** 所有API请求需要在Header中包含：
- `Appid`: 应用ID  
- `Timestamp`: 当前时间戳（毫秒）
- `Sign`: 请求签名

**通用响应格式:**
```json
{
  "code": 10000,
  "msg": "success",
  "data": {}
}
```

## 错误码说明

完整错误码列表请参考proto文件中的Code枚举定义。常见错误码：
- 10000: 成功
- 11000: 参数错误
- 11005: 签名验证失败
- 14000: 余额不足
- 15002: 订单过期

---

## API模块列表

本文档按功能模块组织，共包含13个模块：

### 商家端接口

1. [基础信息模块](./01-basic-info.md) - 代币、法币、链信息查询
2. [商家资产模块](./02-merchant-assets.md) - 商家资产查询
3. [商家充值模块](./03-merchant-deposit.md) - 充值地址生成、充值记录查询
4. [商家提现模块](./04-merchant-withdraw.md) - 提现申请、提现记录查询
5. [商家批量提现模块](./05-merchant-batch-withdraw.md) - 批量提现管理

### 用户端接口

6. [用户资产模块](./06-user-assets.md) - 用户资产查询
7. [用户充值模块](./07-user-deposit.md) - 用户充值地址、充值记录
8. [用户提现模块](./08-user-withdraw.md) - 用户提现申请、提现记录
9. [用户转账模块](./09-user-transfer.md) - 用户转账、批量转账

### 订单与支付

10. [订单模块](./10-orders.md) - Order和Invoice订单
11. [收银台模块](./11-checkout.md) - Checkout和Hosted相关

### 其他功能

12. [换币模块](./12-swap.md) - Swap相关接口
13. [工具接口](./13-utilities.md) - Webhook、地址验证等

---

## 附录

- [数据类型说明](./appendix.md#数据类型说明)
- [状态枚举](./appendix.md#状态枚举)
- [验证规则说明](./appendix.md#验证规则说明)
- [注意事项](./appendix.md#注意事项)

---

**文档版本:** v2.0  
**最后更新:** 2026-04-03  
**基于:** proto文件定义
