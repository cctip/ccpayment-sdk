package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class CheckoutService {
    private final Client client;

    public CheckoutService(Client client) {
        this.client = client;
    }

    public JsonObject createCheckoutUrl(JsonObject data) throws APIError, IOException {
        return client.post("/createCheckoutUrl", data, JsonObject.class);
    }

    public JsonObject hostedOrderInfo(String orderId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("orderId", orderId);
        return client.post("/hostedOrderInfo", data, JsonObject.class);
    }

    public JsonObject hostedOrderInfoFirst(String orderId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("orderId", orderId);
        return client.post("/hostedOrderInfoFirst", data, JsonObject.class);
    }

    public JsonObject createHostedOrderDeposit(JsonObject data) throws APIError, IOException {
        return client.post("/createHostedOrderDeposit", data, JsonObject.class);
    }

    public JsonObject getHostedCoinUSDTPrice(com.google.gson.JsonArray coinIds) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.add("coinIds", coinIds);
        return client.post("/getHostedCoinUSDTPrice", data, JsonObject.class);
    }

    public JsonObject getMainCoinList() throws APIError, IOException {
        return client.post("/getMainCoinList", null, JsonObject.class);
    }

    public JsonObject createAppDemoOrderDeposit(JsonObject data) throws APIError, IOException {
        return client.post("/createAppDemoOrderDeposit", data, JsonObject.class);
    }

    public void confirmTrade(JsonObject data) throws APIError, IOException {
        client.post("/confirmTrade", data, Void.class);
    }
}
