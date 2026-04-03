import { Client } from '../client';

export class UtilitiesService {
  constructor(private client: Client) {}

  async verifyAddress(data: { chain: string; address: string }): Promise<{
    valid: boolean;
    message: string;
  }> {
    return this.client.post('/verifyAddress', data);
  }

  async abandonAddress(data: { chain: string; address: string }): Promise<void> {
    await this.client.post('/abandonAddress', data);
  }

  async hostedInvoiceOrderInfo(orderId: string): Promise<any> {
    return this.client.post('/hostedInvoiceOrderInfo', { orderId });
  }

  async getPayInfo(orderId: string): Promise<{
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
  }> {
    return this.client.post('/getPayInfo', { orderId });
  }

  async health(): Promise<{ status: string; timestamp: number }> {
    return this.client.post('/health', {});
  }
}
