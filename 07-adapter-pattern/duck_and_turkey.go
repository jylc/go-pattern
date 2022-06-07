package adapter_pattern

import "fmt"

type Duck interface {
	quack()
	fly()
}

type MallardDuck struct {
}

func (d MallardDuck) quack() {
	fmt.Println("Quack")
}

func (d MallardDuck) fly() {
	fmt.Println("I'm flying")
}

type Turkey interface {
	gobble()
	fly()
}

type WildTurkey struct {
}

func (t WildTurkey) gobble() {
	fmt.Println("Gobble gobble")
}

func (t WildTurkey) fly() {
	fmt.Println("I'm flying a short distance")
}

type TurkeyAdapter struct {
	turkey Turkey
}

func (ta TurkeyAdapter) quack() {
	ta.turkey.gobble()
}

func (ta TurkeyAdapter) fly() {
	for i := 0; i < 5; i++ {
		ta.turkey.fly()
	}
}

func NewTurkeyAdapter(turkey Turkey) Duck {
	return &TurkeyAdapter{turkey: turkey}
}
