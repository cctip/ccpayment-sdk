HOST = "https://ebc65a6dtestpaymentadmin.cwallet.com"  # dev
# HOST = "https://admin.ccpayment.com" # produce

CREATE_ORDER_URL = HOST + "/ccpayment/v1/bill/create"

CHECKOUT_URL = HOST + "/ccpayment/v1/concise/url/get"

SUPPORT_TOKEN_URL = HOST + "/ccpayment/v1/support/token"

TOKEN_CHAIN_URL = HOST + "/ccpayment/v1/token/chain"

TOKEN_RATE_URL = HOST + "/ccpayment/v1/token/rate"

WITHDRAW_API_URL = HOST + "/ccpayment/v1/withdraw"

CHECK_USER_URL = HOST + "/ccpayment/v1/check/user"

ASSETS_URL = HOST + "/ccpayment/v1/assets"

NETWORK_FEE_URL = HOST + "/ccpayment/v1/network/fee"

# header key
APP_ID_HEADER_KEY = "Appid"
SIGN_HEADER_KEY = "Sign"
TIMESTAMP_HEADER_KEY = "Timestamp"
