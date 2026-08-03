class UserDepositService {
  constructor(client) {
    this.client = client;
  }

  async getOrCreateUserDepositAddress(data) {
    return this.client.post('/getOrCreateUserDepositAddress', data);
  }

  async getUserDepositRecord(recordId) {
    return this.client.post('/getUserDepositRecord', { recordId });
  }

  async getUserDepositRecordList(data) {
    return this.client.post('/getUserDepositRecordList', data);
  }
}

module.exports = UserDepositService;
