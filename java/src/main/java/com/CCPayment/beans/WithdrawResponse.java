package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class WithdrawResponse {
    /**
     * Merchant order ID
     */
    @JsonProperty("order_id")
    private String orderId;

    /**
     * Network fee. Withdrawal to Cwallet address charges 0 network fee.
     */
    @JsonProperty("network_fee")
    private String networkFee;

    /**
     *Trading record: one transaction generates one unique
     */
    @JsonProperty("record_id")
    private String recordId;

    /**
     * Amount to be received from withdrawal
     */
    @JsonProperty("net_receivable")
    private String netReceivable;

    public String getRecordId() {
        return recordId;
    }

    public void setRecordId(String recordId) {
        this.recordId = recordId;
    }

    public String getNetReceivable() {
        return netReceivable;
    }

    public void setNetReceivable(String netReceivable) {
        this.netReceivable = netReceivable;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    public String getNetworkFee() {
        return networkFee;
    }

    public void setNetworkFee(String networkFee) {
        this.networkFee = networkFee;
    }
}
