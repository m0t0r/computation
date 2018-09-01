package expressions

// Multiply is an addition type
type Multiply struct {
	Left  Expression
	Right Expression
}

// Evaluate processes the statement
func (m Multiply) Evaluate(enviroment map[string]Expression) Expression {
	var left = m.Left.Evaluate(enviroment).(Number)
	var right = m.Right.Evaluate(enviroment).(Number)

	l := left.Value.(int)
	r := right.Value.(int)

	return Number{Value: l * r}
}
