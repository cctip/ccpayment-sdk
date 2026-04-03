package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UserDepositService {
    private final Client client;

    public UserDepositService(Client client) {
        this.client = client;
    }

    public JsonObject getOrCreateUserDepositAddress(JsonObject data) throws APIError, IOException {
        return client.post("/getOrCreateUserDepositAddress", data, JsonObject.class);
    }

    public JsonObject getUserDepositRecord(String recordId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("recordId", recordId);
        return client.post("/getUserDepositRecord", data, JsonObject.class);
    }

    public JsonObject getUserDepositRecordList(JsonObject data) throws APIError, IOException {
        return client.post("/getUserDepositRecordList", data, JsonObject.class);
    }
}
