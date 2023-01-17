<?php

function RsaEncrypt($str)
{
    $pi_key = openssl_pkey_get_private(PrivateKey);

    if (!$pi_key) return false;//秘钥不可用

    openssl_private_encrypt($str, $encrypted, $pi_key);

    $encrypted = base64_encode($encrypted);

    return $encrypted;

}



