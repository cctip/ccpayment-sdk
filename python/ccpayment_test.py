import time
import unittest
import ccpayment

app_id = "6cfakPeNnTi0YoHLpbqVLqFw2A#R5EGPH3lJfX75H7x#SO#PXU"
app_secret = "5SQOObrXo#t82ZsVYfrI02Dn@5MLVPX1Fr1gcpCVS2Ca8ClbAgi@AMp64j0pWO#4jyC#zdS3kis1UaNFsaqW5DGjtxA4gYk1ZEH68#S1Z1McmUd@ph8Gbhn#"


# create order
class TestCreateOrder(unittest.TestCase):
    # create order
    oc = ccpayment.OrderClass()

    oc.amount = '6'
    oc.merchant_order_id = str(int(time.time()))
    oc.token_id = '0912e09a-d8e2-41d7-a0bc-a25530892988'
    oc.fiat_currency = 'USD'
    oc.remark = ''

    result, is_verify = oc.create_deposit_order(app_id, app_secret)
    if is_verify:
        print("TestCreateOrder: verify success")
    else:
        print("TestCreateOrder: verify error")

    print("TestCreateOrder:", result)


# webhook validate
class TestWebhookValidate(unittest.TestCase):
    wk = ccpayment.WebhookClass()
    wk.data_str = ''
    wk.signature = ''
    wk.timestamp = ''

    if wk.webhook_validate(app_id, app_secret):
        print('TestWebhookValidate: verify success')
    else:
        print('TestWebhookValidate: verify error')


# get checkout url
class TestCheckoutUrl(unittest.TestCase):
    pu = ccpayment.PaymentUrlClass()
    pu.amount = '12'
    pu.merchant_order_id = str(int(time.time()))
    pu.valid_timestamp = 300  # s
    pu.product_name = 'test'
    pu.return_url = 'http://localhost/index.html'

    result, is_verify = pu.get_checkout_url(app_id, app_secret)
    if is_verify:
        print("TestCheckoutUrl: verify success")
    else:
        print("TestCheckoutUrl :verify error")

    print("TestCheckoutUrl:", result)


# get support token list
class TestSupportTokens(unittest.TestCase):
    result, is_verify = ccpayment.get_support_tokens(app_id, app_secret)
    if is_verify:
        print("TestSupportTokens: verify success")
    else:
        print("TestSupportTokens: verify error")

    print("TestSupportTokens:", result)


# get token chain
class TestTokenChain(unittest.TestCase):
    tc = ccpayment.TokenChainClass()
    tc.token_id = '8e5741cf-6e51-4892-9d04-3d40e1dd0128'
    result, is_verify = tc.get_token_chain(app_id, app_secret)
    if is_verify:
        print("TestTokenChain: verify success")
    else:
        print("TestTokenChain: verify error")

    print("TestTokenChain:", result)


# get token rate
class TestTokenRate(unittest.TestCase):
    tr = ccpayment.TokenRateClass()
    tr.token_id = '8e5741cf-6e51-4892-9d04-3d40e1dd0128'
    tr.amount = '12'
    result, is_verify = tr.get_token_rate(app_id, app_secret)
    if is_verify:
        print("TestTokenRate: verify success")
    else:
        print("TestTokenRate: verify error")

    print("TestTokenRate:", result)


# create api withdrawal
class TestApiWithdraw(unittest.TestCase):
    aw = ccpayment.ApiWithdrawClass()
    aw.token_id = 'f137d42c-f3a6-4f23-9402-76f0395d0cfe'
    aw.address = '9486792'
    aw.memo = ''
    aw.value = '1'
    aw.merchant_order_id = str(int(time.time()))

    result, is_verify = aw.create_withdraw_order(app_id, app_secret)
    if is_verify:
        print("TestApiWithdraw: verify success")
    else:
        print("TestApiWithdraw: verify error")

    print("TestApiWithdraw:", result)


# check user
class TestCheckUser(unittest.TestCase):
    cu = ccpayment.CheckUserClass()
    cu.c_id = '9486792'

    result, is_verify = cu.check_user(app_id, app_secret)
    if is_verify:
        print("TestCheckUser: verify success")
    else:
        print("TestCheckUser: verify error")

    print("TestCheckUser:", result)


# get token assets
class TestTokenAsset(unittest.TestCase):
    ta = ccpayment.TokenAssetClass()
    ta.token_id = '8e5741cf-6e51-4892-9d04-3d40e1dd0128'

    result, is_verify = ta.get_token_assets(app_id, app_secret)
    if is_verify:
        print("TestTokenAsset: verify success")
    else:
        print("TestTokenAsset :verify error")

    print("TestTokenAsset:", result)


# network fee
class TestNetworkFee(unittest.TestCase):
    nf = ccpayment.NetworkFeeClass()
    nf.token_id = 'f137d42c-f3a6-4f23-9402-76f0395d0cfe'
    nf.address = ''
    nf.memo = ''

    result, is_verify = nf.get_network_fee(app_id, app_secret)
    if is_verify:
        print("TestNetworkFee: verify success")
    else:
        print("TestNetworkFee: verify error")

    print("TestNetworkFee:", result)


if __name__ == '__main__':
    unittest.main()
