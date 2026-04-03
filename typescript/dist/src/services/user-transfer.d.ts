import { Client } from '../client';
export declare class UserTransferService {
    private client;
    constructor(client: Client);
    userTransfer(data: {
        orderId: string;
        fromUserId: string;
        toUserId: string;
        coinId: number;
        amount: string;
        remark?: string;
    }): Promise<{
        recordId: string;
    }>;
    getUserTransferRecord(data: {
        recordId?: string;
        orderId?: string;
    }): Promise<{
        record: any;
    }>;
    getUserTransferRecordList(data: {
        orderIds?: string[];
        fromUserId?: string;
        toUserId?: string;
        coinId?: number;
        startAt?: number;
        endAt?: number;
        nextId?: string;
    }): Promise<{
        records: any[];
        nextId?: string;
    }>;
    userBatchTransfer(data: {
        orderId: string;
        userId: string;
        toUsers: {
            userId: string;
            amount: string;
        }[];
        coinId: number;
        remark?: string;
    }): Promise<{
        recordId: string;
    }>;
    getUserBatchTransferRecord(data: {
        recordId?: string;
        orderId?: string;
    }): Promise<{
        record: any;
    }>;
    getUserBatchTransferRecordList(data: {
        orderIds?: string[];
        userId?: string;
        coinId?: number;
        startAt?: number;
        endAt?: number;
        nextId?: string;
    }): Promise<{
        records: any[];
        nextId?: string;
    }>;
    getUserBatchTransferRecordDetail(data: {
        recordId: string;
        nextId?: string;
    }): Promise<{
        records: any[];
        nextId?: string;
    }>;
}
//# sourceMappingURL=user-transfer.d.ts.map