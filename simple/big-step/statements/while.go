package statements

import (
	expr "github.com/m0t0r/computation/simple/big-step/expressions"
)

// While is a looping statement
type While struct {
	Condition expr.Expression
	Body      Statement
}

// Evaluate processes the statement
func (w While) Evaluate(environment map[string]expr.Expression) map[string]expr.Expression {
	switch c := w.Condition.Evaluate(environment); c {
	case expr.Boolean{Value: true}:
		return w.Evaluate(w.Body.Evaluate(environment))
	case expr.Boolean{Value: false}:
		return environment
	default:
		return environment
	}
}
