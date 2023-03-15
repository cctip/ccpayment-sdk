package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class OrderInfoResponse {
    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class Detail {
        @JsonProperty("fiat_amount")
        private String fiatAmount;
        @JsonProperty("fiat_currency")
        private String fiatCurrency;
        @JsonProperty("product_name")
        private String productName;
        @JsonProperty("merchant_order_id")
        private String merchantOrderId;
        @JsonProperty("chain")
        private String chain;
        @JsonProperty("contract")
        private String contract;
        @JsonProperty("crypto")
        private String crypto;
        @JsonProperty("origin_amount")
        private String originAmount;
        @JsonProperty("status")
        private String status;
        @JsonProperty("created")
        private String created;

        public String getCreated() {
            return created;
        }

        public String getFiatAmount() {
            return fiatAmount;
        }

        public void setFiatAmount(String fiatAmount) {
            this.fiatAmount = fiatAmount;
        }

        public String getFiatCurrency() {
            return fiatCurrency;
        }

        public void setFiatCurrency(String fiatCurrency) {
            this.fiatCurrency = fiatCurrency;
        }

        public String getProductName() {
            return productName;
        }

        public void setProductName(String productName) {
            this.productName = productName;
        }

        public String getMerchantOrderId() {
            return merchantOrderId;
        }

        public void setMerchantOrderId(String merchantOrderId) {
            this.merchantOrderId = merchantOrderId;
        }

        public String getChain() {
            return chain;
        }

        public void setChain(String chain) {
            this.chain = chain;
        }

        public String getContract() {
            return contract;
        }

        public void setContract(String contract) {
            this.contract = contract;
        }

        public String getCrypto() {
            return crypto;
        }

        public void setCrypto(String crypto) {
            this.crypto = crypto;
        }

        public String getOriginAmount() {
            return originAmount;
        }

        public void setOriginAmount(String originAmount) {
            this.originAmount = originAmount;
        }

        public String getStatus() {
            return status;
        }

        public void setStatus(String status) {
            this.status = status;
        }
        public void setCreated(String created) {
            this.created = created;
        }
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class Received {
        @JsonProperty("amount")
        private String amount;
        @JsonProperty("chain")
        private String chain;
        @JsonProperty("contract")
        private String contract;
        @JsonProperty("crypto")
        private String crypto;
        @JsonProperty("service_fee")
        private String serviceFee;
        @JsonProperty("pay_time")
        private String payTime;
        @JsonProperty("token_rate")
        private String tokenRate;

        public String getAmount() {
            return amount;
        }

        public void setAmount(String amount) {
            this.amount = amount;
        }

        public String getChain() {
            return chain;
        }

        public void setChain(String chain) {
            this.chain = chain;
        }

        public String getContract() {
            return contract;
        }

        public void setContract(String contract) {
            this.contract = contract;
        }

        public String getCrypto() {
            return crypto;
        }

        public void setCrypto(String crypto) {
            this.crypto = crypto;
        }

        public String getServiceFee() {
            return serviceFee;
        }

        public void setServiceFee(String serviceFee) {
            this.serviceFee = serviceFee;
        }

        public String getTokenRate() {
            return tokenRate;
        }

        public void setTokenRate(String tokenRate) {
            this.tokenRate = tokenRate;
        }
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class Refunds {
        @JsonProperty("amount")
        private String amount;
        @JsonProperty("network_fee")
        private String networkFee;
        @JsonProperty("actual_received_amount")
        private String actualReceivedAmount;
        @JsonProperty("chain")
        private String chain;
        @JsonProperty("contract")
        private String contract;
        @JsonProperty("crypto")
        private String crypto;
        @JsonProperty("txid")
        private String txId;
        @JsonProperty("address")
        private String address;
        @JsonProperty("pay_time")
        private Long payTime;
        @JsonProperty("status")
        private String status;

        public String getAmount() {
            return amount;
        }

        public void setAmount(String amount) {
            this.amount = amount;
        }

        public String getNetworkFee() {
            return networkFee;
        }

        public void setNetworkFee(String networkFee) {
            this.networkFee = networkFee;
        }

        public String getActualReceivedAmount() {
            return actualReceivedAmount;
        }

        public void setActualReceivedAmount(String actualReceivedAmount) {
            this.actualReceivedAmount = actualReceivedAmount;
        }

        public String getChain() {
            return chain;
        }

        public void setChain(String chain) {
            this.chain = chain;
        }

        public String getContract() {
            return contract;
        }

        public void setContract(String contract) {
            this.contract = contract;
        }

        public String getCrypto() {
            return crypto;
        }

        public void setCrypto(String crypto) {
            this.crypto = crypto;
        }

        public String getTxId() {
            return txId;
        }

        public void setTxId(String txId) {
            this.txId = txId;
        }

        public String getAddress() {
            return address;
        }

        public void setAddress(String address) {
            this.address = address;
        }

        public Long getPayTime() {
            return payTime;
        }

        public void setPayTime(Long payTime) {
            this.payTime = payTime;
        }

        public String getStatus() {
            return status;
        }

        public void setStatus(String status) {
            this.status = status;
        }
    }

    @JsonProperty("detail")
    private Detail detail;
    @JsonProperty("received")
    private Received[] received;
    @JsonProperty("refunds")
    private Refunds[] refunds;

    public Detail getDetail() {
        return detail;
    }

    public void setDetail(Detail detail) {
        this.detail = detail;
    }

    public Received[] getReceived() {
        return received;
    }

    public void setReceived(Received[] received) {
        this.received = received;
    }

    public Refunds[] getRefunds() {
        return refunds;
    }

    public void setRefunds(Refunds[] refunds) {
        this.refunds = refunds;
    }
}
