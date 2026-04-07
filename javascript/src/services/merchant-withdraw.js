class MerchantWithdrawService {
  constructor(client) { this.client = client; }
  async getCwalletUserId(cwalletUserId) { return this.client.post('/getCwalletUserId', { cwalletUserId }); }
  async getWithdrawFee(coinId, chain) { return this.client.post('/getWithdrawFee', { coinId, chain }); }
  async applyAppWithdrawToNetwork(data) { return this.client.post('/applyAppWithdrawToNetwork', data); }
  async applyAppWithdrawToCwallet(data) { return this.client.post('/applyAppWithdrawToCwallet', data); }
  async getAppWithdrawRecord(data) { return this.client.post('/getAppWithdrawRecord', data); }
  async getAppWithdrawRecordList(data = {}) { return this.client.post('/getAppWithdrawRecordList', data); }
  async getAutoWithdrawRecordList(data = {}) { return this.client.post('/getAutoWithdrawRecordList', data); }
  async getRiskRefundRecords(data = {}) { return this.client.post('/getRiskyRefundRecordList', data); }
}
module.exports = MerchantWithdrawService;
