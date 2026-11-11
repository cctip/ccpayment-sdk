import { Client } from '../client';

export class SwapService {
  constructor(private client: Client) {}
  /** @deprecated Use getSwapPairList. */
  async getSwapCoinList(): Promise<any> { return this.client.post('/getSwapCoinList', {}); }
  /** @deprecated Use swapEstimateV2 with SwapReqV2 fields. */
  async swapEstimate(data: Record<string, unknown>): Promise<any> { return this.client.post('/estimate', data); }
  /** @deprecated Use getSwapRecordV2. */
  async getSwapRecord(data: Record<string, unknown>): Promise<any> { return this.client.post('/getSwapRecord', data); }
  /** @deprecated Use getSwapRecordListV2. */
  async getSwapRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getSwapRecordList', data); }
  /** @deprecated Use swapV2 with SwapReqV2 fields. */
  async swap(data: Record<string, unknown>): Promise<any> { return this.client.post('/swap', data); }
  async getSwapPairList(): Promise<any> { return this.client.post('/swapv2/getSwapPairList', {}); }
  async swapEstimateV2(data: Record<string, unknown>): Promise<any> { return this.client.post('/swapv2/estimate', data); }
  async swapV2(data: Record<string, unknown>): Promise<any> { return this.client.post('/swapv2/swap', data); }
  async getSwapRecordV2(data: Record<string, unknown>): Promise<any> { return this.client.post('/swapv2/getSwapRecord', data); }
  async getSwapRecordListV2(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/swapv2/getSwapRecordList', data); }
}
