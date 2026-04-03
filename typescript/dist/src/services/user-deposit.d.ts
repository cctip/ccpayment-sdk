import { Client } from '../client';
export declare class UserDepositService {
    private client;
    constructor(client: Client);
    getOrCreateUserDepositAddress(data: {
        userId: string;
        chain: string;
        notifyUrl?: string;
    }): Promise<{
        address: string;
        memo?: string;
    }>;
    getUserDepositRecord(recordId: string): Promise<{
        record: any;
    }>;
    getUserDepositRecordList(data: {
        userId: string;
        chain?: string;
        coinId?: number;
        toAddress?: string;
        startAt?: number;
        endAt?: number;
        nextId?: string;
    }): Promise<{
        records: any[];
        nextId?: string;
    }>;
}
//# sourceMappingURL=user-deposit.d.ts.map