import { Client } from '../client';
export declare class CheckoutService {
    private client;
    constructor(client: Client);
    createCheckoutUrl(data: {
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
        checkoutUrl: string;
    }>;
    hostedOrderInfo(orderId: string): Promise<any>;
    hostedOrderInfoFirst(orderId: string): Promise<any>;
    createHostedOrderDeposit(data: {
        orderId: string;
        coinId: number;
        chain: string;
    }): Promise<{
        address: string;
        memo?: string;
        amountToPay: string;
        confirmsNeeded: number;
    }>;
    getHostedCoinUSDTPrice(coinIds: number[]): Promise<{
        prices: Record<string, string>;
    }>;
    getMainCoinList(): Promise<{
        coins: any[];
    }>;
    createAppDemoOrderDeposit(data: {
        orderId: string;
        coinId: number;
        chain: string;
        price: string;
    }): Promise<{
        address: string;
        amount: string;
        memo?: string;
    }>;
    confirmTrade(data: {
        orderId: string;
        txId: string;
    }): Promise<void>;
}
//# sourceMappingURL=checkout.d.ts.map