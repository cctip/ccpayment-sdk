"""Orders Service"""

from typing import Optional


class OrdersService:
    def __init__(self, client):
        self.client = client

    def get_app_order_info(self, order_id: str) -> dict:
        """Get Order order details"""
        return self.client._post("/getAppOrderInfo", {"orderId": order_id})

    def create_invoice_url(
        self,
        order_id: str,
        product: str,
        price: str,
        price_fiat_id: Optional[int] = None,
        price_coin_id: Optional[int] = None,
        buyer_email: Optional[str] = None,
        return_url: Optional[str] = None,
        expired_at: Optional[int] = None,
        close_url: Optional[str] = None,
        notify_url: Optional[str] = None
    ) -> str:
        """Create an Invoice order and generate a payment URL"""
        data = {
            "orderId": order_id,
            "product": product,
            "price": price
        }
        if price_fiat_id: data["priceFiatId"] = price_fiat_id
        if price_coin_id: data["priceCoinId"] = price_coin_id
        if buyer_email: data["buyerEmail"] = buyer_email
        if return_url: data["returnUrl"] = return_url
        if expired_at: data["expiredAt"] = expired_at
        if close_url: data["closeUrl"] = close_url
        if notify_url: data["notifyUrl"] = notify_url
        result = self.client._post("/createInvoiceUrl", data)
        return result.get("invoiceUrl")

    def get_invoice_order_info(self, order_id: str) -> dict:
        """Get Invoice order details"""
        result = self.client._post("/getInvoiceOrderInfo", {"orderId": order_id})
        return result

    def get_webhook_info(self) -> dict:
        """Get merchant's Webhook configuration information"""
        return self.client._post("/getWebhookInfo", {})
