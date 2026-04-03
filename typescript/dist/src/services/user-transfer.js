"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserTransferService = void 0;
class UserTransferService {
    constructor(client) {
        this.client = client;
    }
    async userTransfer(data) {
        return this.client.post('/userTransfer', data);
    }
    async getUserTransferRecord(data) {
        return this.client.post('/getUserTransferRecord', data);
    }
    async getUserTransferRecordList(data) {
        return this.client.post('/getUserTransferRecordList', data);
    }
    async userBatchTransfer(data) {
        return this.client.post('/userBatchTransfer', data);
    }
    async getUserBatchTransferRecord(data) {
        return this.client.post('/getUserBatchTransferRecord', data);
    }
    async getUserBatchTransferRecordList(data) {
        return this.client.post('/getUserBatchTransferRecordList', data);
    }
    async getUserBatchTransferRecordDetail(data) {
        return this.client.post('/getUserBatchTransferRecordDetail', data);
    }
}
exports.UserTransferService = UserTransferService;
//# sourceMappingURL=user-transfer.js.map