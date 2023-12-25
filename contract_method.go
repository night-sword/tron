package tron

//go:generate enumer --type=ContractMethod --linecomment --extramethod --output=contract_method_enum.go
type ContractMethod uint

const (
	ContractMethod_Unknown ContractMethod = iota // UNKNOWN
	// a9059cbb: transfer(address,uint256)
	ContractUSDTMethod_Transfer // a9059cbb
	// 095ea7b3: approve(address,uint256)
	ContractUSDTMethod_Approve // 095ea7b3
	// 66188463: decreaseApproval(address,uint256)
	ContractUSDTMethod_DecreaseApproval // 66188463
	// 3f4ba83a: unpause()
	ContractUSDTMethod_Unpause // 3f4ba83a
	// cc872b66: issue(uint256)
	ContractUSDTMethod_Issue // cc872b66
	// 0ecb93c0: addBlackList(address)
	ContractUSDTMethod_AddBlackList // 0ecb93c0
	// f2fde38b: transferOwnership(address)
	ContractUSDTMethod_TransferOwnership // f2fde38b
	// 23b872dd: transferFrom(address,address,uint256)
	ContractUSDTMethod_TransferFrom // 23b872dd
	// e4997dc5: removeBlackList(address)
	ContractUSDTMethod_RemoveBlackList // e4997dc5
	// 0753c30c: deprecate(address)
	ContractUSDTMethod_Deprecate // 0753c30c
	// d73dd623: increaseApproval(address,uint256)
	ContractUSDTMethod_IncreaseApproval // d73dd623
	// db006a75: redeem(uint256)
	ContractUSDTMethod_Redeem // db006a75
	// c0324c77: setParams(uint256,uint256)
	ContractUSDTMethod_SetParams // c0324c77
	// f3bdc228: destroyBlackFunds(address)
	ContractUSDTMethod_DestroyBlackFunds // f3bdc228
	// 8456cb59: pause()
	ContractUSDTMethod_Pause // 8456cb59
)
