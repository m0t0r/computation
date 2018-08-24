package statement

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/expression"
)

// While is a looping statement
type While struct {
	Condition expr.Expression
	Body      Statement
}

func (w While) String() string {
	return fmt.Sprintf("while(%v) { %v }", w.Condition, w.Body)
}

// Reducible defines if statement can be simplified
func (w While) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (w While) Reduce(environment map[string]expr.Expression) (Statement, map[string]expr.Expression) {
	return If{Condition: w.Condition, Consequence: Sequence{First: w.Body, Second: w}, Alternative: DoNothing{}}, environment
}
