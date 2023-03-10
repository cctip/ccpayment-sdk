package com.Ccpayment.apis;

import com.Ccpayment.beans.*;
import org.junit.jupiter.api.Test;

class CcpaymentApisTest {

    @Test
    void paymentIntentCreate() throws Exception {
        CreateOrderParam param = new CreateOrderParam();
        param.setNotifyUrl("https://www.xxxx.com/callback");
        param.setAmount("1.34");
        param.setTokenId("1e614496-f9ce-4ec7-8e68-563733deacbf");
        param.setRemark("222222");
        param.setMerchantOrderId("thisisoutordern1o");
        param.setFiatCurrency("USD");
        CreateOrderResponse resp = CcpaymentApis.createOrder(param);
        System.out.println(resp);
    }

    @Test
    void supportTokens() throws Exception{
       GetSupportTokenResponse resp = CcpaymentApis.getSupportToken();
       System.out.println(resp);
    }

    @Test
    void tokenChain() throws Exception{
        GetTokenChainParam param = new GetTokenChainParam();
        param.setTokenId("e8f64d3d-df5b-411d-897f-c6d8d30206b7");
        GetTokenChainResponse resp = CcpaymentApis.getTokenChain(param);
        System.out.println(resp);
    }

    @Test
    void checkUrl() throws Exception{
        CheckoutUrlParam param = new CheckoutUrlParam();
        param.setAmount("1000");
        param.setMerchantOrderId("thisisoutorderno");
        param.setReturnUrl("https://www.xxxx.com/callback");
        param.setValidTimestamp(3600);
        param.setProductName("this is product name");
        CheckoutUrlResponse resp = CcpaymentApis.checkoutUrl(param);
        System.out.println(resp);
    }

    @Test
    void tokenRate() throws Exception{
        GetTokenRateParam param = new GetTokenRateParam();
        param.setAmount("1000");
        param.setTokenId("e8f64d3d-df5b-411d-897f-c6d8d30206b7");
        GetTokenRateResponse resp = CcpaymentApis.getTokenRate(param);
        System.out.println(resp);
    }

    @Test
    void callback() throws Exception {
        String callbackData = "{\"pay_status\":\"success\",\"order_type\":\"Invoice\",\"record_id\":\"202302201213531627642695975706624\",\"order_id\":\"202302200951001627606736097771520\",\"origin_price\":\"25\",\"origin_amount\":\"25\",\"fiat_rate\":\"0.08\",\"paid_amount\":\"10\",\"token_rate\":\"0.08\",\"chain\":\"BSC\",\"contract\":\"0xbA2aE424d960c26247Dd6c32edC70B295c744C43\",\"crypto\":\"DOGE\",\"extend\":{\"invoice_id\":\"202302200950101627606529100480512\",\"user_email\":\"herculew@protonmail.com\",\"merchant_order_id\":\"202211154785795\"}}";
        String appId = "202302010636261620672405236006912";
        long timestamp = 1677653408;
        String sign = "c7c6e98b684f079b71d950ce15033b479a140f6f5aa0dece06eb28f45f04c088";
        CallbackData data = CcpaymentApis.callback(callbackData, appId, timestamp, sign);
        System.out.println(data);
    }
}