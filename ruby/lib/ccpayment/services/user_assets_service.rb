module CCPayment
  module Services
    class UserAssetsService
      def initialize(client)
        @client = client
      end

      def get_user_coin_asset_list(user_id)
        @client.post('/getUserCoinAssetList', { userId: user_id })
      end

      def get_user_coin_asset(user_id, coin_id)
        @client.post('/getUserCoinAsset', { userId: user_id, coinId: coin_id })
      end
    end
  end
end
