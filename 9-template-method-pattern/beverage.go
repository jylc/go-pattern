package template_method_pattern

import "fmt"

type ICaffeineBeverage interface {
	prepareRecipe()
	brew()
	addCondiments()
	boilWater()
	pourInCup()
}
type CaffeineBeverage struct {
	concrete ICaffeineBeverage
}

func (cb CaffeineBeverage) prepareRecipe() {
	cb.concrete.boilWater()
	cb.concrete.brew()
	cb.concrete.pourInCup()
	cb.concrete.addCondiments()
}
func (cb CaffeineBeverage) brew() {
	cb.concrete.brew()
}
func (cb CaffeineBeverage) addCondiments() {
	cb.concrete.addCondiments()
}
func (cb CaffeineBeverage) boilWater() {
	fmt.Println("Boiling water")
}
func (cb CaffeineBeverage) pourInCup() {
	fmt.Println("Pouring into cup")
}

type Tea struct {
	CaffeineBeverage
}

func (t Tea) brew() {
	fmt.Println("Steeping the tea")
}

func (t Tea) addCondiments() {
	fmt.Println("Adding Lemon")
}

type Coffee struct {
	CaffeineBeverage
}

func (c Coffee) brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (c Coffee) addCondiments() {
	fmt.Println("Adding Sugar and Milk")
}
