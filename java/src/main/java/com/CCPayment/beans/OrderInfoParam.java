package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class OrderInfoParam {
    @JsonProperty("merchant_order_ids")
    private String[] merchantOrderIds;

    public String[] getMerchantOrderIds() {
        return merchantOrderIds;
    }

    public void setMerchantOrderIds(String[] merchantOrderIds) {
        this.merchantOrderIds = merchantOrderIds;
    }
}
