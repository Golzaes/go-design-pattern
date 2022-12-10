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
		CardId string
		Name   string
		Money  int
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
	fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardId)
}

func NewPayment(cardId string, name string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			CardId: cardId,
			Name:   name,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() { p.strategy.Pay(p.context) }
