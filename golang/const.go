package golang

const (
	Host = "https://admin.ccpayment.com" // produce
	//Host               = "https://ebc65a6dtestpaymentadmin.cwallet.com" // dev
	CheckoutUrl        = Host + "/ccpayment/v1/concise/url/get"
	CreateOrderUrl     = Host + "/ccpayment/v1/bill/create"
	GetSupportTokenUrl = Host + "/ccpayment/v1/support/token"
	GetTokenChainUrl   = Host + "/ccpayment/v1/token/chain"
	GetTokenRateUrl    = Host + "/ccpayment/v1/token/rate"
	WithdrawApiUrl     = Host + "/ccpayment/v1/withdraw"
	CheckUserUrl       = Host + "/ccpayment/v1/check/user"
	AssetsUrl          = Host + "/ccpayment/v1/assets"
	NetworkFeeUrl      = Host + "/ccpayment/v1/network/fee"

	AppIdHeaderKey     = "Appid"
	SignHeaderKey      = "Sign"
	TimestampHeaderKey = "Timestamp"
)
