require_relative 'ccpayment/client'
require_relative 'ccpayment/api_error'
require_relative 'ccpayment/services/basic_info_service'
require_relative 'ccpayment/services/merchant_assets_service'
require_relative 'ccpayment/services/merchant_deposit_service'
require_relative 'ccpayment/services/merchant_withdraw_service'
require_relative 'ccpayment/services/merchant_batch_withdraw_service'
require_relative 'ccpayment/services/user_assets_service'
require_relative 'ccpayment/services/user_deposit_service'
require_relative 'ccpayment/services/user_withdraw_service'
require_relative 'ccpayment/services/user_transfer_service'
require_relative 'ccpayment/services/orders_service'
require_relative 'ccpayment/services/checkout_service'
require_relative 'ccpayment/services/swap_service'
require_relative 'ccpayment/services/utilities_service'

module CCPayment
  VERSION = '1.0.0'
end
