package com.CCPayment.beans;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class GetChainHeightInfoResponse {

    /**
     * The network's unique identification on the blockchain
     */
    @JsonProperty("chain")
    private String chain;
    /**
     * Current block height of the network
     */
    @JsonProperty("current_chain_height")
    private Long currentChainHeight;
    /**
     * The number of block confirmations required for the trasnaction on the blockchain
     */
    @JsonProperty("tx_confirm_block_num")
    private Long txConfirmBlockNum;
    /**
     * The time gap of every blocks, in seconds
     */
    @JsonProperty("block_rate")
    private Double blockRate;

    public String getChain() {
        return chain;
    }

    public void setChain(String chain) {
        this.chain = chain;
    }

    public Long getCurrentChainHeight() {
        return currentChainHeight;
    }

    public void setCurrentChainHeight(Long currentChainHeight) {
        this.currentChainHeight = currentChainHeight;
    }

    public Long getTxConfirmBlockNum() {
        return txConfirmBlockNum;
    }

    public void setTxConfirmBlockNum(Long txConfirmBlockNum) {
        this.txConfirmBlockNum = txConfirmBlockNum;
    }

    public Double getBlockRate() {
        return blockRate;
    }

    public void setBlockRate(Double blockRate) {
        this.blockRate = blockRate;
    }
}
