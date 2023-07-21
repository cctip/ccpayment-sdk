package com.CCPayment.apis;


import com.CCPayment.beans.*;
import com.CCPayment.constant.Config;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JavaType;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.util.EntityUtils;

import java.math.BigInteger;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;


/**
 * Ccpayment apis
 * document: https://doc.ccpayment.com/ccpayment-for-developer/ccpayment-for-developer
 */
public final class CCPaymentApis {

    /**
     * common response data
     *
     * @param <T>
     */
    public static class Response<T> {
        @JsonProperty("code")
        private int code;
        @JsonProperty("msg")
        private String msg;
        @JsonProperty("data")
        private T data;
    }

    /**
     * Create pay order
     * document: https://doc.ccpayment.com/ccpayment-for-developer/payment-api/native-checkout-integration
     * @param param params
     * @return pay order info
     * @throws Exception
     */
    public static CreateOrderResponse createOrder(CreateOrderParam param) throws Exception {
        return doSend("/bill/create", param, CreateOrderResponse.class);
    }

    /**
     * @deprecated
     * Get tokens ccpayment support list
     *
     * @return token list
     * @throws Exception
     */
    public static GetSupportTokenResponse getSupportToken() throws Exception {
        return doSend("/support/token", new Object(), GetSupportTokenResponse.class);
    }

    /**
     * Get tokens ccpayment support list
     * document: https://doc.ccpayment.com/ccpayment-for-developer/payment-api/token-id-interface
     * @return token list
     * @throws Exception
     */
    public static GetSupportCoinResponse getSupportCoin() throws Exception {
        return doSend("/coin/all", new Object(), GetSupportCoinResponse.class);
    }

    /**
     * Get a payment link
     * document: https://doc.ccpayment.com/ccpayment-for-developer/payment-api/hosted-checkout-page-integration
     * @param param
     * @return pay url
     * @throws Exception
     */
    public static CheckoutUrlResponse checkoutUrl(CheckoutUrlParam param) throws Exception {
        return doSend("/concise/url/get", param, CheckoutUrlResponse.class);
    }

    /**
     * @deprecated
     * Get chains of token
     * @param param
     * @return
     * @throws Exception
     */
    public static GetTokenChainResponse getTokenChain(GetTokenChainParam param) throws Exception {
        return doSend("/token/chain", param, GetTokenChainResponse.class);
    }

    /**
     * Amount equivalent to usd
     * document: https://doc.ccpayment.com/ccpayment-for-developer/resources-document/current-token-rate-interface
     * @param param
     * @return
     * @throws Exception
     */
    public static GetTokenRateResponse getTokenRate(GetTokenRateParam param) throws Exception {
        return doSend("/token/rate", param, GetTokenRateResponse.class);
    }

    /**
     * Api withdraw
     * document: https://doc.ccpayment.com/ccpayment-for-developer/withdrawal-api-integration
     * @param param
     * @return
     * @throws Exception
     */
    public static WithdrawResponse withdraw(WithdrawParam param) throws Exception {
        return doSend("/withdraw", param, WithdrawResponse.class);
    }

    /**
     * Check user exists or not
     * document: https://doc.ccpayment.com/ccpayment-for-developer/resources-document/check-the-validity-of-cwallet-id
     * @param param
     * @return
     * @throws Exception
     */
    public static CheckUserResponse checkUser(CheckUserParam param) throws Exception {
        return doSend("/check/user", param, CheckUserResponse.class);
    }

    /**
     * Get token assets
     * document: https://doc.ccpayment.com/ccpayment-for-developer/resources-document/asset-balance-interface
     * @param param
     * @return
     * @throws Exception
     */
    public static AssetsResponse assets(AssetsParam param) throws Exception {
        return doSend("/assets", param, AssetsResponse.class);
    }

    /**
     * Get network-fee
     * document: https://doc.ccpayment.com/ccpayment-for-developer/resources-document/network-fee-interface
     * @param param
     * @return
     * @throws Exception
     */
    public static NetworkFeeResponse networkFee(NetworkFeeParam param) throws Exception {
        return doSend("/network/fee", param, NetworkFeeResponse.class);
    }

    /**
     * get api order info
     * document: https://doc.ccpayment.com/ccpayment-for-developer/payment-api/order-information-interface
     * @param param
     * @return
     * @throws Exception
     */
    public static OrderInfoResponse[] orderInfo(OrderInfoParam param) throws Exception {
        return doSend("/bill/info", param, OrderInfoResponse[].class);
    }



