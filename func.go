package tron

import (
	"strings"

	"github.com/fbsobreira/gotron-sdk/pkg/common"
)

func EncodeTxId(b []byte) (hex string) {
	hex = strings.TrimLeft(common.BytesToHexString(b), "0x")
	hex = PadLeftStr(hex, 64, '0')
	return
}

func PadLeftStr(str string, length int, pad rune) string {
	if len(str) >= length {
		return str
	}
	return strings.Repeat(string(pad), length-len(str)) + str
}
