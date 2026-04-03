<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class MerchantWithdrawService
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
    public function applyAppWithdrawToNetwork(array $data): array
    {
        return $this->client->post('/applyAppWithdrawToNetwork', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function applyAppWithdrawToCwallet(array $data): array
    {
        return $this->client->post('/applyAppWithdrawToCwallet', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getAppWithdrawRecord(array $data): array
    {
        return $this->client->post('/getAppWithdrawRecord', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getAppWithdrawRecordList(array $data = []): array
    {
        return $this->client->post('/getAppWithdrawRecordList', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getAutoWithdrawRecordList(array $data = []): array
    {
        return $this->client->post('/getAutoWithdrawRecordList', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getRiskyRefundRecordList(array $data = []): array
    {
        return $this->client->post('/getRiskyRefundRecordList', $data);
    }
}
