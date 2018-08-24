package statement

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/expression"
)

// Sequence chains staements eveluation
type Sequence struct {
	First  Statement
	Second Statement
}

func (s Sequence) String() string {
	return fmt.Sprintf("%v; %v", s.First.String(), s.Second.String())
}

// Reducible defines if statement can be simplified
func (s Sequence) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (s Sequence) Reduce(environment map[string]expr.Expression) (Statement, map[string]expr.Expression) {
	if (s.First == DoNothing{}) {
		return s.Second, environment
	}

	reducedFirst, reducedEnvironment := s.First.Reduce(environment)
	return Sequence{First: reducedFirst, Second: s.Second}, reducedEnvironment
}
