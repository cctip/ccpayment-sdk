import { Client } from '../client';
import { WithdrawRecord, AutoWithdrawRecord, RiskyRefundRecord } from '../types';
export declare class MerchantWithdrawService {
    private client;
    constructor(client: Client);
    applyAppWithdrawToNetwork(data: {
        orderId: string;
        coinId: number;
        chain: string;
        address: string;
        amount: string;
        memo?: string;
        merchantPayNetworkFee?: boolean;
        networkFeeInquiryID?: string;
        notifyUrl?: string;
    }): Promise<{
        recordId: string;
    }>;
    applyAppWithdrawToCwallet(data: {
        orderId: string;
        coinId: number;
        cwalletUser: string;
        amount: string;
    }): Promise<{
        recordId: string;
    }>;
    getAppWithdrawRecord(data: {
        orderId?: string;
        recordId?: string;
    }): Promise<{
        record: WithdrawRecord;
    }>;
    getAppWithdrawRecordList(data: {
        chain?: string;
        coinId?: number;
        orderIds?: string[];
        startAt?: number;
        endAt?: number;
        toAddress?: string;
        nextId?: string;
    }): Promise<{
        records: WithdrawRecord[];
        nextId?: string;
    }>;
    getAutoWithdrawRecordList(data: {
        chain?: string;
        coinId?: number;
        recordIds?: string[];
        startAt?: number;
        endAt?: number;
        toAddress?: string;
        nextId?: string;
    }): Promise<{
        records: AutoWithdrawRecord[];
        nextId?: string;
    }>;
    getRiskyRefundRecordList(data: {
        chain?: string;
        coinId?: number;
        recordIds?: string[];
        startAt?: number;
        endAt?: number;
        toAddress?: string;
        nextId?: string;
    }): Promise<{
        records: RiskyRefundRecord[];
        nextId?: string;
    }>;
}
//# sourceMappingURL=merchant-withdraw.d.ts.map