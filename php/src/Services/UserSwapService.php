<?php

namespace CCPayment\Services;

use CCPayment\Client;

class UserSwapService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getUserSwapRecord(array $data): array { return $this->client->post('/getUserSwapRecord', $data); }
    public function getUserSwapRecordList(array $data = []): array { return $this->client->post('/getUserSwapRecordList', $data); }
    public function userSwap(array $data): array { return $this->client->post('/userSwap', $data); }
}
