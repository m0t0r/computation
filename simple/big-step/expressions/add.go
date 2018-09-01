package expressions

// Add is an addition type
type Add struct {
	Left  Expression
	Right Expression
}

// Evaluate processes the statement
func (a Add) Evaluate(enviroment map[string]Expression) Expression {
	var left = a.Left.Evaluate(enviroment).(Number)
	var right = a.Right.Evaluate(enviroment).(Number)

	l := left.Value.(int)
	r := right.Value.(int)

	return Number{Value: l + r}
}
