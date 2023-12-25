package tron

import "github.com/fbsobreira/gotron-sdk/pkg/common"

func IsAddressValid(address string) bool {
	_, err := common.DecodeCheck(address)
	return err == nil
}
