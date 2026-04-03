class APIError extends Error {
  constructor(code, message) {
    super(`API Error ${code}: ${message}`);
    this.code = code;
    this.message = message;
  }
}

module.exports = { APIError };
