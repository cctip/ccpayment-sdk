<?php

class CCPaymentPay
{
    /**
     * 字符串签名
     * @param string $resource 经过urlencode处理过的字符串
     * @return string
     */
    public function SignStr($resource)
    {
        $privateKey = openssl_get_privatekey(PrivateKey);
        $res = openssl_sign($resource, $signature, $privateKey, OPENSSL_ALGO_SHA256);
        openssl_free_key($privateKey);
        if ($res) {
            return base64_encode($signature);
        } else {
            throw new \Exception("String Sign Failed", 10004);
        }
    }

    /**
     * 字符串验签
     * @param string $resource
     * @param string $signature
     * @return bool
     */
    public function VerifyStrMessage($resource, $signature)
    {
        $signature = base64_decode($signature);
        $publicKey = openssl_get_publickey(PublicKey);
        $res = openssl_verify($resource, $signature, $publicKey, OPENSSL_ALGO_SHA256);
        openssl_free_key($publicKey);
        return $res === 1 ? true : false;
    }

}