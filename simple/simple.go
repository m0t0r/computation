package main

import (
	"errors"
	"fmt"
)

// Expression is a base type
type Expression interface {
	String() string
	Value() int
	Reducible() bool
	Reduce() (Expression, error)
}

// ------------------- NUMBER ------------------- //

// Number is number type
type Number struct {
	value int
}

// Value returns the value
func (n Number) Value() int {
	return n.value
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.value)
}

// Reducible defines if expression can be simplified
func (n Number) Reducible() bool {
	return false
}

// Reduce simplifies an expression
func (n Number) Reduce() (Expression, error) {
	return Number{value: -1}, errors.New("Error: Number cannot be reduced")
}

// ------------------- ADD ------------------- //

// Add is an addition type
type Add struct {
	left  Expression
	right Expression
}

// Value returns the value
func (a Add) Value() int {
	return a.left.Value() + a.right.Value()
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.left.String(), a.right.String())
}

// Reducible defines if expression can be simplified
func (a Add) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (a Add) Reduce() (Expression, error) {
	if a.left.Reducible() {
		reduce, _ := a.left.Reduce()
		return Add{left: reduce, right: a.right}, nil
	} else if a.right.Reducible() {
		reduce, _ := a.right.Reduce()
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
func (m Multiply) Value() int {
	return m.left.Value() * m.right.Value()
}

func (m Multiply) String() string {
	return fmt.Sprintf("%s * %s", m.left.String(), m.right.String())
}

// Reducible defines if expression can be simplified
func (m Multiply) Reducible() bool {
	return true
}

// Reduce simplifies an expression
func (m Multiply) Reduce() (Expression, error) {
	if m.left.Reducible() {
		reduce, _ := m.left.Reduce()
		return Multiply{left: reduce, right: m.right}, nil
	} else if m.right.Reducible() {
		reduce, _ := m.right.Reduce()
		return Multiply{left: m.left, right: reduce}, nil
	} else {
		return Number{value: m.Value()}, nil
	}
}

// ------------------- MACHINE ------------------- //

// Machine is a abstract machine type
type Machine struct {
	expression Expression
}

// Step simplifies expression
func (m *Machine) Step() {
	m.expression, _ = m.expression.Reduce()
}

// Run runs the process to simplify expression
func (m Machine) Run() {
	for m.expression.Reducible() {
		fmt.Println(m.expression)
		m.Step()
	}
	fmt.Println(m.expression)
}

func main() {
	var m Machine

	m = Machine{
		expression: Add{
			left:  Multiply{left: Number{value: 1}, right: Number{value: 2}},
			right: Multiply{left: Number{value: 3}, right: Number{value: 4}},
		},
	}

	m.Run()
}
