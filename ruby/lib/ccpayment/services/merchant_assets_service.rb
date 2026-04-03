module CCPayment
  module Services
    class MerchantAssetsService
      def initialize(client)
        @client = client
      end

      def get_app_coin_asset_list
        @client.post('/getAppCoinAssetList', {})
      end

      def get_app_coin_asset(coin_id)
        @client.post('/getAppCoinAsset', { coinId: coin_id })
      end
    end
  end
end
