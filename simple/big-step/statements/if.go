package statements

import expr "github.com/m0t0r/computation/simple/big-step/expressions"

// If is a conditional statement
type If struct {
	Condition   expr.Expression
	Consequence Statement
	Alternative Statement
}

// Evaluate processes the statement
func (_if If) Evaluate(environment map[string]expr.Expression) map[string]expr.Expression {
	switch c := _if.Condition.Evaluate(environment); c {
	case expr.Boolean{Value: true}:
		return _if.Consequence.Evaluate(environment)
	case expr.Boolean{Value: false}:
		return _if.Alternative.Evaluate(environment)
	default:
		return environment
	}
}
