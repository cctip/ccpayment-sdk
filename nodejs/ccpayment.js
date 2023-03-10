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

  /*
   * @param {String} appId
   * @param {String} appSecret
   * @return {void}
   */
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

  /*
   * @param {Function} callback 
   * @return {void}
   */
  async getSupportToken(callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.supportTokenURL, null)
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  /*
   * @param {Object} data
   * @param {String} data.token_id
   * @param {Function} callback 
   * @return {void}
   */
  async getTokenChain(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.tokenChainURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  /*
   * @param {Object} data
   * @param {String} data.remark
   * @param {String} data.token_id
   * @param {String} data.amount
   * @param {String} data.merchant_order_id
   * @param {String} data.fiat_currency
   * @param {Function} callback 
   * @return {void}
   */
  async createOrder(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.createOrderURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  /*
   * @param {Object} data
   * @param {Number} data.valid_timestamp
   * @param {String} data.amount
   * @param {String} data.product_name
   * @param {String} data.return_url
   * @param {Function} callback 
   * @return {void}
   */
  async checkoutURL(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.checkoutURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  /*
   * @param {Object} data
   * @param {String} data.token_id
   * @param {String} data.amount
   * @param {Function} callback 
   * @return {void}
   */
  async getTokenRate(data, callback) {
    const { compareSignture, sign, result } = await this.sendPost(requestAPI.tokenRateURL, {
      ...data
    })
    if (result) {
      callback && callback(result.data.code === 10000 ? compareSignture === sign ? result.data : Error('http code error') : result.data)
    }
  },

  /*
   * @param {Number} timeStamp
   * @param {String} sign
   * @param {Object} data
   * @param {Function} callback 
   * @return {void}
   */
  webhook(timeStamp, sign, data, callback) {
    const compareSignture = this.sha256(`${this.appId}${this.appSecret}${timeStamp}${data ? JSON.stringify({ ...data }) : ''}`)
    callback && callback(compareSignture === sign)
  }
}


