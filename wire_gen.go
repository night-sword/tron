// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package tron

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/go-resty/resty/v2"
)

// Injectors from full_node_wire.go:

func NewFullNode(grpcClient *client.GrpcClient, httpClient *resty.Client) *FullNode {
	tronPrivateKey := newPrivateKey()
	tronTransaction := newTransaction(grpcClient, tronPrivateKey)
	tronResource := newResource(grpcClient, tronTransaction)
	tronTransfer := newTransfer(grpcClient, tronTransaction)
	tronAccount := newAccount(grpcClient, tronTransaction)
	tronNetwork := newNetwork(grpcClient, httpClient)
	fullNode := newFullNode(tronPrivateKey, tronResource, tronTransfer, tronTransaction, tronAccount, tronNetwork)
	return fullNode
}

// Injectors from grid_wire.go:

func NewGrid(client2 *resty.Client) *Grid {
	tronApiKey := newApiKeys()
	grid := newGrid(client2, tronApiKey)
	return grid
}
