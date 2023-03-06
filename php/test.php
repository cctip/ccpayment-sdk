<?php
namespace CCPayment;

use CCPayment\CCPay;
use WpOrg\Requests\Autoload;

require './vendor/autoload.php';

Autoload::register();

header('Content-type: application/json');

//$postStr = $GLOBALS["HTTP_RAW_POST_DATA"] ?? "";
//
//print_r($postStr);
//$data = json_decode(file_get_contents('php://input'), true);
//
//var_dump($data);
//echo 111;



// php -S localhost:8080
// $resp =  CCPay::CreateOrder( [
//      "remark"=>"",
//      "token_id"=>"8e5741cf-6e51-4892-9d04-3d40e1dd0128",
//      "amount"=>"0.5",
//      "merchant_order_id"=>"3735077979050373",
//      "fiat_currency"=> "USD"
//     ],"202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
//   var_dump($resp);

$resp =  CCPay::CheckUrl( [
    "return_url"=>"https://cwallet.com/pay/callback",
    "valid_timestamp"=>4566,
    "amount"=>"0.5",
    "merchant_order_id"=>"3735077979050370",
    "product_name"=> "knowledge is power"
],"202301310325561620262074393440256","c4600b8125b7ed23b5b7b8ee4acb42f4");
var_dump($resp);