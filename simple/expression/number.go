package expression

import (
	"errors"
	"fmt"
)

// Number is number type
type Number struct {
	Value ExpressionValue
}

// GetValue returns the value
func (n Number) GetValue() ExpressionValue {
	return n.Value.(int)
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// Reducible defines if expression can be simplified
func (n Number) Reducible() bool {
	return false
}

// Reduce simplifies an expression
func (n Number) Reduce(environment map[string]Expression) (Expression, error) {
	return nil, errors.New("Error: Number cannot be reduced")
}
