/*
 * This test code can only be displayed as a call method.
 * Related Nodejs Web service support can be run to run
 */
const ccpaymentWidgets = require("./ccpayment.js");

const appId = "202306261609051673241961567358991";
const appSecret = "5af84fd1f5cb2d9ac5e564defb31e7bf";

// init get appid and appsecret
ccpaymentWidgets.init(appId, appSecret);
// select checkout url
ccpaymentWidgets.checkoutURL(
  {
    order_valid_period: 823456,
    product_price: "1",
    merchant_order_id: "012033040550",
    product_name: "test",
    return_url: "https://app.gitbook.com/xxxxx",
  },
  (result) => {
    console.log("aaa:", result);
  }
);
// select token
ccpaymentWidgets.getSupportToken((result) => {
  // console.log("bbb:", JSON.stringify(result, null, 2));
  console.log("bbb:", result);
});

// select chain
ccpaymentWidgets.getTokenChain(
  {
    token_id: "8addd19b-37df-4faf-bd74-e61e214b008a",
  },
  (result) => {
    console.log("ccc:", result);
  }
);
// submit order
ccpaymentWidgets.createOrder(
  {
    remark: "eee",
    token_id: "f36ad1cf-222a-4933-9ad0-86df8069f916",
    product_price: "0.5",
    merchant_order_id: "37350779790509509",
    denominated_currency: "USD",
  },
  (result) => {
    console.log("ddd:", result);
  }
);
// token rate
ccpaymentWidgets.getTokenRate(
  {
    amount: "1000",
    token_id: "f36ad1cf-222a-4933-9ad0-86df8069f916",
  },
  (result) => {
    console.log("eee:", result);
  }
);

// getSupportCoin
ccpaymentWidgets.getSupportCoin((result) => {
  // console.log(
  //   "hhh:",
  //   result.list.find((item) => item.crypto === "USDT")
  // );
  console.log(
    "hhh:",
    result.data.list.find((item) => item.crypto === "USDT")
  );
});

// getOrderInfo
ccpaymentWidgets.getOrderInfo(
  { merchant_order_ids: ["12345678765456720230720"] },
  (result) => {
    console.log("getOrderInfo:", result);
  }
);

/*
 * @param {Object} data
 * @param {String} data.token_id
 * @param {String} data.address
 * @param {String} data.merchant_order_id
 * @param {String} data.value
 * @param {String} data.memo
 * @param {Function} callback
 * @return {void}
 */
// ccpaymentWidgets.withdraw(
//   {
//     token_id: "264f4725-3cfd-4ff6-bc80-ff9d799d5fb2",
//     address: "0xdA2B9571c84b9F3d4d82b0AeCE814B2b2e2C13a2",
//     merchant_order_id: "12345678765456720230720",
//     value: "0.01",
//     memo: "",
//   },
//   (result) => {
//     console.log("withdraw:", result);
//   }
// );

ccpaymentWidgets.paymentAddress(
  {
    user_id: "cc-123",
    chain: "ETH",
    notify_url: "",
  },
  (result) => {
    console.log("paymentAddress:", result);
  }
);
