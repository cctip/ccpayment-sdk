<?php

declare(strict_types=1);

require_once __DIR__ . '/../src/CCPaymentClient.php';

use CCPayment\SDK\CCPaymentClient;

$privateKeyPem = <<<PEM
-----BEGIN PRIVATE KEY-----
YOUR_PRIVATE_KEY
-----END PRIVATE KEY-----
PEM;

$client = new CCPaymentClient(
    'https://admin.ccpayment.com',
    'your-app-id',
    null,
    $privateKeyPem
);

$result = $client->createCheckoutUrlV2([
    'orderId' => 'rsa-order-' . time(),
    'price' => '10',
    'coinId' => 1280,
], 'rsa');

print_r($result);
