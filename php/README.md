# php

## doc: https://doc.ccpayment.com/ccpayment-for-merchant/home
## Install via composer
````
composer require ccpayment/php-sdk
````
## 使用示例
```
<?php

use CCPayment\CCPay;

$resp =  CCPay::CheckUser("9454818","202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
var_dump($resp);

```