package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UserAssetsService {
    private final Client client;

    public UserAssetsService(Client client) {
        this.client = client;
    }

    public JsonObject getUserCoinAssetList(String userId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("userId", userId);
        return client.post("/getUserCoinAssetList", data, JsonObject.class);
    }

    public JsonObject getUserCoinAsset(String userId, long coinId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("userId", userId);
        data.addProperty("coinId", coinId);
        return client.post("/getUserCoinAsset", data, JsonObject.class);
    }
}
