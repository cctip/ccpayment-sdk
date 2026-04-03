import { Client } from '../client';
export declare class UtilitiesService {
    private client;
    constructor(client: Client);
    verifyAddress(data: {
        chain: string;
        address: string;
    }): Promise<{
        valid: boolean;
        message: string;
    }>;
    abandonAddress(data: {
        chain: string;
        address: string;
    }): Promise<void>;
    hostedInvoiceOrderInfo(orderId: string): Promise<any>;
    getPayInfo(orderId: string): Promise<{
        orderId: string;
        product: string;
        price: string;
        priceSymbol: string;
        address: string;
        memo: string;
        amount: string;
        coinSymbol: string;
        chain: string;
        qrCode: string;
        expiredAt: number;
    }>;
    health(): Promise<{
        status: string;
        timestamp: number;
    }>;
}
//# sourceMappingURL=utilities.d.ts.map