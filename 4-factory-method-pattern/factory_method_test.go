package factory_method

import (
	"fmt"
	"testing"
)

func TestFactoryMethodPattern(T *testing.T) {
	nyStore := NewNYPizzaStore()
	pizza := nyStore.orderPizza("cheese")
	fmt.Println("Ethan ordered a ", pizza.getName())
}
