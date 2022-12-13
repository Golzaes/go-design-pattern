package builder

type People struct {
	id     uint64
	name   string
	age    uint8
	gender bool
	height float32
	bmi    float32
	addr   string
}

func NewPeople() *People {
	p := &People{}
	p.Init()
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

func (p *People) WithId(id uint64) *People {
	p.id = id
	return p
}
func (p *People) WithName(name string) *People {
	p.name = name
	return p
}
func (p *People) WithAge(age uint8) *People {
	p.age = age
	return p
}
func (p *People) WithGender(gender bool) *People {
	p.gender = gender
	return p
}
func (p *People) WithHeight(height float32) *People {
	p.height = height
	return p
}
func (p *People) WithBmi(bmi float32) *People {
	p.bmi = bmi
	return p
}
func (p *People) WithAddr(addr string) *People {
	p.addr = addr
	return p
}
