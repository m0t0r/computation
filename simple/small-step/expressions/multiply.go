package expressions

import (
	"errors"
	"fmt"
)

// Multiply is a multiplication type
type Multiply struct {
	Left  Expression
	Right Expression
}

// GetValue returns the value
func (m Multiply) GetValue() ExpressionValue {
	var leftVal interface{} = m.Left.GetValue()
	var rightVal interface{} = m.Right.GetValue()

	l, ok := leftVal.(int)
	r, ok := rightVal.(int)

	if ok {
		return l * r
	}

	return errors.New("Error: Unable to calculate value")
}

func (m Multiply) String() string {
	return fmt.Sprintf("%s * %s", m.Left.String(), m.Right.String())
}

// Reducible defines if expression can be simplified
func (m Multiply) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (m Multiply) Reduce(environment map[string]Expression) (Expression, error) {
	if m.Left.Reducible() {
		reduce, _ := m.Left.Reduce(environment)
		return Multiply{Left: reduce, Right: m.Right}, nil
	} else if m.Right.Reducible() {
		reduce, _ := m.Right.Reduce(environment)
		return Multiply{Left: m.Left, Right: reduce}, nil
	} else {
		return Number{Value: m.GetValue()}, nil
	}
}
