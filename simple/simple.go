package main

import (
	"fmt"
)

// Expression is a base type
type Expression interface {
	String() string
	Reducible() bool
}

// Number is number type
type Number struct {
	value int
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.value)
}

// Reducible defines if expression can be simplified
func (n Number) Reducible() bool {
	return false
}

// Add is an addition type
type Add struct {
	left  Expression
	right Expression
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.left.String(), a.right.String())
}

// Reducible defines if expression can be simplified
func (a Add) Reducible() bool {
	return true
}

// Multiply is a multiplication type
type Multiply struct {
	left  Expression
	right Expression
}

func (m Multiply) String() string {
	return fmt.Sprintf("%s * %s", m.left.String(), m.right.String())
}

// Reducible defines if expression can be simplified
func (m Multiply) Reducible() bool {
	return true
}

func main() {

	n := Number{value: 5}
	a := Add{left: Number{value: 1}, right: Number{value: 2}}

	fmt.Println(n.Reducible())
	fmt.Println(a.Reducible())
}
