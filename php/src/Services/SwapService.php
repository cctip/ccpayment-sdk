<?php

namespace CCPayment\Services;

use CCPayment\Client;

class SwapService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getSwapCoinList(): array { return $this->client->post('/getSwapCoinList', []); }
    public function swapEstimate(array $data): array { return $this->client->post('/estimate', $data); }
    public function getSwapRecord(array $data): array { return $this->client->post('/getSwapRecord', $data); }
    public function getSwapRecordList(array $data = []): array { return $this->client->post('/getSwapRecordList', $data); }
    public function swap(array $data): array { return $this->client->post('/swap', $data); }
}
