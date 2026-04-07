import { Client } from '../client';

export class UtilitiesService {
  constructor(private client: Client) {}
  async webhookResend(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/webhook/resend', data); }
  async getTradeBlockHeight(data: Record<string, unknown>): Promise<any> { return this.client.post('/getTradeBlockHeight', data); }
  async checkWithdrawalAddressValidity(data: Record<string, unknown>): Promise<any> { return this.client.post('/checkWithdrawalAddressValidity', data); }
  async deprecatedAddress(data: Record<string, unknown>): Promise<any> { return this.client.post('/addressUnbinding', data); }
  async rescanLostTransaction(data: Record<string, unknown>): Promise<any> { return this.client.post('/rescanLostTransaction', data); }
}
