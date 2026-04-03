module CCPayment
  module Services
    class OrdersService
      def initialize(client)
        @client = client
      end

      def get_app_order_info(order_id)
        @client.post('/getAppOrderInfo', { orderId: order_id })
      end

      def create_invoice_url(data)
        @client.post('/createInvoiceUrl', data)
      end

      def get_invoice_order_info(order_id)
        @client.post('/getInvoiceOrderInfo', { orderId: order_id })
      end

      def get_webhook_info
        @client.post('/getWebhookInfo', {})
      end
    end
  end
end
