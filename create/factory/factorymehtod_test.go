package factory

import "testing"

func TestOperator(t *testing.T) {
	tests := []struct {
		name       string
		factory    OperatorFactory
		a, b, want int
	}{
		{
			t.Name(),
			&PlusOperatorFactory{},
			71,
			7,
			78,
		},
		{
			t.Name(),
			&MinusOperatorFactory{},
			71,
			7,
			64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.factory, tt.a, tt.b); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})

	}

}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
