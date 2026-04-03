import { Client } from '../client';
import { DepositRecord } from '../types';

export class MerchantDepositService {
  constructor(private client: Client) {}

  async createAppOrderDepositAddress(data: {
    orderId: string;
    coinId: number;
    chain: string;
    price: string;
    fiatId?: number;
    expiredAt?: number;
    buyerEmail?: string;
    generateCheckoutURL?: boolean;
    product?: string;
    returnUrl?: string;
    notifyUrl?: string;
    closeUrl?: string;
  }): Promise<{
    address: string;
    amount: string;
    memo?: string;
    checkoutUrl?: string;
    confirmsNeeded: number;
  }> {
    return this.client.post('/createAppOrderDepositAddress', data);
  }

  async getOrCreateAppDepositAddress(data: {
    referenceId: string;
    chain: string;
    notifyUrl?: string;
  }): Promise<{
    address: string;
    memo?: string;
  }> {
    return this.client.post('/getOrCreateAppDepositAddress', data);
  }

  async getAppDepositRecord(recordId: string): Promise<{ record: DepositRecord }> {
    return this.client.post('/getAppDepositRecord', { recordId });
  }

  async getAppDepositRecordList(data: {
    chain?: string;
    referenceId?: string;
    orderId?: string;
    toAddress?: string;
    coinId?: number;
    startAt?: number;
    endAt?: number;
    nextId?: string;
    recordIds?: string[];
    referenceIds?: string[];
    orderIds?: string[];
    limit?: number;
  }): Promise<{ records: DepositRecord[]; nextId?: string }> {
    return this.client.post('/getAppDepositRecordList', data);
  }
}
