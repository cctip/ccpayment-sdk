#HOST = "https://ebc65a6dtestpaymentadmin.cwallet.com"  # dev
HOST = "https://admin.ccpayment.com" # produce

# create order
CREATE_ORDER_URL = HOST + "/ccpayment/v1/bill/create"

# checkout url
CHECKOUT_URL = HOST + "/ccpayment/v1/concise/url/get"

# support token
SUPPORT_TOKEN_URL = HOST + "/ccpayment/v1/support/token"

# get token chains
TOKEN_CHAIN_URL = HOST + "/ccpayment/v1/token/chain"

# get token rate
TOKEN_RATE_URL = HOST + "/ccpayment/v1/token/rate"

# api - withdrawal
WITHDRAW_API_URL = HOST + "/ccpayment/v1/withdraw"

# c user check
CHECK_USER_URL = HOST + "/ccpayment/v1/check/user"

# get token assets
ASSETS_URL = HOST + "/ccpayment/v1/assets"

# get network fee
NETWORK_FEE_URL = HOST + "/ccpayment/v1/network/fee"

# header key
# app id
APP_ID_HEADER_KEY = "Appid"

# sign str
SIGN_HEADER_KEY = "Sign"

# timestamp
TIMESTAMP_HEADER_KEY = "Timestamp"
