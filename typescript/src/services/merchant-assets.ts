import { Client } from '../client';
import { Asset } from '../types';

export class MerchantAssetsService {
  constructor(private client: Client) {}

  async getAppCoinAssetList(): Promise<{ assets: Asset[] }> {
    return this.client.post('/getAppCoinAssetList', {});
  }

  async getAppCoinAsset(coinId: number): Promise<{ asset: Asset }> {
    return this.client.post('/getAppCoinAsset', { coinId });
  }
}
