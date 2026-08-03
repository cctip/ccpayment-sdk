module CCPayment
  module Services
    class UserSwapService
      def initialize(client); @client = client; end
      def get_user_swap_record(data); @client.post('/getUserSwapRecord', data); end
      def get_user_swap_record_list(data = {}); @client.post('/getUserSwapRecordList', data); end
      def user_swap(data); @client.post('/userSwap', data); end
    end
  end
end
