package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UserTransferService {
    private final Client client;

    public UserTransferService(Client client) {
        this.client = client;
    }

    public JsonObject userTransfer(JsonObject data) throws APIError, IOException {
        return client.post("/userTransfer", data, JsonObject.class);
    }

    public JsonObject getUserTransferRecord(JsonObject data) throws APIError, IOException {
        return client.post("/getUserTransferRecord", data, JsonObject.class);
    }

    public JsonObject getUserTransferRecordList(JsonObject data) throws APIError, IOException {
        return client.post("/getUserTransferRecordList", data, JsonObject.class);
    }

    public JsonObject userBatchTransfer(JsonObject data) throws APIError, IOException {
        return client.post("/userBatchTransfer", data, JsonObject.class);
    }

    public JsonObject getUserBatchTransferRecord(JsonObject data) throws APIError, IOException {
        return client.post("/getUserBatchTransferRecord", data, JsonObject.class);
    }

    public JsonObject getUserBatchTransferRecordList(JsonObject data) throws APIError, IOException {
        return client.post("/getUserBatchTransferRecordList", data, JsonObject.class);
    }

    public JsonObject getUserBatchTransferRecordDetail(JsonObject data) throws APIError, IOException {
        return client.post("/getUserBatchTransferRecordDetail", data, JsonObject.class);
    }
}
