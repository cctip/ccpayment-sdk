# Appendix

## Data Type Reference

- **uint64**: Unsigned 64-bit integer
- **int64**: Signed 64-bit integer
- **int32**: Signed 32-bit integer
- **uint32**: Unsigned 32-bit integer
- **string**: String
- **bool**: Boolean
- **Array**: Array
- **Object**: Object

## Status Enums

### Order Status

- **Pending**: Pending payment
- **Success**: Success
- **Failed**: Failed
- **Expired**: Expired

### Withdrawal Type

- **Network**: Network withdrawal
- **Cwallet**: CCWallet withdrawal

### Hosted Steps

- **SelectCoin**: Select token
- **Pay**: Paying
- **Completed**: Completed

## Validation Rules

- **Length validation**: min_len (minimum length), max_len (maximum length)
- **Numeric validation**: gte (greater than or equal), lte (less than or equal), gt (greater than), lt (less than)
- **Array validation**: min_items (minimum items), max_items (maximum items)
- **Format validation**: email (email format), uri (URI format)

## Important Notes

1. All timestamps are Unix timestamps (seconds)
2. All amount fields are string type to avoid precision loss
3. Pagination queries use nextId for cursor pagination
4. Default query time range is 90 days
5. Batch operations support up to 500 items
6. All interfaces require signature verification

## Signature Algorithm

### Signature Generation Steps

1. Sort request parameters alphabetically
2. Concatenate into key1=value1&key2=value2 format
3. Append AppSecret at the end
4. SHA256 encrypt the entire string
5. Convert the encrypted result to uppercase

### Example Code

```go
// Go Example
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

## Error Handling

### Common Error Codes

| Error Code | Description |
|--------|------|
| 10000 | Success |
| 11000 | Parameter error |
| 11001 | Missing required parameter |
| 11002 | Parameter format error |
| 11003 | Parameter value out of range |
| 11004 | Invalid parameter value |
| 11005 | Signature verification failed |
| 11006 | Timestamp expired |
| 12000 | System error |
| 12001 | Service temporarily unavailable |
| 13000 | Permission error |
| 13001 | AppId does not exist |
| 13002 | AppId disabled |
| 14000 | Insufficient balance |
| 15000 | Order does not exist |
| 15001 | Order status error |
| 15002 | Order expired |
| 16000 | Invalid address |
| 16001 | Chain not supported |
| 16002 | Token not supported |

For a complete list of error codes, please refer to the Code enum definition in the proto file.

## Webhook Notifications

### Notification Format

All Webhook notifications use POST method with Content-Type application/json.

### Notification Signature Verification

1. Get `X-CC-Sign` field from the Header
2. Concatenate the request Body with WebhookSecret
3. SHA256 encrypt the concatenated string
4. Convert the encrypted result to uppercase and compare with X-CC-Sign

### Notification Retry Mechanism

- If notification fails, the system will retry
- Retry intervals: 1 minute, 5 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours
- Maximum 6 retries

### Response Requirements

- The receiver must return HTTP 200 status code within 5 seconds
- The return content can be empty or any content
- Non-200 status codes are considered notification failure and will trigger a retry
