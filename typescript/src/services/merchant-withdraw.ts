import { Client } from '../client';
import { WithdrawRecord, AutoWithdrawRecord, RiskyRefundRecord, WithdrawFee } from '../types';

export class MerchantWithdrawService {
  constructor(private client: Client) {}

  async applyAppWithdrawToNetwork(data: {
    orderId: string;
    coinId: number;
    chain: string;
    address: string;
    amount: string;
    memo?: string;
    merchantPayNetworkFee?: boolean;
    networkFeeInquiryID?: string;
    notifyUrl?: string;
  }): Promise<{ recordId: string }> {
    return this.client.post('/applyAppWithdrawToNetwork', data);
  }

  async applyAppWithdrawToCwallet(data: {
    orderId: string;
    coinId: number;
    cwalletUser: string;
    amount: string;
  }): Promise<{ recordId: string }> {
    return this.client.post('/applyAppWithdrawToCwallet', data);
  }

  async getAppWithdrawRecord(data: { orderId?: string; recordId?: string }): Promise<{ record: WithdrawRecord }> {
    return this.client.post('/getAppWithdrawRecord', data);
  }

  async getAppWithdrawRecordList(data: {
    chain?: string;
    coinId?: number;
    orderIds?: string[];
    startAt?: number;
    endAt?: number;
    toAddress?: string;
    nextId?: string;
  }): Promise<{ records: WithdrawRecord[]; nextId?: string }> {
    return this.client.post('/getAppWithdrawRecordList', data);
  }

  async getAutoWithdrawRecordList(data: {
    chain?: string;
    coinId?: number;
    recordIds?: string[];
    startAt?: number;
    endAt?: number;
    toAddress?: string;
    nextId?: string;
  }): Promise<{ records: AutoWithdrawRecord[]; nextId?: string }> {
    return this.client.post('/getAutoWithdrawRecordList', data);
  }

  async getRiskyRefundRecordList(data: {
    chain?: string;
    coinId?: number;
    recordIds?: string[];
    startAt?: number;
    endAt?: number;
    toAddress?: string;
    nextId?: string;
  }): Promise<{ records: RiskyRefundRecord[]; nextId?: string }> {
    return this.client.post('/getRiskyRefundRecordList', data);
  }
}
