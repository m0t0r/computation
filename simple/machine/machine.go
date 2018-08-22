package machine

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/expression"
	stm "github.com/m0t0r/computation/simple/statement"
)

// Machine is an Abstract Machine type
type Machine struct {
	Statement   stm.Statement
	Environment map[string]expr.Expression
}

// Step simplifies expression
func (m *Machine) Step() {
	m.Statement, m.Environment = m.Statement.Reduce(m.Environment)
}

// Run starts the process to simplify an expression
func (m Machine) Run() {
	for m.Statement.Reducible() {
		fmt.Println(m.Statement, m.Environment)
		m.Step()
	}
	fmt.Println(m.Statement, m.Environment)
}
