"""Merchant Assets Service"""


class MerchantAssetsService:
    def __init__(self, client):
        self.client = client

    def get_app_coin_asset_list(self):
        return self.client._post("/getAppCoinAssetList", {})

    def get_app_coin_asset(self, coin_id: int):
        return self.client._post("/getAppCoinAsset", {"coinId": coin_id})

    def get_app_collect_fee_record_list(self, data=None):
        return self.client._post("/getAggregationFeeRecord", data or {})
