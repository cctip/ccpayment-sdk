---
name: ccpayment.sdk.codegen
version: 2.1.0
description: CCPayment SDK Code Generator - Generate strongly-typed SDK code from API documentation
homepage: https://ccpayment.com
metadata:
  {
    "ccpayment":
      {
        "category": "sdk-generator",
        "api_base": "https://ccpayment.com/ccpayment/v2/",
        "api_docs": ".rules/api/",
        "supported_languages": ["golang", "typescript", "python", "java", "php"]
      },
  }
---

# CCPayment SDK 代码生成器

## 使用方式

### 前置步骤：导入并学习

在使用代码生成器之前，需要先将本文件导入到 AI Agent 中：

**步骤1: 导入技能文件**

在 AI Agent 中使用以下方式导入本文件：

```
@(ccpayment-sdk/.skills/SKILL.md)
```

**步骤2: 让AI Agent学习**

导入后，AI Agent会自动读取并学习本文档中的代码生成规则、类型映射、验证规则等内容。

---

### 生成代码命令

学习完成后，使用以下命令生成SDK代码：

```bash
ccpayment.sdk.codegen <language> [module] [path]
```

**参数说明：**
- `<language>` (必填): 目标语言，支持 `golang`, `typescript`, `python`, `java`, `php`, `ruby`, `javascript`
- `[module]` (可选): 指定生成的模块，不指定则生成所有模块
- `[path]` (可选): 指定生成代码的输出路径，不指定则使用默认路径

**使用示例：**
```bash
# 生成所有模块到默认路径
ccpayment.sdk.codegen golang

# 只生成商家资产模块到默认路径
ccpayment.sdk.codegen golang merchant-assets

# 生成所有模块到指定路径
ccpayment.sdk.codegen golang . /custom/output/path

# 只生成商家资产模块到指定路径
ccpayment.sdk.codegen golang merchant-assets /custom/output/path

# 生成Python SDK到自定义路径
ccpayment.sdk.codegen python merchant-assets ./sdk/python
```

**默认输出路径：**
- Golang: `generated/golang/`
- Python: `generated/python/`
- TypeScript: `generated/typescript/`
- Java: `generated/java/`
- PHP: `generated/php/`

**支持的模块：**
- `basic-info` - 基础信息模块（代币、法币、链信息）
- `merchant-assets` - 商家资产模块
- `merchant-deposit` - 商家充值模块
- `merchant-withdraw` - 商家提现模块
- `merchant-batch-withdraw` - 商家批量提现模块
- `user-assets` - 用户资产模块
- `user-deposit` - 用户充值模块
- `user-withdraw` - 用户提现模块
- `user-transfer` - 用户转账模块
- `orders` - 订单模块
- `checkout` - 收银台模块
- `swap` - 换币模块
- `utilities` - 工具接口

## 工作原理

当AI助手接收到命令时，根据参数执行不同的生成策略：

### 完整生成模式
**命令**: `ccpayment.sdk.codegen golang` 或 `ccpayment.sdk.codegen golang . /custom/path`

生成所有13个模块的完整SDK，包括所有核心组件和服务。

### 模块化生成模式
**命令**: `ccpayment.sdk.codegen golang merchant-assets` 或 `ccpayment.sdk.codegen golang merchant-assets /custom/path`

只生成指定模块的代码，包括：
- 核心组件（Client, Signature, Models, Errors）
- 指定模块的Service
- 该模块相关的数据模型
- 测试程序（只测试该模块）

### 路径参数处理
- 如果提供了 `[path]` 参数，代码将生成到指定路径
- 如果未提供 `[path]` 参数，代码将生成到默认路径 `generated/<language>/`
- 路径可以是相对路径或绝对路径
- 如果目标路径不存在，将自动创建

### 执行步骤

#### 步骤1: 读取API文档

**基础文件（总是需要）：**
- `.rules/api/README.md` - 认证方式、Base URL
- `.rules/api/appendix.md` - 数据类型、错误码、验证规则

**模块文件映射表：**

