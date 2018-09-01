package expressions

import "fmt"

// Variable is a varible type
type Variable struct {
	Name string
}

func (v Variable) String() string {
	return fmt.Sprintf("%s", v.Name)
}

// Evaluate simplifies an expression
func (v Variable) Evaluate(environment map[string]Expression) Expression {
	return environment[v.Name]
}
