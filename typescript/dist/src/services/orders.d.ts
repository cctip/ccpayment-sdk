import { Client } from '../client';
export declare class OrdersService {
    private client;
    constructor(client: Client);
    getAppOrderInfo(orderId: string): Promise<any>;
    createInvoiceUrl(data: {
        orderId: string;
        product: string;
        price: string;
        priceFiatId?: number;
        priceCoinId?: number;
        buyerEmail?: string;
        returnUrl?: string;
        expiredAt?: number;
        closeUrl?: string;
        notifyUrl?: string;
    }): Promise<{
        invoiceUrl: string;
    }>;
    getInvoiceOrderInfo(orderId: string): Promise<any>;
    getWebhookInfo(): Promise<{
        webhookUrl: string;
        webhookSecret: string;
    }>;
}
//# sourceMappingURL=orders.d.ts.map