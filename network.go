package tron

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

type network struct {
	grpcClient *client.GrpcClient
	httpClient *resty.Client
}

func newNetwork(client *client.GrpcClient, httpClient *resty.Client) *network {
	return &network{
		grpcClient: client,
		httpClient: httpClient,
	}
}

func (inst *network) GetBlockByNum(num int64) (*api.BlockExtention, error) {
	return inst.grpcClient.GetBlockByNum(num)
}

func (inst *network) GetNowBlock() (*api.BlockExtention, error) {
	return inst.grpcClient.GetNowBlock()
}

func (inst *network) GetBlockInfoByNum(num int64) (*api.TransactionInfoList, error) {
	return inst.grpcClient.GetBlockInfoByNum(num)
}

// How much energy can be obtained by burning 1 trx
func (inst *network) GetCurrentEnergyPrice(ctx context.Context) (energy uint64, err error) {
	url := "/wallet/getenergyprices"

	rsp, err := inst.get(ctx, url)
	if err != nil {
		return
	}

	energy, err = inst.decodeLatestPrice(rsp)
	return
}

// How much bandwidth can be obtained by burning 1 trx
func (inst *network) GetCurrentBandwidthPrice(ctx context.Context) (bandwidth uint64, err error) {
	url := "/wallet/getbandwidthprices"

	rsp, err := inst.get(ctx, url)
	if err != nil {
		return
	}

	bandwidth, err = inst.decodeLatestPrice(rsp)
	return
}

func (inst *network) decodeLatestPrice(rsp []byte) (price uint64, err error) {
	var response *GetResourcePricesResponse
	err = json.Unmarshal(rsp, &response)
	if err != nil {
		err = errors.New("decode rsp fail")
		return
	}

	prices := strings.Split(response.Prices, ",")
	if len(prices) == 0 {
		err = errors.New("price empty")
		return
	}

	latest := strings.Split(prices[len(prices)-1], ":")
	if len(latest) < 2 {
		err = errors.New("price format error")
		return
	}

	price, err = cvt.Uint64E(latest[1])
	if err != nil {
		err = errors.New("value cannot convert to uint64, err=" + err.Error())
	}

	return
}

func (inst *network) get(ctx context.Context, api string) (body []byte, err error) {
	request := inst.httpClient.R().SetContext(ctx)

	rsp, err := request.Get(api)
	if err != nil {
		return
	}

	body = rsp.Body()
	return
}
