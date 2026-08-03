"""Basic Info Service"""


class BasicInfoService:
    def __init__(self, client):
        self.client = client

    def get_coin_list(self):
        return self.client._post("/getCoinList", {})

    def get_fiat_list(self):
        return self.client._post("/getFiatList", {})

    def get_coin(self, coin_id: int):
        return self.client._post("/getCoin", {"coinId": coin_id})

    def get_coin_usdt_price(self, coin_ids):
        return self.client._post("/getCoinUSDTPrice", {"coinIds": coin_ids})

    def get_chain_list(self, data=None):
        return self.client._post("/getChainList", data or {})

    def get_all_chain_list(self, data=None):
        return self.client._post("/all/chain", data or {})

    def get_main_coin_list(self, app_id: str):
        return self.client._post("/getMainCoinList", {"appId": app_id})
