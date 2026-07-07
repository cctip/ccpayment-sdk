from pathlib import Path


ROOT = Path("/Users/oks/Documents/Projects/MerchantProjects/src/ccpayment-sdk")


def write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content.rstrip() + "\n")


def remove(path: Path) -> None:
    if path.exists():
        path.unlink()


def patch_readme(path: Path) -> None:
    text = path.read_text()
    text = text.replace("- **checkout** - Checkout and Hosted related\n", "")
    text = text.replace("- **Checkout** - Checkout and Hosted related\n", "")
    if "- **user-swap** - User swap management\n" not in text and "- **User Swap** - User swap management\n" not in text:
        text = text.replace("- **swap** - Swap functionality\n", "- **swap** - Swap functionality\n- **user-swap** - User swap management\n")
        text = text.replace("- **Swap** - Swap related interfaces\n", "- **Swap** - Swap related interfaces\n- **User Swap** - User swap management\n")
    path.write_text(text)


def regenerate_python() -> None:
    base = ROOT / "python/ccpayment"
    write(base / "client.py", '''"""CCPayment SDK Client"""

import hashlib
import hmac
import json
import time
from typing import Optional

import requests

from .exceptions import APIError


class Client:
    """CCPayment API Client"""

    DEFAULT_BASE_URL = "https://ccpayment.com/ccpayment/v2"

    def __init__(self, app_id: str, app_secret: str):
        self.app_id = app_id
        self.app_secret = app_secret
        self.base_url = self.DEFAULT_BASE_URL
        self.session = requests.Session()

    def set_base_url(self, base_url: str):
        self.base_url = base_url

    def set_proxy(self, proxy_url: str):
        self.session.proxies = {"http": proxy_url, "https": proxy_url}

    def _generate_sign(self, body: str) -> tuple[str, str]:
        timestamp = str(int(time.time()))
        sign_text = self.app_id + timestamp + body
        sign = hmac.new(self.app_secret.encode(), sign_text.encode(), hashlib.sha256).hexdigest()
        return sign, timestamp

    def _post(self, path: str, data: Optional[dict] = None) -> dict:
        body = json.dumps(data) if data else "{}"
        sign, timestamp = self._generate_sign(body)
        headers = {
            "Content-Type": "application/json",
            "Appid": self.app_id,
            "Sign": sign,
            "Timestamp": timestamp,
        }
        response = self.session.post(f"{self.base_url}{path}", headers=headers, data=body)
        response.raise_for_status()
        result = response.json()
        if result.get("code") != 10000:
            raise APIError(result.get("code"), result.get("msg"))
        return result.get("data", {})

    @property
    def basic_info(self):
        from .services.basic_info import BasicInfoService
        if not hasattr(self, "_basic_info"):
            self._basic_info = BasicInfoService(self)
        return self._basic_info

    @property
    def merchant_assets(self):
        from .services.merchant_assets import MerchantAssetsService
        if not hasattr(self, "_merchant_assets"):
            self._merchant_assets = MerchantAssetsService(self)
        return self._merchant_assets

    @property
    def merchant_deposit(self):
        from .services.merchant_deposit import MerchantDepositService
        if not hasattr(self, "_merchant_deposit"):
            self._merchant_deposit = MerchantDepositService(self)
        return self._merchant_deposit

    @property
    def merchant_withdraw(self):
        from .services.merchant_withdraw import MerchantWithdrawService
        if not hasattr(self, "_merchant_withdraw"):
            self._merchant_withdraw = MerchantWithdrawService(self)
        return self._merchant_withdraw

    @property
    def merchant_batch_withdraw(self):
        from .services.merchant_batch_withdraw import MerchantBatchWithdrawService
        if not hasattr(self, "_merchant_batch_withdraw"):
            self._merchant_batch_withdraw = MerchantBatchWithdrawService(self)
        return self._merchant_batch_withdraw

    @property
    def user_assets(self):
        from .services.user_assets import UserAssetsService
        if not hasattr(self, "_user_assets"):
            self._user_assets = UserAssetsService(self)
        return self._user_assets

    @property
    def user_deposit(self):
        from .services.user_deposit import UserDepositService
        if not hasattr(self, "_user_deposit"):
            self._user_deposit = UserDepositService(self)
        return self._user_deposit

    @property
    def user_withdraw(self):
        from .services.user_withdraw import UserWithdrawService
        if not hasattr(self, "_user_withdraw"):
            self._user_withdraw = UserWithdrawService(self)
        return self._user_withdraw

    @property
    def user_transfer(self):
        from .services.user_transfer import UserTransferService
        if not hasattr(self, "_user_transfer"):
            self._user_transfer = UserTransferService(self)
        return self._user_transfer

    @property
    def orders(self):
        from .services.orders import OrdersService
        if not hasattr(self, "_orders"):
            self._orders = OrdersService(self)
        return self._orders

    @property
    def swap(self):
        from .services.swap import SwapService
        if not hasattr(self, "_swap"):
            self._swap = SwapService(self)
        return self._swap

    @property
    def user_swap(self):
        from .services.user_swap import UserSwapService
        if not hasattr(self, "_user_swap"):
            self._user_swap = UserSwapService(self)
        return self._user_swap

    @property
    def utilities(self):
        from .services.utilities import UtilitiesService
        if not hasattr(self, "_utilities"):
            self._utilities = UtilitiesService(self)
        return self._utilities
''')
    services = base / "services"
    write(services / "basic_info.py", '''"""Basic Info Service"""


class BasicInfoService:
    def __init__(self, client):
        self.client = client

    def get_coin_list(self):
        return self.client._post("/getCoinList", {})

    def get_fiat_list(self):
        return self.client._post("/getFiatList", {})

    def get_coin(self, coin_id: int):
        return self.client._post("/getCoin", {"coinId": coin_id})

    def get_coin_usdt_price(self, coin_ids):
        return self.client._post("/getCoinUSDTPrice", {"coinIds": coin_ids})

    def get_chain_list(self, data=None):
        return self.client._post("/getChainList", data or {})

    def get_all_chain_list(self, data=None):
        return self.client._post("/all/chain", data or {})

    def get_main_coin_list(self, app_id: str):
        return self.client._post("/getMainCoinList", {"appId": app_id})
''')
    write(services / "merchant_assets.py", '''"""Merchant Assets Service"""


class MerchantAssetsService:
    def __init__(self, client):
        self.client = client

    def get_app_coin_asset_list(self):
        return self.client._post("/getAppCoinAssetList", {})

    def get_app_coin_asset(self, coin_id: int):
        return self.client._post("/getAppCoinAsset", {"coinId": coin_id})

    def get_app_collect_fee_record_list(self, data=None):
        return self.client._post("/getAggregationFeeRecord", data or {})
''')
    write(services / "merchant_withdraw.py", '''"""Merchant Withdraw Service"""


class MerchantWithdrawService:
    def __init__(self, client):
        self.client = client

    def get_cwallet_user_id(self, cwallet_user_id: str):
        return self.client._post("/getCwalletUserId", {"cwalletUserId": cwallet_user_id})

    def get_withdraw_fee(self, coin_id: int, chain: str):
        return self.client._post("/getWithdrawFee", {"coinId": coin_id, "chain": chain})

    def apply_app_withdraw_to_network(self, data):
        return self.client._post("/applyAppWithdrawToNetwork", data)

    def apply_app_withdraw_to_cwallet(self, data):
        return self.client._post("/applyAppWithdrawToCwallet", data)

    def get_app_withdraw_record(self, data):
        return self.client._post("/getAppWithdrawRecord", data)

    def get_app_withdraw_record_list(self, data=None):
        return self.client._post("/getAppWithdrawRecordList", data or {})

    def get_auto_withdraw_record_list(self, data=None):
        return self.client._post("/getAutoWithdrawRecordList", data or {})

    def get_risk_refund_records(self, data=None):
        return self.client._post("/getRiskyRefundRecordList", data or {})
''')
    write(services / "orders.py", '''"""Orders Service"""


class OrdersService:
    def __init__(self, client):
        self.client = client

    def get_app_order_info(self, order_id: str):
        return self.client._post("/getAppOrderInfo", {"orderId": order_id})

    def create_invoice_url(self, data):
        return self.client._post("/createInvoiceUrl", data)

    def get_invoice_order_info(self, order_id: str):
        return self.client._post("/getInvoiceOrderInfo", {"orderId": order_id})
''')
    write(services / "swap.py", '''"""Swap Service"""


class SwapService:
    def __init__(self, client):
        self.client = client

    def get_swap_coin_list(self):
        return self.client._post("/getSwapCoinList", {})

    def swap_estimate(self, data):
        return self.client._post("/estimate", data)

    def get_swap_record(self, data):
        return self.client._post("/getSwapRecord", data)

    def get_swap_record_list(self, data=None):
        return self.client._post("/getSwapRecordList", data or {})

    def swap(self, data):
        return self.client._post("/swap", data)
''')
    write(services / "user_swap.py", '''"""User Swap Service"""


class UserSwapService:
    def __init__(self, client):
        self.client = client

    def get_user_swap_record(self, data):
        return self.client._post("/getUserSwapRecord", data)

    def get_user_swap_record_list(self, data=None):
        return self.client._post("/getUserSwapRecordList", data or {})

    def user_swap(self, data):
        return self.client._post("/userSwap", data)
''')
    write(services / "utilities.py", '''"""Utilities Service"""


class UtilitiesService:
    def __init__(self, client):
        self.client = client

    def webhook_resend(self, data=None):
        return self.client._post("/webhook/resend", data or {})

    def get_trade_block_height(self, data):
        return self.client._post("/getTradeBlockHeight", data)

    def check_withdrawal_address_validity(self, data):
        return self.client._post("/checkWithdrawalAddressValidity", data)

    def deprecated_address(self, data):
        return self.client._post("/addressUnbinding", data)

    def rescan_lost_transaction(self, data):
        return self.client._post("/rescanLostTransaction", data)
''')
    remove(services / "checkout.py")
    patch_readme(ROOT / "python/README.md")


