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

    # Service accessors
    def basic_info(self):
        from .services.basic_info import BasicInfoService
        return BasicInfoService(self)

    def merchant_assets(self):
        from .services.merchant_assets import MerchantAssetsService
        return MerchantAssetsService(self)

    def merchant_deposit(self):
        from .services.merchant_deposit import MerchantDepositService
        return MerchantDepositService(self)

    def merchant_withdraw(self):
        from .services.merchant_withdraw import MerchantWithdrawService
        return MerchantWithdrawService(self)

    def merchant_batch_withdraw(self):
        from .services.merchant_batch_withdraw import MerchantBatchWithdrawService
        return MerchantBatchWithdrawService(self)

    def user_assets(self):
        from .services.user_assets import UserAssetsService
        return UserAssetsService(self)

    def user_deposit(self):
        from .services.user_deposit import UserDepositService
        return UserDepositService(self)

    def user_withdraw(self):
        from .services.user_withdraw import UserWithdrawService
        return UserWithdrawService(self)

    def user_transfer(self):
        from .services.user_transfer import UserTransferService
        return UserTransferService(self)

    def orders(self):
        from .services.orders import OrdersService
        return OrdersService(self)

    def checkout(self):
        from .services.checkout import CheckoutService
        return CheckoutService(self)

    def swap(self):
        from .services.swap import SwapService
        return SwapService(self)

    def utilities(self):
        from .services.utilities import UtilitiesService
        return UtilitiesService(self)
