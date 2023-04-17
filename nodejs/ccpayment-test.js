
/* 
 * This test code can only be displayed as a call method. 
 * Related Nodejs Web service support can be run to run
*/
const appId = '202302160812171626132344467566592';
const appSecret = 'a58f572564f7fce44acd66024d6da9b4';

// init get appid and appsecret
ccpaymentWidgets.init(appId, appSecret)
// select checkout url
ccpaymentWidgets.checkoutURL({
  'order_valid_period': 823456,
  'product_price': '1',
  'merchant_order_id': '012033040550',
  'product_name': 'test',
  'return_url': 'https://app.gitbook.com/xxxxx'
}, (result) => {
  console.log('aaa:', result)
})
// select token
ccpaymentWidgets.getSupportToken((result) => {
  console.log('bbb:', result)
})
// select chain
ccpaymentWidgets.getTokenChain({
  "token_id": "8addd19b-37df-4faf-bd74-e61e214b008a"
}, (result) => {
  console.log('ccc:', result)
})
// submit order
ccpaymentWidgets.createOrder({
  "remark": "eee",
  "token_id": "8e5741cf-6e51-4892-9d04-3d40e1dd0128",
  "product_price": "0.5",
  "merchant_order_id": "37350779790509509",
  "denominated_currency": "USD"
}, (result) => {
  console.log('ddd:', result)
})
// token rate
ccpaymentWidgets.getTokenRate({
  "amount": "1000",
  "token_id": "f36ad1cf-222a-4933-9ad0-86df8069f916",
}, (result) => {
  console.log('eee:', result)
})



