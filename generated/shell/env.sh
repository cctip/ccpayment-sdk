#!/bin/bash
# CCPayment API - Environment Configuration
# Usage: source env.sh

# API Base URL
export BASE_URL="https://ccpayment.com/ccpayment/v2"

# Your API Credentials (replace with your actual credentials)
export APP_ID="your_app_id"
export APP_SECRET="your_app_secret"

# Colors for output
export GREEN='\033[0;32m'
export BLUE='\033[0;34m'
export NC='\033[0m' # No Color

# Generate HMAC-SHA256 signature
# Usage: generate_sign '{"json":"body"}'
generate_sign() {
    local body="${1:-}"
    local timestamp=$(date +%s)
    local sign_text="${APP_ID}${timestamp}${body}"
    
    # Generate HMAC-SHA256 signature
    if command -v openssl &> /dev/null; then
        echo -n "$sign_text" | openssl dgst -sha256 -hmac "$APP_SECRET" 2>/dev/null | sed 's/^.* //'
    else
        echo "Error: openssl is required for signature generation" >&2
        return 1
    fi
}

# Export function for use in other scripts
export -f generate_sign

echo -e "${GREEN}✓${NC} Environment configured"
echo "Base URL: ${BASE_URL}"
echo "App ID: ${APP_ID}"
echo ""
echo "Available functions:"
echo "  generate_sign '{json_body}'  - Generate API signature"
