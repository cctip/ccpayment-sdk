require 'net/http'
require 'json'
require 'openssl'

module CCPayment
  class Client
    DEFAULT_BASE_URL = 'https://ccpayment.com/ccpayment/v2'

    attr_accessor :base_url

    def initialize(app_id, app_secret)
      @app_id = app_id
      @app_secret = app_secret
      @base_url = DEFAULT_BASE_URL
      @proxy_url = nil
    end

    def set_proxy(proxy_url)
      @proxy_url = proxy_url
    end

    def generate_sign(body)
      timestamp = Time.now.to_i.to_s
      sign_text = "#{@app_id}#{timestamp}#{body}"
      sign = OpenSSL::HMAC.hexdigest(OpenSSL::Digest.new('sha256'), @app_secret, sign_text)
      [sign, timestamp]
    end

    def post(path, data = {})
      body = data.to_json
      sign, timestamp = generate_sign(body)
      headers = {
        'Content-Type' => 'application/json',
        'Appid' => @app_id,
        'Sign' => sign,
        'Timestamp' => timestamp
      }
      uri = URI.parse("#{@base_url}#{path}")
      response = if @proxy_url
        proxy_uri = URI.parse(@proxy_url)
        http = Net::HTTP.new(uri.host, uri.port, proxy_uri.host, proxy_uri.port)
        http.use_ssl = uri.scheme == 'https'
        http.verify_mode = OpenSSL::SSL::VERIFY_PEER
        request = Net::HTTP::Post.new(uri.path, headers)
        request.body = body
        http.request(request)
      else
        http = Net::HTTP.new(uri.host, uri.port)
        http.use_ssl = uri.scheme == 'https'
        http.verify_mode = OpenSSL::SSL::VERIFY_PEER
        request = Net::HTTP::Post.new(uri.path, headers)
        request.body = body
        http.request(request)
      end
      result = JSON.parse(response.body)
      raise APIError.new(result['code'], result['msg']) unless result['code'] == 10000
      result['data'] || {}
    end

    def basic_info; Services::BasicInfoService.new(self); end
    def merchant_assets; Services::MerchantAssetsService.new(self); end
    def merchant_deposit; Services::MerchantDepositService.new(self); end
    def merchant_withdraw; Services::MerchantWithdrawService.new(self); end
    def merchant_batch_withdraw; Services::MerchantBatchWithdrawService.new(self); end
    def user_assets; Services::UserAssetsService.new(self); end
    def user_deposit; Services::UserDepositService.new(self); end
    def user_withdraw; Services::UserWithdrawService.new(self); end
    def user_transfer; Services::UserTransferService.new(self); end
    def orders; Services::OrdersService.new(self); end
    def swap; Services::SwapService.new(self); end
    def user_swap; Services::UserSwapService.new(self); end
    def utilities; Services::UtilitiesService.new(self); end
  end
end
