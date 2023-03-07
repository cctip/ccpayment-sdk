const axios = require('axios');
const crypto = require('crypto');

const requestAPI = {
  checkoutURL: 'https://admin.ccpayment.com/ccpayment/v1/concise/url/get',
  selectTokenURL: 'https://admin.ccpayment.com/ccpayment/v1/support/token',
  selectChainURL: 'https://admin.ccpayment.com/ccpayment/v1/token/chain',
  submitOrderURL: 'https://admin.ccpayment.com/ccpayment/v1/bill/create',
  tokenRateURL: 'https://admin.ccpayment.com/ccpayment/v1/token/rate'
}


module.exports = {
  appId: null,
  appSecret: null,

  axios: axios,

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

  async selectToken(callback) {

    const { compareSignture, sign, result } = await this.sendPost(requestAPI.selectTokenURL, null)
    if (result) {
      callback && callback(compareSignture === sign ? result.data : Error('http code error'))
    }
  },


  async selectChain(data, callback) {

    const { compareSignture, sign, result } = await this.sendPost(requestAPI.selectChainURL, {
      ...data
    })
    if (result) {
      callback && callback(compareSignture === sign ? result.data : Error('http code error'))
    }
  },

  async submitOrder(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.submitOrderURL, {
      ...data
    })
    if (result) {
      callback && callback(compareSignture === sign ? result.data : Error('http code error'))
    }
  },

  async checkoutURL(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.checkoutURL, {
      ...data
    })
    if (result) {
      callback && callback(compareSignture === sign ? result.data : Error('http code error'))
    }
  },

  async tokenRate(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.tokenRateURL, {
      ...data
    })
    if (result) {
      callback && callback(compareSignture === sign ? result.data : Error('http code error'))
    }
  },

  webHookNotify(timeStamp, sign, data, callback) {
    const compareSignture = this.sha256(`${this.appId}${this.appSecret}${timeStamp}${data ? JSON.stringify({ ...data }) : ''}`)
    callback && callback(compareSignture === sign ? 'success!' : Error('http code error'))
  }
}