| 模块名 | 文档文件 | 生成的服务类 |
|:------|:---------|:------------|
| `basic-info` | `01-basic-info.md` | `BasicInfoService` |
| `merchant-assets` | `02-merchant-assets.md` | `MerchantAssetsService` |
| `merchant-deposit` | `03-merchant-deposit.md` | `MerchantDepositService` |
| `merchant-withdraw` | `04-merchant-withdraw.md` | `MerchantWithdrawService` |
| `merchant-batch-withdraw` | `05-merchant-batch-withdraw.md` | `MerchantBatchWithdrawService` |
| `user-assets` | `06-user-assets.md` | `UserAssetsService` |
| `user-deposit` | `07-user-deposit.md` | `UserDepositService` |
| `user-withdraw` | `08-user-withdraw.md` | `UserWithdrawService` |
| `user-transfer` | `09-user-transfer.md` | `UserTransferService` |
| `orders` | `10-orders.md` | `OrdersService` |
| `checkout` | `11-checkout.md` | `CheckoutService` |
| `swap` | `12-swap.md` | `SwapService` |
| `utilities` | `13-utilities.md` | `UtilitiesService` |

**读取策略：**
- 完整生成：读取所有13个模块文件
- 模块化生成：只读取指定模块的文件 + 基础文件

#### 步骤2: 解析API定义

从每个模块文档中提取：
- **接口路径**: 如 `POST /getCoinList`
- **请求参数表格**: 字段名、类型、必填、说明、验证规则
- **响应数据表格**: 字段名、类型、说明

#### 步骤3: 生成强类型代码

#### Golang 代码结构

**完整生成** (`ccpayment.sdk.codegen golang` 或 `ccpayment.sdk.codegen golang . /output/path`):
```
<output_path>/  (默认: generated/golang/)
├── client.go              # 主客户端（包含所有服务方法）
├── signature.go           # 签名生成
├── models.go              # 所有数据模型
├── errors.go              # 错误处理
├── basic_info.go          # 基础信息服务
├── merchant_assets.go     # 商家资产服务
├── merchant_deposit.go    # 商家充值服务
├── ... (其他10个服务)
├── test_main.go           # 完整测试程序
├── go.mod                 # Go模块配置
└── README.md              # 使用文档
```

**模块化生成** (`ccpayment.sdk.codegen golang merchant-assets` 或 `ccpayment.sdk.codegen golang merchant-assets /output/path`):
```
<output_path>/  (默认: generated/golang/)
├── client.go              # 主客户端（只包含MerchantAssets方法）
├── signature.go           # 签名生成
├── models.go              # 只包含merchant-assets相关的数据模型
├── errors.go              # 错误处理
├── merchant_assets.go     # 商家资产服务
├── test_main.go           # merchant-assets测试程序
├── go.mod                 # Go模块配置
└── README.md              # 使用文档（针对该模块）
```

**自定义路径示例：**
```bash
# 生成到项目的sdk目录
ccpayment.sdk.codegen golang merchant-assets ./sdk/golang

# 生成到绝对路径
ccpayment.sdk.codegen python . /Users/username/projects/my-sdk/python
```

#### 核心组件实现

