package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class PaymentAddressParam {
    /**
     * User ID, unique identification
     */
    @JsonProperty("user_id")
    private String UserId;

    /**
     * Blockchain network, unique identification
     */
    @JsonProperty("chain")
    private String chain;

    /**
     * The URL address will be notified via a POST request when the order status changes.
     * Ensure the URL is accessible to receive notifications from the payment platform.
     */
    @JsonProperty("notify_url")
    private String notifyUrl;

    public String getUserId() {
        return UserId;
    }

    public void setUserId(String userId) {
        UserId = userId;
    }

    public String getChain() {
        return chain;
    }

    public void setChain(String chain) {
        this.chain = chain;
    }

    public String getNotifyUrl() {
        return notifyUrl;
    }

    public void setNotifyUrl(String notifyUrl) {
        this.notifyUrl = notifyUrl;
    }
}
