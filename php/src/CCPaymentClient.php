<?php

declare(strict_types=1);

namespace CCPayment\SDK;

use RuntimeException;

class CCPaymentClient
{
    private string $baseUrl;
    private string $appId;
    private ?string $appSecret;
    private ?string $privateKey;

    public function __construct(string $baseUrl, string $appId, ?string $appSecret = null, ?string $privateKey = null)
    {
        $this->baseUrl = rtrim($baseUrl, '/');
        $this->appId = $appId;
        $this->appSecret = $appSecret;
        $this->privateKey = $privateKey;
    }

    public function postV1(string $path, ?array $payload = null): array
    {
        return $this->request('/ccpayment/v1' . $path, $payload, 'v1');
    }

    public function postV2(string $path, ?array $payload = null, string $signMode = 'hmac'): array
    {
        return $this->request('/ccpayment/v2' . $path, $payload, 'v2', $signMode);
    }

    public function postV1Hosted(string $path, string $billId, ?array $payload = null): array
    {
        $timestamp = (string) time();
        $body = $this->encodeBody($payload);
        $sign = hash('sha256', $billId . $timestamp);

        $headers = [
            'Appid: ' . $this->appId,
            'Timestamp: ' . $timestamp,
            'Sign: ' . $sign,
            'BillId: ' . $billId,
            'Content-Type: application/json',
        ];

        return $this->doHttpPost($this->baseUrl . '/ccpayment/v1' . $path, $body, $headers);
    }

    public function getSupportToken(): array
    {
        return $this->postV1('/support/token', []);
    }

    public function getTokenChain(string $tokenId): array
    {
        return $this->postV1('/token/chain', ['token_id' => $tokenId]);
    }

    public function createOrder(string $tokenId, string $productPrice, string $merchantOrderId, string $denominatedCurrency, ?int $orderValidPeriod = null, ?string $remark = null): array
    {
        $payload = [
            'token_id' => $tokenId,
            'product_price' => $productPrice,
            'merchant_order_id' => $merchantOrderId,
            'denominated_currency' => $denominatedCurrency,
        ];

        if ($orderValidPeriod !== null) {
            $payload['order_valid_period'] = $orderValidPeriod;
        }
        if ($remark !== null && $remark !== '') {
            $payload['remark'] = $remark;
        }

        return $this->postV1('/bill/create', $payload);
    }

    public function checkoutUrl(array $payload): array
    {
        return $this->postV1('/concise/url/get', $payload);
    }

    public function withdraw(array $payload): array
    {
        return $this->postV1('/withdraw', $payload);
    }

    public function checkUser(string $cId): array
    {
        return $this->postV1('/check/user', ['c_id' => $cId]);
    }

    public function assets(?string $coinId = null): array
    {
        $payload = [];
        if ($coinId !== null && $coinId !== '') {
            $payload['coin_id'] = $coinId;
        }

        return $this->postV1('/assets', $payload);
    }

    public function networkFee(string $tokenId, ?string $address = null, ?string $memo = null): array
    {
        $payload = ['token_id' => $tokenId];
        if ($address !== null && $address !== '') {
            $payload['address'] = $address;
        }
        if ($memo !== null && $memo !== '') {
            $payload['memo'] = $memo;
        }

        return $this->postV1('/network/fee', $payload);
    }

    public function getCoinListV2(array $payload = []): array
    {
        return $this->postV2('/getCoinList', $payload, 'hmac');
    }

    public function createCheckoutUrlV2(array $payload, string $signMode = 'hmac'): array
    {
        return $this->postV2('/createCheckoutUrl', $payload, $signMode);
    }

    private function request(string $path, ?array $payload, string $version, string $signMode = 'hmac'): array
    {
        $timestamp = (string) time();
        $body = $this->encodeBody($payload);
        $sign = $this->buildSign($version, $timestamp, $body, $signMode);

        $headers = [
            'Appid: ' . $this->appId,
            'Timestamp: ' . $timestamp,
            'Sign: ' . $sign,
            'Content-Type: application/json',
        ];

        return $this->doHttpPost($this->baseUrl . $path, $body, $headers);
    }

