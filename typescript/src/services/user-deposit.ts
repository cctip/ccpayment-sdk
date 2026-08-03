import { Client } from '../client';

export class UserDepositService {
  constructor(private client: Client) {}

  async getOrCreateUserDepositAddress(data: {
    userId: string;
    chain: string;
    notifyUrl?: string;
  }): Promise<{ address: string; memo?: string }> {
    return this.client.post('/getOrCreateUserDepositAddress', data);
  }

  async getUserDepositRecord(recordId: string): Promise<{ record: any }> {
    return this.client.post('/getUserDepositRecord', { recordId });
  }

  async getUserDepositRecordList(data: {
    userId: string;
    chain?: string;
    coinId?: number;
    toAddress?: string;
    startAt?: number;
    endAt?: number;
    nextId?: string;
  }): Promise<{ records: any[]; nextId?: string }> {
    return this.client.post('/getUserDepositRecordList', data);
  }
}
