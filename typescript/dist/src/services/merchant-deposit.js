"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MerchantDepositService = void 0;
class MerchantDepositService {
    constructor(client) {
        this.client = client;
    }
    async createAppOrderDepositAddress(data) {
        return this.client.post('/createAppOrderDepositAddress', data);
    }
    async getOrCreateAppDepositAddress(data) {
        return this.client.post('/getOrCreateAppDepositAddress', data);
    }
    async getAppDepositRecord(recordId) {
        return this.client.post('/getAppDepositRecord', { recordId });
    }
    async getAppDepositRecordList(data) {
        return this.client.post('/getAppDepositRecordList', data);
    }
}
exports.MerchantDepositService = MerchantDepositService;
//# sourceMappingURL=merchant-deposit.js.map