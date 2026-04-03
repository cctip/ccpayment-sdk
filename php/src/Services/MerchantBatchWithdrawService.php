<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class MerchantBatchWithdrawService
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
    public function checkWithdrawAddress(array $data): array
    {
        return $this->client->post('/checkWithdrawAddress', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function applyBatchWithdraw(array $data): array
    {
        return $this->client->post('/applyBatchWithdraw', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function appendBatchWithdraw(array $data): void
    {
        $this->client->post('/appendBatchWithdraw', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function confirmBatchWithdraw(array $data): array
    {
        return $this->client->post('/confirmBatchWithdraw', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function abortBatchWithdraw(array $data): array
    {
        return $this->client->post('/abortBatchWithdraw', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getBatchWithdrawOrder(array $data): array
    {
        return $this->client->post('/getBatchWithdrawOrder', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getBatchWithdrawOrderRecordList(array $data): array
    {
        return $this->client->post('/getBatchWithdrawOrderRecordList', $data);
    }
}
