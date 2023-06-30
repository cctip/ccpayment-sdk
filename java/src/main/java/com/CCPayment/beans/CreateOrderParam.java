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
     *
     */
    @JsonProperty("remark")
    private String remark;
    /**
     *
     */
    @JsonProperty("token_id")
    private String tokenId;
    @JsonProperty("product_price")
    private String productPrice;
    @JsonProperty("merchant_order_id")
    private String merchantOrderId;
    @JsonProperty("denominated_currency")
    private String denominatedCurrency;

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
