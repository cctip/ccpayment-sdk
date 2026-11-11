<?php

namespace CCPayment\Services;

use CCPayment\Client;

class SwapService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    /** @deprecated Use getSwapPairList(). */
    public function getSwapCoinList(): array { return $this->client->post('/getSwapCoinList', []); }
    /** @deprecated Use swapEstimateV2() with SwapReqV2 fields. */
    public function swapEstimate(array $data): array { return $this->client->post('/estimate', $data); }
    /** @deprecated Use getSwapRecordV2(). */
    public function getSwapRecord(array $data): array { return $this->client->post('/getSwapRecord', $data); }
    /** @deprecated Use getSwapRecordListV2(). */
    public function getSwapRecordList(array $data = []): array { return $this->client->post('/getSwapRecordList', $data); }
    /** @deprecated Use swapV2() with SwapReqV2 fields. */
    public function swap(array $data): array { return $this->client->post('/swap', $data); }
    public function getSwapPairList(): array { return $this->client->post('/swapv2/getSwapPairList', []); }
    public function swapEstimateV2(array $data): array { return $this->client->post('/swapv2/estimate', $data); }
    public function swapV2(array $data): array { return $this->client->post('/swapv2/swap', $data); }
    public function getSwapRecordV2(array $data): array { return $this->client->post('/swapv2/getSwapRecord', $data); }
    public function getSwapRecordListV2(array $data = []): array { return $this->client->post('/swapv2/getSwapRecordList', $data); }
}
