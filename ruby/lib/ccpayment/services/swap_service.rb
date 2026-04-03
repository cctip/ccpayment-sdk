module CCPayment
  module Services
    class SwapService
      def initialize(client)
        @client = client
      end

      def get_swap_coin_list
        @client.post('/getSwapCoinList', {})
      end

      def swap_estimate(data)
        @client.post('/swapEstimate', data)
      end

      def swap(data)
        @client.post('/swap', data)
      end

      def get_swap_record(data)
        @client.post('/getSwapRecord', data)
      end

      def get_swap_record_list(data = {})
        @client.post('/getSwapRecordList', data)
      end

      def select_hosted_invoice_coin(data)
        @client.post('/selectHostedInvoiceCoin', data)
      end
    end
  end
end
