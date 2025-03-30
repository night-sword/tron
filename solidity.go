package tron

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type solidity struct {
	httpClient *resty.Client
}

func newSolidity(httpClient *resty.Client) *solidity {
	return &solidity{
		httpClient: httpClient,
	}
}

func (inst *solidity) IsTxConfirmed(ctx context.Context, txId string) (confirmed bool, err error) {
	url := "/walletsolidity/gettransactionbyid"

	req := map[string]string{
		"value": txId,
	}
	rsp, err := httpPost[map[string]string, GetTransactionByIdResponse](ctx, inst.httpClient, url, &req)
	if err != nil {
		return
	}

	if len(rsp.Ret) == 0 {
		return
	}

	if rsp.Ret[0].ContractRet == "SUCCESS" {
		confirmed = true
	}

	return
}
