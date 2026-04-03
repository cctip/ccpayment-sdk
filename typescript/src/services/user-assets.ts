import { Client } from '../client';
import { Asset } from '../types';

export class UserAssetsService {
  constructor(private client: Client) {}

  async getUserCoinAssetList(userId: string): Promise<{ userId: string; assets: Asset[] }> {
    return this.client.post('/getUserCoinAssetList', { userId });
  }

  async getUserCoinAsset(userId: string, coinId: number): Promise<{ userId: string; asset: Asset }> {
    return this.client.post('/getUserCoinAsset', { userId, coinId });
  }
}
