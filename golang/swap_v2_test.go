package ccpayment

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type swapRoundTripFunc func(*http.Request) (*http.Response, error)

func (fn swapRoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return fn(r) }

func TestSwapV2RoutesAndRequest(t *testing.T) {
	request := &SwapV2Request{
		OrderId: "order-1", FromCoinId: 101, FromAmount: "1.25", ToCoinId: 202,
		MarkupFeeRate: "0.01", SlippageRate: "0.005",
	}
	var path string
	var body map[string]any
	client := NewClient("test-app", "test-secret")
	client.SetBaseURL("https://sdk.test")
	client.SetHTTPClient(&http.Client{Transport: swapRoundTripFunc(func(r *http.Request) (*http.Response, error) {
		path = r.URL.Path
		if r.Method != http.MethodPost {
			t.Errorf("method = %s", r.Method)
		}
		body = nil
		if r.ContentLength > 0 {
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode request: %v", err)
			}
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"code":10000,"msg":"success","data":{}}`)), Header: make(http.Header)}, nil
	})})
	ctx := context.Background()
	tests := []struct {
		name string
		want string
		call func() error
	}{
		{"pairs", "/swapv2/getSwapPairList", func() error { _, err := client.Swap().GetSwapPairList(ctx); return err }},
		{"estimate", "/swapv2/estimate", func() error { _, err := client.Swap().SwapEstimateV2(ctx, request); return err }},
		{"swap", "/swapv2/swap", func() error { _, err := client.Swap().SwapV2(ctx, request); return err }},
		{"record", "/swapv2/getSwapRecord", func() error {
			_, err := client.Swap().GetSwapRecordV2(ctx, &GetSwapRecordV2Request{OrderId: "order-1"})
			return err
		}},
		{"records", "/swapv2/getSwapRecordList", func() error {
			_, err := client.Swap().GetSwapRecordListV2(ctx, &GetSwapRecordListV2Request{OrderIds: []string{"order-1"}})
			return err
		}},
		{"user swap", "/swapv2/userSwap", func() error { _, err := client.UserSwap().UserSwapV2(ctx, request); return err }},
		{"user record", "/swapv2/getUserSwapRecord", func() error {
			_, err := client.UserSwap().GetUserSwapRecordV2(ctx, &GetSwapRecordV2Request{OrderId: "order-1"})
			return err
		}},
		{"user records", "/swapv2/getUserSwapRecordList", func() error {
			_, err := client.UserSwap().GetUserSwapRecordListV2(ctx, &GetSwapRecordListV2Request{OrderIds: []string{"order-1"}})
			return err
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.call(); err != nil {
				t.Fatal(err)
			}
			if path != test.want {
				t.Errorf("path = %q, want %q", path, test.want)
			}
			if test.name == "swap" || test.name == "estimate" || test.name == "user swap" {
				for key, want := range map[string]any{"fromCoinId": float64(101), "fromAmount": "1.25", "toCoinId": float64(202), "markupFeeRate": "0.01", "slippageRate": "0.005"} {
					if body[key] != want {
						t.Errorf("%s = %v, want %v", key, body[key], want)
					}
				}
				if _, exists := body["coinIdIn"]; exists {
					t.Error("V1 field coinIdIn in V2 request")
				}
			}
		})
	}
}
