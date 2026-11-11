package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UserSwapService {
    private final Client client;
    public UserSwapService(Client client) { this.client = client; }
    /** @deprecated Use {@link #getUserSwapRecordV2(JsonObject)}. */
    @Deprecated
    public JsonObject getUserSwapRecord(JsonObject data) throws APIError, IOException { return client.post("/getUserSwapRecord", data, JsonObject.class); }
    /** @deprecated Use {@link #getUserSwapRecordListV2(JsonObject)}. */
    @Deprecated
    public JsonObject getUserSwapRecordList(JsonObject data) throws APIError, IOException { return client.post("/getUserSwapRecordList", data, JsonObject.class); }
    /** @deprecated Use {@link #userSwapV2(JsonObject)}. */
    @Deprecated
    public JsonObject userSwap(JsonObject data) throws APIError, IOException { return client.post("/userSwap", data, JsonObject.class); }
    public JsonObject userSwapV2(JsonObject data) throws APIError, IOException { return client.post("/swapv2/userSwap", data, JsonObject.class); }
    public JsonObject getUserSwapRecordV2(JsonObject data) throws APIError, IOException { return client.post("/swapv2/getUserSwapRecord", data, JsonObject.class); }
    public JsonObject getUserSwapRecordListV2(JsonObject data) throws APIError, IOException { return client.post("/swapv2/getUserSwapRecordList", data, JsonObject.class); }
}
