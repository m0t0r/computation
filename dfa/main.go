package main

import (
	"fmt"

	r "github.com/m0t0r/computation/dfa/rulebook"
)

func main() {

	rules := []r.FARule{
		r.FARule{
			State:     1,
			Character: "a",
			NextState: 2,
		},
		r.FARule{
			State:     1,
			Character: "b",
			NextState: 1,
		},
		r.FARule{
			State:     2,
			Character: "a",
			NextState: 2,
		},
		r.FARule{
			State:     2,
			Character: "b",
			NextState: 3,
		},
		r.FARule{
			State:     3,
			Character: "a",
			NextState: 3,
		},
		r.FARule{
			State:     3,
			Character: "b",
			NextState: 3,
		},
	}

	rulebook := r.DFARulebook{Rules: rules}

	fmt.Println(rulebook.NextState(1, "a"))
	fmt.Println(rulebook.NextState(1, "b"))
	fmt.Println(rulebook.NextState(2, "b"))
}
