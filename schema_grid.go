package tron

type GetContractTransactionByAccountResponse struct {
	Data    []*ContractTransaction `json:"data"`
	Success bool                   `json:"success"`
	Meta    *GridMeta              `json:"meta"`
}

type ContractTransaction struct {
	TransactionId  string     `json:"transaction_id"`
	TokenInfo      *TokenInfo `json:"token_info"`
	BlockTimestamp int64      `json:"block_timestamp"`
	From           string     `json:"from"`
	To             string     `json:"to"`
	Type           string     `json:"type"`
	Value          string     `json:"value"`
}

type TokenInfo struct {
	Symbol   string  `json:"symbol"`
	Address  Address `json:"address"`
	Decimals int     `json:"decimals"`
	Name     string  `json:"name"`
}

type GridMeta struct {
	At       int64 `json:"at"`
	PageSize int   `json:"page_size"`
}

type GetAccountInfoByAddressResponse struct {
	Data    []*AccountInfo `json:"data"`
	Success bool           `json:"success"`
	Meta    *GridMeta      `json:"meta"`
}

type AccountInfo struct {
	OwnerPermission struct {
		Keys []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		PermissionName string `json:"permission_name"`
	} `json:"owner_permission"`
	AccountResource struct {
		EnergyWindowOptimized             bool  `json:"energy_window_optimized"`
		LatestConsumeTimeForEnergy        int64 `json:"latest_consume_time_for_energy"`
		DelegatedFrozenV2BalanceForEnergy int   `json:"delegated_frozenV2_balance_for_energy"`
		EnergyWindowSize                  int   `json:"energy_window_size"`
	} `json:"account_resource"`
	ActivePermission []struct {
		Operations string `json:"operations"`
		Keys       []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		Id             int    `json:"id"`
		Type           string `json:"type"`
		PermissionName string `json:"permission_name"`
	} `json:"active_permission"`
	Address            string `json:"address"`
	CreateTime         int64  `json:"create_time"`
	LatestConsumeTime  int64  `json:"latest_consume_time"`
	Allowance          int    `json:"allowance"`
	NetUsage           int    `json:"net_usage"`
	LatestOprationTime int64  `json:"latest_opration_time"`
	FrozenV2           []struct {
		Amount int64  `json:"amount,omitempty"`
		Type   string `json:"type,omitempty"`
	} `json:"frozenV2"`
	UnfrozenV2 []struct {
		UnfreezeAmount     int    `json:"unfreeze_amount"`
		Type               string `json:"type"`
		UnfreezeExpireTime int64  `json:"unfreeze_expire_time"`
	} `json:"unfrozenV2"`
	Balance               int64 `json:"balance"`
	LatestConsumeFreeTime int64 `json:"latest_consume_free_time"`
	Votes                 []struct {
		VoteAddress string `json:"vote_address"`
		VoteCount   int    `json:"vote_count"`
	} `json:"votes"`
	NetWindowSize      int  `json:"net_window_size"`
	NetWindowOptimized bool `json:"net_window_optimized"`
}

type GetResourcePricesResponse struct {
	Prices string `json:"prices"`
}
