import { Client } from '../client';
import { Asset } from '../types';
export declare class MerchantAssetsService {
    private client;
    constructor(client: Client);
    getAppCoinAssetList(): Promise<{
        assets: Asset[];
    }>;
    getAppCoinAsset(coinId: number): Promise<{
        asset: Asset;
    }>;
}
//# sourceMappingURL=merchant-assets.d.ts.map