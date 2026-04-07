module CCPayment
  module Services
    class UtilitiesService
      def initialize(client)
        @client = client
      end

      def verify_address(chain, address)
        @client.post('/verifyAddress', { chain: chain, address: address })
      end

      def abandon_address(chain, address)
        @client.post('/abandonAddress', { chain: chain, address: address })
      end

      def hosted_invoice_order_info(order_id)
        @client.post('/hostedInvoiceOrderInfo', { orderId: order_id })
      end

      def get_pay_info(order_id)
        @client.post('/getPayInfo', { orderId: order_id })
      end
    end
  end
end
