package decorators

type (
	IPizza         interface{ getPrice() int }
	VegetableMania struct{}
	CheeseTopping  struct{ pizza IPizza }
	TomatoTopping  struct{ pizza IPizza }
)

func (p *VegetableMania) getPrice() int { return 15 }

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
