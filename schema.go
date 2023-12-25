package tron

import "github.com/fbsobreira/gotron-sdk/pkg/proto/core"

type Resource string

const (
	Resource_Bandwidth Resource = "bandwidth"
	Resource_Energy             = "energy"
	Resource_TronPower          = "tron_power"
)

func (inst Resource) ToResourceCode() core.ResourceCode {
	switch inst {
	case "bandwidth":
		return core.ResourceCode_BANDWIDTH
	case "energy":
		return core.ResourceCode_ENERGY
	case "tron_power":
		return core.ResourceCode_TRON_POWER
	}

	return -1
}

type DelegateResourceParams struct {
	Owner        Address  `json:"owner"`
	Receiver     Address  `json:"receiver"`
	Resource     Resource `json:"resource"`
	Balance      int64    `json:"balance"`
	Lock         bool     `json:"lock"`
	LockPeriod   int64    `json:"lock_period"`
	PermissionId int32    `json:"permission_id"`
}

type UnDelegateResourceParams struct {
	Owner        Address  `json:"owner"`
	Receiver     Address  `json:"receiver"`
	Resource     Resource `json:"resource"`
	Balance      int64    `json:"balance"`
	PermissionId int32    `json:"permission_id"`
}

type AccountResource struct {
	FreeNetLimit      int64 `json:"freeNetLimit"`      // Total free bandwidth
	NetUsed           int64 `json:"netUsed"`           // Used amount of bandwidth obtained by staking
	NetLimit          int64 `json:"netLimit"`          // Total bandwidth obtained by staking
	EnergyLimit       int64 `json:"energyLimit"`       // Total energy obtained by staking
	EnergyUsed        int64 `json:"energyUsed"`        // Energy used
	TotalNetWeight    int64 `json:"totalNetWeight"`    // Total TRX staked for bandwidth by the whole network
	TotalNetLimit     int64 `json:"totalNetLimit"`     // Total bandwidth can be obtained by staking by the whole network
	TotalEnergyLimit  int64 `json:"totalEnergyLimit"`  // Total energy can be obtained by staking by the whole network
	TotalEnergyWeight int64 `json:"totalEnergyWeight"` // Total TRX staked for energy by the whole network
	TronPowerUsed     int64 `json:"tronPowerUsed"`     // TRON Power(vote) used
	TronPowerLimit    int64 `json:"tronPowerLimit"`    // TRON Power(vote)
}

type Balance struct {
	Energy      *ResourceBalance `json:"energy"`
	Bandwidth   *ResourceBalance `json:"bandwidth"`
	TRX         int64            `json:"trx"`       // current usable trx
	TRXTotal    int64            `json:"trx_total"` // total trx (usable+frozen+unfrozen trx)
	TRXFrozenV1 int64            `json:"trx_frozen_v1"`
	TRXFrozenV2 *TRXFrozenV2     `json:"trx_frozen_v2"`
	UnfrozenV2  *UnfrozenV2      `json:"unfrozen_v2"`
}

type ResourceBalance struct {
	Current int64 `json:"current"`
	Limit   int64 `json:"limit"`
}

type TRXFrozenV2 struct {
	Total     int64             `json:"total"`
	Energy    *ResourceFrozenV2 `json:"energy"`
	Bandwidth *ResourceFrozenV2 `json:"bandwidth"`
}

type ResourceFrozenV2 struct {
	Total    int64 `json:"total"`
	Self     int64 `json:"self"`
	Delegate int64 `json:"delegate"`
}

type UnfrozenV2 struct {
	Total         int64 `json:"total"`
	Unstacking    int64 `json:"unstacking"`
	ToBeWithdrawn int64 `json:"to_be_withdrawn"`
}

type SmartContractEvent struct {
	Contract Address
	Owner    Address
	To       Address
	Amount   SUN
	Method   ContractMethod
}
