package singleton

import "sync"

type dcSingleton struct{}

func (*dcSingleton) foo() {}

var (
	dcInstance *dcSingleton
	mu         sync.Mutex
)

func dcInstantiation() Singleton {
	if dcInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		dcInstance = &dcSingleton{}
	}
	return dcInstance
}