def regenerate_typescript() -> None:
    base = ROOT / "typescript/src"
    write(base / "client.ts", '''import axios, { AxiosInstance } from 'axios';
import crypto from 'crypto';
import { HttpsProxyAgent } from 'https-proxy-agent';
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
import { SwapService } from './services/swap';
import { UserSwapService } from './services/user-swap';
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
    const agent = new HttpsProxyAgent(proxyUrl);
    this.axiosInstance = axios.create({ httpsAgent: agent, proxy: false });
  }

  private generateSign(body: string): { sign: string; timestamp: string } {
    const timestamp = Math.floor(Date.now() / 1000).toString();
    const signText = this.appId + timestamp + body;
    const sign = crypto.createHmac('sha256', this.appSecret).update(signText).digest('hex');
    return { sign, timestamp };
  }

  public async post<T>(path: string, data?: Record<string, unknown>): Promise<T> {
    const body = data ? JSON.stringify(data) : '{}';
    const { sign, timestamp } = this.generateSign(body);
    const headers = {
      'Content-Type': 'application/json',
      Appid: this.appId,
      Sign: sign,
      Timestamp: timestamp,
    };
    const response = await this.axiosInstance.post(`${this.baseUrl}${path}`, data || {}, { headers });
    const result = response.data as { code: number; msg: string; data: T };
    if (result.code !== 10000) {
      throw new APIError(result.code, result.msg);
    }
    return result.data;
  }

  get basicInfo(): BasicInfoService { return new BasicInfoService(this); }
  get merchantAssets(): MerchantAssetsService { return new MerchantAssetsService(this); }
  get merchantDeposit(): MerchantDepositService { return new MerchantDepositService(this); }
  get merchantWithdraw(): MerchantWithdrawService { return new MerchantWithdrawService(this); }
  get merchantBatchWithdraw(): MerchantBatchWithdrawService { return new MerchantBatchWithdrawService(this); }
  get userAssets(): UserAssetsService { return new UserAssetsService(this); }
  get userDeposit(): UserDepositService { return new UserDepositService(this); }
  get userWithdraw(): UserWithdrawService { return new UserWithdrawService(this); }
  get userTransfer(): UserTransferService { return new UserTransferService(this); }
  get orders(): OrdersService { return new OrdersService(this); }
  get swap(): SwapService { return new SwapService(this); }
  get userSwap(): UserSwapService { return new UserSwapService(this); }
  get utilities(): UtilitiesService { return new UtilitiesService(this); }
}
''')
    svc = base / "services"
    write(svc / "basic-info.ts", '''import { Client } from '../client';

export class BasicInfoService {
  constructor(private client: Client) {}
  async getCoinList(): Promise<any> { return this.client.post('/getCoinList', {}); }
  async getFiatList(): Promise<any> { return this.client.post('/getFiatList', {}); }
  async getCoin(coinId: number): Promise<any> { return this.client.post('/getCoin', { coinId }); }
  async getCoinUSDTPrice(coinIds: number[]): Promise<any> { return this.client.post('/getCoinUSDTPrice', { coinIds }); }
  async getChainList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getChainList', data); }
  async getAllChainList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/all/chain', data); }
  async getMainCoinList(appId: string): Promise<any> { return this.client.post('/getMainCoinList', { appId }); }
}
''')
    write(svc / "merchant-assets.ts", '''import { Client } from '../client';

export class MerchantAssetsService {
  constructor(private client: Client) {}
  async getAppCoinAssetList(): Promise<any> { return this.client.post('/getAppCoinAssetList', {}); }
  async getAppCoinAsset(coinId: number): Promise<any> { return this.client.post('/getAppCoinAsset', { coinId }); }
  async getAppCollectFeeRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getAggregationFeeRecord', data); }
}
''')
    write(svc / "merchant-withdraw.ts", '''import { Client } from '../client';

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
''')
    write(svc / "orders.ts", '''import { Client } from '../client';

export class OrdersService {
  constructor(private client: Client) {}
  async getAppOrderInfo(orderId: string): Promise<any> { return this.client.post('/getAppOrderInfo', { orderId }); }
  async createInvoiceUrl(data: Record<string, unknown>): Promise<any> { return this.client.post('/createInvoiceUrl', data); }
  async getInvoiceOrderInfo(orderId: string): Promise<any> { return this.client.post('/getInvoiceOrderInfo', { orderId }); }
}
''')
    write(svc / "swap.ts", '''import { Client } from '../client';

export class SwapService {
  constructor(private client: Client) {}
  async getSwapCoinList(): Promise<any> { return this.client.post('/getSwapCoinList', {}); }
  async swapEstimate(data: Record<string, unknown>): Promise<any> { return this.client.post('/estimate', data); }
  async getSwapRecord(data: Record<string, unknown>): Promise<any> { return this.client.post('/getSwapRecord', data); }
  async getSwapRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getSwapRecordList', data); }
  async swap(data: Record<string, unknown>): Promise<any> { return this.client.post('/swap', data); }
}
''')
    write(svc / "user-swap.ts", '''import { Client } from '../client';

export class UserSwapService {
  constructor(private client: Client) {}
  async getUserSwapRecord(data: Record<string, unknown>): Promise<any> { return this.client.post('/getUserSwapRecord', data); }
  async getUserSwapRecordList(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/getUserSwapRecordList', data); }
  async userSwap(data: Record<string, unknown>): Promise<any> { return this.client.post('/userSwap', data); }
}
''')
    write(svc / "utilities.ts", '''import { Client } from '../client';

export class UtilitiesService {
  constructor(private client: Client) {}
  async webhookResend(data: Record<string, unknown> = {}): Promise<any> { return this.client.post('/webhook/resend', data); }
  async getTradeBlockHeight(data: Record<string, unknown>): Promise<any> { return this.client.post('/getTradeBlockHeight', data); }
  async checkWithdrawalAddressValidity(data: Record<string, unknown>): Promise<any> { return this.client.post('/checkWithdrawalAddressValidity', data); }
  async deprecatedAddress(data: Record<string, unknown>): Promise<any> { return this.client.post('/addressUnbinding', data); }
  async rescanLostTransaction(data: Record<string, unknown>): Promise<any> { return this.client.post('/rescanLostTransaction', data); }
}
''')
    remove(svc / "checkout.ts")
    patch_readme(ROOT / "typescript/README.md")


