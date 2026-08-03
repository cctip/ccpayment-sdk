module CCPayment
  class APIError < StandardError
    attr_reader :code

    def initialize(code, message)
      @code = code
      super("API Error #{code}: #{message}")
    end
  end
end
