package tron

import "github.com/fbsobreira/gotron-sdk/pkg/client"

type FullNode struct {
	PrivateKey  *privateKey
	Resource    *resource
	Transfer    *transfer
	Transaction *transaction
	Account     *account
	Network     *network
	Solidity    *solidity
	client      *client.GrpcClient
}

func newFullNode(client *client.GrpcClient, privateKey *privateKey, resource *resource, transfer *transfer, transaction *transaction, account *account, network *network, solidity *solidity) *FullNode {
	return &FullNode{
		PrivateKey:  privateKey,
		Resource:    resource,
		Transfer:    transfer,
		Transaction: transaction,
		Account:     account,
		Network:     network,
		Solidity:    solidity,
		client:      client,
	}
}

func (inst *FullNode) Client() *client.GrpcClient {
	return inst.client
}
