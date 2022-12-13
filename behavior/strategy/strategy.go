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
		Pay(*PaymentContext) string
	}
)

type (
	Cash struct{}
	Bank struct{}
	QR   struct{}
)

func (*Cash) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf(`Pay ¥%d to %s by cash`, ctx.Money, ctx.Name)
}

func (*Bank) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf(`Pay ¥%d to %s by bank account %s`, ctx.Money, ctx.Name, ctx.Id)
}

func (*QR) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf(`Pay ¥%d to %s by QR account %s`, ctx.Money, ctx.Name, ctx.Id)
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

func (p *Payment) Pay() string { return p.strategy.Pay(p.context) }
