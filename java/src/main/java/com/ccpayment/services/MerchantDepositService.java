package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class MerchantDepositService {
    private final Client client;

    public MerchantDepositService(Client client) {
        this.client = client;
    }

    public JsonObject createAppOrderDepositAddress(JsonObject data) throws APIError, IOException {
        return client.post("/createAppOrderDepositAddress", data, JsonObject.class);
    }

    public JsonObject getOrCreateAppDepositAddress(JsonObject data) throws APIError, IOException {
        return client.post("/getOrCreateAppDepositAddress", data, JsonObject.class);
    }

    public JsonObject getAppDepositRecord(String recordId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("recordId", recordId);
        return client.post("/getAppDepositRecord", data, JsonObject.class);
    }

    public JsonObject getAppDepositRecordList(JsonObject data) throws APIError, IOException {
        return client.post("/getAppDepositRecordList", data, JsonObject.class);
    }
}
