package tron

type FullNode struct {
	PrivateKey  *privateKey
	Resource    *resource
	Transfer    *transfer
	Transaction *transaction
	Account     *account
	Network     *network
}

func newFullNode(privateKey *privateKey, resource *resource, transfer *transfer, transaction *transaction, account *account, network *network) *FullNode {
	return &FullNode{
		PrivateKey:  privateKey,
		Resource:    resource,
		Transfer:    transfer,
		Transaction: transaction,
		Account:     account,
		Network:     network,
	}
}
