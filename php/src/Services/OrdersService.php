<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class OrdersService
{
    private Client $client;

    public function __construct(Client $client)
    {
        $this->client = $client;
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getAppOrderInfo(string $orderId): array
    {
        return $this->client->post('/getAppOrderInfo', ['orderId' => $orderId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function createInvoiceUrl(array $data): array
    {
        return $this->client->post('/createInvoiceUrl', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getInvoiceOrderInfo(string $orderId): array
    {
        return $this->client->post('/getInvoiceOrderInfo', ['orderId' => $orderId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getWebhookInfo(): array
    {
        return $this->client->post('/getWebhookInfo', []);
    }
}
