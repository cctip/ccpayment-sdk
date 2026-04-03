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

module.exports = UserAssetsService;
