<?php

namespace CCPayment;

use Exception;

class APIError extends Exception
{
    private int $code;

    public function __construct(int $code, string $message)
    {
        parent::__construct("API Error {$code}: {$message}");
        $this->code = $code;
    }

    public function getErrorCode(): int
    {
        return $this->code;
    }
}
