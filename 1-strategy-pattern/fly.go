package strategy

import "fmt"

type FlyWithWings struct {
}

func (f *FlyWithWings) fly() {
	fmt.Println("I'm flying!!!")
}

type FlyNoWay struct {
}

func (f *FlyNoWay) fly() {
	fmt.Println("I can't fly!!!")
}
