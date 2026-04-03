module CCPayment
  module Services
    class MerchantWithdrawService
      def initialize(client)
        @client = client
      end

      def apply_app_withdraw_to_network(data)
        @client.post('/applyAppWithdrawToNetwork', data)
      end

      def apply_app_withdraw_to_cwallet(data)
        @client.post('/applyAppWithdrawToCwallet', data)
      end

      def get_app_withdraw_record(data)
        @client.post('/getAppWithdrawRecord', data)
      end

      def get_app_withdraw_record_list(data = {})
        @client.post('/getAppWithdrawRecordList', data)
      end

      def get_auto_withdraw_record_list(data = {})
        @client.post('/getAutoWithdrawRecordList', data)
      end

      def get_risky_refund_record_list(data = {})
        @client.post('/getRiskyRefundRecordList', data)
      end
    end
  end
end
