package singleton

import "sync"

type lazySingleton struct{}

func (*lazySingleton) foo() {}

var (
	lazyInstance *lazySingleton
	once         sync.Once
)

func LazyInstantiation() Singleton {
	if lazyInstance == nil {
		once.Do(func() {
			lazyInstance = &lazySingleton{}
		})
	}
	return lazyInstance
}
