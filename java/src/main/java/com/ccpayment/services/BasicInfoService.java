package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;
import java.util.List;

public class BasicInfoService {
    private final Client client;

    public BasicInfoService(Client client) {
        this.client = client;
    }

    public JsonObject getCoinList() throws APIError, IOException {
        return client.post("/getCoinList", null, JsonObject.class);
    }

    public JsonObject getCoin(long coinId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("coinId", coinId);
        return client.post("/getCoin", data, JsonObject.class);
    }

    public JsonObject getCoinUSDTPrice(List<Long> coinIds) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.add("coinIds", com.google.gson.JsonParser.parseString(coinIds.toString()));
        return client.post("/getCoinUSDTPrice", data, JsonObject.class);
    }

    public JsonObject getFiatList() throws APIError, IOException {
        return client.post("/getFiatList", null, JsonObject.class);
    }

    public JsonObject getChainList(List<String> chains) throws APIError, IOException {
        JsonObject data = new JsonObject();
        if (chains != null) {
            data.add("chains", com.google.gson.JsonParser.parseString(chains.toString()));
        }
        return client.post("/getChainList", data, JsonObject.class);
    }

    public JsonObject allChain(List<String> chains) throws APIError, IOException {
        JsonObject data = new JsonObject();
        if (chains != null) {
            data.add("chains", com.google.gson.JsonParser.parseString(chains.toString()));
        }
        return client.post("/all/chain", data, JsonObject.class);
    }

    public JsonObject getCwalletUserId(String cwalletUserId) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("cwalletUserId", cwalletUserId);
        return client.post("/getCwalletUserId", data, JsonObject.class);
    }

    public JsonObject getWithdrawFee(long coinId, String chain) throws APIError, IOException {
        JsonObject data = new JsonObject();
        data.addProperty("coinId", coinId);
        data.addProperty("chain", chain);
        return client.post("/getWithdrawFee", data, JsonObject.class);
    }
}
