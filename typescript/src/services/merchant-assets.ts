import { Client } from '../client';

export class MerchantAssetsService {
  constructor(private client: Client) {}
  async getAppCoinAssetList(): Promise<any> { return this.client.post('/getAppCoinAssetList', {}); }
  async getAppCoinAsset(coinId: number): Promise<any> { return this.client.post('/getAppCoinAsset', { coinId }); }
  async getAppCollectFeeRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getAggregationFeeRecord', data); }
}
