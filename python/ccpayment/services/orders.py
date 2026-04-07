"""Orders Service"""


class OrdersService:
    def __init__(self, client):
        self.client = client

    def get_app_order_info(self, order_id: str):
        return self.client._post("/getAppOrderInfo", {"orderId": order_id})

    def create_invoice_url(self, data):
        return self.client._post("/createInvoiceUrl", data)

    def get_invoice_order_info(self, order_id: str):
        return self.client._post("/getInvoiceOrderInfo", {"orderId": order_id})
