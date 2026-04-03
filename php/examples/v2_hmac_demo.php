<?php

declare(strict_types=1);

require_once __DIR__ . '/../src/CCPaymentClient.php';

use CCPayment\SDK\CCPaymentClient;

$client = new CCPaymentClient(
    'https://admin.ccpayment.com',
    'your-app-id',
    'your-app-secret'
);

$result = $client->getCoinListV2([
    'status' => 1,
]);
print_r($result);

$checkoutResult = $client->createCheckoutUrlV2([
    'orderId' => 'demo-order-' . time(),
    'price' => '10',
    'coinId' => 1280,
], 'hmac');
print_r($checkoutResult);
