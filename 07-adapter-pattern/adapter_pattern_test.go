package adapter_pattern

import (
	"fmt"
	"testing"
)

func TestAdapterPattern(T *testing.T) {
	duck := MallardDuck{}
	turkey := WildTurkey{}
	turkeyAdapter := NewTurkeyAdapter(turkey)
	fmt.Println("Turkey says...")
	turkey.gobble()
	turkey.fly()

	fmt.Println("Duck says...")
	duck.quack()
	duck.fly()

	fmt.Println("TurkeyAdapter says...")
	turkeyAdapter.quack()
	turkeyAdapter.fly()
}
