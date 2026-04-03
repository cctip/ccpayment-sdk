module CCPayment
  module Services
    class MerchantDepositService
      def initialize(client)
        @client = client
      end

      def create_app_order_deposit_address(data)
        @client.post('/createAppOrderDepositAddress', data)
      end

      def get_or_create_app_deposit_address(data)
        @client.post('/getOrCreateAppDepositAddress', data)
      end

      def get_app_deposit_record(record_id)
        @client.post('/getAppDepositRecord', { recordId: record_id })
      end

      def get_app_deposit_record_list(data = {})
        @client.post('/getAppDepositRecordList', data)
      end
    end
  end
end
