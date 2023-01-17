<?php

function RsaDecrypt($str)
{
    $pu_key = openssl_pkey_get_public(PublicKey);

    if (!$pu_key) return false;//秘钥不可用

    openssl_public_decrypt(base64_decode($str), $decrypted, $pu_key);

    return $decrypted;

}