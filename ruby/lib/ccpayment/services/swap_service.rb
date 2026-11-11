module CCPayment
  module Services
    class SwapService
      def initialize(client); @client = client; end
      # Deprecated: use get_swap_pair_list.
      def get_swap_coin_list; @client.post('/getSwapCoinList', {}); end
      # Deprecated: use swap_estimate_v2 with SwapReqV2 fields.
      def swap_estimate(data); @client.post('/estimate', data); end
      # Deprecated: use get_swap_record_v2.
      def get_swap_record(data); @client.post('/getSwapRecord', data); end
      # Deprecated: use get_swap_record_list_v2.
      def get_swap_record_list(data = {}); @client.post('/getSwapRecordList', data); end
      # Deprecated: use swap_v2 with SwapReqV2 fields.
      def swap(data); @client.post('/swap', data); end
      def get_swap_pair_list; @client.post('/swapv2/getSwapPairList', {}); end
      def swap_estimate_v2(data); @client.post('/swapv2/estimate', data); end
      def swap_v2(data); @client.post('/swapv2/swap', data); end
      def get_swap_record_v2(data); @client.post('/swapv2/getSwapRecord', data); end
      def get_swap_record_list_v2(data = {}); @client.post('/swapv2/getSwapRecordList', data); end
    end
  end
end
