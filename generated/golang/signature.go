package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// generateSign 生成API签名
// 签名算法: HMAC-SHA256(appID + timestamp + requestBody, appSecret)
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
