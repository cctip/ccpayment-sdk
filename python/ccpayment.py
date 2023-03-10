import hashlib
import time
import urllib.request
import json
import const
import contextlib


class CCPaymentClass:

    def __init__(self, app_id, app_secret):
        self.app_id = app_id
        self.app_secret = app_secret

    """
    * return success
    {
        "code": 10000,
        "msg": "",
        "data": {
            "order_id": "",
            "amount": "",
            "logo": "",
            "network": "",
            "pay_address": "",
            "crypto": ""
        }
    }
    """
    # create api order
    def create_order(self, token_id, amount, merchant_order_id, fiat_currency, remark=None):
        data = {
            "amount": amount,
            "merchant_order_id": merchant_order_id,
            "token_id": token_id,
            "fiat_currency": fiat_currency
        }
        if remark:
            data["remark"] = remark
        return self._send_post(const.CREATE_ORDER_URL, data)

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
    # get checkout url
    def checkout_url(self, amount, merchant_order_id, valid_timestamp, product_name, return_url=None):
        data = {
            "amount": amount,
            "merchant_order_id": merchant_order_id,
            "valid_timestamp": valid_timestamp,
            "product_name": product_name
        }
        if return_url:
            data["return_url"] = return_url
        return self._send_post(const.CHECKOUT_URL, data)

    def webhook(self, data_str, timestamp, signature):
        if self._hash256(data_str, timestamp) == signature and signature != "":
            return True

        return False

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
    def get_support_token(self):
        data = {}
        return self._send_post(const.SUPPORT_TOKEN_URL, data)

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
    def get_token_chain(self, token_id):
        data = {
            "token_id": token_id
        }
        return self._send_post(const.TOKEN_CHAIN_URL, data)

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
    def get_token_rate(self, token_id, amount):
        data = {
            "token_id": token_id,
            "amount": amount
        }
        return self._send_post(const.TOKEN_RATE_URL, data)

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
    def withdraw(self, token_id, address, value, merchant_order_id, memo=None):
        data = {
            "token_id": token_id,
            "address": address,
            "value": value,
            "merchant_order_id": merchant_order_id
        }
        if memo:
            data["memo"] = memo
        return self._send_post(const.WITHDRAW_API_URL, data)

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
    def check_user(self, c_id):
        data = {
            "c_id": c_id
        }
        return self._send_post(const.CHECK_USER_URL, data)

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
    def assets(self, token_id):
        data = {
            "token_id": token_id
        }
        return self._send_post(const.ASSETS_URL, data)

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
    def network_fee(self, token_id, address=None, memo=None):
        data = {
            "token_id": token_id
        }
        if memo:
            data["memo"] = memo
        if address:
            data["address"] = address
        return self._send_post(const.NETWORK_FEE_URL, data)

    def _hash256(self, txt, timestamp):
        txt = self.app_id + self.app_secret + str(timestamp) + txt
        return hashlib.sha256(txt.encode("utf-8")).hexdigest()

    def _send_post(self, url, data):
        timestamp = int(time.time())

        data_str = json.dumps(data)
        sign_str = self._hash256(data_str, timestamp)

        req = urllib.request.Request(url=url, method="POST", data=data_str.encode("utf-8"), headers={
            "Content-Type": "application/json;charset=uf8",
            const.APP_ID_HEADER_KEY: self.app_id,
            const.SIGN_HEADER_KEY: sign_str,
            const.TIMESTAMP_HEADER_KEY: str(timestamp)
        })

        req.timeout = 60  # s

        # post
        resp = urllib.request.urlopen(req)

        data, is_valid = self._deal_and_valid(resp)
        if resp:
            resp.close()

        return data, is_valid

    def _deal_and_valid(self, resp):
        if resp.code != 200:
            return {}, False

        data_str = resp.read().decode('utf-8')
        data = json.loads(data_str)

        if data['code'] != 10000:
            return data, True

        # header
        signature = resp.headers[const.SIGN_HEADER_KEY]
        ts = resp.headers[const.TIMESTAMP_HEADER_KEY]

        # verify
        if self._hash256(data_str, ts) == signature and signature != "":
            return data, True
        return data, False
