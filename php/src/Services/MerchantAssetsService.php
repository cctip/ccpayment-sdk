<?php

namespace CCPayment\Services;

use CCPayment\Client;

class MerchantAssetsService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getAppCoinAssetList(): array { return $this->client->post('/getAppCoinAssetList', []); }
    public function getAppCoinAsset(int $coinId): array { return $this->client->post('/getAppCoinAsset', ['coinId' => $coinId]); }
    public function getAppCollectFeeRecordList(array $data = []): array { return $this->client->post('/getAggregationFeeRecord', $data); }
}
