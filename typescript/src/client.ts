import axios, { AxiosInstance, AxiosProxyConfig } from 'axios';
import crypto from 'crypto';
import { APIError } from './errors';
import { BasicInfoService } from './services/basic-info';
import { MerchantAssetsService } from './services/merchant-assets';
import { MerchantDepositService } from './services/merchant-deposit';
import { MerchantWithdrawService } from './services/merchant-withdraw';
import { MerchantBatchWithdrawService } from './services/merchant-batch-withdraw';
import { UserAssetsService } from './services/user-assets';
import { UserDepositService } from './services/user-deposit';
import { UserWithdrawService } from './services/user-withdraw';
import { UserTransferService } from './services/user-transfer';
import { OrdersService } from './services/orders';
import { CheckoutService } from './services/checkout';
import { SwapService } from './services/swap';
import { UtilitiesService } from './services/utilities';

export class Client {
  private static readonly DEFAULT_BASE_URL = 'https://ccpayment.com/ccpayment/v2';
  private readonly appId: string;
  private readonly appSecret: string;
  private baseUrl: string;
  private axiosInstance: AxiosInstance;

  constructor(appId: string, appSecret: string) {
    this.appId = appId;
    this.appSecret = appSecret;
    this.baseUrl = Client.DEFAULT_BASE_URL;
    this.axiosInstance = axios.create();
  }

  public setBaseUrl(baseUrl: string): void {
    this.baseUrl = baseUrl;
  }

  public setProxy(proxy: AxiosProxyConfig): void {
    this.axiosInstance.defaults.proxy = proxy;
  }

  private generateSign(body: string): { sign: string; timestamp: string } {
    const timestamp = Math.floor(Date.now() / 1000).toString();
    const signText = this.appId + timestamp + body;
    const sign = crypto
      .createHmac('sha256', this.appSecret)
      .update(signText)
      .digest('hex');
    return { sign, timestamp };
  }

  public async post<T>(path: string, data?: Record<string, unknown>): Promise<T> {
    const body = data ? JSON.stringify(data) : '{}';
    const { sign, timestamp } = this.generateSign(body);

    const headers = {
      'Content-Type': 'application/json',
      'Appid': this.appId,
      'Sign': sign,
      'Timestamp': timestamp,
    };

    const response = await this.axiosInstance.post(
      `${this.baseUrl}${path}`,
      data || {},
      { headers }
    );

    const result = response.data as { code: number; msg: string; data: T };
    if (result.code !== 10000) {
      throw new APIError(result.code, result.msg);
    }

    return result.data;
  }

  // Service accessors
  public basicInfo(): BasicInfoService {
    return new BasicInfoService(this);
  }

  public merchantAssets(): MerchantAssetsService {
    return new MerchantAssetsService(this);
  }

  public merchantDeposit(): MerchantDepositService {
    return new MerchantDepositService(this);
  }

  public merchantWithdraw(): MerchantWithdrawService {
    return new MerchantWithdrawService(this);
  }

  public merchantBatchWithdraw(): MerchantBatchWithdrawService {
    return new MerchantBatchWithdrawService(this);
  }

  public userAssets(): UserAssetsService {
    return new UserAssetsService(this);
  }

  public userDeposit(): UserDepositService {
    return new UserDepositService(this);
  }

  public userWithdraw(): UserWithdrawService {
    return new UserWithdrawService(this);
  }

  public userTransfer(): UserTransferService {
    return new UserTransferService(this);
  }

  public orders(): OrdersService {
    return new OrdersService(this);
  }

  public checkout(): CheckoutService {
    return new CheckoutService(this);
  }

  public swap(): SwapService {
    return new SwapService(this);
  }

  public utilities(): UtilitiesService {
    return new UtilitiesService(this);
  }
}
