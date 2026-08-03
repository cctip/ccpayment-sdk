class BasicInfoService {
  constructor(client) { this.client = client; }
  async getCoinList() { return this.client.post('/getCoinList', {}); }
  async getFiatList() { return this.client.post('/getFiatList', {}); }
  async getCoin(coinId) { return this.client.post('/getCoin', { coinId }); }
  async getCoinUSDTPrice(coinIds) { return this.client.post('/getCoinUSDTPrice', { coinIds }); }
  async getChainList(data = {}) { return this.client.post('/getChainList', data); }
  async getAllChainList(data = {}) { return this.client.post('/all/chain', data); }
  async getMainCoinList(appId) { return this.client.post('/getMainCoinList', { appId }); }
}
module.exports = BasicInfoService;
