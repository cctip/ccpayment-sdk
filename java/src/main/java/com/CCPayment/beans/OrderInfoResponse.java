package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class OrderInfoResponse {
    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class OrderDetail {
        @JsonProperty("product_price")
        private String productPrice;
        @JsonProperty("denominated_currency")
        private String denominatedCurrency;
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
        @JsonProperty("order_amount")
        private String orderAmount;
        @JsonProperty("status")
        private String status;
        @JsonProperty("created")
        private Long created;
        @JsonProperty("order_id")
        private String orderId;

        public String getProductPrice() {
            return productPrice;
        }

        public void setProductPrice(String productPrice) {
            this.productPrice = productPrice;
        }

        public String getDenominatedCurrency() {
            return denominatedCurrency;
        }

        public void setDenominatedCurrency(String denominatedCurrency) {
            this.denominatedCurrency = denominatedCurrency;
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

        public String getOrderAmount() {
            return orderAmount;
        }

        public void setOrderAmount(String orderAmount) {
            this.orderAmount = orderAmount;
        }

        public String getStatus() {
            return status;
        }

        public void setStatus(String status) {
            this.status = status;
        }

        public Long getCreated() {
            return created;
        }

        public void setCreated(Long created) {
            this.created = created;
        }

        public String getOrderId() {
            return orderId;
        }

        public void setOrderId(String orderId) {
            this.orderId = orderId;
        }
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class TradeListItem {
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
        @JsonProperty("txid")
        private String txId;
        @JsonProperty("network_fee")
        private String networkFee;
        @JsonProperty("pay_time")
        private Long payTime;
        @JsonProperty("token_rate")
        private String tokenRate;
        @JsonProperty("status")
        private String status;

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

        public String getTxId() {
            return txId;
        }

        public void setTxId(String txId) {
            this.txId = txId;
        }

        public String getNetworkFee() {
            return networkFee;
        }

        public void setNetworkFee(String networkFee) {
            this.networkFee = networkFee;
        }

        public Long getPayTime() {
            return payTime;
        }

        public void setPayTime(Long payTime) {
            this.payTime = payTime;
        }

        public String getTokenRate() {
            return tokenRate;
        }

        public void setTokenRate(String tokenRate) {
            this.tokenRate = tokenRate;
        }

        public String getStatus() {
            return status;
        }

        public void setStatus(String status) {
            this.status = status;
        }
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class RefundListItem {
        @JsonProperty("refund_amount")
        private String refundAmount;
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

        public String getRefundAmount() {
            return refundAmount;
        }

        public void setRefundAmount(String refundAmount) {
            this.refundAmount = refundAmount;
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

    @JsonProperty("order_detail")
    private OrderDetail orderDetail;
    @JsonProperty("trade_list")
    private TradeListItem[] tradeList;
    @JsonProperty("refund_list")
    private RefundListItem[] refundList;

    public OrderDetail getOrderDetail() {
        return orderDetail;
    }

    public void setOrderDetail(OrderDetail orderDetail) {
        this.orderDetail = orderDetail;
    }

    public TradeListItem[] getTradeList() {
        return tradeList;
    }

    public void setTradeList(TradeListItem[] tradeList) {
        this.tradeList = tradeList;
    }

    public RefundListItem[] getRefundList() {
        return refundList;
    }

    public void setRefundList(RefundListItem[] refundList) {
        this.refundList = refundList;
    }
}
