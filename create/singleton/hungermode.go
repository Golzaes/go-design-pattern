package singleton

type hungerSingleton struct{}

func (*hungerSingleton) foo() {}

var hungerInstance *hungerSingleton

func init() {
	hungerInstance = &hungerSingleton{}
}

func HugerInstantiation() Singleton {
	return hungerInstance
}
