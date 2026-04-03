"""Merchant Batch Withdraw Service"""

from typing import List, Optional


class MerchantBatchWithdrawService:
    def __init__(self, client):
        self.client = client

    def check_withdraw_address(self, chain: str, address_info_list: List[dict]) -> List[dict]:
        """Check the validity of withdrawal addresses in batch"""
        result = self.client._post("/checkWithdrawAddress", {
            "chain": chain,
            "addressInfoList": address_info_list
        })
        return result.get("addressInfoResults", [])

    def apply_batch_withdraw(
        self,
        batch_order_id: str,
        coin_id: int,
        chain: str,
        mode: str,
        orders: Optional[List[dict]] = None,
        product_name: Optional[str] = None,
        notify_url: Optional[str] = None
    ) -> dict:
        """Create a batch withdrawal order"""
        data = {
            "batchOrderId": batch_order_id,
            "coinId": coin_id,
            "chain": chain,
            "mode": mode
        }
        if orders: data["orders"] = orders
        if product_name: data["productName"] = product_name
        if notify_url: data["notifyUrl"] = notify_url
        return self.client._post("/applyBatchWithdraw", data)

    def append_batch_withdraw(self, batch_order_id: str, orders: List[dict]):
        """Append tasks to an existing batch withdrawal order"""
        self.client._post("/appendBatchWithdraw", {
            "batchOrderId": batch_order_id,
            "orders": orders
        })

    def confirm_batch_withdraw(self, batch_order_id: str, delay_seconds: Optional[int] = None) -> dict:
        """Confirm and execute the batch withdrawal order"""
        data = {"batchOrderId": batch_order_id}
        if delay_seconds is not None: data["delaySeconds"] = delay_seconds
        return self.client._post("/confirmBatchWithdraw", data)

    def abort_batch_withdraw(self, batch_order_id: str, seqs: Optional[List[int]] = None) -> dict:
        """Cancel the batch withdrawal order or part of the tasks"""
        data = {"batchOrderId": batch_order_id}
        if seqs: data["seqs"] = seqs
        return self.client._post("/abortBatchWithdraw", data)

    def get_batch_withdraw_order(self, batch_order_id: str, verbose: Optional[int] = None) -> dict:
        """Query batch withdrawal order details"""
        data = {"batchOrderId": batch_order_id}
        if verbose is not None: data["verbose"] = verbose
        return self.client._post("/getBatchWithdrawOrder", data)

    def get_batch_withdraw_order_record_list(
        self,
        batch_order_id: str,
        next_id: Optional[str] = None,
        limit: Optional[int] = None
    ) -> dict:
        """Query the task record list of a batch withdrawal order"""
        data = {"batchOrderId": batch_order_id}
        if next_id: data["nextId"] = next_id
        if limit: data["limit"] = limit
        return self.client._post("/getBatchWithdrawOrderRecordList", data)
