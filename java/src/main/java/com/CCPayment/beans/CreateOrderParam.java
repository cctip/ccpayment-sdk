package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Create order params
 * doc: https://doc.ccpayment.com/ccpayment-for-developer/interface-details/create-ccpayment-order
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class CreateOrderParam {

    /**
     * Remark
     */
    @JsonProperty("remark")
    private String remark;

    /**
     * Tells CCPayment's server which coin and which network should be used for the transaction.
     * You can get the token_id by calling the interface or find it on this sheet
     */
    @JsonProperty("token_id")
    private String tokenId;

    /**
     * Amount should be paid for this order (in USD by default, no more than two digits after the dot)
     */
    @JsonProperty("product_price")
    private String productPrice;

    /**
     * Order ID in Merchant's system. A unique ID for every order
     */
    @JsonProperty("merchant_order_id")
    private String merchantOrderId;

    /**
     * Currency for the order.
     * Pass "token" if you want to show the price in crypto.
     * Pass corresponding currency code if you want to show the price in fiat.
     */
    @JsonProperty("denominated_currency")
    private String denominatedCurrency;

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

    public Integer getOrderValidPeriod() {
        return orderValidPeriod;
    }

    public void setOrderValidPeriod(Integer orderValidPeriod) {
        this.orderValidPeriod = orderValidPeriod;
    }

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

    public String getRemark() {
        return remark;
    }

    public void setRemark(String remark) {
        this.remark = remark;
    }

    public String getTokenId() {
        return tokenId;
    }

    public void setTokenId(String tokenId) {
        this.tokenId = tokenId;
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

    public String getDenominatedCurrency() {
        return denominatedCurrency;
    }

    public void setDenominatedCurrency(String denominatedCurrency) {
        this.denominatedCurrency = denominatedCurrency;
    }
}
