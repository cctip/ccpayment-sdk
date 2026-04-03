import { Client } from '../client';

export class UserWithdrawService {
  constructor(private client: Client) {}

  async applyUserWithdrawToNetwork(data: {
    userId: string;
    orderId: string;
    coinId: number;
    chain: string;
    address: string;
    amount: string;
    memo?: string;
    networkFeeInquiryID?: string;
    notifyUrl?: string;
  }): Promise<{ recordId: string }> {
    return this.client.post('/applyUserWithdrawToNetwork', data);
  }

  async applyUserWithdrawToCwallet(data: {
    userId: string;
    orderId: string;
    coinId: number;
    cwalletUser: string;
    amount: string;
  }): Promise<{ recordId: string }> {
    return this.client.post('/applyUserWithdrawToCwallet', data);
  }

  async getUserWithdrawRecord(data: { recordId?: string; orderId?: string }): Promise<{ record: any }> {
    return this.client.post('/getUserWithdrawRecord', data);
  }

  async getUserWithdrawRecordList(data: {
    userId: string;
    orderIds?: string[];
    chain?: string;
    coinId?: number;
    startAt?: number;
    endAt?: number;
    toAddress?: string;
    nextId?: string;
  }): Promise<{ records: any[]; nextId?: string }> {
    return this.client.post('/getUserWithdrawRecordList', data);
  }
}