**1. Client (client.go)**
```go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

const (
    DefaultBaseURL = "https://ccpayment.com/ccpayment/v2"
)

type Client struct {
    appID      string
    appSecret  string
    baseURL    string
    httpClient *http.Client
}

func NewClient(appID, appSecret string) *Client {
    return &Client{
        appID:      appID,
        appSecret:  appSecret,
        baseURL:    DefaultBaseURL,
        httpClient: &http.Client{},
    }
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
    c.httpClient = httpClient
}

// 返回各个服务（根据生成的模块动态添加）
// 如果只生成merchant-assets模块，则只包含此方法
func (c *Client) MerchantAssets() *MerchantAssetsService {
    return &MerchantAssetsService{client: c}
}

// 完整生成时，包含所有服务方法：
// func (c *Client) BasicInfo() *BasicInfoService
// func (c *Client) MerchantDeposit() *MerchantDepositService
// func (c *Client) UserAssets() *UserAssetsService
// ... 等等

// post 内部方法
func (c *Client) post(ctx context.Context, path string, req, resp interface{}) error {
    var bodyBytes []byte
    var err error

    if req != nil {
        bodyBytes, err = json.Marshal(req)
        if err != nil {
            return fmt.Errorf("marshal request: %w", err)
        }
    }

    sign, timestamp, err := c.generateSign(bodyBytes)
    if err != nil {
        return fmt.Errorf("generate sign: %w", err)
    }

    httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+path, bytes.NewReader(bodyBytes))
    if err != nil {
        return fmt.Errorf("create request: %w", err)
    }

    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Appid", c.appID)
    httpReq.Header.Set("Sign", sign)
    httpReq.Header.Set("Timestamp", timestamp)

    httpResp, err := c.httpClient.Do(httpReq)
    if err != nil {
        return fmt.Errorf("do request: %w", err)
    }
    defer httpResp.Body.Close()

    respBody, err := io.ReadAll(httpResp.Body)
    if err != nil {
        return fmt.Errorf("read response: %w", err)
    }

    var apiResp Response
    if err := json.Unmarshal(respBody, &apiResp); err != nil {
        return fmt.Errorf("unmarshal response: %w", err)
    }

    if apiResp.Code != 10000 {
        return &APIError{
            Code:    apiResp.Code,
            Message: apiResp.Msg,
        }
    }

    if resp != nil {
        dataBytes, err := json.Marshal(apiResp.Data)
        if err != nil {
            return fmt.Errorf("marshal data: %w", err)
        }
        if err := json.Unmarshal(dataBytes, resp); err != nil {
            return fmt.Errorf("unmarshal data: %w", err)
        }
    }

    return nil
}
```

**2. Signature (signature.go)**
```go
package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "time"
)

func (c *Client) generateSign(bodyBytes []byte) (sign, timestamp string, err error) {
    timestamp = strconv.FormatInt(time.Now().Unix(), 10)
    
    signText := c.appID + timestamp
    if len(bodyBytes) > 0 {
        signText += string(bodyBytes)
    }
    
    mac := hmac.New(sha256.New, []byte(c.appSecret))
    mac.Write([]byte(signText))
    sign = hex.EncodeToString(mac.Sum(nil))
    
    return sign, timestamp, nil
}
```

**3. Models (models.go)**

为每个API生成请求和响应结构：

```go
package main

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

type APIError struct {
    Code    int
    Message string
}

func (e *APIError) Error() string {
    return e.Message
}

// 资产相关模型
type Asset struct {
    CoinID     uint64 `json:"coinId"`
    CoinSymbol string `json:"coinSymbol"`
    Available  string `json:"available"`
}

type GetAppCoinAssetListResponse struct {
    Assets []Asset `json:"assets"`
}

type GetAppCoinAssetRequest struct {
    CoinID uint64 `json:"coinId"`
}

type GetAppCoinAssetResponse struct {
    Asset Asset `json:"asset"`
}
```

**4. Services (merchant_assets.go)**

```go
package main

import "context"

type MerchantAssetsService struct {
    client *Client
}

func (s *MerchantAssetsService) GetAppCoinAssetList(ctx context.Context) (*GetAppCoinAssetListResponse, error) {
    var result GetAppCoinAssetListResponse
    err := s.client.post(ctx, "/getAppCoinAssetList", nil, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

func (s *MerchantAssetsService) GetAppCoinAsset(ctx context.Context, coinID uint64) (*GetAppCoinAssetResponse, error) {
    req := &GetAppCoinAssetRequest{CoinID: coinID}
    var result GetAppCoinAssetResponse
    err := s.client.post(ctx, "/getAppCoinAsset", req, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}
```

**5. Test Program (test_main.go)**

生成的测试程序会根据指定的模块自动调整测试内容。

