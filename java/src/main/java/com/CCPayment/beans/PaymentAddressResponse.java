package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class PaymentAddressResponse {
    /**
     * The corresponds to the chain address of the network
     */
    @JsonProperty("address")
    private String address;

    /**
     * Tag data of Memo coins, used to label and identify user addresses
     */
    @JsonProperty("memo")
    private String memo;

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getMemo() {
        return memo;
    }

    public void setMemo(String memo) {
        this.memo = memo;
    }
}
