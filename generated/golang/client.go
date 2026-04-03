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

// Client CCPayment API客户端
type Client struct {
	appID      string
	appSecret  string
	baseURL    string
	httpClient *http.Client
}

// NewClient 创建新的CCPayment客户端
func NewClient(appID, appSecret string) *Client {
	return &Client{
		appID:      appID,
		appSecret:  appSecret,
		baseURL:    DefaultBaseURL,
		httpClient: &http.Client{},
	}
}

// SetHTTPClient 设置自定义HTTP客户端
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// MerchantAssets 返回商家资产服务
func (c *Client) MerchantAssets() *MerchantAssetsService {
	return &MerchantAssetsService{client: c}
}

// post 内部POST请求方法
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
