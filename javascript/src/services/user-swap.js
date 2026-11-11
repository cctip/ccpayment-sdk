class UserSwapService {
  constructor(client) { this.client = client; }
  /** @deprecated Use getUserSwapRecordV2. */
  async getUserSwapRecord(data) { return this.client.post('/getUserSwapRecord', data); }
  /** @deprecated Use getUserSwapRecordListV2. */
  async getUserSwapRecordList(data = {}) { return this.client.post('/getUserSwapRecordList', data); }
  /** @deprecated Use userSwapV2 with SwapReqV2 fields. */
  async userSwap(data) { return this.client.post('/userSwap', data); }
  async userSwapV2(data) { return this.client.post('/swapv2/userSwap', data); }
  async getUserSwapRecordV2(data) { return this.client.post('/swapv2/getUserSwapRecord', data); }
  async getUserSwapRecordListV2(data = {}) { return this.client.post('/swapv2/getUserSwapRecordList', data); }
}
module.exports = UserSwapService;
