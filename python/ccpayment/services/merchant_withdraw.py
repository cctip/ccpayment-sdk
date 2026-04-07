"""Merchant Withdraw Service"""


class MerchantWithdrawService:
    def __init__(self, client):
        self.client = client

    def get_cwallet_user_id(self, cwallet_user_id: str):
        return self.client._post("/getCwalletUserId", {"cwalletUserId": cwallet_user_id})

    def get_withdraw_fee(self, coin_id: int, chain: str):
        return self.client._post("/getWithdrawFee", {"coinId": coin_id, "chain": chain})

    def apply_app_withdraw_to_network(self, data):
        return self.client._post("/applyAppWithdrawToNetwork", data)

    def apply_app_withdraw_to_cwallet(self, data):
        return self.client._post("/applyAppWithdrawToCwallet", data)

    def get_app_withdraw_record(self, data):
        return self.client._post("/getAppWithdrawRecord", data)

    def get_app_withdraw_record_list(self, data=None):
        return self.client._post("/getAppWithdrawRecordList", data or {})

    def get_auto_withdraw_record_list(self, data=None):
        return self.client._post("/getAutoWithdrawRecordList", data or {})

    def get_risk_refund_records(self, data=None):
        return self.client._post("/getRiskyRefundRecordList", data or {})
