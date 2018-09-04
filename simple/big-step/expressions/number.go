package expressions

import "fmt"

// Number is an integer type
type Number struct {
	Value ExpressionValue
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// Evaluate processes the statement
func (n Number) Evaluate(environment map[string]Expression) Expression {
	return n
}
