"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.OrdersService = void 0;
class OrdersService {
    constructor(client) {
        this.client = client;
    }
    async getAppOrderInfo(orderId) {
        return this.client.post('/getAppOrderInfo', { orderId });
    }
    async createInvoiceUrl(data) {
        return this.client.post('/createInvoiceUrl', data);
    }
    async getInvoiceOrderInfo(orderId) {
        return this.client.post('/getInvoiceOrderInfo', { orderId });
    }
    async getWebhookInfo() {
        return this.client.post('/getWebhookInfo', {});
    }
}
exports.OrdersService = OrdersService;
//# sourceMappingURL=orders.js.map