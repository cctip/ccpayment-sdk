package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	ccpayment "github.com/cctip/ccpayment-sdk/golang"
)

const (
	appID     = "OOjggt9iLaWHuZps"
	appSecret = "7811db85cafd5e07073ba64cb60e9de8"
	httpProxy = "http://127.0.0.1:10808"
)

func main() {
	client := ccpayment.NewClient()
	client.SetProxy(httpProxy)
	client.SetTimeout(30 * time.Second)

	resp, err := (&ccpayment.SupportTokenReq{}).GetSupportToken(appID, appSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %v\n", err)
		os.Exit(1)
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal response failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
