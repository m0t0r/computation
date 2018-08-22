package main

import (
	"errors"
	"fmt"
)

// Expression is a base type
type Expression interface {
	String() string
	Value() ExpressionValue
	Reducible() bool
	Reduce(environment map[string]Expression) (Expression, error)
}

// ExpressionValue holds the value of expression
type ExpressionValue interface{}

// ------------------- NUMBER ------------------- //

// Number is number type
type Number struct {
	value ExpressionValue
}

// Value returns the value
func (n Number) Value() ExpressionValue {
	return n.value.(int)
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.value)
}

// Reducible defines if expression can be simplified
func (n Number) Reducible() bool {
	return false
}

// Reduce simplifies an expression
func (n Number) Reduce(environment map[string]Expression) (Expression, error) {
	return nil, errors.New("Error: Number cannot be reduced")
}

// ------------------- BOOLEAN ------------------- //

// Boolean is number type
type Boolean struct {
	value ExpressionValue
}

// Value returns the value
func (b Boolean) Value() ExpressionValue {
	return b.value.(bool)
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b.value.(bool))
}

// Reducible defines if expression can be simplified
func (b Boolean) Reducible() bool {
	return false
}

// Reduce simplifies an expression
func (b Boolean) Reduce(environment map[string]Expression) (Expression, error) {
	return nil, errors.New("Error: Boolean cannot be reduced")
}

// ------------------- ADD ------------------- //

// Add is an addition type
type Add struct {
	left  Expression
	right Expression
}

// Value returns the value
func (a Add) Value() ExpressionValue {
	var leftVal interface{} = a.left.Value()
	var rightVal interface{} = a.right.Value()

	l, ok := leftVal.(int)
	r, ok := rightVal.(int)

	if ok {
		return l + r
	}

	return errors.New("Error: Unable to calculate value")
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.left.String(), a.right.String())
}

// Reducible defines if expression can be simplified
func (a Add) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (a Add) Reduce(environment map[string]Expression) (Expression, error) {
	if a.left.Reducible() {
		reduce, _ := a.left.Reduce(environment)
		return Add{left: reduce, right: a.right}, nil
	} else if a.right.Reducible() {
		reduce, _ := a.right.Reduce(environment)
		return Add{left: a.left, right: reduce}, nil
	} else {
		return Number{value: a.Value()}, nil
	}
}

// ------------------- MULTIPLY ------------------- //

// Multiply is a multiplication type
type Multiply struct {
	left  Expression
	right Expression
}

// Value returns the value
func (m Multiply) Value() ExpressionValue {
	var leftVal interface{} = m.left.Value()
	var rightVal interface{} = m.right.Value()

	l, ok := leftVal.(int)
	r, ok := rightVal.(int)

	if ok {
		return l * r
	}

	return errors.New("Error: Unable to calculate value")
}

func (m Multiply) String() string {
	return fmt.Sprintf("%s * %s", m.left.String(), m.right.String())
}

// Reducible defines if expression can be simplified
func (m Multiply) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (m Multiply) Reduce(environment map[string]Expression) (Expression, error) {
	if m.left.Reducible() {
		reduce, _ := m.left.Reduce(environment)
		return Multiply{left: reduce, right: m.right}, nil
	} else if m.right.Reducible() {
		reduce, _ := m.right.Reduce(environment)
		return Multiply{left: m.left, right: reduce}, nil
	} else {
		return Number{value: m.Value()}, nil
	}
}

// ------------------- LESS THAN ------------------- //

// LessThan is a "less than" comparison type
type LessThan struct {
	left  Expression
	right Expression
}

// Value returns the value
func (lt LessThan) Value() ExpressionValue {
	var leftVal interface{} = lt.left.Value()
	var rightVal interface{} = lt.right.Value()

	l, ok := leftVal.(int)
	r, ok := rightVal.(int)

	if ok {
		return l < r
	}

	return errors.New("Error: Unable to calculate value")
}

func (lt LessThan) String() string {
	return fmt.Sprintf("%s < %s", lt.left.String(), lt.right.String())
}

// Reducible defines if expression can be simplified
func (lt LessThan) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (lt LessThan) Reduce(environment map[string]Expression) (Expression, error) {
	if lt.left.Reducible() {
		reduce, _ := lt.left.Reduce(environment)
		return LessThan{left: reduce, right: lt.right}, nil
	} else if lt.right.Reducible() {
		reduce, _ := lt.right.Reduce(environment)
		return LessThan{left: lt.left, right: reduce}, nil
	} else {
		return Boolean{value: lt.Value()}, nil
	}
}

// ------------------- VARIABLE ------------------- //

// Variable is a varible type
type Variable struct {
	name string
}

// Value returns the value
func (v Variable) Value() ExpressionValue {
	return v.Value
}

func (v Variable) String() string {
	return fmt.Sprintf("%s", v.name)
}

// Reducible defines if expression can be simplified
func (v Variable) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (v Variable) Reduce(environment map[string]Expression) (Expression, error) {
	return environment[v.name], nil
}

// ------------------- MACHINE ------------------- //

// Machine is an abstract machine type
type Machine struct {
	expression  Expression
	environment map[string]Expression
}

// Step simplifies expression
func (m *Machine) Step() {
	m.expression, _ = m.expression.Reduce(m.environment)
}

// Run starts the process to simplify an expression
func (m Machine) Run() {
	for m.expression.Reducible() {
		fmt.Println(m.expression)
		m.Step()
	}
	fmt.Println(m.expression)
}

func main() {
	var m1, m2, m3 Machine

	fmt.Print("Machine 1 Test - Expression Reduction\n\n")

	m1 = Machine{
		expression: Add{
			left:  Multiply{left: Number{value: 1}, right: Number{value: 2}},
			right: Multiply{left: Number{value: 3}, right: Number{value: 4}},
		},
	}

	m1.Run()

	fmt.Print("Machine 2 Test - Less Than comparison\n\n")

	m2 = Machine{
		expression: LessThan{
			left:  Number{value: 5},
			right: Add{left: Number{value: 2}, right: Number{value: 2}},
		},
	}

	m2.Run()

	fmt.Print("Machine 3 Test - Variables evaluation\n\n")

	env := make(map[string]Expression)
	env["x"] = Number{value: 3}
	env["y"] = Number{value: 4}

	m3 = Machine{
		expression:  Add{left: Variable{name: "x"}, right: Variable{name: "y"}},
		environment: env,
	}

	m3.Run()

}
