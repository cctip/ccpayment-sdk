package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class GetSupportTokenResponse {
    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class TokenChain{
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
        @JsonProperty("token_id")
        private String tokenId;

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

        public String getTokenId() {
            return tokenId;
        }

        public void setTokenId(String tokenId) {
            this.tokenId = tokenId;
        }

        public String getCrypto() {
            return crypto;
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
