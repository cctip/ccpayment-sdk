package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Create order response
 * doc: https://doc.ccpayment.com/ccpayment-for-developer/interface-details/create-ccpayment-order
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CreateOrderResponse {
    @JsonProperty("amount")
    private String amount;
    @JsonProperty("order_id")
    private String orderId;
    @JsonProperty("logo")
    private String logo;
    @JsonProperty("network")
    private String network;
    @JsonProperty("pay_address")
    private String payAddress;
    @JsonProperty("crypto")
    private String crypto;
    @JsonProperty("token_id")
    private String tokenId;
    @JsonProperty("memo")
    private String memo;
    @JsonProperty("address_valid_period")
    private Integer addressValidPeriod;

    public String getAmount() {
        return amount;
    }

    public void setAmount(String amount) {
        this.amount = amount;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    public String getLogo() {
        return logo;
    }

    public void setLogo(String logo) {
        this.logo = logo;
    }

    public String getNetwork() {
        return network;
    }

    public void setNetwork(String network) {
        this.network = network;
    }

    public String getPayAddress() {
        return payAddress;
    }

    public void setPayAddress(String payAddress) {
        this.payAddress = payAddress;
    }

    public String getCrypto() {
        return crypto;
    }

    public void setCrypto(String crypto) {
        this.crypto = crypto;
    }

    public String getTokenId() {
        return tokenId;
    }

    public void setTokenId(String tokenId) {
        this.tokenId = tokenId;
    }

    public String getMemo() {
        return memo;
    }

    public void setMemo(String memo) {
        this.memo = memo;
    }
    public String getAddressValidPeriod() {
        return addressValidPeriod;
    }

    public void setAddressValidPeriod(Integer addressValidPeriod) {
        this.addressValidPeriod = addressValidPeriod;
    }
}
