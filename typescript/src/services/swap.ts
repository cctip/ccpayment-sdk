import { Client } from '../client';

export class SwapService {
  constructor(private client: Client) {}
  async getSwapCoinList(): Promise<any> { return this.client.post('/getSwapCoinList', {}); }
  async swapEstimate(data: Record<string, unknown>): Promise<any> { return this.client.post('/estimate', data); }
  async getSwapRecord(data: Record<string, unknown>): Promise<any> { return this.client.post('/getSwapRecord', data); }
  async getSwapRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getSwapRecordList', data); }
  async swap(data: Record<string, unknown>): Promise<any> { return this.client.post('/swap', data); }
}
