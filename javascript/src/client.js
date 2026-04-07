const axios = require('axios');
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
    const sign = crypto
      .createHmac('sha256', this.appSecret)
      .update(signText)
      .digest('hex');
    return { sign, timestamp };
  }

  async post(path, data = {}) {
    const body = JSON.stringify(data);
    const { sign, timestamp } = this.generateSign(body);

    const headers = {
      'Content-Type': 'application/json',
      'Appid': this.appId,
      'Sign': sign,
      'Timestamp': timestamp,
    };

    const response = await this.axiosInstance.post(
      `${this.baseUrl}${path}`,
      data,
      { headers }
    );

    const result = response.data;
    if (result.code !== 10000) {
      throw new APIError(result.code, result.msg);
    }

    return result.data;
  }

  basicInfo() {
    return new (require('./services/basic-info'))(this);
  }

  merchantAssets() {
    return new (require('./services/merchant-assets'))(this);
  }

  merchantDeposit() {
    return new (require('./services/merchant-deposit'))(this);
  }

  merchantWithdraw() {
    return new (require('./services/merchant-withdraw'))(this);
  }

  merchantBatchWithdraw() {
    return new (require('./services/merchant-batch-withdraw'))(this);
  }

  userAssets() {
    return new (require('./services/user-assets'))(this);
  }

  userDeposit() {
    return new (require('./services/user-deposit'))(this);
  }

  userWithdraw() {
    return new (require('./services/user-withdraw'))(this);
  }

  userTransfer() {
    return new (require('./services/user-transfer'))(this);
  }

  orders() {
    return new (require('./services/orders'))(this);
  }

  checkout() {
    return new (require('./services/checkout'))(this);
  }

  swap() {
    return new (require('./services/swap'))(this);
  }

  utilities() {
    return new (require('./services/utilities'))(this);
  }
}

module.exports = { Client };
