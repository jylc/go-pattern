package state_pattern

import "fmt"

type GumballMachine struct {
	soldOut    State
	noQuarter  State
	hasQuarter State
	sold       State

	state State
	count int
}

func (g *GumballMachine) insertQuarter() {
	g.state.insertQuarter()
}

func (g *GumballMachine) ejectQuarter() {
	g.state.ejectQuarter()
}

func (g *GumballMachine) turnCrank() {
	g.state.turnCrank()
	g.state.dispense()
}

func (g *GumballMachine) dispense() {

}

func (g *GumballMachine) releaseBall() {
	fmt.Println("A gumball comes rolling out the slot...")
	if g.count != 0 {
		g.count = g.count - 1
	}
}

func (g *GumballMachine) setState(state State) {
	g.state = state
}

func (g *GumballMachine) getNoQuarterState() State {
	return g.noQuarter
}
func (g *GumballMachine) getSoldState() State {
	return g.sold
}

func (g *GumballMachine) getCount() int {
	return g.count
}

func (g *GumballMachine) getHasQuarterState() State {
	return g.hasQuarter
}
func NewGumballMachine(count int) *GumballMachine {
	machine := &GumballMachine{
		count: count,
	}
	machine.sold = NewSoldState(machine)
	machine.noQuarter = NewNoQuarterState(machine)
	machine.soldOut = NewSoldOutState(machine)
	machine.hasQuarter = NewHasQuarterState(machine)
	machine.state = machine.soldOut
	if count > 0 {
		machine.state = machine.noQuarter
	}
	return machine
}
