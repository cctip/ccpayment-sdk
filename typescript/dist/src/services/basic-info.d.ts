import { Client } from '../client';
import { Coin, Chain, ChainSimple, Fiat, Fee } from '../types';
export declare class BasicInfoService {
    private client;
    constructor(client: Client);
    getCoinList(): Promise<{
        coins: Coin[];
    }>;
    getCoin(coinId: number): Promise<Coin>;
    getCoinUSDTPrice(coinIds: number[]): Promise<{
        prices: Record<string, string>;
    }>;
    getFiatList(): Promise<{
        fiats: Fiat[];
    }>;
    getChainList(chains?: string[]): Promise<{
        chains: Chain[];
    }>;
    allChain(chains?: string[]): Promise<{
        chains: ChainSimple[];
    }>;
    getCwalletUserId(cwalletUserId: string): Promise<{
        cwalletUserId: string;
        cwalletUserName: string;
    }>;
    getWithdrawFee(coinId: number, chain: string): Promise<{
        fee: Fee;
        networkFeeInquiryID: string;
    }>;
}
//# sourceMappingURL=basic-info.d.ts.map