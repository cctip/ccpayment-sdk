"""User Withdraw Service"""

from typing import Optional


class UserWithdrawService:
    def __init__(self, client):
        self.client = client

    def apply_user_withdraw_to_network(
        self,
        user_id: str,
        order_id: str,
        coin_id: int,
        chain: str,
        address: str,
        amount: str,
        memo: Optional[str] = None,
        network_fee_inquiry_id: Optional[str] = None,
        notify_url: Optional[str] = None
    ) -> str:
        """User applies for withdrawal to a blockchain network address"""
        data = {
            "userId": user_id,
            "orderId": order_id,
            "coinId": coin_id,
            "chain": chain,
            "address": address,
            "amount": amount
        }
        if memo: data["memo"] = memo
        if network_fee_inquiry_id: data["networkFeeInquiryID"] = network_fee_inquiry_id
        if notify_url: data["notifyUrl"] = notify_url
        result = self.client._post("/applyUserWithdrawToNetwork", data)
        return result.get("recordId")

    def apply_user_withdraw_to_cwallet(
        self,
        user_id: str,
        order_id: str,
        coin_id: int,
        cwallet_user: str,
        amount: str
    ) -> str:
        """User applies for withdrawal to a CCWallet account"""
        data = {
            "userId": user_id,
            "orderId": order_id,
            "coinId": coin_id,
            "cwalletUser": cwallet_user,
            "amount": amount
        }
        result = self.client._post("/applyUserWithdrawToCwallet", data)
        return result.get("recordId")

    def get_user_withdraw_record(
        self,
        record_id: Optional[str] = None,
        order_id: Optional[str] = None
    ) -> dict:
        """Query details of a single user withdrawal record"""
        data = {}
        if record_id: data["recordId"] = record_id
        if order_id: data["orderId"] = order_id
        result = self.client._post("/getUserWithdrawRecord", data)
        return result.get("record", {})

    def get_user_withdraw_record_list(
        self,
        user_id: str,
        order_ids: Optional[list] = None,
        chain: Optional[str] = None,
        coin_id: Optional[int] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        to_address: Optional[str] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query user withdrawal record list"""
        data = {"userId": user_id}
        if order_ids: data["orderIds"] = order_ids
        if chain: data["chain"] = chain
        if coin_id: data["coinId"] = coin_id
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if to_address: data["toAddress"] = to_address
        if next_id: data["nextId"] = next_id
        return self.client._post("/getUserWithdrawRecordList", data)
