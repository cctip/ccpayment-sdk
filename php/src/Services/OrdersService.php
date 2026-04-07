<?php

namespace CCPayment\Services;

use CCPayment\Client;

class OrdersService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getAppOrderInfo(string $orderId): array { return $this->client->post('/getAppOrderInfo', ['orderId' => $orderId]); }
    public function createInvoiceUrl(array $data): array { return $this->client->post('/createInvoiceUrl', $data); }
    public function getInvoiceOrderInfo(string $orderId): array { return $this->client->post('/getInvoiceOrderInfo', ['orderId' => $orderId]); }
}
