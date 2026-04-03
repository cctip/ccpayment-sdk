"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SwapService = void 0;
class SwapService {
    constructor(client) {
        this.client = client;
    }
    async getSwapCoinList() {
        return this.client.post('/getSwapCoinList', {});
    }
    async swapEstimate(data) {
        return this.client.post('/swapEstimate', data);
    }
    async swap(data) {
        return this.client.post('/swap', data);
    }
    async getSwapRecord(data) {
        return this.client.post('/getSwapRecord', data);
    }
    async getSwapRecordList(data) {
        return this.client.post('/getSwapRecordList', data);
    }
    async selectHostedInvoiceCoin(data) {
        return this.client.post('/selectHostedInvoiceCoin', data);
    }
}
exports.SwapService = SwapService;
//# sourceMappingURL=swap.js.map