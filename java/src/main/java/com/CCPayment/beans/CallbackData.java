package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class CallbackData {

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class Extend {
        @JsonProperty("invoice_id")
        private String invoiceId;
        @JsonProperty("user_email")
        private String userEmail;
        @JsonProperty("merchant_order_id")
        private String merchantOrderId;
    }

    @JsonProperty("pay_status")
    private String payStatus;
    @JsonProperty("order_type")
    private String orderType;
    @JsonProperty("record_id")
    private String recordId;
    @JsonProperty("order_id")
    private String orderId;
    @JsonProperty("origin_price")
    private String originPrice;
    @JsonProperty("origin_amount")
    private String originAmount;
    @JsonProperty("fiat_rate")
    private String fiatRate;
    @JsonProperty("paid_amount")
    private String paidAmount;
    @JsonProperty("token_rate")
    private String tokenRate;
    @JsonProperty("chain")
    private String chain;
    @JsonProperty("contract")
    private String contract;
    @JsonProperty("crypto")
    private String crypto;
    @JsonProperty("extend")
    private Extend extend;

    public String getPayStatus() {
        return payStatus;
    }

    public void setPayStatus(String payStatus) {
        this.payStatus = payStatus;
    }

    public String getOrderType() {
        return orderType;
    }

    public void setOrderType(String orderType) {
        this.orderType = orderType;
    }

    public String getRecordId() {
        return recordId;
    }

    public void setRecordId(String recordId) {
        this.recordId = recordId;
    }

    public String getOriginPrice() {
        return originPrice;
    }

    public void setOriginPrice(String originPrice) {
        this.originPrice = originPrice;
    }

    public String getOriginAmount() {
        return originAmount;
    }

    public void setOriginAmount(String originAmount) {
        this.originAmount = originAmount;
    }

    public String getFiatRate() {
        return fiatRate;
    }

    public void setFiatRate(String fiatRate) {
        this.fiatRate = fiatRate;
    }

    public String getPaidAmount() {
        return paidAmount;
    }

    public void setPaidAmount(String paidAmount) {
        this.paidAmount = paidAmount;
    }

    public String getTokenRate() {
        return tokenRate;
    }

    public void setTokenRate(String tokenRate) {
        this.tokenRate = tokenRate;
    }

    public String getChain() {
        return chain;
    }

    public void setChain(String chain) {
        this.chain = chain;
    }

    public String getContract() {
        return contract;
    }

    public void setContract(String contract) {
        this.contract = contract;
    }

    public Extend getExtend() {
        return extend;
    }

    public void setExtend(Extend extend) {
        this.extend = extend;
    }

    public String getCrypto() {
        return crypto;
    }

    public void setCrypto(String crypto) {
        this.crypto = crypto;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }
}
