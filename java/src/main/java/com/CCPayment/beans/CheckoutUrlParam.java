package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class CheckoutUrlParam {
    /**
     * Name of the product
     */
    @JsonProperty("product_name")
    private String productName;

    /**
     * Redirect users to this URL after user makes payment
     */
    @JsonProperty("return_url")
    private String returnUrl;

    /**
     * Amount should be paid for this order (in USD by default, no more than two digits after the dot)
     */
    @JsonProperty("product_price")
    private String productPrice;

    /**
     * Order ID in Merchant's system
     */
    @JsonProperty("merchant_order_id")
    private String merchantOrderId;
    /**
     * he valid duration for the order.
     * The parameter passed by the merchant should be less than the order's valid period in Merchant's system.
     * Cause on-chain transactions may need some time to proceed.
     * BTC will arrive within 24 hours and other tokens will usually arrive within 30 minutes.
     * Orders will be valid for 24 hrs by default. The max valid duration is 10 days (max valid duration for Satoshi is 2 hours).
     */
    @JsonProperty("order_valid_period")
    private Integer orderValidPeriod;

    /**
     * The URL address will be notified via a POST request when the order status changes.
     * Ensure the URL is accessible to receive notifications from the payment platform.
     */
    @JsonProperty("notify_url")
    private String notifyUrl;

    /**
     * Merchant custom field - This custom value field will be returned in transaction status notification.
     */
    @JsonProperty("custom_value")
    private String customValue;

    public String getNotifyUrl() {
        return notifyUrl;
    }

    public void setNotifyUrl(String notifyUrl) {
        this.notifyUrl = notifyUrl;
    }

    public String getCustomValue() {
        return customValue;
    }

    public void setCustomValue(String customValue) {
        this.customValue = customValue;
    }

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
        return productPrice;
    }

    public void setProductPrice(String productPrice) {
        this.productPrice = productPrice;
    }

    public String getMerchantOrderId() {
        return merchantOrderId;
    }

    public void setMerchantOrderId(String merchantOrderId) {
        this.merchantOrderId = merchantOrderId;
    }
}
