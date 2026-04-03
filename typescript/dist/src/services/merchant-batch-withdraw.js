"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MerchantBatchWithdrawService = void 0;
class MerchantBatchWithdrawService {
    constructor(client) {
        this.client = client;
    }
    async checkWithdrawAddress(data) {
        return this.client.post('/checkWithdrawAddress', data);
    }
    async applyBatchWithdraw(data) {
        return this.client.post('/applyBatchWithdraw', data);
    }
    async appendBatchWithdraw(data) {
        await this.client.post('/appendBatchWithdraw', data);
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
exports.MerchantBatchWithdrawService = MerchantBatchWithdrawService;
//# sourceMappingURL=merchant-batch-withdraw.js.map