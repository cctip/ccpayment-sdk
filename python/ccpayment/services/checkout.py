"""Checkout Service"""

from typing import Optional, List


class CheckoutService:
    def __init__(self, client):
        self.client = client

    def create_checkout_url(
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
        """Create a Hosted Checkout payment URL"""
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
        result = self.client._post("/createCheckoutUrl", data)
        return result.get("checkoutUrl")

    def hosted_order_info(self, order_id: str) -> dict:
        """Get Hosted order details"""
        return self.client._post("/hostedOrderInfo", {"orderId": order_id})

    def hosted_order_info_first(self, order_id: str) -> dict:
        """Get the first visit information of a Hosted order"""
        return self.client._post("/hostedOrderInfoFirst", {"orderId": order_id})

    def create_hosted_order_deposit(self, order_id: str, coin_id: int, chain: str) -> dict:
        """Create a deposit address after selecting token and chain for a Hosted order"""
        return self.client._post("/createHostedOrderDeposit", {
            "orderId": order_id,
            "coinId": coin_id,
            "chain": chain
        })

    def get_hosted_coin_usdt_price(self, coin_ids: List[int]) -> dict:
        """Get the USDT price of a token in Hosted mode"""
        result = self.client._post("/getHostedCoinUSDTPrice", {"coinIds": coin_ids})
        return result.get("prices", {})

    def get_main_coin_list(self) -> List[dict]:
        """Get the list of main supported tokens in Hosted mode"""
        result = self.client._post("/getMainCoinList", {})
        return result.get("coins", [])

    def create_app_demo_order_deposit(
        self,
        order_id: str,
        coin_id: int,
        chain: str,
        price: str
    ) -> dict:
        """Create a Demo order deposit address"""
        return self.client._post("/createAppDemoOrderDeposit", {
            "orderId": order_id,
            "coinId": coin_id,
            "chain": chain,
            "price": price
        })

    def confirm_trade(self, order_id: str, tx_id: str):
        """Confirm the trade of a Hosted order"""
        self.client._post("/confirmTrade", {
            "orderId": order_id,
            "txId": tx_id
        })
