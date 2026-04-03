<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class UtilitiesService
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
    public function verifyAddress(string $chain, string $address): array
    {
        return $this->client->post('/verifyAddress', ['chain' => $chain, 'address' => $address]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function abandonAddress(string $chain, string $address): void
    {
        $this->client->post('/abandonAddress', ['chain' => $chain, 'address' => $address]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function hostedInvoiceOrderInfo(string $orderId): array
    {
        return $this->client->post('/hostedInvoiceOrderInfo', ['orderId' => $orderId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getPayInfo(string $orderId): array
    {
        return $this->client->post('/getPayInfo', ['orderId' => $orderId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function health(): array
    {
        return $this->client->post('/health', []);
    }
}
