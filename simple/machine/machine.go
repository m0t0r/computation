package machine

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/expression"
)

// Machine is an abstract machine type
type Machine struct {
	Expression  expr.Expression
	Environment map[string]expr.Expression
}

// Step simplifies expression
func (m *Machine) Step() {
	m.Expression, _ = m.Expression.Reduce(m.Environment)
}

// Run starts the process to simplify an expression
func (m Machine) Run() {
	for m.Expression.Reducible() {
		fmt.Println(m.Expression)
		m.Step()
	}
	fmt.Println(m.Expression)
}
