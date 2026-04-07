import { Client } from '../client';

export class MerchantWithdrawService {
  constructor(private client: Client) {}
  async getCwalletUserId(cwalletUserId: string): Promise<any> { return this.client.post('/getCwalletUserId', { cwalletUserId }); }
  async getWithdrawFee(coinId: number, chain: string): Promise<any> { return this.client.post('/getWithdrawFee', { coinId, chain }); }
  async applyAppWithdrawToNetwork(data: Record<string, unknown>): Promise<any> { return this.client.post('/applyAppWithdrawToNetwork', data); }
  async applyAppWithdrawToCwallet(data: Record<string, unknown>): Promise<any> { return this.client.post('/applyAppWithdrawToCwallet', data); }
  async getAppWithdrawRecord(data: Record<string, unknown>): Promise<any> { return this.client.post('/getAppWithdrawRecord', data); }
  async getAppWithdrawRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getAppWithdrawRecordList', data); }
  async getAutoWithdrawRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getAutoWithdrawRecordList', data); }
  async getRiskRefundRecords(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getRiskyRefundRecordList', data); }
}
