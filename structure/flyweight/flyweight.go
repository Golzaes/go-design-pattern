package flyweight

import "fmt"

func (t *TerroristDress) getColor() string { return t.color }

type (
	game struct {
		terrorists        []*Player
		counterTerrorists []*Player
	}
	Player struct {
		dress      Dress
		playerType string
		lat        int
		long       int
	}
)
type (
	Dress interface{ getColor() string }

	DressFactory struct{ dressMap map[string]Dress }

	TerroristDress        struct{ color string }
	CounterTerroristDress struct{ color string }
)

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	// handle error
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) { p.lat, p.long = lat, long }

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var dressFactorySingleInstance = &DressFactory{dressMap: make(map[string]Dress)}

func getDressFactorySingleInstance() *DressFactory { return dressFactorySingleInstance }

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	case CounterTerrroristDressType:
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("wrong dress type passed")

}

func (c *CounterTerroristDress) getColor() string { return c.color }

func newTerroristDress() *TerroristDress               { return &TerroristDress{color: "red"} }
func newCounterTerroristDress() *CounterTerroristDress { return &CounterTerroristDress{color: "green"} }
