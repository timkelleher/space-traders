package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const url = "https://api.spacetraders.io/v2"

type errorResponse struct {
	Error struct {
		Message string           `json:"message"`
		Code    int              `json:"code"`
		Data    map[string][]any `json:"data"`
	} `json:"error"`
}

type Client struct {
	resty *resty.Client
}

func New() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Client{resty: resty.New()}, nil
}

func traceResponse(resp *resty.Response, err error) {
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}

func traceRequest(resp *resty.Response) {
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}
