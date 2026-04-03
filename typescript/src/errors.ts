export class APIError extends Error {
  public readonly code: number;
  public readonly message: string;

  constructor(code: number, message: string) {
    super(`API Error ${code}: ${message}`);
    this.code = code;
    this.message = message;
    Object.setPrototypeOf(this, APIError.prototype);
  }
}