**示例：merchant-assets 模块测试**
```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "time"
)

func main() {
    // API凭证
    appID := "your_app_id"
    appSecret := "your_app_secret"

    // 创建客户端
    client := NewClient(appID, appSecret)

    // 可选：配置HTTP代理
    // proxyURL, _ := url.Parse("http://127.0.0.1:10808")
    // httpClient := &http.Client{
    //     Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
    //     Timeout: 30 * time.Second,
    // }
    // client.SetHTTPClient(httpClient)

    fmt.Println("=== 商家资产模块测试 ===")
    
    // 测试1: 获取全部资产
    fmt.Println("\n【测试1】获取全部资产...")
    resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("✓ 成功获取资产列表，共 %d 种代币\n", len(resp.Assets))
    
    // 显示前10个
    for i, asset := range resp.Assets {
        if i >= 10 {
            break
        }
        fmt.Printf("%d. %s: %s\n", i+1, asset.CoinSymbol, asset.Available)
    }

    // 测试2: 获取单个币资产
    if len(resp.Assets) > 0 {
        fmt.Println("\n【测试2】获取单个币资产...")
        coinID := resp.Assets[0].CoinID
        assetResp, err := client.MerchantAssets().GetAppCoinAsset(context.Background(), coinID)
        if err != nil {
            log.Printf("获取失败: %v", err)
        } else {
            fmt.Printf("✓ %s 可用余额: %s\n", assetResp.Asset.CoinSymbol, assetResp.Asset.Available)
        }
    }
    
    fmt.Println("\n=== 测试完成 ===")
}
```

**不同模块会生成不同的测试内容：**
- `basic-info`: 测试获取代币列表、法币列表、链列表
- `merchant-deposit`: 测试创建充值地址、查询充值记录
- `user-transfer`: 测试用户转账、批量转账
- 等等...

#### 步骤4: 类型映射

从API文档类型映射到目标语言类型：

| 文档类型 | Golang | TypeScript | Python | Java |
|:--------|:-------|:-----------|:-------|:-----|
| `uint64` | `uint64` | `number` | `int` | `Long` |
| `int64` | `int64` | `number` | `int` | `Long` |
| `int32` | `int32` | `number` | `int` | `Integer` |
| `uint32` | `uint32` | `number` | `int` | `Integer` |
| `string` | `string` | `string` | `str` | `String` |
| `bool` | `bool` | `boolean` | `bool` | `Boolean` |
| `Array` | `[]T` | `T[]` | `List[T]` | `List<T>` |
| `Object` | `struct` | `interface` | `dict` | `Map` |

#### 步骤5: 验证规则映射

从API文档验证规则生成代码验证逻辑：

| 文档规则 | Golang validate tag |
|:---------|:-------------------|
| 必填 | `validate:"required"` |
| 长度≥1 | `validate:"min=1"` |
| 长度3-64 | `validate:"min=3,max=64"` |
| ≥1 | `validate:"gte=1"` |
| ≤100 | `validate:"lte=100"` |
| 邮箱格式 | `validate:"email"` |
| URI格式 | `validate:"uri"` |
| 最大50项 | `validate:"max=50"` |

## 重要注意事项

### 路径参数处理

⚠️ **路径参数规则**：

**路径格式：**
- 相对路径：`./sdk/golang`, `../output/python`
- 绝对路径：`/Users/username/projects/sdk`, `C:/projects/sdk`

**路径处理：**
- 如果目标路径不存在，自动创建目录
- 如果目标路径已存在文件，会覆盖现有文件
- 建议使用相对路径以保持项目可移植性

**示例：**
```bash
# 相对路径（推荐）
ccpayment.sdk.codegen golang merchant-assets ./sdk/golang

# 绝对路径
ccpayment.sdk.codegen python . /Users/oks/projects/ccpayment-sdk/python

# 默认路径（不指定path参数）
ccpayment.sdk.codegen golang merchant-assets
```

### 包名处理

⚠️ **关键问题**：生成测试代码时，所有文件必须使用相同的包名。

**正确做法**：
- 如果生成独立测试程序，所有文件使用 `package main`
- 如果生成SDK库，所有文件使用 `package ccpayment`，测试文件使用 `package ccpayment_test`

**错误示例**（会导致编译失败）：
```
client.go: package main
models.go: package main
example_test.go: package ccpayment  // ❌ 包名不一致
```

### Go工作区配置

如果遇到 "module is not one of the workspace modules" 错误：

```bash
cd generated/golang
go work use .
```

### HTTP代理配置

支持通过自定义HTTP客户端配置代理：

```go
proxyURL, _ := url.Parse("http://127.0.0.1:10808")
httpClient := &http.Client{
    Transport: &http.Transport{
        Proxy: http.ProxyURL(proxyURL),
    },
    Timeout: 30 * time.Second,
}
client.SetHTTPClient(httpClient)
```