    /**
     * get payment address
     * document: https://doc.ccpayment.com/ccpayment-for-developer/get-permanent-deposit-address
     * @param param
     * @return
     * @throws Exception
     */
    public static PaymentAddressResponse paymentAddress(PaymentAddressParam param) throws Exception {
        return doSend("/payment/address/get", param, PaymentAddressResponse.class);
    }


    /**
     * callback when the payment is success
     * document: https://doc.ccpayment.com/ccpayment-for-developer/resources-document/webhook-notification
     * @param callbackData response date
     * @param appId        get from response header Appid
     * @param timestamp    get from response header Timestamp
     * @param sign         get from response header Sign
     * @return callback data
     * @throws Exception if sign verify failed or response date error, will raise an exception
     */
    public static CallbackData callback(String callbackData, String appId, long timestamp, String sign) throws Exception {
        if (System.currentTimeMillis()/ 1000 - timestamp  >  Config.expireDuration) {
            throw new Exception("callback request expired");
        }
        if (!appId.equals(Config.AppId)) {
            throw new Exception("appid compare failed");
        }
        if (!sign.equals(getSign(callbackData, String.format("%d", timestamp)))) {
            throw new Exception("sign error");
        }
        ObjectMapper map = new ObjectMapper();
        map.configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false);
        return map.readValue(callbackData, CallbackData.class);
    }

    /**
     * padding data where length not enough
     *
     * @param hash sha256 hash data
     * @return 64 length data
     */
    public static String toHexString(byte[] hash) {
        BigInteger number = new BigInteger(1, hash);
        StringBuilder hexString = new StringBuilder(number.toString(16));
        while (hexString.length() < 64) {
            hexString.insert(0, '0');
        }
        return hexString.toString();
    }

    /**
     * sign data with timestamp
     *
     * @param data      origin data
     * @param timestamp unix timestamp
     * @return sha256 hashed data
     * @throws NoSuchAlgorithmException exceptions
     */
    public static String getSign(String data, String timestamp) throws NoSuchAlgorithmException {
        String originStr = String.format("%s%s%s%s", Config.AppId, Config.AppSecrete, timestamp, data);
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        return toHexString(digest.digest(
                originStr.getBytes(StandardCharsets.UTF_8)));
    }

    /**
     * Request api to Ccpayment
     *
     * @param api
     * @param params
     * @param resp
     * @param <T>    resp params
     * @return
     * @throws Exception
     */
    protected static <T> T doSend(String api, Object params, Class<T> resp) throws Exception {
        HttpPost post = new HttpPost(String.format("%s%s", Config.ApiUrl, api));
        ObjectMapper mapper = new ObjectMapper();
        mapper.configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false);
        String inputString = mapper.writeValueAsString(params);
        String timestampStr = String.format("%d", System.currentTimeMillis() / 1000);
        String encodedHash = getSign(inputString, timestampStr);
        post.addHeader("Content-Type", "application/json");
        post.addHeader(Config.HeaderAppId, Config.AppId);
        post.addHeader(Config.HeaderTimestamp, timestampStr);
        post.addHeader(Config.HeaderSign, encodedHash);
        final StringEntity entity = new StringEntity(inputString, StandardCharsets.UTF_8);
        post.setEntity(entity);
        CloseableHttpClient httpClient = HttpClientBuilder.create().build();
        CloseableHttpResponse response = httpClient.execute(post);
        if (response.getStatusLine().getStatusCode() == 200) {
            String responseString = EntityUtils.toString(response.getEntity());
            JavaType type = mapper.getTypeFactory().constructParametricType(Response.class, resp);
            Response<T> result = mapper.readValue(responseString, type);
            if (result.code != 10000) {
                throw new Exception(String.format("resp error code is %d, msg is %s", result.code, result.msg));
            }else{
                timestampStr = "";
                String responseSign = "";
                if (response.getHeaders(Config.HeaderTimestamp) != null && response.getHeaders(Config.HeaderTimestamp).length > 0) {
                    timestampStr = response.getHeaders(Config.HeaderTimestamp)[0].getValue();
                }
                if (response.getHeaders(Config.HeaderSign) != null && response.getHeaders(Config.HeaderSign).length > 0) {
                    responseSign = response.getHeaders(Config.HeaderSign)[0].getValue();
                }
                if (timestampStr.equals("") || responseSign.equals("")) {
                    throw new Exception("response data error");
                }
                if (!getSign(responseString, timestampStr).equals(responseSign)) {
                    throw new Exception("resp sign error");
                }
            }
            return result.data;
        } else {
            throw new Exception(String.format("http request err, code is %d", response.getStatusLine().getStatusCode()));
        }
    }
}
