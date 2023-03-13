package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class GetTokenChainResponse {
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
        @JsonProperty("type_symbol")
        private String typeSymbol;
        @JsonProperty("chain_logo")
        private String chainLogo;
        @JsonProperty("token_id")
        private String tokenId;

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

        public String getTypeSymbol() {
            return typeSymbol;
        }

        public void setTypeSymbol(String typeSymbol) {
            this.typeSymbol = typeSymbol;
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

    @JsonProperty("list")
    private TokenChain[] list;

    public TokenChain[] getList() {
        return list;
    }

    public void setList(TokenChain[] list) {
        this.list = list;
    }
}
