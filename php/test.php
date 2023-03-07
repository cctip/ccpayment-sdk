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
// todo create order
// $resp =  CCPay::CreateOrder( [
//      "remark"=>"",
//      "token_id"=>"0912e09a-d8e2-41d7-a0bc-a25530892988",
//      "amount"=>"0.5",
//      "merchant_order_id"=>strval(time()).strval(rand(0,1000)),
//      "fiat_currency"=> "USD"
//     ],"202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
//   var_dump($resp);

// todo url
//$resp =  CCPay::CheckUrl( [
//    "return_url"=>"https://cwallet.com/pay/callback",
//    "valid_timestamp"=>4566,
//    "amount"=>"0.5",
//    "merchant_order_id"=>strval(time()).strval(rand(0,1000)),
//    "product_name"=> "knowledge is power"
//],"202301310325561620262074393440256","c4600b8125b7ed23b5b7b8ee4acb42f4");
//var_dump($resp);

// todo token list
//$resp =  CCPay::GetSupportToken("202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
//var_dump($resp);

// todo token chain
//$resp =  CCPay::GetTokenChain(["token_id"=>"f137d42c-f3a6-4f23-9402-76f0395d0cfe"],"202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
//var_dump($resp);

$resp =  CCPay::GetTokenRate(["token_id"=>"e8f64d3d-df5b-411d-897f-c6d8d30206b7","amount"=>"12"],"202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
var_dump($resp);

//$resp =  CCPay::Withdraw( [
//    "token_id"=>"f137d42c-f3a6-4f23-9402-76f0395d0cfe",//POLYGON  f137d42c-f3a6-4f23-9402-76f0395d0cfe
//    "address"=>"9454818",
//    "value"=>"0.1",
//    "memo"=>"",
//    "merchant_order_id"=>strval(time()).strval(rand(0,1000)),
//],"202302230220521628580622918598656","274edea2db595501770a0ffa8b80b202");
//var_dump($resp);

// todo check user
//$resp =  CCPay::CheckUser("9454818","202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
//var_dump($resp);

// todo fetch assets
//$resp =  CCPay::Assets("202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390","");
//var_dump($resp);

// todo fetch network fee
//$resp =  CCPay::NetworkFee(["token_id"=>"0912e09a-d8e2-41d7-a0bc-a25530892988"],"202302010636261620672405236006912","62fbff1f796c42c50bb44d4d3d065390");
//var_dump($resp);
