package com.Ccpayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class CheckoutUrlParam {
    @JsonProperty("valid_timestamp")
    private Integer validTimestamp;
    @JsonProperty("product_name")
    private String productName;
    @JsonProperty("return_url")
    private String returnUrl;
    @JsonProperty("amount")
    private String amount;
    @JsonProperty("merchant_order_id")
    private String merchantOrderId;

    public Integer getValidTimestamp() {
        return validTimestamp;
    }

    public void setValidTimestamp(Integer validTimestamp) {
        this.validTimestamp = validTimestamp;
    }

    public String getProductName() {
        return productName;
    }

    public void setProductName(String productName) {
        this.productName = productName;
    }

    public String getReturnUrl() {
        return returnUrl;
    }

    public void setReturnUrl(String returnUrl) {
        this.returnUrl = returnUrl;
    }

    public String getAmount() {
        return amount;
    }

    public void setAmount(String amount) {
        this.amount = amount;
    }

    public String getMerchantOrderId() {
        return merchantOrderId;
    }

    public void setMerchantOrderId(String merchantOrderId) {
        this.merchantOrderId = merchantOrderId;
    }
}
