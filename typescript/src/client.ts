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

  public setProxy(proxyUrl: string): void {
    // Parse proxy URL like "http://127.0.0.1:10808" or "127.0.0.1:10808"
    const cleanUrl = proxyUrl.replace(/^http:\/\//, '');
    const parts = cleanUrl.split(':');
    const host = parts[0];
    const port = parts[1] ? parseInt(parts[1]) : 80;
    
    this.axiosInstance = axios.create({
      proxy: {
        host: host,
        port: port,
        protocol: 'http'
      }
    });
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

  // Service accessors - using getters for property-like access
  get basicInfo(): BasicInfoService {
    return new BasicInfoService(this);
  }

  get merchantAssets(): MerchantAssetsService {
    return new MerchantAssetsService(this);
  }

  get merchantDeposit(): MerchantDepositService {
    return new MerchantDepositService(this);
  }

  get merchantWithdraw(): MerchantWithdrawService {
    return new MerchantWithdrawService(this);
  }

  get merchantBatchWithdraw(): MerchantBatchWithdrawService {
    return new MerchantBatchWithdrawService(this);
  }

  get userAssets(): UserAssetsService {
    return new UserAssetsService(this);
  }

  get userDeposit(): UserDepositService {
    return new UserDepositService(this);
  }

  get userWithdraw(): UserWithdrawService {
    return new UserWithdrawService(this);
  }

  get userTransfer(): UserTransferService {
    return new UserTransferService(this);
  }

  get orders(): OrdersService {
    return new OrdersService(this);
  }

  get checkout(): CheckoutService {
    return new CheckoutService(this);
  }

  get swap(): SwapService {
    return new SwapService(this);
  }

  get utilities(): UtilitiesService {
    return new UtilitiesService(this);
  }
}
