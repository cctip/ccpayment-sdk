<?php

namespace CCPayment\Services;

use CCPayment\Client;

class UtilitiesService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function webhookResend(array $data = []): array { return $this->client->post('/webhook/resend', $data); }
    public function getTradeBlockHeight(array $data): array { return $this->client->post('/getTradeBlockHeight', $data); }
    public function checkWithdrawalAddressValidity(array $data): array { return $this->client->post('/checkWithdrawalAddressValidity', $data); }
    public function deprecatedAddress(array $data): array { return $this->client->post('/addressUnbinding', $data); }
    public function rescanLostTransaction(array $data): array { return $this->client->post('/rescanLostTransaction', $data); }
}
