package facade

type (
	action interface {
		TestIt()
	}
	actionImpl struct {
		a moduleA
		b moduleB
	}
	moduleA struct{}
	moduleB struct{}
)

func (*actionImpl) TestIt() {}
