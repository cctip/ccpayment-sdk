module CCPayment
  module Services
    class UtilitiesService
      def initialize(client); @client = client; end
      def webhook_resend(data = {}); @client.post('/webhook/resend', data); end
      def get_trade_block_height(data); @client.post('/getTradeBlockHeight', data); end
      def check_withdrawal_address_validity(data); @client.post('/checkWithdrawalAddressValidity', data); end
      def deprecated_address(data); @client.post('/addressUnbinding', data); end
      def rescan_lost_transaction(data); @client.post('/rescanLostTransaction', data); end
    end
  end
end
