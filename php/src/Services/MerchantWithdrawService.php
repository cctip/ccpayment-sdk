<?php

namespace CCPayment\Services;

use CCPayment\Client;

class MerchantWithdrawService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getCwalletUserId(string $cwalletUserId): array { return $this->client->post('/getCwalletUserId', ['cwalletUserId' => $cwalletUserId]); }
    public function getWithdrawFee(int $coinId, string $chain): array { return $this->client->post('/getWithdrawFee', ['coinId' => $coinId, 'chain' => $chain]); }
    public function applyAppWithdrawToNetwork(array $data): array { return $this->client->post('/applyAppWithdrawToNetwork', $data); }
    public function applyAppWithdrawToCwallet(array $data): array { return $this->client->post('/applyAppWithdrawToCwallet', $data); }
    public function getAppWithdrawRecord(array $data): array { return $this->client->post('/getAppWithdrawRecord', $data); }
    public function getAppWithdrawRecordList(array $data = []): array { return $this->client->post('/getAppWithdrawRecordList', $data); }
    public function getAutoWithdrawRecordList(array $data = []): array { return $this->client->post('/getAutoWithdrawRecordList', $data); }
    public function getRiskRefundRecords(array $data = []): array { return $this->client->post('/getRiskyRefundRecordList', $data); }
}
