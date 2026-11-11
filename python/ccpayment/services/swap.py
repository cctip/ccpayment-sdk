"""Swap Service"""


class SwapService:
    def __init__(self, client):
        self.client = client

    def get_swap_coin_list(self):
        """Deprecated: use get_swap_pair_list."""
        return self.client._post("/getSwapCoinList", {})

    def swap_estimate(self, data):
        """Deprecated: use swap_estimate_v2 with SwapReqV2 fields."""
        return self.client._post("/estimate", data)

    def get_swap_record(self, data):
        """Deprecated: use get_swap_record_v2."""
        return self.client._post("/getSwapRecord", data)

    def get_swap_record_list(self, data=None):
        """Deprecated: use get_swap_record_list_v2."""
        return self.client._post("/getSwapRecordList", data or {})

    def swap(self, data):
        """Deprecated: use swap_v2 with SwapReqV2 fields."""
        return self.client._post("/swap", data)

    def get_swap_pair_list(self):
        return self.client._post("/swapv2/getSwapPairList", {})

    def swap_estimate_v2(self, data):
        return self.client._post("/swapv2/estimate", data)

    def swap_v2(self, data):
        return self.client._post("/swapv2/swap", data)

    def get_swap_record_v2(self, data):
        return self.client._post("/swapv2/getSwapRecord", data)

    def get_swap_record_list_v2(self, data=None):
        return self.client._post("/swapv2/getSwapRecordList", data or {})
