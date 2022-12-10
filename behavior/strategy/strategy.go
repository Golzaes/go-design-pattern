package strategy

import (
	"fmt"
)

type (
	Payment struct {
		context  *PaymentContext
		strategy PaymentStrategy
	}

	PaymentContext struct {
		Id    string
		Name  string
		Money int
	}

	PaymentStrategy interface {
		Pay(*PaymentContext)
	}
)

type (
	Cash struct{}
	Bank struct{}
)

func (*Cash) Pay(ctx *PaymentContext) { fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name) }

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.Id)
}

func NewPayment(id, name string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Id:    id,
			Name:  name,
			Money: money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() { p.strategy.Pay(p.context) }
