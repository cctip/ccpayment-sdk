import { Client } from '../client';
export declare class UserWithdrawService {
    private client;
    constructor(client: Client);
    applyUserWithdrawToNetwork(data: {
        userId: string;
        orderId: string;
        coinId: number;
        chain: string;
        address: string;
        amount: string;
        memo?: string;
        networkFeeInquiryID?: string;
        notifyUrl?: string;
    }): Promise<{
        recordId: string;
    }>;
    applyUserWithdrawToCwallet(data: {
        userId: string;
        orderId: string;
        coinId: number;
        cwalletUser: string;
        amount: string;
    }): Promise<{
        recordId: string;
    }>;
    getUserWithdrawRecord(data: {
        recordId?: string;
        orderId?: string;
    }): Promise<{
        record: any;
    }>;
    getUserWithdrawRecordList(data: {
        userId: string;
        orderIds?: string[];
        chain?: string;
        coinId?: number;
        startAt?: number;
        endAt?: number;
        toAddress?: string;
        nextId?: string;
    }): Promise<{
        records: any[];
        nextId?: string;
    }>;
}
//# sourceMappingURL=user-withdraw.d.ts.map