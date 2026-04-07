package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class BasicInfoService {
    private final Client client;
    public BasicInfoService(Client client) { this.client = client; }
    public JsonObject getCoinList() throws APIError, IOException { return client.post("/getCoinList", null, JsonObject.class); }
    public JsonObject getFiatList() throws APIError, IOException { return client.post("/getFiatList", null, JsonObject.class); }
    public JsonObject getCoin(JsonObject data) throws APIError, IOException { return client.post("/getCoin", data, JsonObject.class); }
    public JsonObject getCoinUSDTPrice(JsonObject data) throws APIError, IOException { return client.post("/getCoinUSDTPrice", data, JsonObject.class); }
    public JsonObject getChainList(JsonObject data) throws APIError, IOException { return client.post("/getChainList", data, JsonObject.class); }
    public JsonObject getAllChainList(JsonObject data) throws APIError, IOException { return client.post("/all/chain", data, JsonObject.class); }
    public JsonObject getMainCoinList(JsonObject data) throws APIError, IOException { return client.post("/getMainCoinList", data, JsonObject.class); }
}
