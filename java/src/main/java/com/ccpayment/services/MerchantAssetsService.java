package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class MerchantAssetsService {
    private final Client client;
    public MerchantAssetsService(Client client) { this.client = client; }
    public JsonObject getAppCoinAssetList() throws APIError, IOException { return client.post("/getAppCoinAssetList", null, JsonObject.class); }
    public JsonObject getAppCoinAsset(JsonObject data) throws APIError, IOException { return client.post("/getAppCoinAsset", data, JsonObject.class); }
    public JsonObject getAppCollectFeeRecordList(JsonObject data) throws APIError, IOException { return client.post("/getAggregationFeeRecord", data, JsonObject.class); }
}
