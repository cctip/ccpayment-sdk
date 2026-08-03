<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class MerchantDepositService
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
    public function createAppOrderDepositAddress(array $data): array
    {
        return $this->client->post('/createAppOrderDepositAddress', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getOrCreateAppDepositAddress(array $data): array
    {
        return $this->client->post('/getOrCreateAppDepositAddress', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getAppDepositRecord(string $recordId): array
    {
        return $this->client->post('/getAppDepositRecord', ['recordId' => $recordId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getAppDepositRecordList(array $data = []): array
    {
        return $this->client->post('/getAppDepositRecordList', $data);
    }
}
