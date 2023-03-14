package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonIgnoreProperties(ignoreUnknown = true)
public class AssetsResponse {
    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class AssetEntity {
        @JsonProperty("token_id")
        private String tokenId;
        @JsonProperty("crypto")
        private String crypto;
        @JsonProperty("name")
        private String name;
        @JsonProperty("value")
        private String value;
        @JsonProperty("price")
        private String price;
        @JsonProperty("logo")
        private String logo;

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

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public String getValue() {
            return value;
        }

        public void setValue(String value) {
            this.value = value;
        }

        public String getPrice() {
            return price;
        }

        public void setPrice(String price) {
            this.price = price;
        }

        public String getLogo() {
            return logo;
        }

        public void setLogo(String logo) {
            this.logo = logo;
        }
    }

    @JsonProperty("list")
    private AssetEntity[] list;

    public AssetEntity[] getList() {
        return list;
    }

    public void setList(AssetEntity[] list) {
        this.list = list;
    }
}
