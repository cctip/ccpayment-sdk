module CCPayment
  module Services
    class UserTransferService
      def initialize(client)
        @client = client
      end

      def user_transfer(data)
        @client.post('/userTransfer', data)
      end

      def get_user_transfer_record(data)
        @client.post('/getUserTransferRecord', data)
      end

      def get_user_transfer_record_list(data = {})
        @client.post('/getUserTransferRecordList', data)
      end

      def user_batch_transfer(data)
        @client.post('/userBatchTransfer', data)
      end

      def get_user_batch_transfer_record(data)
        @client.post('/getUserBatchTransferRecord', data)
      end

      def get_user_batch_transfer_record_list(data = {})
        @client.post('/getUserBatchTransferRecordList', data)
      end

      def get_user_batch_transfer_record_detail(data)
        @client.post('/getUserBatchTransferRecordDetail', data)
      end
    end
  end
end
