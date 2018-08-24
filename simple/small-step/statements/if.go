package statements

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/small-step/expressions"
)

// If is a conditional statement
type If struct {
	Condition   expr.Expression
	Consequence Statement
	Alternative Statement
}

func (_if If) String() string {
	return fmt.Sprintf("if (%v) { %v } else { %v }", _if.Condition, _if.Consequence, _if.Alternative)
}

// Reducible defines if statement can be simplified
func (_if If) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (_if If) Reduce(environment map[string]expr.Expression) (Statement, map[string]expr.Expression) {
	if _if.Condition.Reducible() {
		reduce, _ := _if.Condition.Reduce(environment)
		return If{Condition: reduce, Consequence: _if.Consequence, Alternative: _if.Alternative}, environment
	}

	switch c := _if.Condition.GetValue(); c {
	case true:
		return _if.Consequence, environment
	case false:
		return _if.Alternative, environment
	default:
		return _if.Consequence, environment
	}
}
