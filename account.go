package tron

import (
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
)

type account struct {
	client      *client.GrpcClient
	transaction *transaction
}

func newAccount(client *client.GrpcClient, transaction *transaction) *account {
	return &account{
		client:      client,
		transaction: transaction,
	}
}

// Active an account.
// This operation will cost 1TRX.
func (inst *account) Active(address, helper, operator Address, permissionId int32) (txId string, err error) {
	active, err := inst.IsAccountActive(address)
	if err != nil {
		return
	}
	if active {
		return
	}

	tx, err := inst.client.CreateAccount(helper.String(), address.String())
	if err != nil {
		return
	}

	txId = EncodeTxId(tx.GetTxid())

	err = inst.transaction.BroadcastWithSign(tx, operator, permissionId)
	return
}

func (inst *account) IsAccountActive(address Address) (active bool, err error) {
	_, err = inst.client.GetAccount(address.String())
	if err != nil {
		if err.Error() == "account not found" {
			active, err = false, nil
		}
		return
	}

	active = true
	return
}

// get account current balance (energy bandwidth balance)
//
// PS: unofficial method
func (inst *account) GetAccountBalance(address Address) (balance *Balance, err error) {
	a, err := inst.client.GetAccount(address.String())
	if err != nil {
		return
	}

	rs, err := inst.client.GetAccountResource(address.String())
	if err != nil {
		return
	}

	energy := rs.GetEnergyLimit() - rs.GetEnergyUsed()
	bandwidth := (rs.GetNetLimit() - rs.GetNetUsed()) + (rs.GetFreeNetLimit() - rs.GetFreeNetUsed())
	delegateForEnergy := a.GetAccountResource().GetDelegatedFrozenV2BalanceForEnergy()
	delegateForBandwidth := a.GetDelegatedFrozenV2BalanceForBandwidth()

	energySelf, bandwidthSelf := int64(0), int64(0)
	for _, f := range a.GetFrozenV2() {
		switch f.GetType() {
		case core.ResourceCode_ENERGY:
			energySelf += f.GetAmount()
		case core.ResourceCode_BANDWIDTH:
			bandwidthSelf += f.GetAmount()
		}
	}

	unstacking, toBeWithdrawn := int64(0), int64(0)
	now := time.Now().UnixMilli()
	for _, u := range a.GetUnfrozenV2() {
		if u.GetUnfreezeExpireTime() > now {
			unstacking += u.GetUnfreezeAmount()
		} else {
			toBeWithdrawn += u.GetUnfreezeAmount()
		}
	}

	frozenV2Total := energySelf + delegateForEnergy + bandwidthSelf + delegateForBandwidth
	UnfrozenV2Total := unstacking + toBeWithdrawn

	balance = &Balance{
		Energy: &ResourceBalance{
			Current: energy,
			Limit:   rs.GetEnergyLimit(),
		},
		Bandwidth: &ResourceBalance{
			Current: bandwidth,
			Limit:   rs.GetNetLimit() + rs.GetFreeNetLimit(),
		},
		TRX:         a.Balance,
		TRXTotal:    a.GetBalance() + frozenV2Total + UnfrozenV2Total,
		TRXFrozenV1: 0,
		TRXFrozenV2: &TRXFrozenV2{
			Total: frozenV2Total,
			Energy: &ResourceFrozenV2{
				Total:    energySelf + delegateForEnergy,
				Self:     energySelf,
				Delegate: delegateForEnergy,
			},
			Bandwidth: &ResourceFrozenV2{
				Total:    bandwidthSelf + delegateForBandwidth,
				Self:     bandwidthSelf,
				Delegate: delegateForBandwidth,
			},
		},
		UnfrozenV2: &UnfrozenV2{
			Total:         UnfrozenV2Total,
			Unstacking:    unstacking,
			ToBeWithdrawn: toBeWithdrawn,
		},
	}
	return
}

// Get account USDT balance
func (inst *account) GetUSDTBalance(address Address, contract Address) (balance SUN, err error) {
	if address.String() == "" {
		err = errors.New("address is empty")
		return
	}
	if contract.String() == "" {
		err = errors.New("contract is empty")
		return
	}

	rsp, err := inst.client.TRC20ContractBalance(address.String(), contract.String())
	if err != nil {
		err = errors.Wrap(err, "Call TRC20ContractBalance fail")
		return
	}

	balance = SUN(rsp.Int64())
	return
}
