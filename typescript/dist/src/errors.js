"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.APIError = void 0;
class APIError extends Error {
    constructor(code, message) {
        super(`API Error ${code}: ${message}`);
        this.code = code;
        this.message = message;
        Object.setPrototypeOf(this, APIError.prototype);
    }
}
exports.APIError = APIError;
//# sourceMappingURL=errors.js.map