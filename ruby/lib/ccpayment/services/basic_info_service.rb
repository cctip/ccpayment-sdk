module CCPayment
  module Services
    class BasicInfoService
      def initialize(client)
        @client = client
      end

      def get_coin_list
        @client.post('/getCoinList', {})
      end

      def get_coin(coin_id)
        @client.post('/getCoin', { coinId: coin_id })
      end

      def get_coin_usdt_price(coin_ids)
        @client.post('/getCoinUSDTPrice', { coinIds: coin_ids })
      end

      def get_fiat_list
        @client.post('/getFiatList', {})
      end

      def get_chain_list(chains = nil)
        data = chains ? { chains: chains } : {}
        @client.post('/getChainList', data)
      end

      def all_chain(chains = nil)
        data = chains ? { chains: chains } : {}
        @client.post('/all/chain', data)
      end

      def get_cwallet_user_id(cwallet_user_id)
        @client.post('/getCwalletUserId', { cwalletUserId: cwallet_user_id })
      end

      def get_withdraw_fee(coin_id, chain)
        @client.post('/getWithdrawFee', { coinId: coin_id, chain: chain })
      end
    end
  end
end
