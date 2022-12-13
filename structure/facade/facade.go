package facade

import "fmt"

type (
	//API is facade interface of facade package
	API        interface{ TestIt() string }
	AModuleAPI interface{ TestA() string }
	BModuleAPI interface{ TestB() string }
)

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

type (
	aModuleImpl struct{}
	bModuleImpl struct{}
)

func NewAPI() API {
	return &apiImpl{
		a: &aModuleImpl{},
		b: &bModuleImpl{},
	}
}

func (i *apiImpl) TestIt() string { return fmt.Sprintf(`%s\n%s`, i.a.TestA(), i.b.TestB()) }

func (*aModuleImpl) TestA() string { return `A model is tested` }
func (*bModuleImpl) TestB() string { return `B model is tested` }
