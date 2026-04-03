# 附录

## 数据类型说明

- **uint64**: 无符号64位整数
- **int64**: 有符号64位整数
- **int32**: 有符号32位整数
- **uint32**: 无符号32位整数
- **string**: 字符串
- **bool**: 布尔值
- **Array**: 数组
- **Object**: 对象

## 状态枚举

### 订单状态

- **Pending**: 待支付
- **Success**: 成功
- **Failed**: 失败
- **Expired**: 已过期

### 提现类型

- **Network**: 网络提现
- **Cwallet**: CCWallet提现

### Hosted步骤

- **SelectCoin**: 选择代币
- **Pay**: 支付中
- **Completed**: 已完成

## 验证规则说明

- **长度验证**: min_len（最小长度）、max_len（最大长度）
- **数值验证**: gte（大于等于）、lte（小于等于）、gt（大于）、lt（小于）
- **数组验证**: min_items（最小项数）、max_items（最大项数）
- **格式验证**: email（邮箱格式）、uri（URI格式）

## 注意事项

1. 所有时间戳均为Unix时间戳（秒）
2. 所有金额字段均为字符串类型，避免精度丢失
3. 分页查询使用nextId进行游标分页
4. 默认查询时间范围为90天
5. 批量操作最大支持500项
6. 所有接口均需要签名验证

## 签名算法

### 签名生成步骤

1. 将请求参数按照字母顺序排序
2. 拼接成 key1=value1&key2=value2 格式
3. 在末尾追加 AppSecret
4. 对整个字符串进行 SHA256 加密
5. 将加密结果转为大写

### 示例代码

```go
// Go 示例
func GenerateSign(params map[string]string, appSecret string) string {
    keys := make([]string, 0, len(params))
    for k := range params {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    var builder strings.Builder
    for _, k := range keys {
        builder.WriteString(k)
        builder.WriteString("=")
        builder.WriteString(params[k])
        builder.WriteString("&")
    }
    builder.WriteString(appSecret)
    
    hash := sha256.Sum256([]byte(builder.String()))
    return strings.ToUpper(hex.EncodeToString(hash[:]))
}
```

## 错误处理

### 通用错误码

| 错误码 | 说明 |
|--------|------|
| 10000 | 成功 |
| 11000 | 参数错误 |
| 11001 | 缺少必填参数 |
| 11002 | 参数格式错误 |
| 11003 | 参数值超出范围 |
| 11004 | 无效的参数值 |
| 11005 | 签名验证失败 |
| 11006 | 时间戳过期 |
| 12000 | 系统错误 |
| 12001 | 服务暂时不可用 |
| 13000 | 权限错误 |
| 13001 | AppId不存在 |
| 13002 | AppId已禁用 |
| 14000 | 余额不足 |
| 15000 | 订单不存在 |
| 15001 | 订单状态错误 |
| 15002 | 订单已过期 |
| 16000 | 地址无效 |
| 16001 | 链不支持 |
| 16002 | 代币不支持 |

完整错误码列表请参考proto文件中的Code枚举定义。

## Webhook通知

### 通知格式

所有Webhook通知均采用POST方式，Content-Type为application/json。

### 通知签名验证

1. 从Header中获取 `X-CC-Sign` 字段
2. 将请求Body与WebhookSecret拼接
3. 对拼接后的字符串进行SHA256加密
4. 将加密结果转为大写并与X-CC-Sign比对

### 通知重试机制

- 如果通知失败，系统会进行重试
- 重试间隔：1分钟、5分钟、15分钟、30分钟、1小时、2小时
- 最多重试6次

### 响应要求

- 接收方需在5秒内返回HTTP 200状态码
- 返回内容可以为空或任意内容
- 非200状态码视为通知失败，会触发重试
