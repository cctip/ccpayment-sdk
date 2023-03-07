package golang

const (
	//Host           = "https://admin.ccpayment.com" // produce
	Host               = "https://ebc65a6dtestpaymentadmin.cwallet.com" // dev
	CheckoutUrl        = Host + "/ccpayment/v1/concise/url/get"
	CreateOrderUrl     = Host + "/ccpayment/v1/bill/create"
	GetSupportTokenUrl = Host + "/ccpayment/v1/support/token"
	GetTokenChainUrl   = Host + "/ccpayment/v1/token/chain"
	GetTokenRateUrl    = Host + "/ccpayment/v1/token/rate"

	AppIdHeaderKey     = "Appid"
	SignHeaderKey      = "Sign"
	TimestampHeaderKey = "Timestamp"
)
