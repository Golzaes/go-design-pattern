package factory

import "testing"

func TestOperator(t *testing.T) {
	var factory OperatorFactory

	factory = &PlusOperatorFactory{}
	if compute(factory, 77, 21) != 98 {
		t.Fatal("error with factory method pattern")
	}

	factory = &MinusOperatorFactory{}
	if compute(factory, 92, 22) != 70 {
		t.Fatal("error with factory method pattern")
	}
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
