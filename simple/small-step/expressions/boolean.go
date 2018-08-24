package expressions

import (
	"errors"
	"fmt"
)

// Boolean is a boolean type
type Boolean struct {
	Value ExpressionValue
}

// GetValue returns the value
func (b Boolean) GetValue() ExpressionValue {
	return b.Value.(bool)
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b.Value.(bool))
}

// Reducible defines if expression can be simplified
func (b Boolean) Reducible() bool {
	return false
}

// Reduce simplifies an expression
func (b Boolean) Reduce(environment map[string]Expression) (Expression, error) {
	return nil, errors.New("Error: Boolean cannot be reduced")
}
