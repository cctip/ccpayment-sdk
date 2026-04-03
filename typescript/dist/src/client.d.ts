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
export declare class Client {
    private static readonly DEFAULT_BASE_URL;
    private readonly appId;
    private readonly appSecret;
    private baseUrl;
    private axiosInstance;
    constructor(appId: string, appSecret: string);
    setBaseUrl(baseUrl: string): void;
    setProxy(proxyUrl: string): void;
    private generateSign;
    post<T>(path: string, data?: Record<string, unknown>): Promise<T>;
    get basicInfo(): BasicInfoService;
    get merchantAssets(): MerchantAssetsService;
    get merchantDeposit(): MerchantDepositService;
    get merchantWithdraw(): MerchantWithdrawService;
    get merchantBatchWithdraw(): MerchantBatchWithdrawService;
    get userAssets(): UserAssetsService;
    get userDeposit(): UserDepositService;
    get userWithdraw(): UserWithdrawService;
    get userTransfer(): UserTransferService;
    get orders(): OrdersService;
    get checkout(): CheckoutService;
    get swap(): SwapService;
    get utilities(): UtilitiesService;
}
//# sourceMappingURL=client.d.ts.map