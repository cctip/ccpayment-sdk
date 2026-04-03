"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CheckoutService = void 0;
class CheckoutService {
    constructor(client) {
        this.client = client;
    }
    async createCheckoutUrl(data) {
        return this.client.post('/createCheckoutUrl', data);
    }
    async hostedOrderInfo(orderId) {
        return this.client.post('/hostedOrderInfo', { orderId });
    }
    async hostedOrderInfoFirst(orderId) {
        return this.client.post('/hostedOrderInfoFirst', { orderId });
    }
    async createHostedOrderDeposit(data) {
        return this.client.post('/createHostedOrderDeposit', data);
    }
    async getHostedCoinUSDTPrice(coinIds) {
        return this.client.post('/getHostedCoinUSDTPrice', { coinIds });
    }
    async getMainCoinList() {
        return this.client.post('/getMainCoinList', {});
    }
    async createAppDemoOrderDeposit(data) {
        return this.client.post('/createAppDemoOrderDeposit', data);
    }
    async confirmTrade(data) {
        await this.client.post('/confirmTrade', data);
    }
}
exports.CheckoutService = CheckoutService;
//# sourceMappingURL=checkout.js.map