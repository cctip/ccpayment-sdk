"""Swap Service"""

from typing import Optional, List


class SwapService:
    def __init__(self, client):
        self.client = client

    def get_swap_coin_list(self) -> List[dict]:
        """Get the list of tokens supported for swapping"""
        result = self.client._post("/getSwapCoinList", {})
        return result.get("coins", [])

    def swap_estimate(self, from_coin_id: int, to_coin_id: int, from_amount: str) -> dict:
        """Estimate the swap amount"""
        return self.client._post("/swapEstimate", {
            "fromCoinId": from_coin_id,
            "toCoinId": to_coin_id,
            "fromAmount": from_amount
        })

    def swap(self, order_id: str, from_coin_id: int, to_coin_id: int, from_amount: str) -> str:
        """Execute a swap operation"""
        result = self.client._post("/swap", {
            "orderId": order_id,
            "fromCoinId": from_coin_id,
            "toCoinId": to_coin_id,
            "fromAmount": from_amount
        })
        return result.get("recordId")

    def get_swap_record(self, record_id: Optional[str] = None, order_id: Optional[str] = None) -> dict:
        """Query details of a single swap record"""
        data = {}
        if record_id: data["recordId"] = record_id
        if order_id: data["orderId"] = order_id
        result = self.client._post("/getSwapRecord", data)
        return result.get("record", {})

    def get_swap_record_list(
        self,
        order_ids: Optional[List[str]] = None,
        from_coin_id: Optional[int] = None,
        to_coin_id: Optional[int] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        next_id: Optional[str] = None
    ) -> dict:
        """Query swap record list"""
        data = {}
        if order_ids: data["orderIds"] = order_ids
        if from_coin_id: data["fromCoinId"] = from_coin_id
        if to_coin_id: data["toCoinId"] = to_coin_id
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if next_id: data["nextId"] = next_id
        return self.client._post("/getSwapRecordList", data)

    def select_hosted_invoice_coin(self, order_id: str, coin_id: int, chain: str) -> dict:
        """Select the payment token for a Hosted Invoice order"""
        return self.client._post("/selectHostedInvoiceCoin", {
            "orderId": order_id,
            "coinId": coin_id,
            "chain": chain
        })
