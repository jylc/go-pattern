package template_method_pattern

import "testing"

func TestBeveragePattern(T *testing.T) {
	tea := &Tea{}
	caffeineBeverage := new(CaffeineBeverage)
	caffeineBeverage.concrete = tea
	tea.CaffeineBeverage = *caffeineBeverage
	tea.prepareRecipe()

	coffee := &Coffee{}
	caffeineBeverage.concrete = coffee
	coffee.CaffeineBeverage = *caffeineBeverage
	coffee.prepareRecipe()
}
