package decorator

import "fmt"

type Beverage interface {
	getDescription() string
	cost() float64
}

// Espresso 一种饮料
type Espresso struct {
	description string
}

func (e *Espresso) getDescription() string {
	return e.description
}

func (e *Espresso) cost() float64 {
	return 1.99
}

func NewEspresso() *Espresso {
	return &Espresso{description: "Espresso"}
}

// HouseBlend 一种饮料
type HouseBlend struct {
	description string
}

func (h *HouseBlend) getDescription() string {
	return h.description
}

func (h *HouseBlend) cost() float64 {
	return 0.89
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{
		description: "HouseBlend",
	}
}

// CondimentDecorator 调料
type CondimentDecorator interface {
	Beverage
}

// Mocha 一种调料
type Mocha struct {
	beverage Beverage
}

func (m *Mocha) getDescription() string {
	return fmt.Sprintf("%s, %s", m.beverage.getDescription(), " Mocha")
}

func (m *Mocha) cost() float64 {
	return m.beverage.cost() + 0.20
}

func NewMocha(beverage Beverage) *Mocha {
	return &Mocha{beverage: beverage}
}

// Whip 一种调料
type Whip struct {
	beverage Beverage
}

func (m *Whip) getDescription() string {
	return fmt.Sprintf("%s, %s", m.beverage.getDescription(), " Whip")
}

func (m *Whip) cost() float64 {
	return m.beverage.cost() + 0.20
}

func NewWhip(beverage Beverage) *Whip {
	return &Whip{beverage: beverage}
}
