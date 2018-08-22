package expression

import (
	"errors"
	"fmt"
)

// Add is an addition type
type Add struct {
	Left  Expression
	Right Expression
}

// GetValue returns the value
func (a Add) GetValue() ExpressionValue {
	var leftVal interface{} = a.Left.GetValue()
	var rightVal interface{} = a.Right.GetValue()

	l, ok := leftVal.(int)
	r, ok := rightVal.(int)

	if ok {
		return l + r
	}

	return errors.New("Error: Unable to calculate value")
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.Left.String(), a.Right.String())
}

// Reducible defines if expression can be simplified
func (a Add) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (a Add) Reduce(environment map[string]Expression) (Expression, error) {
	if a.Left.Reducible() {
		reduce, _ := a.Left.Reduce(environment)
		return Add{Left: reduce, Right: a.Right}, nil
	} else if a.Right.Reducible() {
		reduce, _ := a.Right.Reduce(environment)
		return Add{Left: a.Left, Right: reduce}, nil
	} else {
		return Number{Value: a.GetValue()}, nil
	}
}
