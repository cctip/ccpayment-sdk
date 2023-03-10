package com.Ccpayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Create order params
 * doc: https://doc.ccpayment.com/ccpayment-for-developer/interface-details/create-ccpayment-order
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class PaymentIntentCreateParam {
    /**
     * the server notify callback url
     * when payment order is success, we'll notify to your server by this url
     */
    @JsonProperty("notify_url")
    private String notifyUrl;
    /**
     *
     */
    @JsonProperty("remark")
    private String remark;
    /**
     *
     */
    @JsonProperty("token_id")
    private String tokenId;
    @JsonProperty("amount")
    private String amount;
    @JsonProperty("merchant_order_id")
    private String merchantOrderId;
    @JsonProperty("fiat_currency")
    private String fiatCurrency;


    public String getNotifyUrl() {
        return notifyUrl;
    }

    public void setNotifyUrl(String notifyUrl) {
        this.notifyUrl = notifyUrl;
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

    public String getFiatCurrency() {
        return fiatCurrency;
    }

    public void setFiatCurrency(String fiatCurrency) {
        this.fiatCurrency = fiatCurrency;
    }
}
