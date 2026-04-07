class SwapService {
  constructor(client) { this.client = client; }
  async getSwapCoinList() { return this.client.post('/getSwapCoinList', {}); }
  async swapEstimate(data) { return this.client.post('/estimate', data); }
  async getSwapRecord(data) { return this.client.post('/getSwapRecord', data); }
  async getSwapRecordList(data = {}) { return this.client.post('/getSwapRecordList', data); }
  async swap(data) { return this.client.post('/swap', data); }
}
module.exports = SwapService;
