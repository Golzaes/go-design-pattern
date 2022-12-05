package factory

type (
	// OperatorFactory Factory interface
	OperatorFactory interface{ Create() Operator }
	// Operator The actual class interface that encapsulates
	Operator interface {
		SetA(int)
		SetB(int)
		Result() int
	}
)
type (
	PlusOperatorFactory  struct{}
	MinusOperatorFactory struct{}
)

type (
	// OperatorBase is the base class for the Operator interface implementation, encapsulating common methods
	OperatorBase struct{ a, b int }
	// PlusOperator The actual addition implementation of the operator
	PlusOperator struct{ *OperatorBase }
	// MinusOperator The actual subtraction implementation of the operator
	MinusOperator struct{ *OperatorBase }
)

func (*MinusOperatorFactory) Create() Operator { return &MinusOperator{OperatorBase: &OperatorBase{}} }
func (*PlusOperatorFactory) Create() Operator  { return &PlusOperator{OperatorBase: &OperatorBase{}} }

func (o *OperatorBase) SetA(a int)  { o.a = a }
func (o *OperatorBase) SetB(b int)  { o.b = b }
func (o PlusOperator) Result() int  { return o.a + o.b }
func (o MinusOperator) Result() int { return o.a - o.b }
