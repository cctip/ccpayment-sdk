import { Client } from '../client';

export class BasicInfoService {
  constructor(private client: Client) {}
  async getCoinList(): Promise<any> { return this.client.post('/getCoinList', {}); }
  async getFiatList(): Promise<any> { return this.client.post('/getFiatList', {}); }
  async getCoin(coinId: number): Promise<any> { return this.client.post('/getCoin', { coinId }); }
  async getCoinUSDTPrice(coinIds: number[]): Promise<any> { return this.client.post('/getCoinUSDTPrice', { coinIds }); }
  async getChainList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getChainList', data); }
  async getAllChainList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/all/chain', data); }
  async getMainCoinList(appId: string): Promise<any> { return this.client.post('/getMainCoinList', { appId }); }
}
