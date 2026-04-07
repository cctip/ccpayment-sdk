import { Client } from '../client';

export class OrdersService {
  constructor(private client: Client) {}
  async getAppOrderInfo(orderId: string): Promise<any> { return this.client.post('/getAppOrderInfo', { orderId }); }
  async createInvoiceUrl(data: Record<string, unknown>): Promise<any> { return this.client.post('/createInvoiceUrl', data); }
  async getInvoiceOrderInfo(orderId: string): Promise<any> { return this.client.post('/getInvoiceOrderInfo', { orderId }); }
}
