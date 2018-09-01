package expressions

// Expression is a base type
type Expression interface {
	Evaluate(enviroment map[string]Expression) Expression
}

// ExpressionValue holds the value of expression
type ExpressionValue interface {
}
