"""Merchant Deposit Service"""

from typing import Optional


class MerchantDepositService:
    def __init__(self, client):
        self.client = client

    def create_app_order_deposit_address(
        self,
        order_id: str,
        coin_id: int,
        chain: str,
        price: str,
        fiat_id: Optional[int] = None,
        expired_at: Optional[int] = None,
        buyer_email: Optional[str] = None,
        generate_checkout_url: Optional[bool] = None,
        product: Optional[str] = None,
        return_url: Optional[str] = None,
        notify_url: Optional[str] = None,
        close_url: Optional[str] = None
    ) -> dict:
        """Create a deposit address for a specific order"""
        data = {
            "orderId": order_id,
            "coinId": coin_id,
            "chain": chain,
            "price": price
        }
        if fiat_id: data["fiatId"] = fiat_id
        if expired_at: data["expiredAt"] = expired_at
        if buyer_email: data["buyerEmail"] = buyer_email
        if generate_checkout_url: data["generateCheckoutURL"] = generate_checkout_url
        if product: data["product"] = product
        if return_url: data["returnUrl"] = return_url
        if notify_url: data["notifyUrl"] = notify_url
        if close_url: data["closeUrl"] = close_url
        return self.client._post("/createAppOrderDepositAddress", data)

    def get_or_create_app_deposit_address(
        self,
        reference_id: str,
        chain: str,
        notify_url: Optional[str] = None
    ) -> dict:
        """Get or create a direct deposit address"""
        data = {
            "referenceId": reference_id,
            "chain": chain
        }
        if notify_url: data["notifyUrl"] = notify_url
        return self.client._post("/getOrCreateAppDepositAddress", data)

    def get_app_deposit_record(self, record_id: str) -> dict:
        """Query details of a single deposit record"""
        result = self.client._post("/getAppDepositRecord", {"recordId": record_id})
        return result.get("record", {})

    def get_app_deposit_record_list(
        self,
        chain: Optional[str] = None,
        reference_id: Optional[str] = None,
        order_id: Optional[str] = None,
        to_address: Optional[str] = None,
        coin_id: Optional[int] = None,
        start_at: Optional[int] = None,
        end_at: Optional[int] = None,
        next_id: Optional[str] = None,
        record_ids: Optional[List[str]] = None,
        reference_ids: Optional[List[str]] = None,
        order_ids: Optional[List[str]] = None,
        limit: Optional[int] = None
    ) -> dict:
        """Query deposit record list"""
        data = {}
        if chain: data["chain"] = chain
        if reference_id: data["referenceId"] = reference_id
        if order_id: data["orderId"] = order_id
        if to_address: data["toAddress"] = to_address
        if coin_id: data["coinId"] = coin_id
        if start_at: data["startAt"] = start_at
        if end_at: data["endAt"] = end_at
        if next_id: data["nextId"] = next_id
        if record_ids: data["recordIds"] = record_ids
        if reference_ids: data["referenceIds"] = reference_ids
        if order_ids: data["orderIds"] = order_ids
        if limit: data["limit"] = limit
        return self.client._post("/getAppDepositRecordList", data)
