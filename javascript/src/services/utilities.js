class UtilitiesService {
  constructor(client) {
    this.client = client;
  }

  async verifyAddress(data) {
    return this.client.post('/verifyAddress', data);
  }

  async abandonAddress(data) {
    return this.client.post('/abandonAddress', data);
  }

  async hostedInvoiceOrderInfo(orderId) {
    return this.client.post('/hostedInvoiceOrderInfo', { orderId });
  }

  async getPayInfo(orderId) {
    return this.client.post('/getPayInfo', { orderId });
  }

  async health() {
    return this.client.post('/health', {});
  }
}

module.exports = UtilitiesService;
