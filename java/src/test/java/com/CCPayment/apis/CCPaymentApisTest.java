package com.CCPayment.apis;

import com.CCPayment.beans.*;
import org.junit.jupiter.api.Test;

import java.util.Arrays;

class CCPaymentApisTest {

    @Test
    void createOrder() throws Exception {
        CreateOrderParam param = new CreateOrderParam();
        param.setProductPrice("0.07");
        param.setTokenId("2e6cfa7d-f658-455d-89cd-31ebbcfdfa2c");
        param.setRemark("222222");
        param.setMerchantOrderId("thisisoutorderno1100");
        param.setDenominatedCurrency("TOKEN");
        param.setNotifyUrl("https://xxxxxxxx.com/notify.url");
        param.setOrderValidPeriod(1800);
        param.setCustomValue("custom value");
        CreateOrderResponse resp = CCPaymentApis.createOrder(param);
        System.out.println(resp);
    }

    @Test
    void getSupportToken() throws Exception{
       GetSupportTokenResponse resp = CCPaymentApis.getSupportToken();
       System.out.println(resp);
    }

    @Test
    void getTokenChain() throws Exception{
        GetTokenChainParam param = new GetTokenChainParam();
        param.setTokenId("217d4c3b-4f84-4416-bf1a-a1d15d7d2f50");
        GetTokenChainResponse resp = CCPaymentApis.getTokenChain(param);
        System.out.println(resp);
    }

    @Test
    void checkoutUrl() throws Exception{
        CheckoutUrlParam param = new CheckoutUrlParam();
        param.setProductPrice("1000");
        param.setMerchantOrderId("thisisoutorderno");
        param.setReturnUrl("https://www.xxxx.com/callback");
        param.setProductName("this is product name");
        param.setNotifyUrl("https://xxxxxxxx.com/notify.url");
        param.setOrderValidPeriod(1800);
        param.setCustomValue("custom value");
        CheckoutUrlResponse resp = CCPaymentApis.checkoutUrl(param);
        System.out.println(resp);
    }

    @Test
    void getTokenRate() throws Exception{
        GetTokenRateParam param = new GetTokenRateParam();
        param.setAmount("1000");
        param.setTokenId("e8f64d3d-df5b-411d-897f-c6d8d30206b7");
        GetTokenRateResponse resp = CCPaymentApis.getTokenRate(param);
        System.out.println(resp);
    }

    @Test
    void callback() throws Exception {
        String callbackData = "{\"pay_status\":\"success\",\"order_type\":\"Invoice\",\"record_id\":\"202302201213531627642695975706624\",\"order_id\":\"202302200951001627606736097771520\",\"origin_price\":\"25\",\"origin_amount\":\"25\",\"fiat_rate\":\"0.08\",\"paid_amount\":\"10\",\"token_rate\":\"0.08\",\"chain\":\"BSC\",\"contract\":\"0xbA2aE424d960c26247Dd6c32edC70B295c744C43\",\"crypto\":\"DOGE\",\"extend\":{\"invoice_id\":\"202302200950101627606529100480512\",\"user_email\":\"herculew@protonmail.com\",\"merchant_order_id\":\"202211154785795\"}}";
        String appId = "202302010636261620672405236006912";
        long timestamp = 1677653408;
        String sign = "c7c6e98b684f079b71d950ce15033b479a140f6f5aa0dece06eb28f45f04c088";
        CallbackData data = CCPaymentApis.callback(callbackData, appId, timestamp, sign);
        System.out.println(data);
    }

    @Test
    void withdraw() throws Exception{
        WithdrawParam param = new WithdrawParam();
        param.setTokenId("2e6cfa7d-f658-455d-89cd-31ebbcfdfa2c");
        param.setAddress("TJXTfwqepEZwTbXgZQoDLgqHt5bfQUN1HX");
        param.setMerchantOrderId("withdraw0001");
        param.setValue("8");
        WithdrawResponse resp = CCPaymentApis.withdraw(param);
        System.out.println(resp);
    }

    @Test
    void checkUser() throws Exception{
        CheckUserParam param = new CheckUserParam();
        param.setCId("9454818");
        CheckUserResponse resp = CCPaymentApis.checkUser(param);
        System.out.println(resp);
    }

    @Test
    void assets() throws Exception{
        AssetsParam param = new AssetsParam();
        param.setTokenId("e8f64d3d-df5b-411d-897f-c6d8d30206b7");
        AssetsResponse resp = CCPaymentApis.assets(param);
        System.out.println(resp);
    }

    @Test
    void networkFee() throws Exception{
        NetworkFeeParam param = new NetworkFeeParam();
        param.setTokenId("f137d42c-f3a6-4f23-9402-76f0395d0cfe");
        NetworkFeeResponse resp = CCPaymentApis.networkFee(param);
        System.out.println(resp);
    }

    @Test
    void orderInfo() throws Exception{
        OrderInfoParam param = new OrderInfoParam();
        param.setMerchantOrderIds(new String[]{"thisisoutorderno1100"});
        OrderInfoResponse[] resp = CCPaymentApis.orderInfo(param);
        System.out.println(Arrays.toString(resp));
    }

    @Test
    void  getSupportCoin() throws Exception {
        GetSupportCoinResponse resp = CCPaymentApis.getSupportCoin();
        System.out.println(resp);
    }

    @Test
    void paymentAddressGet() throws Exception {
        PaymentAddressParam param = new PaymentAddressParam();
        param.setChain("TRX");
        param.setUserId("59586869696");
        param.setNotifyUrl("https://xxxxxxx.com/notify.html");
        PaymentAddressResponse resp = CCPaymentApis.paymentAddressGet(param);
        System.out.println(resp);
    }

    @Test
    void GetChainHeightInfo() throws Exception {
        GetChainHeightInfoResponse[] resp = CCPaymentApis.getChainHeightInfo();
        System.out.println(Arrays.toString(resp));
    }
}