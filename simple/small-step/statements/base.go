package statements

import (
	expr "github.com/m0t0r/computation/simple/small-step/expressions"
)

// Statement is a base type
type Statement interface {
	String() string
	Reducible() bool
	Reduce(environment map[string]expr.Expression) (Statement, map[string]expr.Expression)
}
