package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class OrdersService {
    private final Client client;

    public OrdersService(Client client) {
        this.client = client;
    }

    public JsonObject getAppOrderInfo(String orderId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("orderId", orderId);
        return client.post("/getAppOrderInfo", data, JsonObject.class);
    }

    public JsonObject createInvoiceUrl(JsonObject data) throws APIError, IOException {
        return client.post("/createInvoiceUrl", data, JsonObject.class);
    }

    public JsonObject getInvoiceOrderInfo(String orderId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("orderId", orderId);
        return client.post("/getInvoiceOrderInfo", data, JsonObject.class);
    }

    public JsonObject getWebhookInfo() throws APIError, IOException {
        return client.post("/getWebhookInfo", null, JsonObject.class);
    }
}
