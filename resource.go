package tron

import (
	"context"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/pkg/errors"
)

type resource struct {
	client      *client.GrpcClient
	transaction *transaction
}

func newResource(client *client.GrpcClient, transaction *transaction) *resource {
	return &resource{
		client:      client,
		transaction: transaction,
	}
}

func (inst *resource) GetByAccount(address Address) (r *AccountResource, err error) {
	rs, err := inst.client.GetAccountResource(address.String())
	if err != nil {
		return
	}

	r = &AccountResource{
		FreeNetUsed:       rs.GetFreeNetUsed(),
		FreeNetLimit:      rs.GetFreeNetLimit(),
		NetUsed:           rs.GetNetUsed(),
		NetLimit:          rs.GetNetLimit(),
		EnergyLimit:       rs.GetEnergyLimit(),
		EnergyUsed:        rs.GetEnergyUsed(),
		TotalNetWeight:    rs.GetTotalNetWeight(),
		TotalNetLimit:     rs.GetTotalNetLimit(),
		TotalEnergyLimit:  rs.GetTotalEnergyLimit(),
		TotalEnergyWeight: rs.GetTotalEnergyWeight(),
		TronPowerUsed:     rs.GetTronPowerUsed(),
		TronPowerLimit:    rs.GetTronPowerLimit(),
	}
	return
}

func (inst *resource) Delegate(params *DelegateResourceParams, operator Address) (txId string, err error) {
	tx, err := inst.client.DelegateResource(params.Owner.String(), params.Receiver.String(), params.Resource.ToResourceCode(), params.Balance, params.Lock, params.LockPeriod)
	if err != nil {
		return
	}

	if !tx.GetResult().GetResult() {
		err = errors.New(string(tx.GetResult().GetMessage()))
		return
	}

	err = inst.transaction.BroadcastWithSign(tx, operator, params.PermissionId)
	if err != nil {
		return
	}

	txId = EncodeTxId(tx.GetTxid())
	return
}

func (inst *resource) UnDelegate(params *UnDelegateResourceParams, operator Address) (txId string, err error) {
	tx, err := inst.client.UnDelegateResource(params.Owner.String(), params.Receiver.String(), params.Resource.ToResourceCode(), params.Balance, false)
	if err != nil {
		return
	}

	if !tx.GetResult().GetResult() {
		err = errors.New(string(tx.GetResult().GetMessage()))
		return
	}

	err = inst.transaction.BroadcastWithSign(tx, operator, params.PermissionId)
	if err != nil {
		return
	}

	txId = EncodeTxId(tx.GetTxid())
	return
}

// How much energy can be obtained by staking N TRX current.
//
// PS: TRXNum uint is not SUN.
//
// PS: The number of TRX staked must be at least 1 and must be an integer.
//
// Energy can be obtained by staking or burning TRX. The total Energy supply on TRON each day is 90,000,000,000.
// Therefore, the Energy you can get by staking TRX is calculated with the following formula: Energy obtained = TRX staked for the Energy / Total amount of TRX staked for Energy on TRON * 90,000,000,000.
func (inst *resource) CalcEnergyFromTRXStake(trx TRX) (energy uint64, err error) {
	rs, err := inst.client.GetAccountResource(SPECIAL_ADDRESS_SUN)
	if err != nil {
		return
	}

	energy = uint64(trx.Float64()) * uint64(rs.GetTotalEnergyLimit()) / uint64(rs.GetTotalEnergyWeight())
	return
}

// How many TRX need to be staked to obtain a specific amount of energy.
//
// PS: The amount of TRX returned is an integer.
//
// Due to TRON staking mechanism for energy, the minimum allowable value is 1 TRX; therefore, the calculation result of this method will be rounded up.
func (inst *resource) CalcTRXStakeForEnergy(energy uint64) (trx TRX, err error) {
	energyFrom10000TRX, err := inst.CalcEnergyFromTRXStake(10000)
	if err != nil {
		return
	}

	value := float64(energy) * 10000 / float64(energyFrom10000TRX)
	trx = TRX(value).Ceil()
	return
}

// How much bandwidth can be obtained by staking N TRX current.
//
// PS: TRXNum uint is not SUN.
//
// PS: The number of TRX staked must be at least 1 and must be an integer.
func (inst *resource) CalcBandwidthFromTRXStake(trx TRX) (bandwidth uint64, err error) {
	rs, err := inst.client.GetAccountResource(SPECIAL_ADDRESS_SUN)
	if err != nil {
		return
	}

	bandwidth = uint64(trx.Float64()) * uint64(rs.GetTotalNetLimit()) / uint64(rs.GetTotalNetWeight())
	return
}

// How many TRX need to be staked to obtain a specific amount of bandwidth.
//
// PS: The amount of TRX returned is an integer.
func (inst *resource) CalcTRXStakeForBandwidth(bandwidth uint64) (trx TRX, err error) {
	bandwidthFrom10000TRX, err := inst.CalcBandwidthFromTRXStake(10000)
	if err != nil {
		return
	}

	value := float64(bandwidth) * 10000 / float64(bandwidthFrom10000TRX)
	trx = TRX(value).Ceil()
	return
}

// In Stake2.0, query the resource delegation index by an account. Two lists will return, one is the list of addresses the account has delegated its resources(toAddress), and the other is the list of addresses that have delegated resources to the account(fromAddress).
func (inst *resource) GetDelegatedResourceAccountIndexV2(ctx context.Context, account Address) (result *DelegatedResourceAccountIndex, err error) {
	_account, err := common.DecodeCheck(account.String())
	if err != nil {
		return
	}

	rsp, err := inst.client.Client.GetDelegatedResourceAccountIndexV2(ctx, client.GetMessageBytes(_account))
	if err != nil {
		return
	}

	result = NewDelegatedResourceAccountIndex(rsp)
	return
}

// In Stake2.0, query the detail of resource share delegated from fromAddress to toAddress
func (inst *resource) GetDelegatedResourceV2(ctx context.Context, from, to Address) (list DelegatedResourceSlice, err error) {
	_from, err := common.DecodeCheck(from.String())
	if err != nil {
		return
	}
	_to, err := common.DecodeCheck(to.String())
	if err != nil {
		return
	}

	req := &api.DelegatedResourceMessage{
		FromAddress: _from,
		ToAddress:   _to,
	}
	rsp, err := inst.client.Client.GetDelegatedResourceV2(ctx, req)
	if err != nil {
		return
	}

	list = NewDelegatedResourceSlice(rsp.GetDelegatedResource())
	return
}
