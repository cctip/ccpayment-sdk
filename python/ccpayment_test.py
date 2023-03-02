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
    oc.chain = 'TRX'
    oc.contract = 'TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t'
    oc.merchant_order_id = '1640938137795209'
    oc.token_id = '8e5741cf-6e51-4892-9d04-3d40e1dd0128'
    oc.fiat_currency = 'USD'
    oc.remark = ''

    result, status = oc.create_order(app_id, app_secret)
    if status == 0:
        print("verify success")
    else:
        print("verify error")

    print(result)


# webhook validate
class TestWebhookValidateClass(unittest.TestCase):
    wk = ccpayment.WebhookClass()
    wk.data_str = ''
    wk.signature = ''
    wk.timestamp = ''

    if wk.webhook_validate(app_id, app_secret):
        print('verify success')
    else:
        print('verify error')


# get checkout url
class TestCheckoutUrlClass(unittest.TestCase):
    pu = ccpayment.PaymentUrlClass()
    pu.amount = '12'
    pu.merchant_order_id = '1234'
    pu.valid_timestamp = 300  # s
    pu.product_name = 'test'
    pu.return_url = 'http://localhost/index.html'

    result, status = pu.get_checkout_url(app_id, app_secret)
    if status == 0:
        print("verify success")
    else:
        print("verify error")

    print(result)


# get support token list
class TestSupportTokens(unittest.TestCase):
    result, status = ccpayment.get_support_tokens(app_id, app_secret)
    if status == 0:
        print("verify success")
    else:
        print("verify error")

    print(result)


# get token chain
class TestTokenChain(unittest.TestCase):
    tc = ccpayment.TokenChainClass()
    tc.token_id = '8e5741cf-6e51-4892-9d04-3d40e1dd0128'
    result, status = tc.get_token_chain(app_id, app_secret)
    if status == 0:
        print("verify success")
    else:
        print("verify error")

    print(result)


# get token rate
class TestTokenRate(unittest.TestCase):
    tr = ccpayment.TokenRateClass()
    tr.token_id = '8e5741cf-6e51-4892-9d04-3d40e1dd0128'
    tr.amount = '12'
    result, status = tr.get_token_rate(app_id, app_secret)
    if status == 0:
        print("verify success")
    else:
        print("verify error")

    print(result)


if __name__ == '__main__':
    unittest.main()
