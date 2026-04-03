"""User Deposit Service"""

from typing import Optional


class UserDepositService:
    def __init__(self, client):
        self.client = client

    def get_or_create_user_deposit_address(
        self,
        user_id: str,
        chain: str,
        notify_url: Optional[str] = None
    ) -> dict:
        """Get or create a user's deposit address"""
        data = {"userId": user_id, "chain": chain}
        if notify_url: data["notifyUrl"] = notify_url
        return self.client._post("/getOrCreateUserDepositAddress", data)

    def get_user_deposit_record(self, record_id: str) -> dict:
        """Query details of a single user deposit record"""
        result = self.client._post("/getUserDepositRecord", {"recordId": record_id})
        return result.get("record", {})

    def get_user_deposit_record_list(
        self,
        user_id: str,
        chain: Optional[str] = None,
        coin_id: Optional[int] = None,
        to_address: Optional[str] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query user deposit record list"""
        data = {"userId": user_id}
        if chain: data["chain"] = chain
        if coin_id: data["coinId"] = coin_id
        if to_address: data["toAddress"] = to_address
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if next_id: data["nextId"] = next_id
        return self.client._post("/getUserDepositRecordList", data)
