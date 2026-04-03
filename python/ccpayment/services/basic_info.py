"""Basic Info Service"""

from typing import List, Dict, Optional


class BasicInfoService:
    def __init__(self, client):
        self.client = client

    def get_coin_list(self) -> List[dict]:
        """Get all supported token information"""
        result = self.client._post("/getCoinList", {})
        return result.get("coins", [])

    def get_coin(self, coin_id: int) -> dict:
        """Get detailed information for a specified token"""
        return self.client._post("/getCoin", {"coinId": coin_id})

    def get_coin_usdt_price(self, coin_ids: List[int]) -> Dict[str, str]:
        """Get USD prices for tokens in batch"""
        result = self.client._post("/getCoinUSDTPrice", {"coinIds": coin_ids})
        return result.get("prices", {})

    def get_fiat_list(self) -> List[dict]:
        """Get all supported fiat currency information"""
        result = self.client._post("/getFiatList", {})
        return result.get("fiats", [])

    def get_chain_list(self, chains: Optional[List[str]] = None) -> List[dict]:
        """Get chain information"""
        data = {"chains": chains} if chains else {}
        result = self.client._post("/getChainList", data)
        return result.get("chains", [])

    def all_chain(self, chains: Optional[List[str]] = None) -> List[dict]:
        """Get information for all chains (simplified version)"""
        data = {"chains": chains} if chains else {}
        result = self.client._post("/all/chain", data)
        return result.get("chains", [])

    def get_cwallet_user_id(self, cwallet_user_id: str) -> dict:
        """Verify and get CCWallet user information"""
        return self.client._post("/getCwalletUserId", {"cwalletUserId": cwallet_user_id})

    def get_withdraw_fee(self, coin_id: int, chain: str) -> dict:
        """Get the withdrawal fee for a specified token and chain"""
        return self.client._post("/getWithdrawFee", {"coinId": coin_id, "chain": chain})
