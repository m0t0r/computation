package main

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/expression"
	"github.com/m0t0r/computation/simple/machine"
)

func main() {
	var m1, m2, m3 machine.Machine

	fmt.Print("Machine 1 Test - Expression Reduction\n\n")

	m1 = machine.Machine{
		Expression: expr.Add{
			Left:  expr.Multiply{Left: expr.Number{Value: 1}, Right: expr.Number{Value: 2}},
			Right: expr.Multiply{Left: expr.Number{Value: 3}, Right: expr.Number{Value: 4}},
		},
	}

	m1.Run()

	fmt.Print("Machine 2 Test - Less Than comparison\n\n")

	m2 = machine.Machine{
		Expression: expr.LessThan{
			Left:  expr.Number{Value: 5},
			Right: expr.Add{Left: expr.Number{Value: 2}, Right: expr.Number{Value: 2}},
		},
	}

	m2.Run()

	fmt.Print("Machine 3 Test - Variables evaluation\n\n")

	env := make(map[string]expr.Expression)
	env["x"] = expr.Number{Value: 3}
	env["y"] = expr.Number{Value: 4}

	m3 = machine.Machine{
		Expression:  expr.Add{Left: expr.Variable{Name: "x"}, Right: expr.Variable{Name: "y"}},
		Environment: env,
	}

	m3.Run()

}
