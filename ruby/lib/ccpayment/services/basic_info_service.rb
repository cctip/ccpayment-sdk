module CCPayment
  module Services
    class BasicInfoService
      def initialize(client); @client = client; end
      def get_coin_list; @client.post('/getCoinList', {}); end
      def get_fiat_list; @client.post('/getFiatList', {}); end
      def get_coin(coin_id); @client.post('/getCoin', { coinId: coin_id }); end
      def get_coin_usdt_price(coin_ids); @client.post('/getCoinUSDTPrice', { coinIds: coin_ids }); end
      def get_chain_list(data = {}); @client.post('/getChainList', data); end
      def get_all_chain_list(data = {}); @client.post('/all/chain', data); end
      def get_main_coin_list(app_id); @client.post('/getMainCoinList', { appId: app_id }); end
    end
  end
end
