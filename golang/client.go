package ccpayment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// Service accessors
func (c *Client) BasicInfo() *BasicInfoService {
	return &BasicInfoService{client: c}
}

func (c *Client) MerchantAssets() *MerchantAssetsService {
	return &MerchantAssetsService{client: c}
}

func (c *Client) MerchantDeposit() *MerchantDepositService {
	return &MerchantDepositService{client: c}
}

func (c *Client) MerchantWithdraw() *MerchantWithdrawService {
	return &MerchantWithdrawService{client: c}
}

func (c *Client) MerchantBatchWithdraw() *MerchantBatchWithdrawService {
	return &MerchantBatchWithdrawService{client: c}
}

func (c *Client) UserAssets() *UserAssetsService {
	return &UserAssetsService{client: c}
}

func (c *Client) UserDeposit() *UserDepositService {
	return &UserDepositService{client: c}
}

func (c *Client) UserWithdraw() *UserWithdrawService {
	return &UserWithdrawService{client: c}
}

func (c *Client) UserTransfer() *UserTransferService {
	return &UserTransferService{client: c}
}

func (c *Client) Orders() *OrdersService {
	return &OrdersService{client: c}
}

func (c *Client) Checkout() *CheckoutService {
	return &CheckoutService{client: c}
}

func (c *Client) Swap() *SwapService {
	return &SwapService{client: c}
}

func (c *Client) Utilities() *UtilitiesService {
	return &UtilitiesService{client: c}
}

// post internal method
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

	if resp != nil && apiResp.Data != nil {
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

// postWithStringResponse handles endpoints that may return plain text (like /health)
func (c *Client) postWithStringResponse(ctx context.Context, path string, req interface{}, resp interface{}) error {
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

	// Try to parse as JSON first
	var apiResp Response
	if err := json.Unmarshal(respBody, &apiResp); err == nil {
		// JSON response
		if apiResp.Code != 10000 {
			return &APIError{
				Code:    apiResp.Code,
				Message: apiResp.Msg,
			}
		}

		if resp != nil && apiResp.Data != nil {
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

	// Plain text response (like "pong" for health check)
	if resp != nil {
		// Try to set status field if it's a HealthResponse
		if healthResp, ok := resp.(*HealthResponse); ok {
			healthResp.Status = string(respBody)
			healthResp.Timestamp = int64(time.Now().Unix())
			return nil
		}
	}

	return fmt.Errorf("unmarshal response: %w", err)
}
