"""Merchant Withdraw Service"""

from typing import Optional, List


class MerchantWithdrawService:
    def __init__(self, client):
        self.client = client

    def apply_app_withdraw_to_network(
        self,
        order_id: str,
        coin_id: int,
        chain: str,
        address: str,
        amount: str,
        memo: Optional[str] = None,
        merchant_pay_network_fee: Optional[bool] = None,
        network_fee_inquiry_id: Optional[str] = None,
        notify_url: Optional[str] = None
    ) -> str:
        """Apply for withdrawal to a blockchain network address"""
        data = {
            "orderId": order_id,
            "coinId": coin_id,
            "chain": chain,
            "address": address,
            "amount": amount
        }
        if memo: data["memo"] = memo
        if merchant_pay_network_fee: data["merchantPayNetworkFee"] = merchant_pay_network_fee
        if network_fee_inquiry_id: data["networkFeeInquiryID"] = network_fee_inquiry_id
        if notify_url: data["notifyUrl"] = notify_url
        result = self.client._post("/applyAppWithdrawToNetwork", data)
        return result.get("recordId")

    def apply_app_withdraw_to_cwallet(
        self,
        order_id: str,
        coin_id: int,
        cwallet_user: str,
        amount: str
    ) -> str:
        """Apply for withdrawal to a CCWallet account"""
        data = {
            "orderId": order_id,
            "coinId": coin_id,
            "cwalletUser": cwallet_user,
            "amount": amount
        }
        result = self.client._post("/applyAppWithdrawToCwallet", data)
        return result.get("recordId")

    def get_app_withdraw_record(self, order_id: Optional[str] = None, record_id: Optional[str] = None) -> dict:
        """Query details of a single withdrawal record"""
        data = {}
        if order_id: data["orderId"] = order_id
        if record_id: data["recordId"] = record_id
        result = self.client._post("/getAppWithdrawRecord", data)
        return result.get("record", {})

    def get_app_withdraw_record_list(
        self,
        chain: Optional[str] = None,
        coin_id: Optional[int] = None,
        order_ids: Optional[List[str]] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        to_address: Optional[str] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query withdrawal record list"""
        data = {}
        if chain: data["chain"] = chain
        if coin_id: data["coinId"] = coin_id
        if order_ids: data["orderIds"] = order_ids
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if to_address: data["toAddress"] = to_address
        if next_id: data["nextId"] = next_id
        return self.client._post("/getAppWithdrawRecordList", data)

    def get_auto_withdraw_record_list(
        self,
        chain: Optional[str] = None,
        coin_id: Optional[int] = None,
        record_ids: Optional[List[str]] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        to_address: Optional[str] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query automatic withdrawal record list"""
        data = {}
        if chain: data["chain"] = chain
        if coin_id: data["coinId"] = coin_id
        if record_ids: data["recordIds"] = record_ids
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if to_address: data["toAddress"] = to_address
        if next_id: data["nextId"] = next_id
        return self.client._post("/getAutoWithdrawRecordList", data)

    def get_risky_refund_record_list(
        self,
        chain: Optional[str] = None,
        coin_id: Optional[int] = None,
        record_ids: Optional[List[str]] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        to_address: Optional[str] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query risky refund record list"""
        data = {}
        if chain: data["chain"] = chain
        if coin_id: data["coinId"] = coin_id
        if record_ids: data["recordIds"] = record_ids
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if to_address: data["toAddress"] = to_address
        if next_id: data["nextId"] = next_id
        return self.client._post("/getRiskyRefundRecordList", data)
