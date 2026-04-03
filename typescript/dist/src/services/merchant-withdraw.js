"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MerchantWithdrawService = void 0;
class MerchantWithdrawService {
    constructor(client) {
        this.client = client;
    }
    async applyAppWithdrawToNetwork(data) {
        return this.client.post('/applyAppWithdrawToNetwork', data);
    }
    async applyAppWithdrawToCwallet(data) {
        return this.client.post('/applyAppWithdrawToCwallet', data);
    }
    async getAppWithdrawRecord(data) {
        return this.client.post('/getAppWithdrawRecord', data);
    }
    async getAppWithdrawRecordList(data) {
        return this.client.post('/getAppWithdrawRecordList', data);
    }
    async getAutoWithdrawRecordList(data) {
        return this.client.post('/getAutoWithdrawRecordList', data);
    }
    async getRiskyRefundRecordList(data) {
        return this.client.post('/getRiskyRefundRecordList', data);
    }
}
exports.MerchantWithdrawService = MerchantWithdrawService;
//# sourceMappingURL=merchant-withdraw.js.map