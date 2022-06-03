package factory_method

type IFactory interface {
	createPizza(t string) IPizza
}
type PizzaStore struct {
	createPizza func(item string) IPizza
}

func (ps *PizzaStore) orderPizza(t string) IPizza {
	pizza := ps.createPizza(t)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return pizza
}

type NYPizzaStore struct {
	PizzaStore
}

func (ny *NYPizzaStore) createPizza(item string) IPizza {
	ingredientFactory := NewNYPizzaIngredientFactory()
	if item == "cheese" {
		pizza := NewNYStyleCheesePizza(ingredientFactory)
		pizza.setName("New York Style Cheese Pizza")
		return pizza
	} else if item == "veggie" {
		return nil
	} else if item == "clam" {
		return nil

	} else if item == "pepperoni" {
		return nil
	} else {
		return nil
	}
}
func NewNYPizzaStore() *NYPizzaStore {
	nyStore := &NYPizzaStore{}
	nyStore.PizzaStore.createPizza = nyStore.createPizza //模拟java的抽象类
	return nyStore
}
