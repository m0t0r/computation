package rulebook

import (
	"fmt"
)

// FARule ...
type FARule struct {
	State     int
	Character string
	NextState int
}

func (f FARule) String() string {
	return fmt.Sprintf("#<FARule #{%d} --#{%s}--> #{%d}>", f.State, f.Character, f.NextState)
}

// AppliesTo ...
func (f FARule) AppliesTo(state int, character string) bool {
	return f.State == state && f.Character == character
}

// Follow ...
func (f FARule) Follow() int {
	return f.NextState
}

// DFARulebook ...
type DFARulebook struct {
	Rules []FARule
}

// RuleFor ...
func (d DFARulebook) RuleFor(state int, character string) FARule {
	var result FARule
	for _, r := range d.Rules {
		if r.AppliesTo(state, character) {
			result = r
			break
		}
	}

	return result
}

// NextState ...
func (d DFARulebook) NextState(state int, character string) int {
	return d.RuleFor(state, character).Follow()
}
