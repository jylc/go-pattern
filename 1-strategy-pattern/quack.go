package strategy

import "fmt"

type MuteQuack struct {
}

func (q *MuteQuack) quack() {
	fmt.Println("<<Silence>>")
}

type Squeak struct {
}

func (q *Squeak) quack() {
	fmt.Println("Squeak")
}

type Quack struct {
}

func (q *Quack) quack() {
	fmt.Println("Quack")
}
