"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.BasicInfoService = void 0;
class BasicInfoService {
    constructor(client) {
        this.client = client;
    }
    async getCoinList() {
        return this.client.post('/getCoinList', {});
    }
    async getCoin(coinId) {
        return this.client.post('/getCoin', { coinId });
    }
    async getCoinUSDTPrice(coinIds) {
        return this.client.post('/getCoinUSDTPrice', { coinIds });
    }
    async getFiatList() {
        return this.client.post('/getFiatList', {});
    }
    async getChainList(chains) {
        return this.client.post('/getChainList', chains ? { chains } : {});
    }
    async allChain(chains) {
        return this.client.post('/all/chain', chains ? { chains } : {});
    }
    async getCwalletUserId(cwalletUserId) {
        return this.client.post('/getCwalletUserId', { cwalletUserId });
    }
    async getWithdrawFee(coinId, chain) {
        return this.client.post('/getWithdrawFee', { coinId, chain });
    }
}
exports.BasicInfoService = BasicInfoService;
//# sourceMappingURL=basic-info.js.map