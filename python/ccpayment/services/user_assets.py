"""User Assets Service"""


class UserAssetsService:
    def __init__(self, client):
        self.client = client

    def get_user_coin_asset_list(self, user_id: str) -> dict:
        """Get all token assets of a specified user"""
        result = self.client._post("/getUserCoinAssetList", {"userId": user_id})
        return result

    def get_user_coin_asset(self, user_id: str, coin_id: int) -> dict:
        """Get a specific token asset of a specified user"""
        result = self.client._post("/getUserCoinAsset", {
            "userId": user_id,
            "coinId": coin_id
        })
        return result.get("asset", {})
