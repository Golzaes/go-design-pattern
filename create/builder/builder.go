package builder

type (
	Director struct{ Builder }
	Builder  interface {
		f1()
		f2()
		f3()
	}
)

type (
	Build1 struct{ result string }
	Build2 struct{ result int }
)

func NewDirector(builder Builder) *Director { return &Director{builder} }

// Construct Product
func (d *Director) Construct() {
	d.Builder.f1()
	d.Builder.f2()
	d.Builder.f3()
}

func (b *Build1) f1()            { b.result += `1` }
func (b *Build1) f2()            { b.result += `2` }
func (b *Build1) f3()            { b.result += `3` }
func (b *Build1) Result() string { return b.result }

func (b *Build2) f1()         { b.result += 1 }
func (b *Build2) f2()         { b.result += 2 }
func (b *Build2) f3()         { b.result += 3 }
func (b *Build2) Result() int { return b.result }
