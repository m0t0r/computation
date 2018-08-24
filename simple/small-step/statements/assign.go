package statements

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/small-step/expressions"
)

// Assign is an assignment statement
type Assign struct {
	Name       string
	Expression expr.Expression
}

func (a Assign) String() string {
	return fmt.Sprintf("%s = %v", a.Name, a.Expression)
}

// Reducible defines if statement can be simplified
func (a Assign) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (a Assign) Reduce(environment map[string]expr.Expression) (Statement, map[string]expr.Expression) {
	if a.Expression.Reducible() {
		reduce, _ := a.Expression.Reduce(environment)
		return Assign{Name: a.Name, Expression: reduce}, environment
	}

	newEnvironment := make(map[string]expr.Expression)
	// copy key-values from original environment
	for key, value := range environment {
		newEnvironment[key] = value
	}

	newEnvironment[a.Name] = a.Expression

	return DoNothing{}, newEnvironment
}
