class MerchantBatchWithdrawService {
  constructor(client) {
    this.client = client;
  }

  /**
   * Request fields: taskName, orders[].orderId, orderIds, confirmExecution, stats.execOrderId.
   */

  async checkWithdrawAddress(data) {
    return this.client.post('/checkWithdrawAddress', data);
  }

  async applyBatchWithdraw(data) {
    return this.client.post('/applyBatchWithdraw', data);
  }

  async appendBatchWithdraw(data) {
    return this.client.post('/appendBatchWithdraw', data);
  }

  async confirmBatchWithdraw(data) {
    return this.client.post('/confirmBatchWithdraw', data);
  }

  async abortBatchWithdraw(data) {
    return this.client.post('/abortBatchWithdraw', data);
  }

  async getBatchWithdrawOrder(data) {
    return this.client.post('/getBatchWithdrawOrder', data);
  }

  async getBatchWithdrawOrderRecordList(data) {
    return this.client.post('/getBatchWithdrawOrderRecordList', data);
  }
}

module.exports = MerchantBatchWithdrawService;
