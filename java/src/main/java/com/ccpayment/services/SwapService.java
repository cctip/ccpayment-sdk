package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class SwapService {
    private final Client client;

    public SwapService(Client client) {
        this.client = client;
    }

    public JsonObject getSwapCoinList() throws APIError, IOException {
        return client.post("/getSwapCoinList", null, JsonObject.class);
    }

    public JsonObject swapEstimate(JsonObject data) throws APIError, IOException {
        return client.post("/swapEstimate", data, JsonObject.class);
    }

    public JsonObject swap(JsonObject data) throws APIError, IOException {
        return client.post("/swap", data, JsonObject.class);
    }

    public JsonObject getSwapRecord(JsonObject data) throws APIError, IOException {
        return client.post("/getSwapRecord", data, JsonObject.class);
    }

    public JsonObject getSwapRecordList(JsonObject data) throws APIError, IOException {
        return client.post("/getSwapRecordList", data, JsonObject.class);
    }

    public JsonObject selectHostedInvoiceCoin(JsonObject data) throws APIError, IOException {
        return client.post("/selectHostedInvoiceCoin", data, JsonObject.class);
    }
}
