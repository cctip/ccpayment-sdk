import { Client } from '../client';

export class OrdersService {
  constructor(private client: Client) {}

  async getAppOrderInfo(orderId: string): Promise<any> {
    return this.client.post('/getAppOrderInfo', { orderId });
  }

  async createInvoiceUrl(data: {
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
  }): Promise<{ invoiceUrl: string }> {
    return this.client.post('/createInvoiceUrl', data);
  }

  async getInvoiceOrderInfo(orderId: string): Promise<any> {
    return this.client.post('/getInvoiceOrderInfo', { orderId });
  }

  async getWebhookInfo(): Promise<{ webhookUrl: string; webhookSecret: string }> {
    return this.client.post('/getWebhookInfo', {});
  }
}
