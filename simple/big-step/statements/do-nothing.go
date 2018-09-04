package statements

import expr "github.com/m0t0r/computation/simple/big-step/expressions"

// DoNothing is noop statement
type DoNothing struct{}

// Evaluate processes the statement
func (dn DoNothing) Evaluate(environment map[string]expr.Expression) map[string]expr.Expression {
	return environment
}