### IP白名单

⚠️ **首次测试前必须配置**

API调用需要在CCPayment开发者后台配置IP白名单：

1. 登录 https://console.ccpayment.com/developer/config
2. 找到"IP白名单"设置
3. 添加您的公网IP地址
4. 保存配置

如果未配置，会收到错误：
```
ip not in whitelist, please check the ip whitelist settings on the developer page
```

## 生成README.md

为生成的SDK创建使用文档：

```markdown
# CCPayment Go SDK

## 安装

\`\`\`bash
go get github.com/yourusername/ccpayment-sdk/generated/golang
\`\`\`

## 快速开始

\`\`\`go
package main

import (
    "context"
    "log"
)

func main() {
    client := NewClient("your_app_id", "your_app_secret")
    
    // 获取全部资产
    resp, err := client.MerchantAssets().GetAppCoinAssetList(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    
    for _, asset := range resp.Assets {
        log.Printf("%s: %s\n", asset.CoinSymbol, asset.Available)
    }
}
\`\`\`

## 配置HTTP代理

\`\`\`go
import (
    "net/http"
    "net/url"
    "time"
)

proxyURL, _ := url.Parse("http://127.0.0.1:10808")
httpClient := &http.Client{
    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
    Timeout: 30 * time.Second,
}
client.SetHTTPClient(httpClient)
\`\`\`

## API认证

所有API请求都需要在Header中包含：
- \`Appid\`: 应用ID
- \`Timestamp\`: 当前时间戳（秒）
- \`Sign\`: HMAC-SHA256签名

SDK会自动处理签名生成。

## 错误处理

\`\`\`go
resp, err := client.MerchantAssets().GetAppCoinAssetList(ctx)
if err != nil {
    if apiErr, ok := err.(*APIError); ok {
        log.Printf("API错误: code=%d, message=%s", apiErr.Code, apiErr.Message)
    } else {
        log.Printf("其他错误: %v", err)
    }
    return
}
\`\`\`
```

## 测试验证清单

生成代码后，按以下步骤验证：

1. ✅ **检查输出路径**
   ```bash
   # 确认文件已生成到指定路径
   ls -la <output_path>
   
   # 默认路径示例
   ls -la generated/golang
   
   # 自定义路径示例
   ls -la ./sdk/python
   ```

2. ✅ **编译检查**
   ```bash
   # Golang
   cd <output_path>
   go build .
   
   # Python
   cd <output_path>
   python3 -m py_compile *.py
   ```

3. ✅ **配置IP白名单**
   - 登录开发者后台
   - 添加公网IP到白名单

4. ✅ **运行测试**
   ```bash
   # Golang
   cd <output_path>
   go run .
   
   # Python
   cd <output_path>
   python3 test_main.py
   ```

5. ✅ **验证输出**
   
   **完整生成时**：
   - 测试所有13个模块的主要接口
   - 验证所有数据类型正确解析
   
   **模块化生成时**（如 `merchant-assets`）：
   - ✓ 成功获取资产列表，共 X 种代币
   - ✓ 成功获取单个币资产信息
   - 确认数据类型正确解析
   - 验证错误处理正常

6. ✅ **代码质量检查**
   ```bash
   # Golang
   go fmt .
   go vet .
   
   # Python
   python3 -m pylint *.py  # 需要安装pylint
   python3 -m black *.py   # 需要安装black
   ```

## 常见问题

### 1. 包名冲突
**问题**: `found packages main and ccpayment`
**解决**: 删除或重命名冲突的测试文件，确保所有文件使用相同包名

### 2. IP白名单错误
**问题**: `ip not in whitelist`
**解决**: 在开发者后台配置IP白名单

### 3. 工作区错误
**问题**: `module is not one of the workspace modules`
**解决**: 运行 `go work use .`

### 4. 代理连接失败
**问题**: 无法通过代理访问API
**解决**: 检查代理地址和端口，确保代理服务正常运行

## 参考

- API文档: `.rules/api/README.md`
- Base URL: `https://ccpayment.com/ccpayment/v2/`
- 官网: https://ccpayment.com
- 开发者控制台: https://console.ccpayment.com
