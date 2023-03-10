const axios = require('axios');
const crypto = require('crypto');

const HOST = 'https://admin.ccpayment.com'

const requestAPI = {
  checkoutURL: `${HOST}/ccpayment/v1/concise/url/get`,
  supportTokenURL: `${HOST}/ccpayment/v1/support/token`,
  tokenChainURL: `${HOST}/ccpayment/v1/token/chain`,
  createOrderURL: `${HOST}/ccpayment/v1/bill/create`,
  tokenRateURL: `${HOST}/ccpayment/v1/token/rate`
}


module.exports = {

  appId: null,
  appSecret: null,

  createTimestamp(threshold = 0) {
    return parseInt(Date.now() / 1000, 10) + threshold
  },


  sha256(data) {
    let hash = crypto.createHash('sha256');
    return hash.update(data, 'utf-8').digest('hex');
  },


  init(appId, appSecret) {
    this.appId = appId
    this.appSecret = appSecret
  },

  async sendPost(url, data = null) {
    const timeStamp = this.createTimestamp()
    try {
      const result = await axios.post(url, data, {
        headers: {
          'Appid': this.appId,
          'Timestamp': timeStamp,
          'Sign': this.sha256(`${this.appId}${this.appSecret}${timeStamp}${data ? JSON.stringify({ ...data }) : ''}`)
        }
      })
      const { appid, timestamp, sign } = result.headers;
      const compareSignture = this.sha256(`${appid}${this.appSecret}${timestamp}${result.data ? JSON.stringify(result.data) : ''}`)
      return {
        result,
        sign,
        compareSignture
      }
    }
    catch (err) {
      throw Error(err)
    }

  },

  async getSupportToken(callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.supportTokenURL, null)
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },


  async getTokenChain(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.tokenChainURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  async createOrder(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.createOrderURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  async checkoutURL(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.checkoutURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  async getTokenRate(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.tokenRateURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  webhook(timeStamp, sign, data, callback) {
    const compareSignture = this.sha256(`${this.appId}${this.appSecret}${timeStamp}${data ? JSON.stringify({ ...data }) : ''}`)
    callback && callback(compareSignture === sign)
  }
}


