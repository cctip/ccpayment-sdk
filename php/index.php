<?php
include('vendor/rmccue/requests/library/Requests.php');


const  TestCreateOrderUrl = "http://74ab25e1merchant.cwallet.com/ccpayment/v1/pay/CreateTokenTradeOrder";
const    PublicKey = "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAvpbwh1QidzeChzTcp/3P\nrupdq80cSuVi093/jg+t1kvBE8lQLwfWVfTmjfHQXKPSZ7jjmgu5Z1AJErePoksp\nnJs/zdtnfly4Kelp/VynhXgolVjgbrdJ8IzNIShmvCT+xrgQ4mC6w3eibnkWMCzk\nfulGRFCnUPOhhVevvnVpaEOp105TeBxXTiS+TNzcgOjq8m7nBHYraEYor8COs85K\n/gbEQSc+qdWVNpR9ZzPvwl5CE8k1dNnZdHY3/2y215KeQLEuIe1Yhd2/OI6sMg40\n8++f1ogiEOdIoRfSfiCCETUlHsNxZDRCt3sho4VNuuqUXJoroJUhavpqkfQXlgJB\n4x+NFo76/PZJKFJ5iY2Ob7SlY3R5w1pYAzxk5WrqmTjwyUhEbVX25yEeioVkwqNz\ngl7cUbR9JvJQ/Is0+eA9apVG5AysIfNBnkBXct47hOVI+FR+XUwvNwaRbn43ydsu\n7oqYGHA8OLhIk52aUS6aT+0CBHBB35NrRRDaq252EXGINg7utkjUuLGN8hCLmZv2\nSMhvb5G3g7btBKs1hNtZpnQXj+T/bviP6Wgsy2b1ZSJLr3iv/1JNcH6780NSygSo\ngVkBlKt0hkenLWdGRMbCfwQ7LSRUmQqzjkgjePkJMTXVdwESZ02jI5NSkLOi8+kW\nmcWh0rSpoX41HC8ZLwmFJq0CAwEAAQ==\n-----END PUBLIC KEY-----";
const    PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----";


function Pay()
{

    Requests::register_autoloader();
    $headers = array(
        'User-Agent' => 'Apipost client Runtime/+https://www.apipost.cn/',
        'Content-Type' => 'application/json'
    );
    $sign = '';
    $data = '{ "sign": "' + $sign + '", "merchant_id": "CP10001", "app_id": "202211181420251593489282956267520", "timestamp": 1672261484, "notify_url": "https://cwallet.com/pay/callback", "remark": "", "device": "web", "noncestr": "werwer", "json_content": { "out_order_no": "202234304340343512", "amount": "1002.00", "chain": "ETH", "contract": "0x2170ed0880ac9a755fd29b2688956bd959f933f8", "token_id": "e8f64d3d-df5b-411d-897f-c6d8d30206b7", "fiat_name": "USD" } }';
    $response = Requests::post(TestCreateOrderUrl, $headers, $data);

}

function Notify() {

}






