package expressions

import "fmt"

// Boolean is an integer type
type Boolean struct {
	Value ExpressionValue
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b.Value.(bool))
}

// Evaluate processes the statement
func (b Boolean) Evaluate(environment map[string]Expression) Expression {
	return b
}
