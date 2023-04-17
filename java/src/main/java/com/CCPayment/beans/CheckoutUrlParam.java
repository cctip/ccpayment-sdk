package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class CheckoutUrlParam {
    @JsonProperty("order_valid_period")
    private Integer orderValidPeriod;
    @JsonProperty("product_name")
    private String productName;
    @JsonProperty("return_url")
    private String returnUrl;
    @JsonProperty("product_price")
    private String productPrice;
    @JsonProperty("merchant_order_id")
    private String merchantOrderId;

    public Integer getOrderValidPeriod() {
        return orderValidPeriod;
    }

    public void setOrderValidPeriod(Integer orderValidPeriod) {
        this.orderValidPeriod = orderValidPeriod;
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

    public String getProductPrice() {
        return ProductPrice;
    }

    public void setProductPrice(String productPrice) {
        this.amount = productPrice;
    }

    public String getMerchantOrderId() {
        return merchantOrderId;
    }

    public void setMerchantOrderId(String merchantOrderId) {
        this.merchantOrderId = merchantOrderId;
    }
}
