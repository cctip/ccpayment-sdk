# CCPayment SDK Code Generator

AI-powered multi-language SDK code generator that automatically generates strongly-typed SDK code from API documentation.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Languages](https://img.shields.io/badge/languages-7-blue.svg)](#supported-languages)

## Features

✨ **AI-Powered Code Generation** - Uses AI Agent to automatically generate SDK code from API documentation  
🔒 **Strong Type Support** - Generates strong type definitions and validation for all languages  
🎯 **Modular Design** - Supports generating complete SDKs or individual modules  
🌍 **Multi-Language Support** - Supports 7 mainstream programming languages  
🚀 **Ready to Use** - Generated code includes test programs and complete documentation  
🔧 **Flexible Configuration** - Supports custom output paths and proxy configuration

## Quick Start

### Import to AI Agent

This project uses the **Skills mechanism**, requiring you to first import the `.skills/SKILL.md` file into your AI Agent:

**Step 1: Import Skill File**

Import using the following method in your AI Agent:

```
@(ccpayment-sdk/.skills/SKILL.md)
```

**Step 2: Let AI Agent Learn**

After import, the AI Agent will automatically read and learn the code generation rules in `SKILL.md`.

### Usage

After learning is complete, enter the following commands in your AI Agent to generate SDK code:

```bash
# Generate complete Golang SDK
ccpayment.sdk.codegen golang

# Only generate merchant assets module
ccpayment.sdk.codegen golang merchant-assets

# Generate to specified path
ccpayment.sdk.codegen python merchant-assets ./sdk/python
```

### Shell Script Example

```bash
# Generate Shell script for merchant assets module
ccpayment.sdk.codegen shell merchant-assets

# Use after generation
source generated/shell/env.sh && ./generated/shell/merchant_assets.sh
```

## Supported Languages

| Language | Status | Default Output Path |
|------|------|-------------|
| Golang | ✅ Supported | `generated/golang/` |
| Python | ✅ Supported | `generated/python/` |
| TypeScript | ✅ Supported | `generated/typescript/` |
| JavaScript | ✅ Supported | `generated/javascript/` |
| Java | ✅ Supported | `generated/java/` |
| PHP | ✅ Supported | `generated/php/` |
| Ruby | ✅ Supported | `generated/ruby/` |
| Shell | ✅ Supported | `generated/shell/` |

## Supported Modules

The SDK is organized by functional modules, totaling 13 modules:

### Merchant Modules
- `basic-info` - Basic information (tokens, fiat currencies, chain information)
- `merchant-assets` - Merchant asset queries
- `merchant-deposit` - Merchant deposit management
- `merchant-withdraw` - Merchant withdrawal management
- `merchant-batch-withdraw` - Batch withdrawal management

### User Modules
- `user-assets` - User asset queries
- `user-deposit` - User deposit management
- `user-withdraw` - User withdrawal management
- `user-transfer` - User transfer management

### Other Modules
- `orders` - Order management
- `checkout` - Checkout management
- `swap` - Swap functionality
- `utilities` - Utility interfaces

## Usage Examples

### Golang SDK

```go
package main

import (
    "context"
    "log"
)

func main() {
    // Create client
    client := NewClient("your_app_id", "your_app_secret")
    
    // Get all assets
    resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    
    // Print assets
    for _, asset := range resp.Assets {
        log.Printf("%s: %s", asset.CoinSymbol, asset.Available)
    }
}
```

### Python SDK

```python
from client import Client

# Create client
client = Client("your_app_id", "your_app_secret")

# Get all assets
response = client.merchant_assets().get_app_coin_asset_list()

# Print assets
for asset in response.assets:
    print(f"{asset.coin_symbol}: {asset.available}")
```

### Shell Example

```bash
# Set environment variables
export BASE_URL="https://ccpayment.com/ccpayment/v2"
export APP_ID="your_app_id"
export APP_SECRET="your_app_secret"

# Generate signature and call API
generate_sign() {
    local body="$1"
    local timestamp=$(date +%s)
    local sign_text="${APP_ID}${timestamp}${body}"
    echo -n "$sign_text" | openssl dgst -sha256 -hmac "$APP_SECRET" | sed 's/^.* //'
}

# Get all assets
curl -X POST "${BASE_URL}/getAppCoinAssetList" \
  -H "Content-Type: application/json" \
  -H "Appid: ${APP_ID}" \
  -H "Sign: $(generate_sign '{}')" \
  -H "Timestamp: $(date +%s)" \
  -d '{}'
```

## Generated Code Structure

Each SDK contains the following components:

```
generated/<language>/
├── client.*           # Core client
├── signature.*        # HMAC-SHA256 signature generation
├── models.*           # Strong type data models
├── errors.*           # Error handling
├── <module>.*         # Service classes for each module
├── test_main.*        # Test programs
└── README.md          # Usage documentation
```

## Configuration

### API Authentication

All API requests require the following authentication information:

1. Log in to [CCPayment Developer Console](https://console.ccpayment.com)
2. Obtain `App ID` and `App Secret`
3. Configure IP whitelist (required)

### HTTP Proxy

Generated SDKs support HTTP proxy configuration:

**Golang:**
```go
proxyURL, _ := url.Parse("http://127.0.0.1:10808")
httpClient := &http.Client{
    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
}
client.SetHTTPClient(httpClient)
```

**Shell:**
```bash
source generated/shell/env.sh && ./generated/shell/merchant_assets.sh
```

## Testing

Generated SDKs include complete test programs:

```bash
# Golang
cd generated/golang
go run .

# Python
cd generated/python
python3 test_main.py

# Shell
source generated/shell/env.sh && ./generated/shell/merchant_assets.sh
```

## Documentation

- 📖 [API Documentation](.rules/api/README.md) - Complete API interface documentation
- 🛠️ [Code Generator Instructions](.skills/SKILL.md) - Detailed instructions for the code generator
- 🌐 [Official Documentation](https://doc.ccpayment.com) - CCPayment official documentation

## FAQ

### 1. How to use the code generator?

**Step 1: Import Skill File**

Copy the `.skills/SKILL.md` file to your AI Agent's skills directory:

- Windows: `%USERPROFILE%\.{agent}\skills\ccpayment.sdk.codegen.md`
- macOS/Linux: `~/.{agent}/skills/ccpayment.sdk.codegen.md`

Supported AI Agents: Windsurf, Cursor, Cline, Continue.dev, Aider, Gemini, Codex, DeepSeek

**Step 2: Let AI Agent Learn**

After import, the AI Agent will automatically read and learn the code generation rules in `SKILL.md`.

**Step 3: Use Commands to Generate Code**

After learning is complete, enter in your AI Agent:
```bash
ccpayment.sdk.codegen golang merchant-assets
```

### 2. Can the generated code be used directly?

Yes! The generated code includes complete implementation, test programs, and documentation, and can be run directly. Just configure the API credentials.

### 3. How to configure the IP whitelist?

Log in to the [Developer Console](https://console.ccpayment.com/developer/config) and add your public IP address in the "IP Whitelist" settings.

### 4. Which AI Agents are supported?

Currently supported AI Agents:
- **Windsurf** - Codeium's AI programming assistant
- **Cursor** - AI code editor
- **Cline** - VS Code's Claude extension
- **Continue.dev** - Open source AI programming assistant
- **Aider** - Command-line AI programming tool
- **Gemini** - Google's AI code assistant
- **Codex** - GitHub Copilot/OpenAI Codex
- **DeepSeek** - DeepSeek Coder AI assistant

### 5. Can the output path be customized?

Yes! Add the path parameter when using the command:
```bash
ccpayment.sdk.codegen golang merchant-assets ./custom/path

# Or generate Shell scripts to specified path
ccpayment.sdk.codegen shell merchant-assets ./shell-examples
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Related Links

- 🌐 [CCPayment Official Website](https://ccpayment.com)
- 📚 [API Documentation](https://doc.ccpayment.com)
- 💻 [Developer Console](https://console.ccpayment.com)
- 🐛 [Issue Feedback](https://github.com/yourusername/ccpayment-sdk/issues)

---

**Made with ❤️ by CCPayment Team**
