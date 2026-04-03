# CCPayment SDK Code Generator

AI驱动的多语言SDK代码生成器，从API文档自动生成强类型的SDK代码。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Languages](https://img.shields.io/badge/languages-7-blue.svg)](#支持的语言)

## 特性

✨ **AI驱动代码生成** - 使用AI Agent自动从API文档生成SDK代码  
🔒 **强类型支持** - 为所有语言生成强类型定义和验证  
🎯 **模块化设计** - 支持生成完整SDK或单个模块  
🌍 **多语言支持** - 支持7种主流编程语言  
🚀 **开箱即用** - 生成的代码包含测试程序和完整文档  
🔧 **灵活配置** - 支持自定义输出路径和代理配置

## 快速开始

### 导入到AI Agent

本项目使用 **Skills 机制**，需要先将 `.skills/SKILL.md` 文件导入到您的 AI Agent 中：

**步骤1: 导入技能文件**

在 AI Agent 中使用以下方式导入：

```
@(ccpayment-sdk/.skills/SKILL.md)
```

**步骤2: 让AI Agent学习**

导入后，AI Agent会自动读取并学习 `SKILL.md` 中的代码生成规则。

### 使用方法

学习完成后，在AI Agent中输入以下命令即可生成SDK代码：

```bash
# 生成完整的Golang SDK
ccpayment.sdk.codegen golang

# 只生成商家资产模块
ccpayment.sdk.codegen golang merchant-assets

# 生成到指定路径
ccpayment.sdk.codegen python merchant-assets ./sdk/python
```

## 支持的语言

| 语言 | 状态 | 默认输出路径 |
|------|------|-------------|
| Golang | ✅ 已测试 | `generated/golang/` |
| Python | ✅ 已测试 | `generated/python/` |
| TypeScript | 🚧 开发中 | `generated/typescript/` |
| JavaScript | 🚧 开发中 | `generated/javascript/` |
| Java | 🚧 开发中 | `generated/java/` |
| PHP | 🚧 开发中 | `generated/php/` |
| Ruby | 🚧 开发中 | `generated/ruby/` |

## 支持的模块

SDK按功能模块组织，共13个模块：

### 商家端模块
- `basic-info` - 基础信息（代币、法币、链信息）
- `merchant-assets` - 商家资产查询
- `merchant-deposit` - 商家充值管理
- `merchant-withdraw` - 商家提现管理
- `merchant-batch-withdraw` - 批量提现管理

### 用户端模块
- `user-assets` - 用户资产查询
- `user-deposit` - 用户充值管理
- `user-withdraw` - 用户提现管理
- `user-transfer` - 用户转账管理

### 其他模块
- `orders` - 订单管理
- `checkout` - 收银台管理
- `swap` - 换币功能
- `utilities` - 工具接口

## 使用示例

### Golang SDK

```go
package main

import (
    "context"
    "log"
)

func main() {
    // 创建客户端
    client := NewClient("your_app_id", "your_app_secret")
    
    // 获取全部资产
    resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    
    // 打印资产
    for _, asset := range resp.Assets {
        log.Printf("%s: %s", asset.CoinSymbol, asset.Available)
    }
}
```

### Python SDK

```python
from client import Client

# 创建客户端
client = Client("your_app_id", "your_app_secret")

# 获取全部资产
response = client.merchant_assets().get_app_coin_asset_list()

# 打印资产
for asset in response.assets:
    print(f"{asset.coin_symbol}: {asset.available}")
```

## 生成的代码结构

每个SDK都包含以下组件：

```
generated/<language>/
├── client.*           # 核心客户端
├── signature.*        # HMAC-SHA256签名生成
├── models.*           # 强类型数据模型
├── errors.*           # 错误处理
├── <module>.*         # 各个模块的服务类
├── test_main.*        # 测试程序
└── README.md          # 使用文档
```

## 配置说明

### API认证

所有API请求需要以下认证信息：

1. 登录 [CCPayment开发者控制台](https://console.ccpayment.com)
2. 获取 `App ID` 和 `App Secret`
3. 配置IP白名单（必须）

### HTTP代理

生成的SDK支持HTTP代理配置：

**Golang:**
```go
proxyURL, _ := url.Parse("http://127.0.0.1:10808")
httpClient := &http.Client{
    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
}
client.SetHTTPClient(httpClient)
```

**Python:**
```python
client.set_proxy("http://127.0.0.1:10808")
```

## 测试

生成的SDK包含完整的测试程序：

```bash
# Golang
cd generated/golang
go run .

# Python
cd generated/python
python3 test_main.py
```

## 文档

- 📖 [API文档](.rules/api/README.md) - 完整的API接口文档
- 🛠️ [代码生成器说明](.skills/SKILL.md) - 代码生成器的详细说明
- 🌐 [官方文档](https://doc.ccpayment.com) - CCPayment官方文档

## 常见问题

### 1. 如何使用代码生成器？

**步骤1: 导入技能文件**

将 `.skills/SKILL.md` 文件复制到您使用的 AI Agent 的 skills 目录中：

- Windows: `%USERPROFILE%\.{agent}\skills\ccpayment.sdk.codegen.md`
- macOS/Linux: `~/.{agent}/skills/ccpayment.sdk.codegen.md`

支持的 AI Agent: Windsurf, Cursor, Cline, Continue.dev, Aider, Gemini, Codex, DeepSeek

**步骤2: 让AI Agent学习**

导入后，AI Agent会自动读取并学习 `SKILL.md` 中的代码生成规则。

**步骤3: 使用命令生成代码**

学习完成后，在AI Agent中输入：
```bash
ccpayment.sdk.codegen golang merchant-assets
```

### 2. 生成的代码可以直接使用吗？

是的！生成的代码包含完整的实现、测试程序和文档，可以直接运行。只需配置API凭证即可。

### 3. 如何配置IP白名单？

登录 [开发者控制台](https://console.ccpayment.com/developer/config)，在"IP白名单"设置中添加您的公网IP地址。

### 4. 支持哪些AI Agent？

目前支持以下AI Agent：
- **Windsurf** - Codeium的AI编程助手
- **Cursor** - AI代码编辑器
- **Cline** - VS Code的Claude扩展
- **Continue.dev** - 开源AI编程助手
- **Aider** - 命令行AI编程工具
- **Gemini** - Google的AI代码助手
- **Codex** - GitHub Copilot/OpenAI Codex
- **DeepSeek** - DeepSeek Coder AI助手

### 5. 可以自定义输出路径吗？

可以！使用命令时添加路径参数：
```bash
ccpayment.sdk.codegen golang merchant-assets ./custom/path
```

## 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 相关链接

- 🌐 [CCPayment官网](https://ccpayment.com)
- 📚 [API文档](https://doc.ccpayment.com)
- 💻 [开发者控制台](https://console.ccpayment.com)
- 🐛 [问题反馈](https://github.com/yourusername/ccpayment-sdk/issues)

---

**Made with ❤️ by CCPayment Team**