def regenerate_javascript() -> None:
    base = ROOT / "javascript/src"
    write(base / "client.js", '''const axios = require('axios');
const crypto = require('crypto');
const { HttpsProxyAgent } = require('https-proxy-agent');
const { APIError } = require('./errors');

const DEFAULT_BASE_URL = 'https://ccpayment.com/ccpayment/v2';

class Client {
  constructor(appId, appSecret) {
    this.appId = appId;
    this.appSecret = appSecret;
    this.baseUrl = DEFAULT_BASE_URL;
    this.axiosInstance = axios.create();
  }

  setBaseUrl(baseUrl) {
    this.baseUrl = baseUrl;
  }

  setProxy(proxyUrl) {
    const agent = new HttpsProxyAgent(proxyUrl);
    this.axiosInstance.defaults.httpsAgent = agent;
    this.axiosInstance.defaults.proxy = false;
  }

  generateSign(body) {
    const timestamp = Math.floor(Date.now() / 1000).toString();
    const signText = this.appId + timestamp + body;
    const sign = crypto.createHmac('sha256', this.appSecret).update(signText).digest('hex');
    return { sign, timestamp };
  }

  async post(path, data = {}) {
    const body = JSON.stringify(data);
    const { sign, timestamp } = this.generateSign(body);
    const headers = {
      'Content-Type': 'application/json',
      Appid: this.appId,
      Sign: sign,
      Timestamp: timestamp,
    };
    const response = await this.axiosInstance.post(`${this.baseUrl}${path}`, data, { headers });
    const result = response.data;
    if (result.code !== 10000) {
      throw new APIError(result.code, result.msg);
    }
    return result.data;
  }

  basicInfo() { return new (require('./services/basic-info'))(this); }
  merchantAssets() { return new (require('./services/merchant-assets'))(this); }
  merchantDeposit() { return new (require('./services/merchant-deposit'))(this); }
  merchantWithdraw() { return new (require('./services/merchant-withdraw'))(this); }
  merchantBatchWithdraw() { return new (require('./services/merchant-batch-withdraw'))(this); }
  userAssets() { return new (require('./services/user-assets'))(this); }
  userDeposit() { return new (require('./services/user-deposit'))(this); }
  userWithdraw() { return new (require('./services/user-withdraw'))(this); }
  userTransfer() { return new (require('./services/user-transfer'))(this); }
  orders() { return new (require('./services/orders'))(this); }
  swap() { return new (require('./services/swap'))(this); }
  userSwap() { return new (require('./services/user-swap'))(this); }
  utilities() { return new (require('./services/utilities'))(this); }
}

module.exports = { Client };
''')
    svc = base / "services"
    write(svc / "basic-info.js", '''class BasicInfoService {
  constructor(client) { this.client = client; }
  async getCoinList() { return this.client.post('/getCoinList', {}); }
  async getFiatList() { return this.client.post('/getFiatList', {}); }
  async getCoin(coinId) { return this.client.post('/getCoin', { coinId }); }
  async getCoinUSDTPrice(coinIds) { return this.client.post('/getCoinUSDTPrice', { coinIds }); }
  async getChainList(data = {}) { return this.client.post('/getChainList', data); }
  async getAllChainList(data = {}) { return this.client.post('/all/chain', data); }
  async getMainCoinList(appId) { return this.client.post('/getMainCoinList', { appId }); }
}
module.exports = BasicInfoService;
''')
    write(svc / "merchant-assets.js", '''class MerchantAssetsService {
  constructor(client) { this.client = client; }
  async getAppCoinAssetList() { return this.client.post('/getAppCoinAssetList', {}); }
  async getAppCoinAsset(coinId) { return this.client.post('/getAppCoinAsset', { coinId }); }
  async getAppCollectFeeRecordList(data = {}) { return this.client.post('/getAggregationFeeRecord', data); }
}
module.exports = MerchantAssetsService;
''')
    write(svc / "merchant-withdraw.js", '''class MerchantWithdrawService {
  constructor(client) { this.client = client; }
  async getCwalletUserId(cwalletUserId) { return this.client.post('/getCwalletUserId', { cwalletUserId }); }
  async getWithdrawFee(coinId, chain) { return this.client.post('/getWithdrawFee', { coinId, chain }); }
  async applyAppWithdrawToNetwork(data) { return this.client.post('/applyAppWithdrawToNetwork', data); }
  async applyAppWithdrawToCwallet(data) { return this.client.post('/applyAppWithdrawToCwallet', data); }
  async getAppWithdrawRecord(data) { return this.client.post('/getAppWithdrawRecord', data); }
  async getAppWithdrawRecordList(data = {}) { return this.client.post('/getAppWithdrawRecordList', data); }
  async getAutoWithdrawRecordList(data = {}) { return this.client.post('/getAutoWithdrawRecordList', data); }
  async getRiskRefundRecords(data = {}) { return this.client.post('/getRiskyRefundRecordList', data); }
}
module.exports = MerchantWithdrawService;
''')
    write(svc / "orders.js", '''class OrdersService {
  constructor(client) { this.client = client; }
  async getAppOrderInfo(orderId) { return this.client.post('/getAppOrderInfo', { orderId }); }
  async createInvoiceUrl(data) { return this.client.post('/createInvoiceUrl', data); }
  async getInvoiceOrderInfo(orderId) { return this.client.post('/getInvoiceOrderInfo', { orderId }); }
}
module.exports = OrdersService;
''')
    write(svc / "swap.js", '''class SwapService {
  constructor(client) { this.client = client; }
  async getSwapCoinList() { return this.client.post('/getSwapCoinList', {}); }
  async swapEstimate(data) { return this.client.post('/estimate', data); }
  async getSwapRecord(data) { return this.client.post('/getSwapRecord', data); }
  async getSwapRecordList(data = {}) { return this.client.post('/getSwapRecordList', data); }
  async swap(data) { return this.client.post('/swap', data); }
}
module.exports = SwapService;
''')
    write(svc / "user-swap.js", '''class UserSwapService {
  constructor(client) { this.client = client; }
  async getUserSwapRecord(data) { return this.client.post('/getUserSwapRecord', data); }
  async getUserSwapRecordList(data = {}) { return this.client.post('/getUserSwapRecordList', data); }
  async userSwap(data) { return this.client.post('/userSwap', data); }
}
module.exports = UserSwapService;
''')
    write(svc / "utilities.js", '''class UtilitiesService {
  constructor(client) { this.client = client; }
  async webhookResend(data = {}) { return this.client.post('/webhook/resend', data); }
  async getTradeBlockHeight(data) { return this.client.post('/getTradeBlockHeight', data); }
  async checkWithdrawalAddressValidity(data) { return this.client.post('/checkWithdrawalAddressValidity', data); }
  async deprecatedAddress(data) { return this.client.post('/addressUnbinding', data); }
  async rescanLostTransaction(data) { return this.client.post('/rescanLostTransaction', data); }
}
module.exports = UtilitiesService;
''')
    remove(svc / "checkout.js")
    patch_readme(ROOT / "javascript/README.md")


