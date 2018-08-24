package expressions

// Expression is a base type
type Expression interface {
	String() string
	GetValue() ExpressionValue
	Reducible() bool
	Reduce(environment map[string]Expression) (Expression, error)
}

// ExpressionValue holds the value of expression
type ExpressionValue interface {
}
