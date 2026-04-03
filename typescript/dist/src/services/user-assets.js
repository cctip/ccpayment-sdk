"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserAssetsService = void 0;
class UserAssetsService {
    constructor(client) {
        this.client = client;
    }
    async getUserCoinAssetList(userId) {
        return this.client.post('/getUserCoinAssetList', { userId });
    }
    async getUserCoinAsset(userId, coinId) {
        return this.client.post('/getUserCoinAsset', { userId, coinId });
    }
}
exports.UserAssetsService = UserAssetsService;
//# sourceMappingURL=user-assets.js.map