def regenerate_java() -> None:
    base = ROOT / "java/src/main/java/com/ccpayment"
    write(base / "Client.java", '''package com.ccpayment;

import com.ccpayment.services.*;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import okhttp3.*;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.TimeUnit;

public class Client {
    private static final String DEFAULT_BASE_URL = "https://ccpayment.com/ccpayment/v2";
    private static final Gson gson = new Gson();
    private static final MediaType JSON = MediaType.parse("application/json");

    private final String appId;
    private final String appSecret;
    private String baseUrl;
    private OkHttpClient httpClient;

    public Client(String appId, String appSecret) {
        this.appId = appId;
        this.appSecret = appSecret;
        this.baseUrl = DEFAULT_BASE_URL;
        this.httpClient = new OkHttpClient.Builder()
                .connectTimeout(30, TimeUnit.SECONDS)
                .readTimeout(30, TimeUnit.SECONDS)
                .build();
    }

    public void setBaseUrl(String baseUrl) { this.baseUrl = baseUrl; }

    public void setProxy(java.net.Proxy proxy) {
        this.httpClient = new OkHttpClient.Builder()
                .connectTimeout(30, TimeUnit.SECONDS)
                .readTimeout(30, TimeUnit.SECONDS)
                .proxy(proxy)
                .build();
    }

    private String generateSign(String body, String timestamp) throws Exception {
        String signText = appId + timestamp + body;
        Mac mac = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKey = new SecretKeySpec(appSecret.getBytes(StandardCharsets.UTF_8), "HmacSHA256");
        mac.init(secretKey);
        byte[] hash = mac.doFinal(signText.getBytes(StandardCharsets.UTF_8));
        StringBuilder hexString = new StringBuilder();
        for (byte b : hash) {
            String hex = Integer.toHexString(0xff & b);
            if (hex.length() == 1) hexString.append('0');
            hexString.append(hex);
        }
        return hexString.toString();
    }

    public <T> T post(String path, Object data, Class<T> responseClass) throws APIError, IOException {
        try {
            String body = data != null ? gson.toJson(data) : "{}";
            String timestamp = String.valueOf(System.currentTimeMillis() / 1000);
            String sign = generateSign(body, timestamp);
            RequestBody requestBody = RequestBody.create(body, JSON);
            Request request = new Request.Builder()
                    .url(baseUrl + path)
                    .post(requestBody)
                    .addHeader("Content-Type", "application/json")
                    .addHeader("Appid", appId)
                    .addHeader("Sign", sign)
                    .addHeader("Timestamp", timestamp)
                    .build();
            try (Response response = httpClient.newCall(request).execute()) {
                if (!response.isSuccessful()) {
                    throw new IOException("Unexpected code " + response);
                }
                String responseBody = response.body().string();
                JsonObject jsonObject = gson.fromJson(responseBody, JsonObject.class);
                int code = jsonObject.get("code").getAsInt();
                if (code != 10000) {
                    throw new APIError(code, jsonObject.get("msg").getAsString());
                }
                if (responseClass == Void.class) {
                    return null;
                }
                if (jsonObject.has("data") && !jsonObject.get("data").isJsonNull()) {
                    return gson.fromJson(jsonObject.get("data"), responseClass);
                }
                return null;
            }
        } catch (APIError | IOException e) {
            throw e;
        } catch (Exception e) {
            throw new IOException("Request failed", e);
        }
    }

    public BasicInfoService basicInfo() { return new BasicInfoService(this); }
    public MerchantAssetsService merchantAssets() { return new MerchantAssetsService(this); }
    public MerchantDepositService merchantDeposit() { return new MerchantDepositService(this); }
    public MerchantWithdrawService merchantWithdraw() { return new MerchantWithdrawService(this); }
    public MerchantBatchWithdrawService merchantBatchWithdraw() { return new MerchantBatchWithdrawService(this); }
    public UserAssetsService userAssets() { return new UserAssetsService(this); }
    public UserDepositService userDeposit() { return new UserDepositService(this); }
    public UserWithdrawService userWithdraw() { return new UserWithdrawService(this); }
    public UserTransferService userTransfer() { return new UserTransferService(this); }
    public OrdersService orders() { return new OrdersService(this); }
    public SwapService swap() { return new SwapService(this); }
    public UserSwapService userSwap() { return new UserSwapService(this); }
    public UtilitiesService utilities() { return new UtilitiesService(this); }
}
''')
    svc = base / "services"
    write(svc / "BasicInfoService.java", '''package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class BasicInfoService {
    private final Client client;
    public BasicInfoService(Client client) { this.client = client; }
    public JsonObject getCoinList() throws APIError, IOException { return client.post("/getCoinList", null, JsonObject.class); }
    public JsonObject getFiatList() throws APIError, IOException { return client.post("/getFiatList", null, JsonObject.class); }
    public JsonObject getCoin(JsonObject data) throws APIError, IOException { return client.post("/getCoin", data, JsonObject.class); }
    public JsonObject getCoinUSDTPrice(JsonObject data) throws APIError, IOException { return client.post("/getCoinUSDTPrice", data, JsonObject.class); }
    public JsonObject getChainList(JsonObject data) throws APIError, IOException { return client.post("/getChainList", data, JsonObject.class); }
    public JsonObject getAllChainList(JsonObject data) throws APIError, IOException { return client.post("/all/chain", data, JsonObject.class); }
    public JsonObject getMainCoinList(JsonObject data) throws APIError, IOException { return client.post("/getMainCoinList", data, JsonObject.class); }
}
''')
    write(svc / "MerchantAssetsService.java", '''package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class MerchantAssetsService {
    private final Client client;
    public MerchantAssetsService(Client client) { this.client = client; }
    public JsonObject getAppCoinAssetList() throws APIError, IOException { return client.post("/getAppCoinAssetList", null, JsonObject.class); }
    public JsonObject getAppCoinAsset(JsonObject data) throws APIError, IOException { return client.post("/getAppCoinAsset", data, JsonObject.class); }
    public JsonObject getAppCollectFeeRecordList(JsonObject data) throws APIError, IOException { return client.post("/getAggregationFeeRecord", data, JsonObject.class); }
}
''')
    write(svc / "MerchantWithdrawService.java", '''package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class MerchantWithdrawService {
    private final Client client;
    public MerchantWithdrawService(Client client) { this.client = client; }
    public JsonObject getCwalletUserId(JsonObject data) throws APIError, IOException { return client.post("/getCwalletUserId", data, JsonObject.class); }
    public JsonObject getWithdrawFee(JsonObject data) throws APIError, IOException { return client.post("/getWithdrawFee", data, JsonObject.class); }
    public JsonObject applyAppWithdrawToNetwork(JsonObject data) throws APIError, IOException { return client.post("/applyAppWithdrawToNetwork", data, JsonObject.class); }
    public JsonObject applyAppWithdrawToCwallet(JsonObject data) throws APIError, IOException { return client.post("/applyAppWithdrawToCwallet", data, JsonObject.class); }
    public JsonObject getAppWithdrawRecord(JsonObject data) throws APIError, IOException { return client.post("/getAppWithdrawRecord", data, JsonObject.class); }
    public JsonObject getAppWithdrawRecordList(JsonObject data) throws APIError, IOException { return client.post("/getAppWithdrawRecordList", data, JsonObject.class); }
    public JsonObject getAutoWithdrawRecordList(JsonObject data) throws APIError, IOException { return client.post("/getAutoWithdrawRecordList", data, JsonObject.class); }
    public JsonObject getRiskRefundRecords(JsonObject data) throws APIError, IOException { return client.post("/getRiskyRefundRecordList", data, JsonObject.class); }
}
''')
    write(svc / "OrdersService.java", '''package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class OrdersService {
    private final Client client;
    public OrdersService(Client client) { this.client = client; }
    public JsonObject getAppOrderInfo(JsonObject data) throws APIError, IOException { return client.post("/getAppOrderInfo", data, JsonObject.class); }
    public JsonObject createInvoiceUrl(JsonObject data) throws APIError, IOException { return client.post("/createInvoiceUrl", data, JsonObject.class); }
    public JsonObject getInvoiceOrderInfo(JsonObject data) throws APIError, IOException { return client.post("/getInvoiceOrderInfo", data, JsonObject.class); }
}
''')
    write(svc / "SwapService.java", '''package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class SwapService {
    private final Client client;
    public SwapService(Client client) { this.client = client; }
    public JsonObject getSwapCoinList() throws APIError, IOException { return client.post("/getSwapCoinList", null, JsonObject.class); }
    public JsonObject swapEstimate(JsonObject data) throws APIError, IOException { return client.post("/estimate", data, JsonObject.class); }
    public JsonObject getSwapRecord(JsonObject data) throws APIError, IOException { return client.post("/getSwapRecord", data, JsonObject.class); }
    public JsonObject getSwapRecordList(JsonObject data) throws APIError, IOException { return client.post("/getSwapRecordList", data, JsonObject.class); }
    public JsonObject swap(JsonObject data) throws APIError, IOException { return client.post("/swap", data, JsonObject.class); }
}
''')
    write(svc / "UserSwapService.java", '''package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UserSwapService {
    private final Client client;
    public UserSwapService(Client client) { this.client = client; }
    public JsonObject getUserSwapRecord(JsonObject data) throws APIError, IOException { return client.post("/getUserSwapRecord", data, JsonObject.class); }
    public JsonObject getUserSwapRecordList(JsonObject data) throws APIError, IOException { return client.post("/getUserSwapRecordList", data, JsonObject.class); }
    public JsonObject userSwap(JsonObject data) throws APIError, IOException { return client.post("/userSwap", data, JsonObject.class); }
}
''')
    write(svc / "UtilitiesService.java", '''package com.ccpayment.services;

import com.ccpayment.APIError;
import com.ccpayment.Client;
import com.google.gson.JsonObject;

import java.io.IOException;

public class UtilitiesService {
    private final Client client;
    public UtilitiesService(Client client) { this.client = client; }
    public JsonObject webhookResend(JsonObject data) throws APIError, IOException { return client.post("/webhook/resend", data, JsonObject.class); }
    public JsonObject getTradeBlockHeight(JsonObject data) throws APIError, IOException { return client.post("/getTradeBlockHeight", data, JsonObject.class); }
    public JsonObject checkWithdrawalAddressValidity(JsonObject data) throws APIError, IOException { return client.post("/checkWithdrawalAddressValidity", data, JsonObject.class); }
    public JsonObject deprecatedAddress(JsonObject data) throws APIError, IOException { return client.post("/addressUnbinding", data, JsonObject.class); }
    public JsonObject rescanLostTransaction(JsonObject data) throws APIError, IOException { return client.post("/rescanLostTransaction", data, JsonObject.class); }
}
''')
    remove(svc / "CheckoutService.java")
    patch_readme(ROOT / "java/README.md")


