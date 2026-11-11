module CCPayment
  module Services
    class UserSwapService
      def initialize(client); @client = client; end
      # Deprecated: use get_user_swap_record_v2.
      def get_user_swap_record(data); @client.post('/getUserSwapRecord', data); end
      # Deprecated: use get_user_swap_record_list_v2.
      def get_user_swap_record_list(data = {}); @client.post('/getUserSwapRecordList', data); end
      # Deprecated: use user_swap_v2 with SwapReqV2 fields.
      def user_swap(data); @client.post('/userSwap', data); end
      def user_swap_v2(data); @client.post('/swapv2/userSwap', data); end
      def get_user_swap_record_v2(data); @client.post('/swapv2/getUserSwapRecord', data); end
      def get_user_swap_record_list_v2(data = {}); @client.post('/swapv2/getUserSwapRecordList', data); end
    end
  end
end
