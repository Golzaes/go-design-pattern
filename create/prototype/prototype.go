package prototype

type (
	// Cloneable  interface that the prototype object needs to implement
	Cloneable interface {
		Clone() Cloneable
	}
	PrototypeManager struct {
		prototypes map[string]Cloneable
	}
)

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
