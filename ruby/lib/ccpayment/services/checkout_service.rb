module CCPayment
  module Services
    class CheckoutService
      def initialize(client)
        @client = client
      end

      def create_checkout_url(data)
        @client.post('/createCheckoutUrl', data)
      end

      def hosted_order_info(order_id)
        @client.post('/hostedOrderInfo', { orderId: order_id })
      end

      def hosted_order_info_first(order_id)
        @client.post('/hostedOrderInfoFirst', { orderId: order_id })
      end

      def create_hosted_order_deposit(data)
        @client.post('/createHostedOrderDeposit', data)
      end

      def get_hosted_coin_usdt_price(coin_ids)
        @client.post('/getHostedCoinUSDTPrice', { coinIds: coin_ids })
      end

      def get_main_coin_list
        @client.post('/getMainCoinList', {})
      end

      def create_app_demo_order_deposit(data)
        @client.post('/createAppDemoOrderDeposit', data)
      end

      def confirm_trade(data)
        @client.post('/confirmTrade', data)
      end
    end
  end
end
