<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class SwapService
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
    public function getSwapCoinList(): array
    {
        return $this->client->post('/getSwapCoinList', []);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function swapEstimate(array $data): array
    {
        return $this->client->post('/swapEstimate', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function swap(array $data): array
    {
        return $this->client->post('/swap', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getSwapRecord(array $data): array
    {
        return $this->client->post('/getSwapRecord', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getSwapRecordList(array $data = []): array
    {
        return $this->client->post('/getSwapRecordList', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function selectHostedInvoiceCoin(array $data): array
    {
        return $this->client->post('/selectHostedInvoiceCoin', $data);
    }
}
