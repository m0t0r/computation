package expressions

// Multiply is an addition type
type Multiply struct {
	Left  Expression
	Right Expression
}

// Evaluate processes the statement
func (m Multiply) Evaluate(environment map[string]Expression) Expression {
	var left = m.Left.Evaluate(environment).(Number)
	var right = m.Right.Evaluate(environment).(Number)

	l := left.Value.(int)
	r := right.Value.(int)

	return Number{Value: l * r}
}
