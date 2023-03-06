import hashlib
import time
import urllib.request
import json
import const


# crate order
class OrderClass:
    chain = ''
    contract = ''
    amount = ''
    merchant_order_id = ''
    token_id = ''
    fiat_currency = ''
    remark = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "bill_id": "",
            "amount": "",
            "logo": "",
            "network": "",
            "pay_address": "",
            "crypto": ""
        }
    }
    """
    def create_order(self, app_id, app_secret):
        timestamp = int(time.time())

        data = {
            "chain": self.chain,
            "contract": self.contract,
            "amount": self.amount,
            "merchant_order_id": self.merchant_order_id,
            "token_id": self.token_id,
            "fiat_currency": self.fiat_currency,
            "remark": self.remark
        }

        data_str = json.dumps(data)
        sign_str = _hash256(data_str, app_id, app_secret, timestamp)

        return _send_post(const.CREATE_ORDER_URL, data_str, app_id, app_secret, sign_str, timestamp)


# get checkout url
class PaymentUrlClass:
    amount = ''
    merchant_order_id = ''
    valid_timestamp = 0  # s
    product_name = ''
    return_url = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "payment_url": ""
        }
    }
    """
    def get_checkout_url(self, app_id, app_secret):
        timestamp = int(time.time())

        data = {
            "amount": self.amount,
            "merchant_order_id": self.merchant_order_id,
            "valid_timestamp": self.valid_timestamp,
            "product_name": self.product_name,
            "return_url": self.return_url
        }

        data_str = json.dumps(data)
        sign_str = _hash256(data_str, app_id, app_secret, timestamp)

        return _send_post(const.CHECKOUT_URL, data_str, app_id, app_secret, sign_str, timestamp)


# webhook validate
class WebhookClass:
    data_str = ''  # from header key: Appid
    signature = ''  # from header key: Sign
    timestamp = ''  # from header key: Timestamp

    def webhook_validate(self, app_id, app_secret):
        if _hash256(self.data_str, app_id, app_secret, self.timestamp) == self.signature and self.signature != "":
            return True

        return False


# get support token list
def get_support_tokens(app_id, app_secret):
    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "list": [
                {
                    "crypto": "",
                    "name": "",
                    "logo": "",
                    "min": "",
                    "price": "",
                    "token_id": ""
                }
            ]
        }
    }
    """

    timestamp = int(time.time())

    data = {}

    data_str = json.dumps(data)
    sign_str = _hash256(data_str, app_id, app_secret, timestamp)

    return _send_post(const.SUPPORT_TOKEN_URL, data_str, app_id, app_secret, sign_str, timestamp)


# get token chains
class TokenChainClass:
    token_id = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "list": [
                {
                    "id": "",
                    "chain": "",
                    "contract": "",
                    "crypto": "",
                    "logo": "",
                    "name": "",
                    "network": "",
                    "chain_logo": ""
                }
            ]
        }
    }
    """
    def get_token_chain(self, app_id, app_secret):
        timestamp = int(time.time())

        data = {
            "token_id": self.token_id
        }

        data_str = json.dumps(data)
        sign_str = _hash256(data_str, app_id, app_secret, timestamp)

        return _send_post(const.TOKEN_CHAIN_URL, data_str, app_id, app_secret, sign_str, timestamp)


# get token rate
class TokenRateClass:
    amount = ''
    token_id = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "price": "",
            "value": ""
        }
    }
    """
    def get_token_rate(self, app_id, app_secret):
        timestamp = int(time.time())

        data = {
            "token_id": self.token_id,
            "amount": self.amount
        }

        data_str = json.dumps(data)
        sign_str = _hash256(data_str, app_id, app_secret, timestamp)

        return _send_post(const.TOKEN_RATE_URL, data_str, app_id, app_secret, sign_str, timestamp)


def _hash256(txt, app_id, app_secret, timestamp):
    txt = app_id + app_secret + str(timestamp) + txt

    return hashlib.sha256(txt.encode("utf-8")).hexdigest()


def _send_post(url, data_str, app_id, app_secret, sign_str, timestamp):
    req = urllib.request.Request(url=url, method="POST", data=data_str.encode("utf-8"))

    req.add_header(const.APP_ID_HEADER_KEY, app_id)
    req.add_header(const.SIGN_HEADER_KEY, sign_str)
    req.add_header(const.TIMESTAMP_HEADER_KEY, str(timestamp))

    req.timeout = 30  # s

    # post
    resp = urllib.request.urlopen(req)
    if resp.code != 200:
        return {}, False

    data_str = resp.read().decode('utf-8')

    data = json.loads(data_str)

    if data['code'] != 10000:
        return data, False

    # header
    signature = resp.headers[const.SIGN_HEADER_KEY]
    ts = resp.headers[const.TIMESTAMP_HEADER_KEY]

    # verify
    if _hash256(data_str, app_id, app_secret, ts) == signature and signature != "":
        return data, True

    return data, False
