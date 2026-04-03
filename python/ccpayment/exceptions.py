"""CCPayment SDK Exceptions"""


class APIError(Exception):
    """API Error"""

    def __init__(self, code: int, message: str):
        self.code = code
        self.message = message
        super().__init__(f"API Error {code}: {message}")
