package tron

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
	"github.com/tyler-smith/go-bip39"
)

func IsAddressValid(address string) bool {
	_, err := common.DecodeCheck(address)
	return err == nil
}

func ValidAddressPk(addr Address, pk string) (ok bool) {
	// Decode the private key from hex string
	privateKeyBytes, err := hex.DecodeString(pk)
	if err != nil {
		return
	}

	// Convert to ECDSA private key
	pKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return
	}

	// Get the public key from a private key
	publicKey := pKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}

	// Convert public key to address
	_address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Convert to TRON address format
	addressBytes := _address.Bytes()
	// TRON addresses use a different prefix (0x41 instead of 0x00)
	tronAddressBytes := append([]byte{0x41}, addressBytes...)
	tronAddress := Address(common.EncodeCheck(tronAddressBytes))

	return tronAddress == addr
}

func CreateAccount() (addr Address, pk string, err error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return
	}

	private, public := keys.FromMnemonicSeedAndPassphrase(mnemonic, "", 0)

	pub := public.ToECDSA()
	add := address.PubkeyToAddress(*pub)

	pk = private.Key.String()
	addr = Address(add.String())
	return
}
