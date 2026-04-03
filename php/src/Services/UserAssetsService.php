<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class UserAssetsService
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
    public function getUserCoinAssetList(string $userId): array
    {
        return $this->client->post('/getUserCoinAssetList', ['userId' => $userId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getUserCoinAsset(string $userId, int $coinId): array
    {
        return $this->client->post('/getUserCoinAsset', ['userId' => $userId, 'coinId' => $coinId]);
    }
}
