package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class MerchantWithdrawService {
    private final Client client;
    public MerchantWithdrawService(Client client) { this.client = client; }
    public JsonObject getCwalletUserId(JsonObject data) throws APIError, IOException { return client.post("/getCwalletUserId", data, JsonObject.class); }
    public JsonObject getWithdrawFee(JsonObject data) throws APIError, IOException { return client.post("/getWithdrawFee", data, JsonObject.class); }
    public JsonObject applyAppWithdrawToNetwork(JsonObject data) throws APIError, IOException { return client.post("/applyAppWithdrawToNetwork", data, JsonObject.class); }
    public JsonObject applyAppWithdrawToCwallet(JsonObject data) throws APIError, IOException { return client.post("/applyAppWithdrawToCwallet", data, JsonObject.class); }
    public JsonObject getAppWithdrawRecord(JsonObject data) throws APIError, IOException { return client.post("/getAppWithdrawRecord", data, JsonObject.class); }
    public JsonObject getAppWithdrawRecordList(JsonObject data) throws APIError, IOException { return client.post("/getAppWithdrawRecordList", data, JsonObject.class); }
    public JsonObject getAutoWithdrawRecordList(JsonObject data) throws APIError, IOException { return client.post("/getAutoWithdrawRecordList", data, JsonObject.class); }
    public JsonObject getRiskRefundRecords(JsonObject data) throws APIError, IOException { return client.post("/getRiskyRefundRecordList", data, JsonObject.class); }
}
