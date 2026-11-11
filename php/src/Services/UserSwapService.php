<?php

namespace CCPayment\Services;

use CCPayment\Client;

class UserSwapService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    /** @deprecated Use getUserSwapRecordV2(). */
    public function getUserSwapRecord(array $data): array { return $this->client->post('/getUserSwapRecord', $data); }
    /** @deprecated Use getUserSwapRecordListV2(). */
    public function getUserSwapRecordList(array $data = []): array { return $this->client->post('/getUserSwapRecordList', $data); }
    /** @deprecated Use userSwapV2() with SwapReqV2 fields. */
    public function userSwap(array $data): array { return $this->client->post('/userSwap', $data); }
    public function userSwapV2(array $data): array { return $this->client->post('/swapv2/userSwap', $data); }
    public function getUserSwapRecordV2(array $data): array { return $this->client->post('/swapv2/getUserSwapRecord', $data); }
    public function getUserSwapRecordListV2(array $data = []): array { return $this->client->post('/swapv2/getUserSwapRecordList', $data); }
}
