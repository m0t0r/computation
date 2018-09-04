package main

import (
	"fmt"

	expr "github.com/m0t0r/computation/simple/big-step/expressions"
	stm "github.com/m0t0r/computation/simple/big-step/statements"
)

func main() {

	fmt.Println("Test 1 - Evaluating Expressions")

	var env1 = make(map[string]expr.Expression)
	env1["x"] = expr.Number{Value: 2}
	env1["y"] = expr.Number{Value: 5}

	left := expr.Add{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 2}}

	fmt.Println(expr.LessThan{
		Left:  left,
		Right: expr.Variable{Name: "y"},
	}.Evaluate(env1))

	fmt.Println("Test 2 - Evaluating Statements")

	var env2 = make(map[string]expr.Expression)

	fmt.Println(stm.Sequence{
		First:  stm.Assign{Name: "x", Expression: expr.Add{Left: expr.Number{Value: 1}, Right: expr.Number{Value: 1}}},
		Second: stm.Assign{Name: "y", Expression: expr.Add{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 3}}},
	}.Evaluate(env2))

	fmt.Println("Test 3 - Evaluating While Statement")

	var env3 = make(map[string]expr.Expression)
	env3["x"] = expr.Number{Value: 1}

	fmt.Println(stm.While{
		Condition: expr.LessThan{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 5}},
		Body:      stm.Assign{Name: "x", Expression: expr.Multiply{Left: expr.Variable{Name: "x"}, Right: expr.Number{Value: 3}}},
	}.Evaluate(env3))
}
