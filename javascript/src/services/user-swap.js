class UserSwapService {
  constructor(client) { this.client = client; }
  async getUserSwapRecord(data) { return this.client.post('/getUserSwapRecord', data); }
  async getUserSwapRecordList(data = {}) { return this.client.post('/getUserSwapRecordList', data); }
  async userSwap(data) { return this.client.post('/userSwap', data); }
}
module.exports = UserSwapService;
