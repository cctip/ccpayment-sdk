"""Swap Service"""


class SwapService:
    def __init__(self, client):
        self.client = client

    def get_swap_coin_list(self):
        return self.client._post("/getSwapCoinList", {})

    def swap_estimate(self, data):
        return self.client._post("/estimate", data)

    def get_swap_record(self, data):
        return self.client._post("/getSwapRecord", data)

    def get_swap_record_list(self, data=None):
        return self.client._post("/getSwapRecordList", data or {})

    def swap(self, data):
        return self.client._post("/swap", data)
