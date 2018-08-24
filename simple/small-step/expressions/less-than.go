package expressions

import (
	"errors"
	"fmt"
)

// LessThan is a "less than" comparison type
type LessThan struct {
	Left  Expression
	Right Expression
}

// GetValue returns the value
func (lt LessThan) GetValue() ExpressionValue {
	var leftVal interface{} = lt.Left.GetValue()
	var rightVal interface{} = lt.Right.GetValue()

	l, ok := leftVal.(int)
	r, ok := rightVal.(int)

	if ok {
		return l < r
	}

	return errors.New("Error: Unable to calculate value")
}

func (lt LessThan) String() string {
	return fmt.Sprintf("%s < %s", lt.Left.String(), lt.Right.String())
}

// Reducible defines if expression can be simplified
func (lt LessThan) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (lt LessThan) Reduce(environment map[string]Expression) (Expression, error) {
	if lt.Left.Reducible() {
		reduce, _ := lt.Left.Reduce(environment)
		return LessThan{Left: reduce, Right: lt.Right}, nil
	} else if lt.Right.Reducible() {
		reduce, _ := lt.Right.Reduce(environment)
		return LessThan{Left: lt.Left, Right: reduce}, nil
	} else {
		return Boolean{Value: lt.GetValue()}, nil
	}
}
