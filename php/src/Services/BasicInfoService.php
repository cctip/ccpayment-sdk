<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class BasicInfoService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getCoinList(): array { return $this->client->post('/getCoinList', []); }
    public function getFiatList(): array { return $this->client->post('/getFiatList', []); }
    public function getCoin(int $coinId): array { return $this->client->post('/getCoin', ['coinId' => $coinId]); }
    public function getCoinUSDTPrice(array $coinIds): array { return $this->client->post('/getCoinUSDTPrice', ['coinIds' => $coinIds]); }
    public function getChainList(array $data = []): array { return $this->client->post('/getChainList', $data); }
    public function getAllChainList(array $data = []): array { return $this->client->post('/all/chain', $data); }
    public function getMainCoinList(string $appId): array { return $this->client->post('/getMainCoinList', ['appId' => $appId]); }
}
