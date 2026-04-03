import { Client } from '../client';
import { Coin, Chain, ChainSimple, Fiat, Fee } from '../types';

export class BasicInfoService {
  constructor(private client: Client) {}

  async getCoinList(): Promise<{ coins: Coin[] }> {
    return this.client.post('/getCoinList', {});
  }

  async getCoin(coinId: number): Promise<Coin> {
    return this.client.post('/getCoin', { coinId });
  }

  async getCoinUSDTPrice(coinIds: number[]): Promise<{ prices: Record<string, string> }> {
    return this.client.post('/getCoinUSDTPrice', { coinIds });
  }

  async getFiatList(): Promise<{ fiats: Fiat[] }> {
    return this.client.post('/getFiatList', {});
  }

  async getChainList(chains?: string[]): Promise<{ chains: Chain[] }> {
    return this.client.post('/getChainList', chains ? { chains } : {});
  }

  async allChain(chains?: string[]): Promise<{ chains: ChainSimple[] }> {
    return this.client.post('/all/chain', chains ? { chains } : {});
  }

  async getCwalletUserId(cwalletUserId: string): Promise<{ cwalletUserId: string; cwalletUserName: string }> {
    return this.client.post('/getCwalletUserId', { cwalletUserId });
  }

  async getWithdrawFee(coinId: number, chain: string): Promise<{ fee: Fee; networkFeeInquiryID: string }> {
    return this.client.post('/getWithdrawFee', { coinId, chain });
  }
}
