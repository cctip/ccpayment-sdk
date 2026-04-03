class UserWithdrawService {
  constructor(client) {
    this.client = client;
  }

  async applyUserWithdrawToNetwork(data) {
    return this.client.post('/applyUserWithdrawToNetwork', data);
  }

  async applyUserWithdrawToCwallet(data) {
    return this.client.post('/applyUserWithdrawToCwallet', data);
  }

  async getUserWithdrawRecord(data) {
    return this.client.post('/getUserWithdrawRecord', data);
  }

  async getUserWithdrawRecordList(data) {
    return this.client.post('/getUserWithdrawRecordList', data);
  }
}

module.exports = UserWithdrawService;
