class SwapService {
  constructor(client) { this.client = client; }
  /** @deprecated Use getSwapPairList. */
  async getSwapCoinList() { return this.client.post('/getSwapCoinList', {}); }
  /** @deprecated Use swapEstimateV2 with SwapReqV2 fields. */
  async swapEstimate(data) { return this.client.post('/estimate', data); }
  /** @deprecated Use getSwapRecordV2. */
  async getSwapRecord(data) { return this.client.post('/getSwapRecord', data); }
  /** @deprecated Use getSwapRecordListV2. */
  async getSwapRecordList(data = {}) { return this.client.post('/getSwapRecordList', data); }
  /** @deprecated Use swapV2 with SwapReqV2 fields. */
  async swap(data) { return this.client.post('/swap', data); }
  async getSwapPairList() { return this.client.post('/swapv2/getSwapPairList', {}); }
  async swapEstimateV2(data) { return this.client.post('/swapv2/estimate', data); }
  async swapV2(data) { return this.client.post('/swapv2/swap', data); }
  async getSwapRecordV2(data) { return this.client.post('/swapv2/getSwapRecord', data); }
  async getSwapRecordListV2(data = {}) { return this.client.post('/swapv2/getSwapRecordList', data); }
}
module.exports = SwapService;
