package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class WithdrawParam {
    /**
     * Tell CCPayment's server which coin and network should be used for the transaction.
     * You can get the token_id by calling the interface or finding it on this sheet.
     * If the withdrawal is to Cwallet, any token_id of the coin will be ok.
     * If the withdrawal is to external addresses, only the corresponding token_id will work.
     */
    @JsonProperty("token_id")
    private String tokenId;

    /**
     *Case 1: Pass the Cwallet ID or Email address linked to Cwallet account If a user wants to withdraw to his Cwallet account
     *Case 2: Pass the receiving address if the user wants to withdraw to his Web3 wallet.
     *Note: SATS uses the invoice as a receiving address. Each invoice can only accept one payment.
     */
    @JsonProperty("address")
    private String address;

    /**
     *Pass memo parameter if the receiving address requires a memo.
     *If a memo is required and it is not filled, or filled incorrectly, the asset may be lost
     */
    @JsonProperty("memo")
    private String memo;

    /**
     *Withdrawal amount
     */
    @JsonProperty("value")
    private String value;

    /**
     *Order ID in Merchant's system. A unique ID for every order
     */
    @JsonProperty("merchant_order_id")
    private String merchantOrderId;

    /**
     *True: merchant pays the network fee
     *False: not specified: uer pays the network fee, amount received by user = withdrawal amount - network fee
     */
    @JsonProperty("merchant_pays_fee")
    private Boolean merchantPaysFee;


    public Boolean getMerchantPaysFee() {
        return merchantPaysFee;
    }

    public void setMerchantPaysFee(Boolean merchantPaysFee) {
        this.merchantPaysFee = merchantPaysFee;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public String getTokenId() {
        return tokenId;
    }

    public void setTokenId(String tokenId) {
        this.tokenId = tokenId;
    }

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

    public String getMerchantOrderId() {
        return merchantOrderId;
    }

    public void setMerchantOrderId(String merchantOrderId) {
        this.merchantOrderId = merchantOrderId;
    }
}
