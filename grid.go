package tron

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type Grid struct {
	apiKeyHeaderName string

	Apikey *apiKey
	client *resty.Client
}

func newGrid(client *resty.Client, keys *apiKey) *Grid {
	return &Grid{
		apiKeyHeaderName: "TRON-PRO-API-KEY",
		Apikey:           keys,
		client:           client,
	}
}

// get account info by address
func (inst *Grid) GetAccountInfo(ctx context.Context, address Address) (account *AccountInfo, err error) {
	api := "/v1/accounts/%s"
	api = fmt.Sprintf(api, address)

	rsp, err := inst.get(ctx, api)
	if err != nil {
		return
	}

	var response *GetAccountInfoByAddressResponse
	err = json.Unmarshal(rsp, &response)
	if err != nil {
		err = errors.New("decode rsp fail")
		return
	}

	if len(response.Data) == 0 {
		err = errors.New("account info empty")
		return
	}

	account = response.Data[0]
	return
}

func (inst *Grid) GetContractRecentTxsByTs(ctx context.Context, contract, address Address, lastTs int64, only DIRECTION) (txs []*ContractTransaction, last int64, err error) {
	list, err := inst.GetContractTxsByTs(ctx, address, lastTs, only)
	if err != nil {
		return
	}

	txs = make([]*ContractTransaction, 0, len(list))
	for _, tx := range list {
		last = tx.BlockTimestamp

		if tx.Type != "Transfer" {
			continue
		}
		if tx.TokenInfo.Address != contract {
			continue
		}
		txs = append(txs, tx)
	}

	return
}

// get contract transaction info by account address
//
//	only : from | to
func (inst *Grid) GetContractTxsByTs(ctx context.Context, address Address, lastTs int64, only DIRECTION) (txs []*ContractTransaction, err error) {
	api := "/v1/accounts/%s/transactions/trc20"
	api = fmt.Sprintf(api, address)

	query := "?limit=200&order_by=block_timestamp%2Casc"
	query += fmt.Sprintf("&min_timestamp=%d", lastTs)

	if only == "to" {
		query += "&only_to=true"
	} else if only == "from" {
		query += "&only_from=true"
	}

	rsp, err := inst.get(ctx, api+query)
	if err != nil {
		return
	}

	var response *GetContractTransactionByAccountResponse
	err = json.Unmarshal(rsp, &response)
	if err != nil {
		err = errors.New("decode rsp fail")
		return
	}

	txs = response.Data
	return
}

func (inst *Grid) get(ctx context.Context, api string) (body []byte, err error) {
	client := inst.client.R().SetContext(ctx)

	key, ok := inst.Apikey.GetRandom()
	if ok {
		client = client.SetHeader(inst.apiKeyHeaderName, key)
	}

	rsp, err := client.Get(api)
	if err != nil {
		return
	}

	body = rsp.Body()
	return
}

func (inst *Grid) post(ctx context.Context, api string, input any) (body []byte, err error) {
	client := inst.client.R().SetContext(ctx)

	key, ok := inst.Apikey.GetRandom()
	if ok {
		client = client.SetHeader(inst.apiKeyHeaderName, key)
	}

	client.SetBody(input)

	rsp, err := client.Post(api)
	if err != nil {
		return
	}

	body = rsp.Body()
	return
}
