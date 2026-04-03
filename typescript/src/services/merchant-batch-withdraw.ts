import { Client } from '../client';
import { AddressInfo, BatchWithdrawOrder } from '../types';

export class MerchantBatchWithdrawService {
  constructor(private client: Client) {}

  async checkWithdrawAddress(data: {
    chain: string;
    addressInfoList: AddressInfo[];
  }): Promise<{ addressInfoResults: { seq: number; codes: number[] }[] }> {
    return this.client.post('/checkWithdrawAddress', data);
  }

  async applyBatchWithdraw(data: {
    batchOrderId: string;
    coinId: number;
    chain: string;
    mode: string;
    orders?: BatchWithdrawOrder[];
    productName?: string;
    notifyUrl?: string;
  }): Promise<{ batchOrderId: string; billId: string }> {
    return this.client.post('/applyBatchWithdraw', data);
  }

  async appendBatchWithdraw(data: {
    batchOrderId: string;
    orders: BatchWithdrawOrder[];
  }): Promise<void> {
    await this.client.post('/appendBatchWithdraw', data);
  }

  async confirmBatchWithdraw(data: {
    batchOrderId: string;
    delaySeconds?: number;
  }): Promise<{
    batchOrderId: string;
    productName: string;
    billId: string;
    coin: { coin_id: number; coin_symbol: string; coin_price: string };
    amount: string;
    networkFee: string;
    networkFeeCoin: { coin_id: number; coin_symbol: string; coin_price: string };
    status: string;
    reason?: string;
    mode: string;
    stats: {
      total: number;
      succeeded: number;
      failed: number;
      canceled: number;
      processing: number;
      execSeq: number;
    };
    createdAt: number;
    updatedAt: number;
  }> {
    return this.client.post('/confirmBatchWithdraw', data);
  }

  async abortBatchWithdraw(data: {
    batchOrderId: string;
    seqs?: number[];
  }): Promise<{ batchOrderId: string; canceledSeqs: number[]; ignoredSeqs: number[] }> {
    return this.client.post('/abortBatchWithdraw', data);
  }

  async getBatchWithdrawOrder(data: {
    batchOrderId: string;
    verbose?: number;
  }): Promise<any> {
    return this.client.post('/getBatchWithdrawOrder', data);
  }

  async getBatchWithdrawOrderRecordList(data: {
    batchOrderId: string;
    nextId?: string;
    limit?: number;
  }): Promise<{ records: BatchWithdrawOrder[]; nextId?: string }> {
    return this.client.post('/getBatchWithdrawOrderRecordList', data);
  }
}
