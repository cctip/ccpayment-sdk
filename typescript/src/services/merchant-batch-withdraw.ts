import { Client } from '../client';
import { AddressInfo, BatchWithdrawOrder, BatchWithdrawOrderDetail } from '../types';

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
    taskName?: string;
    notifyUrl?: string;
  }): Promise<{ batchOrderId: string; billId: string }> {
    return this.client.post('/applyBatchWithdraw', data);
  }

  async appendBatchWithdraw(data: {
    batchOrderId: string;
    orders: BatchWithdrawOrder[];
  }): Promise<{ appended: number; total: number; batch_order_id: string }> {
    return this.client.post('/appendBatchWithdraw', data);
  }

  async confirmBatchWithdraw(data: {
    batchOrderId: string;
    delaySeconds?: number;
    confirmExecution?: boolean;
  }): Promise<BatchWithdrawOrderDetail> {
    return this.client.post('/confirmBatchWithdraw', data);
  }

  async abortBatchWithdraw(data: {
    batchOrderId: string;
    orderIds?: string[];
  }): Promise<{
    batchOrderId: string;
    status?: string;
    canceledOrderIds?: string[];
    ignoredOrderIds?: string[];
  }> {
    return this.client.post('/abortBatchWithdraw', data);
  }

  async getBatchWithdrawOrder(data: {
    batchOrderId: string;
    verbose?: number;
  }): Promise<BatchWithdrawOrderDetail> {
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
