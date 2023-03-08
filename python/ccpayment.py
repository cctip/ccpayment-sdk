import hashlib
import time
import urllib.request
import json
import const


# crate order
class OrderClass:
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
    def create_deposit_order(self, app_id, app_secret):

        data = {
            "amount": self.amount,
            "merchant_order_id": self.merchant_order_id,
            "token_id": self.token_id,
            "fiat_currency": self.fiat_currency,
            "remark": self.remark
        }

        return _send_post(const.CREATE_ORDER_URL, data, app_id, app_secret)


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

        data = {
            "amount": self.amount,
            "merchant_order_id": self.merchant_order_id,
            "valid_timestamp": self.valid_timestamp,
            "product_name": self.product_name,
            "return_url": self.return_url
        }

        return _send_post(const.CHECKOUT_URL, data, app_id, app_secret)


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

    data = {}

    return _send_post(const.SUPPORT_TOKEN_URL, data, app_id, app_secret)


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
                    "token_id": "",
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

        data = {
            "token_id": self.token_id
        }

        return _send_post(const.TOKEN_CHAIN_URL, data, app_id, app_secret)


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

        data = {
            "token_id": self.token_id,
            "amount": self.amount
        }

        return _send_post(const.TOKEN_RATE_URL, data, app_id, app_secret)


# create api withdrawal
class ApiWithdrawClass:
    token_id = ''
    address = ''
    memo = ''
    value = ''
    merchant_order_id = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "bill_id": "",
            "type": "",
            "network_fee": ""
        }
    }
    """
    def create_withdraw_order(self, app_id, app_secret):

        data = {
            "token_id": self.token_id,
            "address": self.address,
            "memo": self.memo,
            "value": self.value,
            "merchant_order_id": self.merchant_order_id
        }

        return _send_post(const.WITHDRAW_API_URL, data, app_id, app_secret)


# check user
class CheckUserClass:
    c_id = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "c_id": "",
            "nickname": ""
        }
    }
    """
    def check_user(self, app_id, app_secret):

        data = {
            "c_id": self.c_id
        }

        return _send_post(const.CHECK_USER_URL, data, app_id, app_secret)


# get token assets
class TokenAssetClass:
    token_id = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": [
            {
                "token_id": "",
                "crypto": "",
                "name": "",
                "value": "",
                "price": "",
                "logo": ""
            }
        ]
    }
    """
    def get_token_assets(self, app_id, app_secret):

        data = {
            "token_id": self.token_id
        }

        return _send_post(const.ASSETS_URL, data, app_id, app_secret)


# network fee
class NetworkFeeClass:
    token_id = ''
    address = ''
    memo = ''

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "token_id": "",
            "crypto": "",
            "fee": ""
        }
    }
    """
    def get_network_fee(self, app_id, app_secret):

        data = {
            "token_id": self.token_id,
            "address": self.address,
            "memo": self.memo
        }

        return _send_post(const.ASSETS_URL, data, app_id, app_secret)


def _hash256(txt, app_id, app_secret, timestamp):
    txt = app_id + app_secret + str(timestamp) + txt

    return hashlib.sha256(txt.encode("utf-8")).hexdigest()


def _send_post(url, data, app_id, app_secret):
    timestamp = int(time.time())

    data_str = json.dumps(data)
    sign_str = _hash256(data_str, app_id, app_secret, timestamp)

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
