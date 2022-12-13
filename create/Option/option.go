package Option

type PeopleOption func(*People)

type People struct {
	id     uint64
	name   string
	age    uint8
	gender bool
	height float32
	bmi    float32
	addr   string
}

func NewPeople(option ...PeopleOption) *People {
	p := &People{}
	p.Init()
	for _, f := range option {
		f(p)
	}
	return p
}

func (p *People) Init() {
	p.id = 10
	p.name = `payne`
	p.age = 31
	p.gender = false
	p.height = 178.37
	p.bmi = 28.31
	p.addr = `xxx,xxx`
}

func Id(id uint64) PeopleOption {
	return func(p *People) {
		p.id = id
	}
}

func Name(name string) PeopleOption {
	return func(p *People) {
		if name != `` {
			p.name = name
		}
	}
}

func Age(age uint8) PeopleOption {
	return func(p *People) {
		p.age = age
	}
}
func Gender(gender bool) PeopleOption {
	return func(p *People) {
		p.gender = gender
	}
}
func Height(height float32) PeopleOption {
	return func(p *People) {
		p.height = height
	}
}
func Bmi(bmi float32) PeopleOption {
	return func(p *People) {
		p.bmi = bmi
	}
}
func Addr(addr string) PeopleOption {
	return func(p *People) {
		p.addr = addr
	}
}
