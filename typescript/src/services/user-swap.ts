import { Client } from '../client';

export class UserSwapService {
  constructor(private client: Client) {}
  async getUserSwapRecord(data: Record<string, unknown>): Promise<any> { return this.client.post('/getUserSwapRecord', data); }
  async getUserSwapRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getUserSwapRecordList', data); }
  async userSwap(data: Record<string, unknown>): Promise<any> { return this.client.post('/userSwap', data); }
}
