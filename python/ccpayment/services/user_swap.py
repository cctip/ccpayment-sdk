"""User Swap Service"""


class UserSwapService:
    def __init__(self, client):
        self.client = client

    def get_user_swap_record(self, data):
        return self.client._post("/getUserSwapRecord", data)

    def get_user_swap_record_list(self, data=None):
        return self.client._post("/getUserSwapRecordList", data or {})

    def user_swap(self, data):
        return self.client._post("/userSwap", data)
