package state_pattern

import "fmt"

type State interface {
	insertQuarter()
	ejectQuarter()
	turnCrank()
	dispense()
}

type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func (s *NoQuarterState) insertQuarter() {
	fmt.Println("You inserted a quarter")
	s.gumballMachine.setState(s.gumballMachine.hasQuarter)
}
func (s *NoQuarterState) ejectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}
func (s *NoQuarterState) turnCrank() {
	fmt.Println("You turned, but there's no quarter")
}
func (s *NoQuarterState) dispense() {
	fmt.Println("You need to pay first")
}

func NewNoQuarterState(gumballMachine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{gumballMachine: gumballMachine}
}

type HasQuarterState struct {
	gumballMachine *GumballMachine
}

func (s *HasQuarterState) insertQuarter() {
	fmt.Println("You can't inserted a quarter")
}
func (s *HasQuarterState) ejectQuarter() {
	fmt.Println("Quarter returned")
	s.gumballMachine.setState(s.gumballMachine.getNoQuarterState())
}
func (s *HasQuarterState) turnCrank() {
	fmt.Println("You turned...")
	s.gumballMachine.setState(s.gumballMachine.getSoldState())
}
func (s *HasQuarterState) dispense() {
	fmt.Println("No gumball dispensed")
}

func NewHasQuarterState(gumballMachine *GumballMachine) *HasQuarterState {
	return &HasQuarterState{gumballMachine: gumballMachine}
}

type SoldState struct {
	gumballMachine *GumballMachine
}

func (s *SoldState) insertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}
func (s *SoldState) ejectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}
func (s *SoldState) turnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball")
}
func (s *SoldState) dispense() {
	s.gumballMachine.releaseBall()
	if s.gumballMachine.getCount() > 0 {
		s.gumballMachine.setState(s.gumballMachine.getNoQuarterState())
	} else {
		fmt.Println("Oops, out of gumballs")
		s.gumballMachine.setState(s.gumballMachine.getSoldState())
	}
}

func NewSoldState(gumballMachine *GumballMachine) *SoldState {
	return &SoldState{gumballMachine: gumballMachine}
}

type SoldOutState struct {
	gumballMachine *GumballMachine
}

func (s *SoldOutState) insertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}
func (s *SoldOutState) ejectQuarter() {
	fmt.Println("You can't eject, you haven.t inserted a quarter yet")
}
func (s *SoldOutState) turnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}
func (s *SoldOutState) dispense() {
	fmt.Println("No gumball dispensed")
}

func NewSoldOutState(gumballMachine *GumballMachine) *SoldOutState {
	return &SoldOutState{gumballMachine: gumballMachine}
}
