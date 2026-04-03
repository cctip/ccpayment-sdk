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

module.exports = MerchantAssetsService;
