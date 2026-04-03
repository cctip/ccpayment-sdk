import { Client } from '../client';
export declare class SwapService {
    private client;
    constructor(client: Client);
    getSwapCoinList(): Promise<{
        coins: any[];
    }>;
    swapEstimate(data: {
        fromCoinId: number;
        toCoinId: number;
        fromAmount: string;
    }): Promise<{
        fromCoinId: number;
        fromCoinSymbol: string;
        fromAmount: string;
        toCoinId: number;
        toCoinSymbol: string;
        toAmount: string;
        rate: string;
    }>;
    swap(data: {
        orderId: string;
        fromCoinId: number;
        toCoinId: number;
        fromAmount: string;
    }): Promise<{
        recordId: string;
    }>;
    getSwapRecord(data: {
        recordId?: string;
        orderId?: string;
    }): Promise<{
        record: any;
    }>;
    getSwapRecordList(data: {
        orderIds?: string[];
        fromCoinId?: number;
        toCoinId?: number;
        startAt?: number;
        endAt?: number;
        nextId?: string;
    }): Promise<{
        records: any[];
        nextId?: string;
    }>;
    selectHostedInvoiceCoin(data: {
        orderId: string;
        coinId: number;
        chain: string;
    }): Promise<{
        address: string;
        memo?: string;
        amountToPay: string;
    }>;
}
//# sourceMappingURL=swap.d.ts.map