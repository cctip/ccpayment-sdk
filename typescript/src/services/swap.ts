import { Client } from '../client';

export class SwapService {
  constructor(private client: Client) {}

  async getSwapCoinList(): Promise<{ coins: any[] }> {
    return this.client.post('/getSwapCoinList', {});
  }

  async swapEstimate(data: {
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
  }> {
    return this.client.post('/swapEstimate', data);
  }

  async swap(data: {
    orderId: string;
    fromCoinId: number;
    toCoinId: number;
    fromAmount: string;
  }): Promise<{ recordId: string }> {
    return this.client.post('/swap', data);
  }

  async getSwapRecord(data: { recordId?: string; orderId?: string }): Promise<{ record: any }> {
    return this.client.post('/getSwapRecord', data);
  }

  async getSwapRecordList(data: {
    orderIds?: string[];
    fromCoinId?: number;
    toCoinId?: number;
    startAt?: number;
    endAt?: number;
    nextId?: string;
  }): Promise<{ records: any[]; nextId?: string }> {
    return this.client.post('/getSwapRecordList', data);
  }

  async selectHostedInvoiceCoin(data: {
    orderId: string;
    coinId: number;
    chain: string;
  }): Promise<{ address: string; memo?: string; amountToPay: string }> {
    return this.client.post('/selectHostedInvoiceCoin', data);
  }
}
