module CCPayment
  module Services
    class UserDepositService
      def initialize(client)
        @client = client
      end

      def get_or_create_user_deposit_address(data)
        @client.post('/getOrCreateUserDepositAddress', data)
      end

      def get_user_deposit_record(record_id)
        @client.post('/getUserDepositRecord', { recordId: record_id })
      end

      def get_user_deposit_record_list(data)
        @client.post('/getUserDepositRecordList', data)
      end
    end
  end
end
