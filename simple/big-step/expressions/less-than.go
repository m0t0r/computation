package expressions

// LessThan is a "less than" comparison type
type LessThan struct {
	Left  Expression
	Right Expression
}

// Evaluate processes the statement
func (lt LessThan) Evaluate(enviroment map[string]Expression) Expression {
	var left = lt.Left.Evaluate(enviroment).(Number)
	var right = lt.Right.Evaluate(enviroment).(Number)

	l := left.Value.(int)
	r := right.Value.(int)

	return Boolean{Value: l < r}
}