def regenerate_php() -> None:
    base = ROOT / "php/src"
    write(base / "Client.php", '''<?php

namespace CCPayment;

use GuzzleHttp\\Client as GuzzleClient;
use GuzzleHttp\\Exception\\GuzzleException;

class Client
{
    private const DEFAULT_BASE_URL = 'https://ccpayment.com/ccpayment/v2';

    private string $appId;
    private string $appSecret;
    private string $baseUrl;
    private GuzzleClient $httpClient;

    public function __construct(string $appId, string $appSecret)
    {
        $this->appId = $appId;
        $this->appSecret = $appSecret;
        $this->baseUrl = self::DEFAULT_BASE_URL;
        $this->httpClient = new GuzzleClient();
    }

    public function setBaseUrl(string $baseUrl): void { $this->baseUrl = $baseUrl; }
    public function setProxy(string $proxyUrl): void { $this->httpClient = new GuzzleClient(['proxy' => $proxyUrl]); }

    private function generateSign(string $body): array
    {
        $timestamp = (string) time();
        $signText = $this->appId . $timestamp . $body;
        $sign = hash_hmac('sha256', $signText, $this->appSecret);
        return ['sign' => $sign, 'timestamp' => $timestamp];
    }

    public function post(string $path, array $data = []): array
    {
        $body = json_encode($data);
        ['sign' => $sign, 'timestamp' => $timestamp] = $this->generateSign($body);
        $response = $this->httpClient->post($this->baseUrl . $path, [
            'headers' => [
                'Content-Type' => 'application/json',
                'Appid' => $this->appId,
                'Sign' => $sign,
                'Timestamp' => $timestamp,
            ],
            'body' => $body,
        ]);
        $result = json_decode($response->getBody()->getContents(), true);
        if ($result['code'] !== 10000) {
            throw new APIError($result['code'], $result['msg']);
        }
        return $result['data'] ?? [];
    }

    public function basicInfo(): Services\\BasicInfoService { return new Services\\BasicInfoService($this); }
    public function merchantAssets(): Services\\MerchantAssetsService { return new Services\\MerchantAssetsService($this); }
    public function merchantDeposit(): Services\\MerchantDepositService { return new Services\\MerchantDepositService($this); }
    public function merchantWithdraw(): Services\\MerchantWithdrawService { return new Services\\MerchantWithdrawService($this); }
    public function merchantBatchWithdraw(): Services\\MerchantBatchWithdrawService { return new Services\\MerchantBatchWithdrawService($this); }
    public function userAssets(): Services\\UserAssetsService { return new Services\\UserAssetsService($this); }
    public function userDeposit(): Services\\UserDepositService { return new Services\\UserDepositService($this); }
    public function userWithdraw(): Services\\UserWithdrawService { return new Services\\UserWithdrawService($this); }
    public function userTransfer(): Services\\UserTransferService { return new Services\\UserTransferService($this); }
    public function orders(): Services\\OrdersService { return new Services\\OrdersService($this); }
    public function swap(): Services\\SwapService { return new Services\\SwapService($this); }
    public function userSwap(): Services\\UserSwapService { return new Services\\UserSwapService($this); }
    public function utilities(): Services\\UtilitiesService { return new Services\\UtilitiesService($this); }
}
''')
    svc = base / "Services"
    write(svc / "BasicInfoService.php", '''<?php

namespace CCPayment\\Services;

use CCPayment\\Client;
use CCPayment\\APIError;
use GuzzleHttp\\Exception\\GuzzleException;

class BasicInfoService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getCoinList(): array { return $this->client->post('/getCoinList', []); }
    public function getFiatList(): array { return $this->client->post('/getFiatList', []); }
    public function getCoin(int $coinId): array { return $this->client->post('/getCoin', ['coinId' => $coinId]); }
    public function getCoinUSDTPrice(array $coinIds): array { return $this->client->post('/getCoinUSDTPrice', ['coinIds' => $coinIds]); }
    public function getChainList(array $data = []): array { return $this->client->post('/getChainList', $data); }
    public function getAllChainList(array $data = []): array { return $this->client->post('/all/chain', $data); }
    public function getMainCoinList(string $appId): array { return $this->client->post('/getMainCoinList', ['appId' => $appId]); }
}
''')
    write(svc / "MerchantAssetsService.php", '''<?php

namespace CCPayment\\Services;

use CCPayment\\Client;

class MerchantAssetsService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getAppCoinAssetList(): array { return $this->client->post('/getAppCoinAssetList', []); }
    public function getAppCoinAsset(int $coinId): array { return $this->client->post('/getAppCoinAsset', ['coinId' => $coinId]); }
    public function getAppCollectFeeRecordList(array $data = []): array { return $this->client->post('/getAggregationFeeRecord', $data); }
}
''')
    write(svc / "MerchantWithdrawService.php", '''<?php

namespace CCPayment\\Services;

use CCPayment\\Client;

class MerchantWithdrawService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getCwalletUserId(string $cwalletUserId): array { return $this->client->post('/getCwalletUserId', ['cwalletUserId' => $cwalletUserId]); }
    public function getWithdrawFee(int $coinId, string $chain): array { return $this->client->post('/getWithdrawFee', ['coinId' => $coinId, 'chain' => $chain]); }
    public function applyAppWithdrawToNetwork(array $data): array { return $this->client->post('/applyAppWithdrawToNetwork', $data); }
    public function applyAppWithdrawToCwallet(array $data): array { return $this->client->post('/applyAppWithdrawToCwallet', $data); }
    public function getAppWithdrawRecord(array $data): array { return $this->client->post('/getAppWithdrawRecord', $data); }
    public function getAppWithdrawRecordList(array $data = []): array { return $this->client->post('/getAppWithdrawRecordList', $data); }
    public function getAutoWithdrawRecordList(array $data = []): array { return $this->client->post('/getAutoWithdrawRecordList', $data); }
    public function getRiskRefundRecords(array $data = []): array { return $this->client->post('/getRiskyRefundRecordList', $data); }
}
''')
    write(svc / "OrdersService.php", '''<?php

namespace CCPayment\\Services;

use CCPayment\\Client;

class OrdersService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getAppOrderInfo(string $orderId): array { return $this->client->post('/getAppOrderInfo', ['orderId' => $orderId]); }
    public function createInvoiceUrl(array $data): array { return $this->client->post('/createInvoiceUrl', $data); }
    public function getInvoiceOrderInfo(string $orderId): array { return $this->client->post('/getInvoiceOrderInfo', ['orderId' => $orderId]); }
}
''')
    write(svc / "SwapService.php", '''<?php

namespace CCPayment\\Services;

use CCPayment\\Client;

class SwapService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getSwapCoinList(): array { return $this->client->post('/getSwapCoinList', []); }
    public function swapEstimate(array $data): array { return $this->client->post('/estimate', $data); }
    public function getSwapRecord(array $data): array { return $this->client->post('/getSwapRecord', $data); }
    public function getSwapRecordList(array $data = []): array { return $this->client->post('/getSwapRecordList', $data); }
    public function swap(array $data): array { return $this->client->post('/swap', $data); }
}
''')
    write(svc / "UserSwapService.php", '''<?php

namespace CCPayment\\Services;

use CCPayment\\Client;

class UserSwapService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function getUserSwapRecord(array $data): array { return $this->client->post('/getUserSwapRecord', $data); }
    public function getUserSwapRecordList(array $data = []): array { return $this->client->post('/getUserSwapRecordList', $data); }
    public function userSwap(array $data): array { return $this->client->post('/userSwap', $data); }
}
''')
    write(svc / "UtilitiesService.php", '''<?php

namespace CCPayment\\Services;

use CCPayment\\Client;

class UtilitiesService
{
    private Client $client;
    public function __construct(Client $client) { $this->client = $client; }
    public function webhookResend(array $data = []): array { return $this->client->post('/webhook/resend', $data); }
    public function getTradeBlockHeight(array $data): array { return $this->client->post('/getTradeBlockHeight', $data); }
    public function checkWithdrawalAddressValidity(array $data): array { return $this->client->post('/checkWithdrawalAddressValidity', $data); }
    public function deprecatedAddress(array $data): array { return $this->client->post('/addressUnbinding', $data); }
    public function rescanLostTransaction(array $data): array { return $this->client->post('/rescanLostTransaction', $data); }
}
''')
    remove(svc / "CheckoutService.php")
    patch_readme(ROOT / "php/README.md")


