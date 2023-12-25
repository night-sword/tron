package tron

import (
	"sync"

	"github.com/samber/lo"
)

type apiKey struct {
	keys  []string
	mutex *sync.RWMutex
}

func newApiKeys() *apiKey {
	return &apiKey{
		keys:  make([]string, 0, 8),
		mutex: &sync.RWMutex{},
	}
}

func (inst *apiKey) GetRandom() (key string, ok bool) {
	inst.mutex.RLock()
	defer inst.mutex.RUnlock()

	if len(inst.keys) == 0 {
		return
	}

	inst.keys = lo.Shuffle(inst.keys)
	key = inst.keys[0]
	ok = true

	return
}

func (inst *apiKey) Set(keys []string) {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	inst.keys = keys

	return
}

func (inst *apiKey) Append(key string) {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	inst.keys = append(inst.keys, key)
	return
}
