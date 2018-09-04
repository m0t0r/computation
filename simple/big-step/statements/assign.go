package statements

import (
	expr "github.com/m0t0r/computation/simple/big-step/expressions"
)

// Assign is an assignment statement
type Assign struct {
	Name       string
	Expression expr.Expression
}

// Evaluate processes the statement
func (a Assign) Evaluate(environment map[string]expr.Expression) map[string]expr.Expression {

	newEnvironment := make(map[string]expr.Expression)
	// copy key-values from original environment
	for name, expr := range environment {
		newEnvironment[name] = expr.Evaluate(environment)
	}

	newEnvironment[a.Name] = a.Expression.Evaluate(environment)

	return newEnvironment
}
