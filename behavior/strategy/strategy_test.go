package strategy

import (
	"testing"
)

func TestBank_Pay(t *testing.T) {
	tests := []struct {
		name string
		Pay  *Payment
		want string
	}{
		{
			t.Name(),
			NewPayment(``, `Ada`, 213, &Cash{}),
			`Pay ¥213 to Ada by cash`,
		},
		{
			t.Name(),
			NewPayment(`0021`, `Payne`, 213, &Bank{}),
			`Pay ¥213 to Payne by bank account 0021`,
		},
		{
			t.Name(),
			NewPayment(`0021`, `Payne`, 213, &QR{}),
			`Pay ¥213 to Payne by QR account 0021`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Pay.Pay(); got != tt.want {
				t.Errorf("Pay() = %v, want %v", got, tt.want)
			}
		})
	}
}
