<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class UserTransferService
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
    public function userTransfer(array $data): array
    {
        return $this->client->post('/userTransfer', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserTransferRecord(array $data): array
    {
        return $this->client->post('/getUserTransferRecord', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserTransferRecordList(array $data = []): array
    {
        return $this->client->post('/getUserTransferRecordList', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function userBatchTransfer(array $data): array
    {
        return $this->client->post('/userBatchTransfer', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserBatchTransferRecord(array $data): array
    {
        return $this->client->post('/getUserBatchTransferRecord', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserBatchTransferRecordList(array $data = []): array
    {
        return $this->client->post('/getUserBatchTransferRecordList', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserBatchTransferRecordDetail(array $data): array
    {
        return $this->client->post('/getUserBatchTransferRecordDetail', $data);
    }
}
