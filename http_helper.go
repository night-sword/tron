package tron

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

func decodeJsonResponse[T any](response *resty.Response) (rsp *T, err error) {
	if response.StatusCode() != 200 {
		err = errors.New(fmt.Sprintf("response code is %d", response.StatusCode()))
		return
	}

	err = json.Unmarshal(response.Body(), &rsp)
	if err != nil {
		err = errors.New(fmt.Sprintf("decode response fail"))
	}
	return
}

func httpGet[Response any](ctx context.Context, client *resty.Client, api string) (rsp *Response, err error) {
	r, err := client.R().
		SetContext(ctx).
		Get(api)
	if err != nil {
		return
	}

	return decodeJsonResponse[Response](r)
}

func httpGetParams[Request, Response any](ctx context.Context, client *resty.Client, api string, req *Request) (rsp *Response, err error) {
	if req != nil {
		r, e := query.Values(req)
		if err = e; err != nil {
			err = errors.New(fmt.Sprintf("build query params error:%s", err.Error()))
			return
		}

		api = api + "?" + r.Encode()
	}

	r, err := client.R().
		SetContext(ctx).
		Get(api)
	if err != nil {
		return
	}

	return decodeJsonResponse[Response](r)
}

func httpPost[Request, Response any](ctx context.Context, client *resty.Client, api string, req *Request) (rsp *Response, err error) {
	r, err := client.R().
		SetContext(ctx).
		SetBody(req).
		Post(api)
	if err != nil {
		return
	}

	return decodeJsonResponse[Response](r)
}
