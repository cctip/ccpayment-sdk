class OrdersService {
  constructor(client) { this.client = client; }
  async getAppOrderInfo(orderId) { return this.client.post('/getAppOrderInfo', { orderId }); }
  async createInvoiceUrl(data) { return this.client.post('/createInvoiceUrl', data); }
  async getInvoiceOrderInfo(orderId) { return this.client.post('/getInvoiceOrderInfo', { orderId }); }
}
module.exports = OrdersService;
