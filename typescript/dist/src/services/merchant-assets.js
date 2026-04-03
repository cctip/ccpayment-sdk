"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MerchantAssetsService = void 0;
class MerchantAssetsService {
    constructor(client) {
        this.client = client;
    }
    async getAppCoinAssetList() {
        return this.client.post('/getAppCoinAssetList', {});
    }
    async getAppCoinAsset(coinId) {
        return this.client.post('/getAppCoinAsset', { coinId });
    }
}
exports.MerchantAssetsService = MerchantAssetsService;
//# sourceMappingURL=merchant-assets.js.map