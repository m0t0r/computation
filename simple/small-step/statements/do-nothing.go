package statements

import expr "github.com/m0t0r/computation/simple/small-step/expressions"

// DoNothing is noop statement
type DoNothing struct{}

func (dn DoNothing) String() string {
	return "do-nothing"
}

// Reducible defines if statement can be simplified
func (dn DoNothing) Reducible() bool {
	return false
}

// Reduce simplifies an expression
func (dn DoNothing) Reduce(environment map[string]expr.Expression) (Statement, map[string]expr.Expression) {
	return DoNothing{}, environment
}
