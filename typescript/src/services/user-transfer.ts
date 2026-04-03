import { Client } from '../client';

export class UserTransferService {
  constructor(private client: Client) {}

  async userTransfer(data: {
    orderId: string;
    fromUserId: string;
    toUserId: string;
    coinId: number;
    amount: string;
    remark?: string;
  }): Promise<{ recordId: string }> {
    return this.client.post('/userTransfer', data);
  }

  async getUserTransferRecord(data: { recordId?: string; orderId?: string }): Promise<{ record: any }> {
    return this.client.post('/getUserTransferRecord', data);
  }

  async getUserTransferRecordList(data: {
    orderIds?: string[];
    fromUserId?: string;
    toUserId?: string;
    coinId?: number;
    startAt?: number;
    endAt?: number;
    nextId?: string;
  }): Promise<{ records: any[]; nextId?: string }> {
    return this.client.post('/getUserTransferRecordList', data);
  }

  async userBatchTransfer(data: {
    orderId: string;
    userId: string;
    toUsers: { userId: string; amount: string }[];
    coinId: number;
    remark?: string;
  }): Promise<{ recordId: string }> {
    return this.client.post('/userBatchTransfer', data);
  }

  async getUserBatchTransferRecord(data: { recordId?: string; orderId?: string }): Promise<{ record: any }> {
    return this.client.post('/getUserBatchTransferRecord', data);
  }

  async getUserBatchTransferRecordList(data: {
    orderIds?: string[];
    userId?: string;
    coinId?: number;
    startAt?: number;
    endAt?: number;
    nextId?: string;
  }): Promise<{ records: any[]; nextId?: string }> {
    return this.client.post('/getUserBatchTransferRecordList', data);
  }

  async getUserBatchTransferRecordDetail(data: {
    recordId: string;
    nextId?: string;
  }): Promise<{ records: any[]; nextId?: string }> {
    return this.client.post('/getUserBatchTransferRecordDetail', data);
  }
}
