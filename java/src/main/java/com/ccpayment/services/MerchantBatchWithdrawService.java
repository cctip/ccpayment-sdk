package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

/**
 * Merchant batch withdraw service.
 * Request fields: taskName, orders[].orderId, orderIds, confirmExecution, stats.execOrderId.
 */
public class MerchantBatchWithdrawService {
    private final Client client;

    public MerchantBatchWithdrawService(Client client) {
        this.client = client;
    }

    public JsonObject checkWithdrawAddress(JsonObject data) throws APIError, IOException {
        return client.post("/checkWithdrawAddress", data, JsonObject.class);
    }

    public JsonObject applyBatchWithdraw(JsonObject data) throws APIError, IOException {
        return client.post("/applyBatchWithdraw", data, JsonObject.class);
    }

    public JsonObject appendBatchWithdraw(JsonObject data) throws APIError, IOException {
        return client.post("/appendBatchWithdraw", data, JsonObject.class);
    }

    public JsonObject confirmBatchWithdraw(JsonObject data) throws APIError, IOException {
        return client.post("/confirmBatchWithdraw", data, JsonObject.class);
    }

    public JsonObject abortBatchWithdraw(JsonObject data) throws APIError, IOException {
        return client.post("/abortBatchWithdraw", data, JsonObject.class);
    }

    public JsonObject getBatchWithdrawOrder(JsonObject data) throws APIError, IOException {
        return client.post("/getBatchWithdrawOrder", data, JsonObject.class);
    }

    public JsonObject getBatchWithdrawOrderRecordList(JsonObject data) throws APIError, IOException {
        return client.post("/getBatchWithdrawOrderRecordList", data, JsonObject.class);
    }
}
