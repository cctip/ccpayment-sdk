<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class MerchantAssetsService
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
    public function getAppCoinAssetList(): array
    {
        return $this->client->post('/getAppCoinAssetList', []);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getAppCoinAsset(int $coinId): array
    {
        return $this->client->post('/getAppCoinAsset', ['coinId' => $coinId]);
    }
}
