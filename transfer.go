package tron

import (
	"math/big"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/pkg/errors"
)

type transfer struct {
	client      *client.GrpcClient
	transaction *transaction
}

func newTransfer(client *client.GrpcClient, transaction *transaction) *transfer {
	return &transfer{
		client:      client,
		transaction: transaction,
	}
}

func (inst *transfer) TransferUSDT(from, to Address, contract Address, amount SUN, operator Address, feeLimit int64, permissionId int32) (txId string, err error) {
	i := big.NewInt(amount.Int64())
	tx, err := inst.client.TRC20Send(from.String(), to.String(), contract.String(), i, feeLimit)
	if err != nil {
		return
	}

	if !tx.GetResult().GetResult() {
		err = errors.New(string(tx.GetResult().GetMessage()))
		return
	}

	err = inst.transaction.BroadcastWithSign(tx, operator, permissionId)
	if err != nil {
		return
	}

	txId = EncodeTxId(tx.GetTxid())
	return
}

func (inst *transfer) Transfer(from, to Address, amount int64, operator Address, permissionId int32) (txId string, err error) {
	tx, err := inst.client.Transfer(from.String(), to.String(), amount)
	if err != nil {
		return
	}

	if !tx.GetResult().GetResult() {
		err = errors.New(string(tx.GetResult().GetMessage()))
		return
	}

	err = inst.transaction.BroadcastWithSign(tx, operator, permissionId)
	if err != nil {
		return
	}

	txId = EncodeTxId(tx.GetTxid())
	return
}

func (inst *transfer) Approve(from, to Address, contract Address, amount int64, operator Address, feeLimit int64, permissionId int32) (txId string, err error) {
	i := big.NewInt(amount)
	tx, err := inst.client.TRC20Approve(from.String(), to.String(), contract.String(), i, feeLimit)
	if err != nil {
		return
	}

	if !tx.GetResult().GetResult() {
		err = errors.New(string(tx.GetResult().GetMessage()))
		return
	}

	err = inst.transaction.BroadcastWithSign(tx, operator, permissionId)
	if err != nil {
		return
	}

	txId = EncodeTxId(tx.GetTxid())
	return
}

func (inst *transfer) TransferTRX(from, to Address, amount int64, operator Address, permissionId int32) (txId string, err error) {
	return inst.Transfer(from, to, amount*SUN_VALUE, operator, permissionId)
}
