package strategy

func ExampleCash_Pay() {
	// output:
	// Pay $21 to Ada by cash

	payment := NewPayment(``, `Ada`, 21, &Cash{})
	payment.Pay()
}

func ExampleBank_Pay() {
	// output:
	// Pay $213 to Payne by bank account 0021

	payment := NewPayment(`0021`, `Payne`, 213, &Bank{})
	payment.Pay()
}
