package singleton

import "sync"

type dcSingleton struct{}

func (*dcSingleton) foo() {}

var (
	// dcInstance double check Instance
	dcInstance *dcSingleton
	mu         sync.Mutex
)

func dcInstantiation() Singleton {
	if dcInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if dcInstance != nil {
			dcInstance = &dcSingleton{}
		}
	}
	return dcInstance
}
