package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UtilitiesService {
    private final Client client;

    public UtilitiesService(Client client) {
        this.client = client;
    }

    public JsonObject verifyAddress(String chain, String address) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("chain", chain);
        data.addProperty("address", address);
        return client.post("/verifyAddress", data, JsonObject.class);
    }

    public void abandonAddress(String chain, String address) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("chain", chain);
        data.addProperty("address", address);
        client.post("/abandonAddress", data, Void.class);
    }

    public JsonObject hostedInvoiceOrderInfo(String orderId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("orderId", orderId);
        return client.post("/hostedInvoiceOrderInfo", data, JsonObject.class);
    }

    public JsonObject getPayInfo(String orderId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("orderId", orderId);
        return client.post("/getPayInfo", data, JsonObject.class);
    }
}
