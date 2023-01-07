package bridge

import "fmt"

type (
	Computer interface {
		Print()
		SetPrinter(Printer)
	}

	Printer interface {
		PrintFile()
	}
)

type (
	Mac struct{ printer Printer }
	Win struct{ printer Printer }
	EP  struct{}
	HP  struct{}
)

func (*HP) PrintFile() { fmt.Println("Printing by a HP Printer") }
func (*EP) PrintFile() { fmt.Println("Printing by a EP Printer") }

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (w *Win) Print() {
	fmt.Println("Print request for win")
	w.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) { m.printer = p }
func (w *Win) SetPrinter(p Printer) { w.printer = p }
