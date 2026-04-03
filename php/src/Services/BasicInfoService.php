<?php

namespace CCPayment\Services;

use CCPayment\Client;
use CCPayment\APIError;
use GuzzleHttp\Exception\GuzzleException;

class BasicInfoService
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
    public function getCoinList(): array
    {
        return $this->client->post('/getCoinList', []);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getCoin(int $coinId): array
    {
        return $this->client->post('/getCoin', ['coinId' => $coinId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getCoinUSDTPrice(array $coinIds): array
    {
        return $this->client->post('/getCoinUSDTPrice', ['coinIds' => $coinIds]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getFiatList(): array
    {
        return $this->client->post('/getFiatList', []);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getChainList(?array $chains = null): array
    {
        $data = $chains ? ['chains' => $chains] : [];
        return $this->client->post('/getChainList', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function allChain(?array $chains = null): array
    {
        $data = $chains ? ['chains' => $chains] : [];
        return $this->client->post('/all/chain', $data);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getCwalletUserId(string $cwalletUserId): array
    {
        return $this->client->post('/getCwalletUserId', ['cwalletUserId' => $cwalletUserId]);
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function getWithdrawFee(int $coinId, string $chain): array
    {
        return $this->client->post('/getWithdrawFee', ['coinId' => $coinId, 'chain' => $chain]);
    }
}
