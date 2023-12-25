package tron

import (
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc"
)

func NewGrpcClient(endpoint string) (c *client.GrpcClient, err error) {
	c = client.NewGrpcClient(endpoint)
	err = c.Start(grpc.WithInsecure())
	return
}

func NewHttpClient(endpoint string) (c *resty.Client) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
		"accept":     "application/json",
	}
	c = resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(time.Second).
		SetBaseURL(endpoint).
		SetHeaders(headers)

	return
}

func SetHttpClientApiKey(client *resty.Client, key string) (c *resty.Client) {
	return client.SetHeader("TRON-PRO-API-KEY", key)
}
