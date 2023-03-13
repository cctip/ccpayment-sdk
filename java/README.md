# ccpayment

## 1、 Modify file: src/main/java/com/Ccpayment/constant/Config.java,  set AppId, AppSecrete of your own

## 2、 User cases: src/test/java/com/Ccpayment/apis/CcpaymentApisTest.java

## Reference of functions
Specific parameters: Refer to the API documentation<br>
Document Address : https://doc.ccpayment.com/ccpayment-for-merchant/home


| Name            | Description                                                                                                                                                                                                |
|-----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| getSupportToken | Obtain the token list supported by merchants                                                                                                                                                               |
| getTokenChain   | Obtain the list of the available networks for a certain token                                                                                                                                              |
| createOrder     | Manage 100% of your front-end interactions and use our APIs to build your own checkout page.                                                                                                               |
| checkoutUrl     | Simplify your integration by utilizing CCPayment's hosted checkout page. CCPayment will generate a checkout URL to merchant, it will direct customers to CCPayment hosted checkout page to make payment.   |
| withdraw        | Call the withdrawal API to initiate withdrawals                                                                                                                                                            |
| checkUser       | Check the Validity of Cwallet ID                                                                                                                                                                           |
| getTokenRate    | The amount of USD converted into tokens                                                                                                                                                                    |
| assets          | Obtain details of merchant's assets                                                                                                                                                                        |
| networkFee      | Obtain the network fee of a certain network                                                                                                                                                                |
| webhook         | Notification of order callbacks                                                                                                                                                                            |
