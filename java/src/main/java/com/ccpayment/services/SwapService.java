package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class SwapService {
    private final Client client;
    public SwapService(Client client) { this.client = client; }
    /** @deprecated Use {@link #getSwapPairList()}. */
    @Deprecated
    public JsonObject getSwapCoinList() throws APIError, IOException { return client.post("/getSwapCoinList", null, JsonObject.class); }
    /** @deprecated Use {@link #swapEstimateV2(JsonObject)}. */
    @Deprecated
    public JsonObject swapEstimate(JsonObject data) throws APIError, IOException { return client.post("/estimate", data, JsonObject.class); }
    /** @deprecated Use {@link #getSwapRecordV2(JsonObject)}. */
    @Deprecated
    public JsonObject getSwapRecord(JsonObject data) throws APIError, IOException { return client.post("/getSwapRecord", data, JsonObject.class); }
    /** @deprecated Use {@link #getSwapRecordListV2(JsonObject)}. */
    @Deprecated
    public JsonObject getSwapRecordList(JsonObject data) throws APIError, IOException { return client.post("/getSwapRecordList", data, JsonObject.class); }
    /** @deprecated Use {@link #swapV2(JsonObject)}. */
    @Deprecated
    public JsonObject swap(JsonObject data) throws APIError, IOException { return client.post("/swap", data, JsonObject.class); }
    public JsonObject getSwapPairList() throws APIError, IOException { return client.post("/swapv2/getSwapPairList", null, JsonObject.class); }
    public JsonObject swapEstimateV2(JsonObject data) throws APIError, IOException { return client.post("/swapv2/estimate", data, JsonObject.class); }
    public JsonObject swapV2(JsonObject data) throws APIError, IOException { return client.post("/swapv2/swap", data, JsonObject.class); }
    public JsonObject getSwapRecordV2(JsonObject data) throws APIError, IOException { return client.post("/swapv2/getSwapRecord", data, JsonObject.class); }
    public JsonObject getSwapRecordListV2(JsonObject data) throws APIError, IOException { return client.post("/swapv2/getSwapRecordList", data, JsonObject.class); }
}