def regenerate_ruby() -> None:
    base = ROOT / "ruby/lib/ccpayment"
    write(base / "client.rb", '''require 'net/http'
require 'json'
require 'openssl'

module CCPayment
  class Client
    DEFAULT_BASE_URL = 'https://ccpayment.com/ccpayment/v2'

    attr_accessor :base_url

    def initialize(app_id, app_secret)
      @app_id = app_id
      @app_secret = app_secret
      @base_url = DEFAULT_BASE_URL
      @proxy_url = nil
    end

    def set_proxy(proxy_url)
      @proxy_url = proxy_url
    end

    def generate_sign(body)
      timestamp = Time.now.to_i.to_s
      sign_text = "#{@app_id}#{timestamp}#{body}"
      sign = OpenSSL::HMAC.hexdigest(OpenSSL::Digest.new('sha256'), @app_secret, sign_text)
      [sign, timestamp]
    end

    def post(path, data = {})
      body = data.to_json
      sign, timestamp = generate_sign(body)
      headers = {
        'Content-Type' => 'application/json',
        'Appid' => @app_id,
        'Sign' => sign,
        'Timestamp' => timestamp
      }
      uri = URI.parse("#{@base_url}#{path}")
      response = if @proxy_url
        proxy_uri = URI.parse(@proxy_url)
        http = Net::HTTP.new(uri.host, uri.port, proxy_uri.host, proxy_uri.port)
        http.use_ssl = uri.scheme == 'https'
        http.verify_mode = OpenSSL::SSL::VERIFY_PEER
        request = Net::HTTP::Post.new(uri.path, headers)
        request.body = body
        http.request(request)
      else
        http = Net::HTTP.new(uri.host, uri.port)
        http.use_ssl = uri.scheme == 'https'
        http.verify_mode = OpenSSL::SSL::VERIFY_PEER
        request = Net::HTTP::Post.new(uri.path, headers)
        request.body = body
        http.request(request)
      end
      result = JSON.parse(response.body)
      raise APIError.new(result['code'], result['msg']) unless result['code'] == 10000
      result['data'] || {}
    end

    def basic_info; Services::BasicInfoService.new(self); end
    def merchant_assets; Services::MerchantAssetsService.new(self); end
    def merchant_deposit; Services::MerchantDepositService.new(self); end
    def merchant_withdraw; Services::MerchantWithdrawService.new(self); end
    def merchant_batch_withdraw; Services::MerchantBatchWithdrawService.new(self); end
    def user_assets; Services::UserAssetsService.new(self); end
    def user_deposit; Services::UserDepositService.new(self); end
    def user_withdraw; Services::UserWithdrawService.new(self); end
    def user_transfer; Services::UserTransferService.new(self); end
    def orders; Services::OrdersService.new(self); end
    def swap; Services::SwapService.new(self); end
    def user_swap; Services::UserSwapService.new(self); end
    def utilities; Services::UtilitiesService.new(self); end
  end
end
''')
    write(ROOT / "ruby/lib/ccpayment.rb", '''require_relative 'ccpayment/client'
require_relative 'ccpayment/api_error'
require_relative 'ccpayment/services/basic_info_service'
require_relative 'ccpayment/services/merchant_assets_service'
require_relative 'ccpayment/services/merchant_deposit_service'
require_relative 'ccpayment/services/merchant_withdraw_service'
require_relative 'ccpayment/services/merchant_batch_withdraw_service'
require_relative 'ccpayment/services/user_assets_service'
require_relative 'ccpayment/services/user_deposit_service'
require_relative 'ccpayment/services/user_withdraw_service'
require_relative 'ccpayment/services/user_transfer_service'
require_relative 'ccpayment/services/orders_service'
require_relative 'ccpayment/services/swap_service'
require_relative 'ccpayment/services/user_swap_service'
require_relative 'ccpayment/services/utilities_service'

module CCPayment
  VERSION = '1.0.0'
end
''')
    svc = base / "services"
    write(svc / "basic_info_service.rb", '''module CCPayment
  module Services
    class BasicInfoService
      def initialize(client); @client = client; end
      def get_coin_list; @client.post('/getCoinList', {}); end
      def get_fiat_list; @client.post('/getFiatList', {}); end
      def get_coin(coin_id); @client.post('/getCoin', { coinId: coin_id }); end
      def get_coin_usdt_price(coin_ids); @client.post('/getCoinUSDTPrice', { coinIds: coin_ids }); end
      def get_chain_list(data = {}); @client.post('/getChainList', data); end
      def get_all_chain_list(data = {}); @client.post('/all/chain', data); end
      def get_main_coin_list(app_id); @client.post('/getMainCoinList', { appId: app_id }); end
    end
  end
end
''')
    write(svc / "merchant_assets_service.rb", '''module CCPayment
  module Services
    class MerchantAssetsService
      def initialize(client); @client = client; end
      def get_app_coin_asset_list; @client.post('/getAppCoinAssetList', {}); end
      def get_app_coin_asset(coin_id); @client.post('/getAppCoinAsset', { coinId: coin_id }); end
      def get_app_collect_fee_record_list(data = {}); @client.post('/getAggregationFeeRecord', data); end
    end
  end
end
''')
    write(svc / "merchant_withdraw_service.rb", '''module CCPayment
  module Services
    class MerchantWithdrawService
      def initialize(client); @client = client; end
      def get_cwallet_user_id(cwallet_user_id); @client.post('/getCwalletUserId', { cwalletUserId: cwallet_user_id }); end
      def get_withdraw_fee(coin_id, chain); @client.post('/getWithdrawFee', { coinId: coin_id, chain: chain }); end
      def apply_app_withdraw_to_network(data); @client.post('/applyAppWithdrawToNetwork', data); end
      def apply_app_withdraw_to_cwallet(data); @client.post('/applyAppWithdrawToCwallet', data); end
      def get_app_withdraw_record(data); @client.post('/getAppWithdrawRecord', data); end
      def get_app_withdraw_record_list(data = {}); @client.post('/getAppWithdrawRecordList', data); end
      def get_auto_withdraw_record_list(data = {}); @client.post('/getAutoWithdrawRecordList', data); end
      def get_risk_refund_records(data = {}); @client.post('/getRiskyRefundRecordList', data); end
    end
  end
end
''')
    write(svc / "orders_service.rb", '''module CCPayment
  module Services
    class OrdersService
      def initialize(client); @client = client; end
      def get_app_order_info(order_id); @client.post('/getAppOrderInfo', { orderId: order_id }); end
      def create_invoice_url(data); @client.post('/createInvoiceUrl', data); end
      def get_invoice_order_info(order_id); @client.post('/getInvoiceOrderInfo', { orderId: order_id }); end
    end
  end
end
''')
    write(svc / "swap_service.rb", '''module CCPayment
  module Services
    class SwapService
      def initialize(client); @client = client; end
      def get_swap_coin_list; @client.post('/getSwapCoinList', {}); end
      def swap_estimate(data); @client.post('/estimate', data); end
      def get_swap_record(data); @client.post('/getSwapRecord', data); end
      def get_swap_record_list(data = {}); @client.post('/getSwapRecordList', data); end
      def swap(data); @client.post('/swap', data); end
    end
  end
end
''')
    write(svc / "user_swap_service.rb", '''module CCPayment
  module Services
    class UserSwapService
      def initialize(client); @client = client; end
      def get_user_swap_record(data); @client.post('/getUserSwapRecord', data); end
      def get_user_swap_record_list(data = {}); @client.post('/getUserSwapRecordList', data); end
      def user_swap(data); @client.post('/userSwap', data); end
    end
  end
end
''')
    write(svc / "utilities_service.rb", '''module CCPayment
  module Services
    class UtilitiesService
      def initialize(client); @client = client; end
      def webhook_resend(data = {}); @client.post('/webhook/resend', data); end
      def get_trade_block_height(data); @client.post('/getTradeBlockHeight', data); end
      def check_withdrawal_address_validity(data); @client.post('/checkWithdrawalAddressValidity', data); end
      def deprecated_address(data); @client.post('/addressUnbinding', data); end
      def rescan_lost_transaction(data); @client.post('/rescanLostTransaction', data); end
    end
  end
end
''')
    remove(svc / "checkout_service.rb")
    patch_readme(ROOT / "ruby/README.md")


def main() -> None:
    regenerate_python()
    regenerate_typescript()
    regenerate_javascript()
    regenerate_java()
    regenerate_php()
    regenerate_ruby()


if __name__ == "__main__":
    main()
