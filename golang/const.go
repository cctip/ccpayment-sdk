package golang

const (
	Host = "https://admin.ccpayment.com" // produce
	//Host               = "https://ebc65a6dtestpaymentadmin.ccpayment.com" // dev
	CheckoutUrl        = Host + "/ccpayment/v1/concise/url/get"
	CreateOrderUrl     = Host + "/ccpayment/v1/bill/create"
	GetSupportTokenUrl = Host + "/ccpayment/v1/support/token"
	GetTokenChainUrl   = Host + "/ccpayment/v1/token/chain"
	GetTokenRateUrl    = Host + "/ccpayment/v1/token/rate"
	WithdrawApiUrl     = Host + "/ccpayment/v1/withdraw"
	CheckUserUrl       = Host + "/ccpayment/v1/check/user"
	AssetsUrl          = Host + "/ccpayment/v1/assets"
	NetworkFeeUrl      = Host + "/ccpayment/v1/network/fee"
	ApiOrderInfoUrl    = Host + "/ccpayment/v1/bill/info"

	AppIdHeaderKey     = "Appid"
	SignHeaderKey      = "Sign"
	TimestampHeaderKey = "Timestamp"

	appId     = "6cfakPeNnTi0YoHLpbqVLqFw2A#R5EGPH3lJfX75H7x#SO#PXU"
	appSecret = `5SQOObrXo#t82ZsVYfrI02Dn@5MLVPX1Fr1gcpCVS2Ca8ClbAgi@AMp64j0pWO#4jyC#zdS3kis1UaNFsaqW5DGjtxA4gYk1ZEH68#S1Z1McmUd@ph8Gbhn#`
)
