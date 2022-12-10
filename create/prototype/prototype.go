package prototype

type (
	// Cloneable  interface that the prototype object needs to implement
	Cloneable interface {
		Clone() Cloneable
	}

	ProtoManager struct {
		prototypes map[string]Cloneable
	}
)

func NewPrototypeManager() *ProtoManager {
	return &ProtoManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *ProtoManager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

func (p *ProtoManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
