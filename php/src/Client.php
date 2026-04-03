<?php

namespace CCPayment;

use GuzzleHttp\Client as GuzzleClient;
use GuzzleHttp\Exception\GuzzleException;

class Client
{
    private const DEFAULT_BASE_URL = 'https://ccpayment.com/ccpayment/v2';

    private string $appId;
    private string $appSecret;
    private string $baseUrl;
    private GuzzleClient $httpClient;

    public function __construct(string $appId, string $appSecret)
    {
        $this->appId = $appId;
        $this->appSecret = $appSecret;
        $this->baseUrl = self::DEFAULT_BASE_URL;
        $this->httpClient = new GuzzleClient();
    }

    public function setBaseUrl(string $baseUrl): void
    {
        $this->baseUrl = $baseUrl;
    }

    public function setProxy(string $proxyUrl): void
    {
        $this->httpClient = new GuzzleClient(['proxy' => $proxyUrl]);
    }

    private function generateSign(string $body): array
    {
        $timestamp = (string)time();
        $signText = $this->appId . $timestamp . $body;
        $sign = hash_hmac('sha256', $signText, $this->appSecret);
        return ['sign' => $sign, 'timestamp' => $timestamp];
    }

    /**
     * @throws APIError
     * @throws GuzzleException
     */
    public function post(string $path, array $data = []): array
    {
        $body = json_encode($data);
        ['sign' => $sign, 'timestamp' => $timestamp] = $this->generateSign($body);

        $headers = [
            'Content-Type' => 'application/json',
            'Appid' => $this->appId,
            'Sign' => $sign,
            'Timestamp' => $timestamp,
        ];

        $response = $this->httpClient->post($this->baseUrl . $path, [
            'headers' => $headers,
            'body' => $body,
        ]);

        $result = json_decode($response->getBody()->getContents(), true);

        if ($result['code'] !== 10000) {
            throw new APIError($result['code'], $result['msg']);
        }

        return $result['data'] ?? [];
    }

    public function basicInfo(): Services\BasicInfoService
    {
        return new Services\BasicInfoService($this);
    }

    public function merchantAssets(): Services\MerchantAssetsService
    {
        return new Services\MerchantAssetsService($this);
    }

    public function merchantDeposit(): Services\MerchantDepositService
    {
        return new Services\MerchantDepositService($this);
    }

    public function merchantWithdraw(): Services\MerchantWithdrawService
    {
        return new Services\MerchantWithdrawService($this);
    }

    public function merchantBatchWithdraw(): Services\MerchantBatchWithdrawService
    {
        return new Services\MerchantBatchWithdrawService($this);
    }

    public function userAssets(): Services\UserAssetsService
    {
        return new Services\UserAssetsService($this);
    }

    public function userDeposit(): Services\UserDepositService
    {
        return new Services\UserDepositService($this);
    }

    public function userWithdraw(): Services\UserWithdrawService
    {
        return new Services\UserWithdrawService($this);
    }

    public function userTransfer(): Services\UserTransferService
    {
        return new Services\UserTransferService($this);
    }

    public function orders(): Services\OrdersService
    {
        return new Services\OrdersService($this);
    }

    public function checkout(): Services\CheckoutService
    {
        return new Services\CheckoutService($this);
    }

    public function swap(): Services\SwapService
    {
        return new Services\SwapService($this);
    }

    public function utilities(): Services\UtilitiesService
    {
        return new Services\UtilitiesService($this);
    }
}
