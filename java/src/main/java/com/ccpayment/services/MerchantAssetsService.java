package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class MerchantAssetsService {
    private final Client client;

    public MerchantAssetsService(Client client) {
        this.client = client;
    }

    public JsonObject getAppCoinAssetList() throws APIError, IOException {
        return client.post("/getAppCoinAssetList", null, JsonObject.class);
    }

    public JsonObject getAppCoinAsset(long coinId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("coinId", coinId);
        return client.post("/getAppCoinAsset", data, JsonObject.class);
    }
}
