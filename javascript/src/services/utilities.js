class UtilitiesService {
  constructor(client) { this.client = client; }
  async webhookResend(data = {}) { return this.client.post('/webhook/resend', data); }
  async getTradeBlockHeight(data) { return this.client.post('/getTradeBlockHeight', data); }
  async checkWithdrawalAddressValidity(data) { return this.client.post('/checkWithdrawalAddressValidity', data); }
  async deprecatedAddress(data) { return this.client.post('/addressUnbinding', data); }
  async rescanLostTransaction(data) { return this.client.post('/rescanLostTransaction', data); }
}
module.exports = UtilitiesService;
