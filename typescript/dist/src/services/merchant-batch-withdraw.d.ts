import { Client } from '../client';
import { AddressInfo, BatchWithdrawOrder } from '../types';
export declare class MerchantBatchWithdrawService {
    private client;
    constructor(client: Client);
    checkWithdrawAddress(data: {
        chain: string;
        addressInfoList: AddressInfo[];
    }): Promise<{
        addressInfoResults: {
            seq: number;
            codes: number[];
        }[];
    }>;
    applyBatchWithdraw(data: {
        batchOrderId: string;
        coinId: number;
        chain: string;
        mode: string;
        orders?: BatchWithdrawOrder[];
        productName?: string;
        notifyUrl?: string;
    }): Promise<{
        batchOrderId: string;
        billId: string;
    }>;
    appendBatchWithdraw(data: {
        batchOrderId: string;
        orders: BatchWithdrawOrder[];
    }): Promise<void>;
    confirmBatchWithdraw(data: {
        batchOrderId: string;
        delaySeconds?: number;
    }): Promise<{
        batchOrderId: string;
        productName: string;
        billId: string;
        coin: {
            coin_id: number;
            coin_symbol: string;
            coin_price: string;
        };
        amount: string;
        networkFee: string;
        networkFeeCoin: {
            coin_id: number;
            coin_symbol: string;
            coin_price: string;
        };
        status: string;
        reason?: string;
        mode: string;
        stats: {
            total: number;
            succeeded: number;
            failed: number;
            canceled: number;
            processing: number;
            execSeq: number;
        };
        createdAt: number;
        updatedAt: number;
    }>;
    abortBatchWithdraw(data: {
        batchOrderId: string;
        seqs?: number[];
    }): Promise<{
        batchOrderId: string;
        canceledSeqs: number[];
        ignoredSeqs: number[];
    }>;
    getBatchWithdrawOrder(data: {
        batchOrderId: string;
        verbose?: number;
    }): Promise<any>;
    getBatchWithdrawOrderRecordList(data: {
        batchOrderId: string;
        nextId?: string;
        limit?: number;
    }): Promise<{
        records: BatchWithdrawOrder[];
        nextId?: string;
    }>;
}
//# sourceMappingURL=merchant-batch-withdraw.d.ts.map