package state_pattern

import "testing"

func TestStatePattern(T *testing.T) {
	machine := NewGumballMachine(5)
	machine.insertQuarter()
	machine.turnCrank()
}
