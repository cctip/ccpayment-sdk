package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UserSwapService {
    private final Client client;
    public UserSwapService(Client client) { this.client = client; }
    public JsonObject getUserSwapRecord(JsonObject data) throws APIError, IOException { return client.post("/getUserSwapRecord", data, JsonObject.class); }
    public JsonObject getUserSwapRecordList(JsonObject data) throws APIError, IOException { return client.post("/getUserSwapRecordList", data, JsonObject.class); }
    public JsonObject userSwap(JsonObject data) throws APIError, IOException { return client.post("/userSwap", data, JsonObject.class); }
}
