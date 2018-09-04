package statements

import (
	expr "github.com/m0t0r/computation/simple/big-step/expressions"
)

// Sequence chains staements eveluation
type Sequence struct {
	First  Statement
	Second Statement
}

// Evaluate processes the statement
func (s Sequence) Evaluate(environment map[string]expr.Expression) map[string]expr.Expression {
	first := s.First.Evaluate(environment)
	return s.Second.Evaluate(first)
}
