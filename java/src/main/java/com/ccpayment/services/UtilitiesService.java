package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UtilitiesService {
    private final Client client;
    public UtilitiesService(Client client) { this.client = client; }
    public JsonObject webhookResend(JsonObject data) throws APIError, IOException { return client.post("/webhook/resend", data, JsonObject.class); }
    public JsonObject getTradeBlockHeight(JsonObject data) throws APIError, IOException { return client.post("/getTradeBlockHeight", data, JsonObject.class); }
    public JsonObject checkWithdrawalAddressValidity(JsonObject data) throws APIError, IOException { return client.post("/checkWithdrawalAddressValidity", data, JsonObject.class); }
    public JsonObject deprecatedAddress(JsonObject data) throws APIError, IOException { return client.post("/addressUnbinding", data, JsonObject.class); }
    public JsonObject rescanLostTransaction(JsonObject data) throws APIError, IOException { return client.post("/rescanLostTransaction", data, JsonObject.class); }
}
