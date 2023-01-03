package bridge

import "testing"

func TestP(t *testing.T) {
	HpPinter, EpPinter := &HP{}, &EP{}
	macComputer, winComputer := &Mac{}, &Win{}

	// Mac HP
	macComputer.SetPrinter(HpPinter)
	macComputer.Print()
	// Mac EP
	macComputer.SetPrinter(EpPinter)
	macComputer.Print()
	// Win HP
	winComputer.SetPrinter(HpPinter)
	winComputer.Print()
	// Win EP
	winComputer.SetPrinter(EpPinter)
	winComputer.Print()
}
