<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class UserDepositService
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
    public function getOrCreateUserDepositAddress(array $data): array
    {
        return $this->client->post('/getOrCreateUserDepositAddress', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserDepositRecord(string $recordId): array
    {
        return $this->client->post('/getUserDepositRecord', ['recordId' => $recordId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserDepositRecordList(array $data): array
    {
        return $this->client->post('/getUserDepositRecordList', $data);
    }
}
