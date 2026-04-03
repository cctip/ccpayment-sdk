package com.ccpayment;

import com.google.gson.Gson;
import com.google.gson.JsonObject;
import okhttp3.*;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.TimeUnit;

public class Client {
    private static final String DEFAULT_BASE_URL = "https://ccpayment.com/ccpayment/v2";
    private static final Gson gson = new Gson();
    private static final MediaType JSON = MediaType.parse("application/json");

    private final String appId;
    private final String appSecret;
    private String baseUrl;
    private final OkHttpClient httpClient;

    public Client(String appId, String appSecret) {
        this.appId = appId;
        this.appSecret = appSecret;
        this.baseUrl = DEFAULT_BASE_URL;
        this.httpClient = new OkHttpClient.Builder()
                .connectTimeout(30, TimeUnit.SECONDS)
                .readTimeout(30, TimeUnit.SECONDS)
                .build();
    }

    public void setBaseUrl(String baseUrl) {
        this.baseUrl = baseUrl;
    }

    public void setProxy(java.net.Proxy proxy) {
        // OkHttpClient proxy configuration would go here
    }

    private String generateSign(String body, String timestamp) throws Exception {
        String signText = appId + timestamp + body;
        Mac mac = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKey = new SecretKeySpec(appSecret.getBytes(StandardCharsets.UTF_8), "HmacSHA256");
        mac.init(secretKey);
        byte[] hash = mac.doFinal(signText.getBytes(StandardCharsets.UTF_8));
        StringBuilder hexString = new StringBuilder();
        for (byte b : hash) {
            String hex = Integer.toHexString(0xff & b);
            if (hex.length() == 1) hexString.append('0');
            hexString.append(hex);
        }
        return hexString.toString();
    }

    public <T> T post(String path, Object data, Class<T> responseClass) throws APIError, IOException {
        try {
            String body = data != null ? gson.toJson(data) : "{}";
            String timestamp = String.valueOf(System.currentTimeMillis() / 1000);
            String sign = generateSign(body, timestamp);

            RequestBody requestBody = RequestBody.create(body, JSON);
            Request request = new Request.Builder()
                    .url(baseUrl + path)
                    .post(requestBody)
                    .addHeader("Content-Type", "application/json")
                    .addHeader("Appid", appId)
                    .addHeader("Sign", sign)
                    .addHeader("Timestamp", timestamp)
                    .build();

            try (Response response = httpClient.newCall(request).execute()) {
                if (!response.isSuccessful()) {
                    throw new IOException("Unexpected code " + response);
                }

                String responseBody = response.body().string();
                JsonObject jsonObject = gson.fromJson(responseBody, JsonObject.class);

                int code = jsonObject.get("code").getAsInt();
                if (code != 10000) {
                    String msg = jsonObject.get("msg").getAsString();
                    throw new APIError(code, msg);
                }

                if (responseClass == Void.class) {
                    return null;
                }

                if (jsonObject.has("data") && !jsonObject.get("data").isJsonNull()) {
                    return gson.fromJson(jsonObject.get("data"), responseClass);
                }

                return null;
            }
        } catch (APIError | IOException e) {
            throw e;
        } catch (Exception e) {
            throw new IOException("Request failed", e);
        }
    }

    // Service accessors
    public BasicInfoService basicInfo() {
        return new BasicInfoService(this);
    }

    public MerchantAssetsService merchantAssets() {
        return new MerchantAssetsService(this);
    }

    public MerchantDepositService merchantDeposit() {
        return new MerchantDepositService(this);
    }

    public MerchantWithdrawService merchantWithdraw() {
        return new MerchantWithdrawService(this);
    }

    public MerchantBatchWithdrawService merchantBatchWithdraw() {
        return new MerchantBatchWithdrawService(this);
    }

    public UserAssetsService userAssets() {
        return new UserAssetsService(this);
    }

    public UserDepositService userDeposit() {
        return new UserDepositService(this);
    }

    public UserWithdrawService userWithdraw() {
        return new UserWithdrawService(this);
    }

    public UserTransferService userTransfer() {
        return new UserTransferService(this);
    }

    public OrdersService orders() {
        return new OrdersService(this);
    }

    public CheckoutService checkout() {
        return new CheckoutService(this);
    }

    public SwapService swap() {
        return new SwapService(this);
    }

    public UtilitiesService utilities() {
        return new UtilitiesService(this);
    }
}
