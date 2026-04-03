"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Client = void 0;
const axios_1 = __importDefault(require("axios"));
const crypto_1 = __importDefault(require("crypto"));
const errors_1 = require("./errors");
const basic_info_1 = require("./services/basic-info");
const merchant_assets_1 = require("./services/merchant-assets");
const merchant_deposit_1 = require("./services/merchant-deposit");
const merchant_withdraw_1 = require("./services/merchant-withdraw");
const merchant_batch_withdraw_1 = require("./services/merchant-batch-withdraw");
const user_assets_1 = require("./services/user-assets");
const user_deposit_1 = require("./services/user-deposit");
const user_withdraw_1 = require("./services/user-withdraw");
const user_transfer_1 = require("./services/user-transfer");
const orders_1 = require("./services/orders");
const checkout_1 = require("./services/checkout");
const swap_1 = require("./services/swap");
const utilities_1 = require("./services/utilities");
class Client {
    constructor(appId, appSecret) {
        this.appId = appId;
        this.appSecret = appSecret;
        this.baseUrl = Client.DEFAULT_BASE_URL;
        this.axiosInstance = axios_1.default.create();
    }
    setBaseUrl(baseUrl) {
        this.baseUrl = baseUrl;
    }
    setProxy(proxyUrl) {
        // Parse proxy URL like "http://127.0.0.1:10808" or "127.0.0.1:10808"
        const cleanUrl = proxyUrl.replace(/^http:\/\//, '');
        const parts = cleanUrl.split(':');
        const host = parts[0];
        const port = parts[1] ? parseInt(parts[1]) : 80;
        this.axiosInstance = axios_1.default.create({
            proxy: {
                host: host,
                port: port,
                protocol: 'http'
            }
        });
    }
    generateSign(body) {
        const timestamp = Math.floor(Date.now() / 1000).toString();
        const signText = this.appId + timestamp + body;
        const sign = crypto_1.default
            .createHmac('sha256', this.appSecret)
            .update(signText)
            .digest('hex');
        return { sign, timestamp };
    }
    async post(path, data) {
        const body = data ? JSON.stringify(data) : '{}';
        const { sign, timestamp } = this.generateSign(body);
        const headers = {
            'Content-Type': 'application/json',
            'Appid': this.appId,
            'Sign': sign,
            'Timestamp': timestamp,
        };
        const response = await this.axiosInstance.post(`${this.baseUrl}${path}`, data || {}, { headers });
        const result = response.data;
        if (result.code !== 10000) {
            throw new errors_1.APIError(result.code, result.msg);
        }
        return result.data;
    }
    // Service accessors - using getters for property-like access
    get basicInfo() {
        return new basic_info_1.BasicInfoService(this);
    }
    get merchantAssets() {
        return new merchant_assets_1.MerchantAssetsService(this);
    }
    get merchantDeposit() {
        return new merchant_deposit_1.MerchantDepositService(this);
    }
    get merchantWithdraw() {
        return new merchant_withdraw_1.MerchantWithdrawService(this);
    }
    get merchantBatchWithdraw() {
        return new merchant_batch_withdraw_1.MerchantBatchWithdrawService(this);
    }
    get userAssets() {
        return new user_assets_1.UserAssetsService(this);
    }
    get userDeposit() {
        return new user_deposit_1.UserDepositService(this);
    }
    get userWithdraw() {
        return new user_withdraw_1.UserWithdrawService(this);
    }
    get userTransfer() {
        return new user_transfer_1.UserTransferService(this);
    }
    get orders() {
        return new orders_1.OrdersService(this);
    }
    get checkout() {
        return new checkout_1.CheckoutService(this);
    }
    get swap() {
        return new swap_1.SwapService(this);
    }
    get utilities() {
        return new utilities_1.UtilitiesService(this);
    }
}
exports.Client = Client;
Client.DEFAULT_BASE_URL = 'https://ccpayment.com/ccpayment/v2';
//# sourceMappingURL=client.js.map