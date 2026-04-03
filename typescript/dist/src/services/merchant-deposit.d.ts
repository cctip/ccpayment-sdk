import { Client } from '../client';
import { DepositRecord } from '../types';
export declare class MerchantDepositService {
    private client;
    constructor(client: Client);
    createAppOrderDepositAddress(data: {
        orderId: string;
        coinId: number;
        chain: string;
        price: string;
        fiatId?: number;
        expiredAt?: number;
        buyerEmail?: string;
        generateCheckoutURL?: boolean;
        product?: string;
        returnUrl?: string;
        notifyUrl?: string;
        closeUrl?: string;
    }): Promise<{
        address: string;
        amount: string;
        memo?: string;
        checkoutUrl?: string;
        confirmsNeeded: number;
    }>;
    getOrCreateAppDepositAddress(data: {
        referenceId: string;
        chain: string;
        notifyUrl?: string;
    }): Promise<{
        address: string;
        memo?: string;
    }>;
    getAppDepositRecord(recordId: string): Promise<{
        record: DepositRecord;
    }>;
    getAppDepositRecordList(data: {
        chain?: string;
        referenceId?: string;
        orderId?: string;
        toAddress?: string;
        coinId?: number;
        startAt?: number;
        endAt?: number;
        nextId?: string;
        recordIds?: string[];
        referenceIds?: string[];
        orderIds?: string[];
        limit?: number;
    }): Promise<{
        records: DepositRecord[];
        nextId?: string;
    }>;
}
//# sourceMappingURL=merchant-deposit.d.ts.map