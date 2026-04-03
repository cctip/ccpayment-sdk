<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class CheckoutService
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
    public function createCheckoutUrl(array $data): array
    {
        return $this->client->post('/createCheckoutUrl', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function hostedOrderInfo(string $orderId): array
    {
        return $this->client->post('/hostedOrderInfo', ['orderId' => $orderId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function hostedOrderInfoFirst(string $orderId): array
    {
        return $this->client->post('/hostedOrderInfoFirst', ['orderId' => $orderId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function createHostedOrderDeposit(array $data): array
    {
        return $this->client->post('/createHostedOrderDeposit', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getHostedCoinUSDTPrice(array $coinIds): array
    {
        return $this->client->post('/getHostedCoinUSDTPrice', ['coinIds' => $coinIds]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getMainCoinList(): array
    {
        return $this->client->post('/getMainCoinList', []);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function createAppDemoOrderDeposit(array $data): array
    {
        return $this->client->post('/createAppDemoOrderDeposit', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function confirmTrade(array $data): void
    {
        $this->client->post('/confirmTrade', $data);
    }
}
