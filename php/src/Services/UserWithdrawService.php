<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class UserWithdrawService
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
    public function applyUserWithdrawToNetwork(array $data): array
    {
        return $this->client->post('/applyUserWithdrawToNetwork', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function applyUserWithdrawToCwallet(array $data): array
    {
        return $this->client->post('/applyUserWithdrawToCwallet', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserWithdrawRecord(array $data): array
    {
        return $this->client->post('/getUserWithdrawRecord', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserWithdrawRecordList(array $data): array
    {
        return $this->client->post('/getUserWithdrawRecordList', $data);
    }
}
