"""User Transfer Service"""

from typing import Optional, List


class UserTransferService:
    def __init__(self, client):
        self.client = client

    def user_transfer(
        self,
        order_id: str,
        from_user_id: str,
        to_user_id: str,
        coin_id: int,
        amount: str,
        remark: Optional[str] = None
    ) -> str:
        """Initiate a transfer between users"""
        data = {
            "orderId": order_id,
            "fromUserId": from_user_id,
            "toUserId": to_user_id,
            "coinId": coin_id,
            "amount": amount
        }
        if remark: data["remark"] = remark
        result = self.client._post("/userTransfer", data)
        return result.get("recordId")

    def get_user_transfer_record(
        self,
        record_id: Optional[str] = None,
        order_id: Optional[str] = None
    ) -> dict:
        """Query details of a single user transfer record"""
        data = {}
        if record_id: data["recordId"] = record_id
        if order_id: data["orderId"] = order_id
        result = self.client._post("/getUserTransferRecord", data)
        return result.get("record", {})

    def get_user_transfer_record_list(
        self,
        order_ids: Optional[List[str]] = None,
        from_user_id: Optional[str] = None,
        to_user_id: Optional[str] = None,
        coin_id: Optional[int] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query user transfer record list"""
        data = {}
        if order_ids: data["orderIds"] = order_ids
        if from_user_id: data["fromUserId"] = from_user_id
        if to_user_id: data["toUserId"] = to_user_id
        if coin_id: data["coinId"] = coin_id
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if next_id: data["nextId"] = next_id
        return self.client._post("/getUserTransferRecordList", data)

    def user_batch_transfer(
        self,
        order_id: str,
        user_id: str,
        to_users: List[dict],
        coin_id: int,
        remark: Optional[str] = None
    ) -> str:
        """Initiate a batch transfer for users"""
        data = {
            "orderId": order_id,
            "userId": user_id,
            "toUsers": to_users,
            "coinId": coin_id
        }
        if remark: data["remark"] = remark
        result = self.client._post("/userBatchTransfer", data)
        return result.get("recordId")

    def get_user_batch_transfer_record(
        self,
        record_id: Optional[str] = None,
        order_id: Optional[str] = None
    ) -> dict:
        """Query details of a single user batch transfer record"""
        data = {}
        if record_id: data["recordId"] = record_id
        if order_id: data["orderId"] = order_id
        result = self.client._post("/getUserBatchTransferRecord", data)
        return result.get("record", {})

    def get_user_batch_transfer_record_list(
        self,
        order_ids: Optional[List[str]] = None,
        user_id: Optional[str] = None,
        coin_id: Optional[int] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query user batch transfer record list"""
        data = {}
        if order_ids: data["orderIds"] = order_ids
        if user_id: data["userId"] = user_id
        if coin_id: data["coinId"] = coin_id
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if next_id: data["nextId"] = next_id
        return self.client._post("/getUserBatchTransferRecordList", data)

    def get_user_batch_transfer_record_detail(
        self,
        record_id: str,
        next_id: Optional[str] = None
    ) -> dict:
        """Query the detail list of a user batch transfer record"""
        data = {"recordId": record_id}
        if next_id: data["nextId"] = next_id
        return self.client._post("/getUserBatchTransferRecordDetail", data)
