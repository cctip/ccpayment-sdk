"""Utilities Service"""


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
