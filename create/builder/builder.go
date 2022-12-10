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
	Builder1 struct{ result string }
	Builder2 struct{ result int }
)

func NewDirector(builder Builder) *Director { return &Director{Builder: builder} }

// Construct Product
func (d *Director) Construct() {
	d.Builder.f1()
	d.Builder.f2()
	d.Builder.f3()
}

func (b *Builder1) f1()            { b.result += `1` }
func (b *Builder1) f2()            { b.result += `2` }
func (b *Builder1) f3()            { b.result += `3` }
func (b *Builder1) Result() string { return b.result }

func (b *Builder2) f1()         { b.result += 1 }
func (b *Builder2) f2()         { b.result += 2 }
func (b *Builder2) f3()         { b.result += 3 }
func (b *Builder2) Result() int { return b.result }
