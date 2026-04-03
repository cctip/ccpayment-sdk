import { Client } from '../client';
import { Asset } from '../types';
export declare class UserAssetsService {
    private client;
    constructor(client: Client);
    getUserCoinAssetList(userId: string): Promise<{
        userId: string;
        assets: Asset[];
    }>;
    getUserCoinAsset(userId: string, coinId: number): Promise<{
        userId: string;
        asset: Asset;
    }>;
}
//# sourceMappingURL=user-assets.d.ts.map