package expressions

import "fmt"

// Variable is a varible type
type Variable struct {
	Name string
}

// GetValue returns the value
func (v Variable) GetValue() ExpressionValue {
	return v.Name
}

func (v Variable) String() string {
	return fmt.Sprintf("%s", v.Name)
}

// Reducible defines if expression can be simplified
func (v Variable) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (v Variable) Reduce(environment map[string]Expression) (Expression, error) {
	return environment[v.Name], nil
}
