"""CCPayment SDK Client"""

import hashlib
import hmac
import time
import json
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
        """Set custom base URL"""
        self.base_url = base_url

    def set_proxy(self, proxy_url: str):
        """Set HTTP proxy"""
        self.session.proxies = {
            "http": proxy_url,
            "https": proxy_url
        }

    def _generate_sign(self, body: str) -> tuple[str, str]:
        """Generate HMAC-SHA256 signature"""
        timestamp = str(int(time.time()))
        sign_text = self.app_id + timestamp + body
        sign = hmac.new(
            self.app_secret.encode(),
            sign_text.encode(),
            hashlib.sha256
        ).hexdigest()
        return sign, timestamp

    def _post(self, path: str, data: Optional[dict] = None) -> dict:
        """Make POST request to API"""
        body = json.dumps(data) if data else "{}"
        sign, timestamp = self._generate_sign(body)

        headers = {
            "Content-Type": "application/json",
            "Appid": self.app_id,
            "Sign": sign,
            "Timestamp": timestamp
        }

        url = f"{self.base_url}{path}"
        response = self.session.post(url, headers=headers, data=body)
        response.raise_for_status()

        result = response.json()
        if result.get("code") != 10000:
            raise APIError(result.get("code"), result.get("msg"))

        return result.get("data", {})

    @property
    def basic_info(self):
        from .services.basic_info import BasicInfoService
        if not hasattr(self, '_basic_info'):
            self._basic_info = BasicInfoService(self)
        return self._basic_info

    @property
    def merchant_assets(self):
        from .services.merchant_assets import MerchantAssetsService
        if not hasattr(self, '_merchant_assets'):
            self._merchant_assets = MerchantAssetsService(self)
        return self._merchant_assets

    @property
    def merchant_deposit(self):
        from .services.merchant_deposit import MerchantDepositService
        if not hasattr(self, '_merchant_deposit'):
            self._merchant_deposit = MerchantDepositService(self)
        return self._merchant_deposit

    @property
    def merchant_withdraw(self):
        from .services.merchant_withdraw import MerchantWithdrawService
        if not hasattr(self, '_merchant_withdraw'):
            self._merchant_withdraw = MerchantWithdrawService(self)
        return self._merchant_withdraw

    @property
    def merchant_batch_withdraw(self):
        from .services.merchant_batch_withdraw import MerchantBatchWithdrawService
        if not hasattr(self, '_merchant_batch_withdraw'):
            self._merchant_batch_withdraw = MerchantBatchWithdrawService(self)
        return self._merchant_batch_withdraw

    @property
    def user_assets(self):
        from .services.user_assets import UserAssetsService
        if not hasattr(self, '_user_assets'):
            self._user_assets = UserAssetsService(self)
        return self._user_assets

    @property
    def user_deposit(self):
        from .services.user_deposit import UserDepositService
        if not hasattr(self, '_user_deposit'):
            self._user_deposit = UserDepositService(self)
        return self._user_deposit

    @property
    def user_withdraw(self):
        from .services.user_withdraw import UserWithdrawService
        if not hasattr(self, '_user_withdraw'):
            self._user_withdraw = UserWithdrawService(self)
        return self._user_withdraw

    @property
    def user_transfer(self):
        from .services.user_transfer import UserTransferService
        if not hasattr(self, '_user_transfer'):
            self._user_transfer = UserTransferService(self)
        return self._user_transfer

    @property
    def orders(self):
        from .services.orders import OrdersService
        if not hasattr(self, '_orders'):
            self._orders = OrdersService(self)
        return self._orders

    @property
    def checkout(self):
        from .services.checkout import CheckoutService
        if not hasattr(self, '_checkout'):
            self._checkout = CheckoutService(self)
        return self._checkout

    @property
    def swap(self):
        from .services.swap import SwapService
        if not hasattr(self, '_swap'):
            self._swap = SwapService(self)
        return self._swap

    @property
    def utilities(self):
        from .services.utilities import UtilitiesService
        if not hasattr(self, '_utilities'):
            self._utilities = UtilitiesService(self)
        return self._utilities
