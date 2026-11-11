"""User Swap Service"""


class UserSwapService:
    def __init__(self, client):
        self.client = client

    def get_user_swap_record(self, data):
        """Deprecated: use get_user_swap_record_v2."""
        return self.client._post("/getUserSwapRecord", data)

    def get_user_swap_record_list(self, data=None):
        """Deprecated: use get_user_swap_record_list_v2."""
        return self.client._post("/getUserSwapRecordList", data or {})

    def user_swap(self, data):
        """Deprecated: use user_swap_v2 with SwapReqV2 fields."""
        return self.client._post("/userSwap", data)

    def user_swap_v2(self, data):
        return self.client._post("/swapv2/userSwap", data)

    def get_user_swap_record_v2(self, data):
        return self.client._post("/swapv2/getUserSwapRecord", data)

    def get_user_swap_record_list_v2(self, data=None):
        return self.client._post("/swapv2/getUserSwapRecordList", data or {})
