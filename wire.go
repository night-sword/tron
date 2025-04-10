//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package tron

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
)

func NewFullNode(grpcClient *client.GrpcClient, httpClient *resty.Client) *FullNode {
	panic(
		wire.Build(
			newAccount,
			newNetwork,
			newResource,
			newTransfer,
			newSolidity,
			newPrivateKey,
			newTransaction,
			newFullNode,
		),
	)
}

func NewGrid(client *resty.Client) *Grid {
	panic(
		wire.Build(
			newApiKeys,
			newGrid,
		),
	)
}
