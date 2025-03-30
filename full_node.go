package tron

type FullNode struct {
	PrivateKey  *privateKey
	Resource    *resource
	Transfer    *transfer
	Transaction *transaction
	Account     *account
	Network     *network
	Solidity    *solidity
}

func newFullNode(privateKey *privateKey, resource *resource, transfer *transfer, transaction *transaction, account *account, network *network, solidity *solidity) *FullNode {
	return &FullNode{
		PrivateKey:  privateKey,
		Resource:    resource,
		Transfer:    transfer,
		Transaction: transaction,
		Account:     account,
		Network:     network,
		Solidity:    solidity,
	}
}
