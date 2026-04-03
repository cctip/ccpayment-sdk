<?php

declare(strict_types=1);

require_once __DIR__ . '/../src/CCPaymentClient.php';

use CCPayment\SDK\CCPaymentClient;

$client = new CCPaymentClient(
    'https://admin.ccpayment.com',
    'your-app-id',
    'your-app-secret'
);

$result = $client->getSupportToken();
print_r($result);

$createOrderResult = $client->createOrder(
    '1280',
    '10.00',
    'merchant-order-' . time(),
    'USD'
);
print_r($createOrderResult);
