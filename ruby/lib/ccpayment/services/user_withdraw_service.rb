module CCPayment
  module Services
    class UserWithdrawService
      def initialize(client)
        @client = client
      end

      def apply_user_withdraw_to_network(data)
        @client.post('/applyUserWithdrawToNetwork', data)
      end

      def apply_user_withdraw_to_cwallet(data)
        @client.post('/applyUserWithdrawToCwallet', data)
      end

      def get_user_withdraw_record(data)
        @client.post('/getUserWithdrawRecord', data)
      end

      def get_user_withdraw_record_list(data)
        @client.post('/getUserWithdrawRecordList', data)
      end
    end
  end
end