    private function buildSign(string $version, string $timestamp, string $body, string $signMode): string
    {
        $signBody = $body === '' ? '' : $body;

        if ($version === 'v1') {
            if ($this->appSecret === null || $this->appSecret === '') {
                throw new RuntimeException('appSecret is required for v1 signing');
            }
            return hash('sha256', $this->appId . $this->appSecret . $timestamp . $signBody);
        }

        if ($version !== 'v2') {
            throw new RuntimeException('unsupported version: ' . $version);
        }

        $signText = $this->appId . $timestamp . $signBody;

        if ($signMode === 'hmac') {
            if ($this->appSecret === null || $this->appSecret === '') {
                throw new RuntimeException('appSecret is required for v2 hmac signing');
            }
            return hash_hmac('sha256', $signText, $this->appSecret);
        }

        if ($signMode === 'rsa') {
            if ($this->privateKey === null || $this->privateKey === '') {
                throw new RuntimeException('privateKey is required for v2 rsa signing');
            }

            $privateKey = openssl_pkey_get_private($this->privateKey);
            if ($privateKey === false) {
                throw new RuntimeException('invalid RSA private key');
            }

            $signature = '';
            $ok = openssl_sign($signText, $signature, $privateKey, OPENSSL_ALGO_SHA256);
            openssl_free_key($privateKey);

            if (!$ok) {
                throw new RuntimeException('failed to sign request with RSA private key');
            }

            return base64_encode($signature);
        }

        throw new RuntimeException('unsupported v2 sign mode: ' . $signMode);
    }

    private function encodeBody(?array $payload): string
    {
        if ($payload === null) {
            return '';
        }

        $json = json_encode($payload, JSON_UNESCAPED_UNICODE | JSON_UNESCAPED_SLASHES);
        if ($json === false) {
            throw new RuntimeException('json encode failed: ' . json_last_error_msg());
        }

        return $json;
    }

    private function doHttpPost(string $url, string $body, array $headers): array
    {
        $ch = curl_init();
        if ($ch === false) {
            throw new RuntimeException('failed to init curl');
        }

        curl_setopt_array($ch, [
            CURLOPT_URL => $url,
            CURLOPT_POST => true,
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_HTTPHEADER => $headers,
            CURLOPT_POSTFIELDS => $body,
            CURLOPT_TIMEOUT => 30,
            CURLOPT_CONNECTTIMEOUT => 10,
            CURLOPT_HEADER => true,
        ]);

        $raw = curl_exec($ch);
        if ($raw === false) {
            $err = curl_error($ch);
            curl_close($ch);
            throw new RuntimeException('curl request failed: ' . $err);
        }

        $status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        $headerSize = curl_getinfo($ch, CURLINFO_HEADER_SIZE);
        curl_close($ch);

        $headerText = substr($raw, 0, $headerSize);
        $bodyText = substr($raw, $headerSize);

        $decoded = json_decode($bodyText, true);
        if (!is_array($decoded)) {
            $decoded = [
                'raw' => $bodyText,
            ];
        }

        return [
            'status' => $status,
            'headers' => $this->parseHeaders($headerText),
            'body' => $decoded,
        ];
    }

    private function parseHeaders(string $headerText): array
    {
        $lines = preg_split('/\r\n|\r|\n/', trim($headerText));
        $headers = [];

        if ($lines === false) {
            return $headers;
        }

        foreach ($lines as $line) {
            $pos = strpos($line, ':');
            if ($pos === false) {
                continue;
            }

            $key = trim(substr($line, 0, $pos));
            $value = trim(substr($line, $pos + 1));
            if ($key === '') {
                continue;
            }

            $headers[$key] = $value;
        }

        return $headers;
    }
}
