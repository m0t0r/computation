package main

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/small-step/expressions"
	"github.com/m0t0r/computation/simple/small-step/machine"
	stm "github.com/m0t0r/computation/simple/small-step/statements"
)

func main() {
	var m1, m2, m3, m4 machine.Machine

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

	fmt.Print("Machine 3 Test - Sequence Statement Evaluation\n\n")

	env = make(map[string]expr.Expression)

	m3 = machine.Machine{
		Statement: stm.Sequence{
			First:  stm.Assign{Name: "x", Expression: expr.Add{Left: expr.Number{Value: 1}, Right: expr.Number{Value: 1}}},
			Second: stm.Assign{Name: "y", Expression: expr.Add{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 3}}},
		},
		Environment: env,
	}

	m3.Run()

	fmt.Print("Machine 4 Test - While Statement Evaluation\n\n")

	env = make(map[string]expr.Expression)
	env["x"] = expr.Number{Value: 1}

	m4 = machine.Machine{
		Statement: stm.While{
			Condition: expr.LessThan{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 5}},
			Body:      stm.Assign{Name: "x", Expression: expr.Multiply{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 3}}},
		},
		Environment: env,
	}

	m4.Run()
}
