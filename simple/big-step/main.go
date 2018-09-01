package main

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/big-step/expressions"
)

func main() {
	var env = make(map[string]expr.Expression)
	env["x"] = expr.Number{Value: 2}
	env["y"] = expr.Number{Value: 5}

	fmt.Println(expr.LessThan{
		Left:  expr.Add{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 2}},
		Right: expr.Variable{Name: "y"},
	}.Evaluate(env))
}
