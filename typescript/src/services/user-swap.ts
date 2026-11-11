import { Client } from '../client';

export class UserSwapService {
  constructor(private client: Client) {}
  /** @deprecated Use getUserSwapRecordV2. */
  async getUserSwapRecord(data: Record<string, unknown>): Promise<any> { return this.client.post('/getUserSwapRecord', data); }
  /** @deprecated Use getUserSwapRecordListV2. */
  async getUserSwapRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getUserSwapRecordList', data); }
  /** @deprecated Use userSwapV2 with SwapReqV2 fields. */
  async userSwap(data: Record<string, unknown>): Promise<any> { return this.client.post('/userSwap', data); }
  async userSwapV2(data: Record<string, unknown>): Promise<any> { return this.client.post('/swapv2/userSwap', data); }
  async getUserSwapRecordV2(data: Record<string, unknown>): Promise<any> { return this.client.post('/swapv2/getUserSwapRecord', data); }
  async getUserSwapRecordListV2(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/swapv2/getUserSwapRecordList', data); }
}
