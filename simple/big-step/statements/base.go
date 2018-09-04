package statements

import (
	expr "github.com/m0t0r/computation/simple/big-step/expressions"
)

// Statement is a base type
type Statement interface {
	Evaluate(environment map[string]expr.Expression) map[string]expr.Expression
}
