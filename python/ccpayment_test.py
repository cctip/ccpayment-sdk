import time
import unittest
import ccpayment

app_id = "6cfakPeNnTi0YoHLpbqVLqFw2A#R5EGPH3lJfX75H7x#SO#PXU"
app_secret = "5SQOObrXo#t82ZsVYfrI02Dn@5MLVPX1Fr1gcpCVS2Ca8ClbAgi@AMp64j0pWO#4jyC#zdS3kis1UaNFsaqW5DGjtxA4gYk1ZEH68#S1Z1McmUd@ph8Gbhn#"


class TestCCPaymentClass(unittest.TestCase):
    cp = ccpayment.CCPaymentClass(app_id, app_secret)

    # create api deposit order
    data, is_verify = cp.create_order(token_id='0912e09a-d8e2-41d7-a0bc-a25530892988',
                                      amount='6',
                                      merchant_order_id=str(int(time.time())),
                                      fiat_currency='USD')
    if is_verify:
        print("TestCreateOrder: verify success")
    else:
        print("TestCreateOrder: verify error")
    print("TestCreateOrder:", data)

    # webhook validate
    # data_str body;
    # timestamp header
    # signature header
    if cp.webhook(data_str='', timestamp='', signature=''):
        print('TestWebhookValidate: verify success')
    else:
        print('TestWebhookValidate: verify error')

    # get checkout url
    data, is_verify = cp.checkout_url(amount='12',
                                      merchant_order_id=str(int(time.time())),
                                      valid_timestamp=300,
                                      product_name='product_name',
                                      return_url='return_url')
    if is_verify:
        print("TestCheckoutUrl: verify success")
    else:
        print("TestCheckoutUrl :verify error")
    print("TestCheckoutUrl:", data)

    # get support tokens
    data, is_verify = cp.get_support_token()
    if is_verify:
        print("TestSupportTokens: verify success")
    else:
        print("TestSupportTokens: verify error")
    print("TestSupportTokens:", data)

    # get token chains
    data, is_verify = cp.get_token_chain(token_id='8e5741cf-6e51-4892-9d04-3d40e1dd0128')
    if is_verify:
        print("TestTokenChain: verify success")
    else:
        print("TestTokenChain: verify error")
    print("TestTokenChain:", data)

    # get token rate
    data, is_verify = cp.get_token_rate(token_id='8e5741cf-6e51-4892-9d04-3d40e1dd0128',
                                        amount='12')
    if is_verify:
        print("TestTokenRate: verify success")
    else:
        print("TestTokenRate: verify error")
    print("TestTokenRate:", data)

    # create api withdraw order
    data, is_verify = cp.withdraw(token_id='8e5741cf-6e51-4892-9d04-3d40e1dd0128',
                                  address='9486792',
                                  value='1',
                                  merchant_order_id=str(int(time.time())))
    if is_verify:
        print("TestApiWithdraw: verify success")
    else:
        print("TestApiWithdraw: verify error")
    print("TestApiWithdraw:", data)

    # check user
    data, is_verify = cp.check_user(c_id='9486792')
    if is_verify:
        print("TestCheckUser: verify success")
    else:
        print("TestCheckUser: verify error")
    print("TestCheckUser:", data)

    # get token assets
    data, is_verify = cp.assets(token_id='8e5741cf-6e51-4892-9d04-3d40e1dd0128')
    if is_verify:
        print("TestTokenAsset: verify success")
    else:
        print("TestTokenAsset :verify error")
    print("TestTokenAsset:", data)

    # get network fee
    data, is_verify = cp.network_fee(token_id='f137d42c-f3a6-4f23-9402-76f0395d0cfe')
    if is_verify:
        print("TestNetworkFee: verify success")
    else:
        print("TestNetworkFee: verify error")
    print("TestNetworkFee:", data)


if __name__ == '__main__':
    unittest.main()
