package main

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/expression"
	"github.com/m0t0r/computation/simple/machine"
	stm "github.com/m0t0r/computation/simple/statement"
)

func main() {
	var m1, m2 machine.Machine

	fmt.Print("Machine 1 Test - Statements Evaluation\n\n")

	env := make(map[string]expr.Expression)
	env["x"] = expr.Number{Value: 2}

	m1 = machine.Machine{
		Statement: stm.Assign{
			Name:       "x",
			Expression: expr.Add{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 1}},
		},
		Environment: env,
	}

	m1.Run()

	fmt.Print("Machine 2 Test - If Statement Evaluation\n\n")

	env = make(map[string]expr.Expression)
	env["x"] = expr.Boolean{Value: true}

	m2 = machine.Machine{
		Statement: stm.If{
			Condition:   expr.Variable{Name: "x"},
			Consequence: stm.Assign{Name: "y", Expression: expr.Number{Value: 1}},
			Alternative: stm.Assign{Name: "y", Expression: expr.Number{Value: 2}},
		},
		Environment: env,
	}

	m2.Run()
}
