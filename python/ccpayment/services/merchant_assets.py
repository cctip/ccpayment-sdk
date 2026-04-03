"""Merchant Assets Service"""


class MerchantAssetsService:
    def __init__(self, client):
        self.client = client

    def get_app_coin_asset_list(self) -> List[dict]:
        """Get all token assets of the merchant"""
        result = self.client._post("/getAppCoinAssetList", {})
        return result.get("assets", [])

    def get_app_coin_asset(self, coin_id: int) -> dict:
        """Get the asset of a specified token for the merchant"""
        result = self.client._post("/getAppCoinAsset", {"coinId": coin_id})
        return result.get("asset", {})
