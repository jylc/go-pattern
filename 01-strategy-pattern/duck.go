package strategy

import "fmt"

type MallardDuck struct {
	*Duck
}

func (d MallardDuck) display() {
	fmt.Println("I'm a MallardDuck!!!")
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{NewDuck()}
}

type ModelDuck struct {
	*Duck
}

func (d ModelDuck) display() {
	fmt.Println("I'm a ModelDuck!!!")
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{NewDuck()}
}
