package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class GetSupportCoinResponse {
    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class TokenChain {
        @JsonProperty("crypto")
        private String crypto;
        @JsonProperty("logo")
        private String logo;
        @JsonProperty("name")
        private String name;
        @JsonProperty("network")
        private String network;
        @JsonProperty("chain")
        private String chain;
        @JsonProperty("contract")
        private String contract;
        @JsonProperty("is_support_memo")
        private Boolean isSupportMemo;
        @JsonProperty("chain_logo")
        private String chainLogo;
        @JsonProperty("token_id")
        private String tokenId;
        @JsonProperty("status")
        private int status;
        @JsonProperty("precision")
        private int precision;

        public Boolean getSupportMemo() {
            return isSupportMemo;
        }

        public void setSupportMemo(Boolean supportMemo) {
            isSupportMemo = supportMemo;
        }

        public int getStatus() {
            return status;
        }

        public void setStatus(int status) {
            this.status = status;
        }

        public int getPrecision() {
            return precision;
        }

        public void setPrecision(int precision) {
            this.precision = precision;
        }

        public String getLogo() {
            return logo;
        }

        public void setLogo(String logo) {
            this.logo = logo;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public String getNetwork() {
            return network;
        }

        public void setNetwork(String network) {
            this.network = network;
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

        public Boolean isSupportMemo() {
            return isSupportMemo;
        }

        public void setIsSupportMemo(Boolean isSupportMemo) {
            this.isSupportMemo = isSupportMemo;
        }

        public String getChainLogo() {
            return chainLogo;
        }

        public void setChainLogo(String chainLogo) {
            this.chainLogo = chainLogo;
        }

        public String getCrypto() {
            return crypto;
        }

        public String getTokenId() {
            return tokenId;
        }

        public void setTokenId(String tokenId) {
            this.tokenId = tokenId;
        }

        public void setCrypto(String crypto) {
            this.crypto = crypto;
        }
    }
    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class TokenCoin{
        @JsonProperty("symbol")
        private String symbol;
        @JsonProperty("crypto")
        private String crypto;
        @JsonProperty("name")
        private String name;
        @JsonProperty("logo")
        private String logo;
        @JsonProperty("min")
        private String min;
        @JsonProperty("price")
        private String price;
        @JsonProperty("coin_id")
        private String coinId;
        @JsonProperty("tokens")
        private TokenChain[] tokens;
        @JsonProperty("status")
        private int status;

        public String getSymbol() {
            return symbol;
        }

        public void setSymbol(String symbol) {
            this.symbol = symbol;
        }

        public int getStatus() {
            return status;
        }

        public void setStatus(int status) {
            this.status = status;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public String getLogo() {
            return logo;
        }

        public void setLogo(String logo) {
            this.logo = logo;
        }

        public String getMin() {
            return min;
        }

        public void setMin(String min) {
            this.min = min;
        }

        public String getPrice() {
            return price;
        }

        public void setPrice(String price) {
            this.price = price;
        }

        public String getCoinId() {
            return coinId;
        }

        public void setCoinId(String coinId) {
            this.coinId = coinId;
        }

        public String getCrypto() {
            return crypto;
        }

        public void setCrypto(String crypto) {
            this.crypto = crypto;
        }

        public TokenChain[] getTokens() {
            return tokens;
        }

        public void setTokens(TokenChain[] tokens) {
            this.tokens = tokens;
        }
    }
    @JsonProperty("list")
    private TokenCoin[] list;

    public TokenCoin[] getList() {
        return list;
    }

    public void setList(TokenCoin[] list) {
        this.list = list;
    }
}
