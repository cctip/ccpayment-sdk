package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UserWithdrawService {
    private final Client client;

    public UserWithdrawService(Client client) {
        this.client = client;
    }

    public JsonObject applyUserWithdrawToNetwork(JsonObject data) throws APIError, IOException {
        return client.post("/applyUserWithdrawToNetwork", data, JsonObject.class);
    }

    public JsonObject applyUserWithdrawToCwallet(JsonObject data) throws APIError, IOException {
        return client.post("/applyUserWithdrawToCwallet", data, JsonObject.class);
    }

    public JsonObject getUserWithdrawRecord(JsonObject data) throws APIError, IOException {
        return client.post("/getUserWithdrawRecord", data, JsonObject.class);
    }

    public JsonObject getUserWithdrawRecordList(JsonObject data) throws APIError, IOException {
        return client.post("/getUserWithdrawRecordList", data, JsonObject.class);
    }
}
