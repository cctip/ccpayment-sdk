"""Utilities Service"""


class UtilitiesService:
    def __init__(self, client):
        self.client = client

    def verify_address(self, chain: str, address: str) -> dict:
        """Verify the validity of a blockchain address"""
        return self.client._post("/verifyAddress", {
            "chain": chain,
            "address": address
        })

    def abandon_address(self, chain: str, address: str):
        """Abandon a deposit address that is no longer in use"""
        self.client._post("/abandonAddress", {
            "chain": chain,
            "address": address
        })

    def hosted_invoice_order_info(self, order_id: str) -> dict:
        """Get Hosted Invoice order details"""
        return self.client._post("/hostedInvoiceOrderInfo", {"orderId": order_id})

    def get_pay_info(self, order_id: str) -> dict:
        """Get the payment information of an order"""
        return self.client._post("/getPayInfo", {"orderId": order_id})

    def health(self) -> dict:
        """API health check interface"""
        try:
            return self.client._post("/health", {})
        except Exception as e:
            # Health endpoint may return plain text or 404
            if "404" in str(e):
                return {"status": "404 Not Found", "message": "Health endpoint not available"}
            # Try to return the response text if available
            return {"status": str(e), "message": "Health check completed with exception"}
