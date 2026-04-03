module CCPayment
  module Services
    class MerchantBatchWithdrawService
      def initialize(client)
        @client = client
      end

      def check_withdraw_address(data)
        @client.post('/checkWithdrawAddress', data)
      end

      def apply_batch_withdraw(data)
        @client.post('/applyBatchWithdraw', data)
      end

      def append_batch_withdraw(data)
        @client.post('/appendBatchWithdraw', data)
      end

      def confirm_batch_withdraw(data)
        @client.post('/confirmBatchWithdraw', data)
      end

      def abort_batch_withdraw(data)
        @client.post('/abortBatchWithdraw', data)
      end

      def get_batch_withdraw_order(data)
        @client.post('/getBatchWithdrawOrder', data)
      end

      def get_batch_withdraw_order_record_list(data)
        @client.post('/getBatchWithdrawOrderRecordList', data)
      end
    end
  end
end
