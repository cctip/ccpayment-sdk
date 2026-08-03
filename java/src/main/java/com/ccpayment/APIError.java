package com.ccpayment;

public class APIError extends Exception {
    private final int code;
    private final String message;

    public APIError(int code, String message) {
        super("API Error " + code + ": " + message);
        this.code = code;
        this.message = message;
    }

    public int getCode() {
        return code;
    }

    @Override
    public String getMessage() {
        return message;
    }
}
