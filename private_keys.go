package tron

import (
	"encoding/hex"
	"sync"

	"github.com/pkg/errors"
)

type privateKey struct {
	address2keys map[Address][]byte
	mutex        *sync.RWMutex
}

func newPrivateKey() *privateKey {
	return &privateKey{
		address2keys: make(map[Address][]byte),
		mutex:        &sync.RWMutex{},
	}
}

func (inst *privateKey) Get(address Address) (key []byte, err error) {
	inst.mutex.RLock()
	defer inst.mutex.RUnlock()

	key, ok := inst.address2keys[address]
	if !ok {
		err = errors.New("address key not found")
		return
	}

	return
}

func (inst *privateKey) Set(address2key map[Address]string) (err error) {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	for address, key := range address2key {
		bytes, e := hex.DecodeString(key)
		if err = e; err != nil {
			return
		}

		inst.address2keys[address] = bytes
	}

	return
}

func (inst *privateKey) Append(address Address, key string) (err error) {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	bytes, err := hex.DecodeString(key)
	if err != nil {
		return
	}

	inst.address2keys[address] = bytes
	return
}
