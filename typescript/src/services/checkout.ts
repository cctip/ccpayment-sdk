import { Client } from '../client';

export class CheckoutService {
  constructor(private client: Client) {}

  async createCheckoutUrl(data: {
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
  }): Promise<{ checkoutUrl: string }> {
    return this.client.post('/createCheckoutUrl', data);
  }

  async hostedOrderInfo(orderId: string): Promise<any> {
    return this.client.post('/hostedOrderInfo', { orderId });
  }

  async hostedOrderInfoFirst(orderId: string): Promise<any> {
    return this.client.post('/hostedOrderInfoFirst', { orderId });
  }

  async createHostedOrderDeposit(data: {
    orderId: string;
    coinId: number;
    chain: string;
  }): Promise<{ address: string; memo?: string; amountToPay: string; confirmsNeeded: number }> {
    return this.client.post('/createHostedOrderDeposit', data);
  }

  async getHostedCoinUSDTPrice(coinIds: number[]): Promise<{ prices: Record<string, string> }> {
    return this.client.post('/getHostedCoinUSDTPrice', { coinIds });
  }

  async getMainCoinList(): Promise<{ coins: any[] }> {
    return this.client.post('/getMainCoinList', {});
  }

  async createAppDemoOrderDeposit(data: {
    orderId: string;
    coinId: number;
    chain: string;
    price: string;
  }): Promise<{ address: string; amount: string; memo?: string }> {
    return this.client.post('/createAppDemoOrderDeposit', data);
  }

  async confirmTrade(data: { orderId: string; txId: string }): Promise<void> {
    await this.client.post('/confirmTrade', data);
  }
}
