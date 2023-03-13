# ccpayment

## 1、 Modify file: src/main/java/com/Ccpayment/constant/Config.java,  set AppId, AppSecrete of your own

## 2、 User cases: src/test/java/com/Ccpayment/apis/CcpaymentApisTest.java

## Reference of functions
Specific parameters: Refer to the API documentation<br>
Document Address : https://doc.ccpayment.com/ccpayment-for-merchant/home
```azure
Due to the different command rules in different programming languages, there may be small hump, underscore, and other command methods, but the words are the same.
Use the function name of php as a reference:
```

| Name            | Description                                                                                                                                                                                                |
|-----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GetSupportToken | Obtain the token list supported by merchants                                                                                                                                                               |
| GetTokenChain   | Obtain the list of the available networks for a certain token                                                                                                                                              |
| CreateOrder     | Manage 100% of your front-end interactions and use our APIs to build your own checkout page.                                                                                                               |
| CheckoutUrl     | Simplify your integration by utilizing CCPayment's hosted checkout page. CCPayment will generate a checkout URL to merchant, it will direct customers to CCPayment hosted checkout page to make payment.   |
| Withdraw        | Call the withdrawal API to initiate withdrawals                                                                                                                                                            |
| CheckUser       | Check the Validity of Cwallet ID                                                                                                                                                                           |
| GetTokenRate    | The amount of USD converted into tokens                                                                                                                                                                    |
| Assets          | Obtain details of merchant's assets                                                                                                                                                                        |
| NetworkFee      | Obtain the network fee of a certain network                                                                                                                                                                |
| Webhook         | Notification of order callbacks                                                                                                                                                                            |
