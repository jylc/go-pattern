package factory_method

import "fmt"

type IPizza interface {
	prepare()
	bake()
	cut()
	box()
	getName() string
	setName(name string)
}

type Pizza struct {
	IPizza
	name      string
	dough     Dough
	sauce     Sauce
	veggies   []Veggies
	cheese    Cheese
	pepperoni Pepperoni
	clams     Clams
}

func (p *Pizza) bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *Pizza) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}
func (p *Pizza) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *Pizza) getName() string {
	return p.name
}

func (p *Pizza) setName(name string) {
	p.name = name
}

func NewPizza() *Pizza {
	return &Pizza{}
}

type NYStyleCheesePizza struct {
	*Pizza
	ingredientFactory IPizzaIngredientFactory
}

func (ny *NYStyleCheesePizza) prepare() {
	fmt.Println("Preparing ", ny.name)
	ny.dough = ny.ingredientFactory.createDough()
	ny.sauce = ny.ingredientFactory.createSauce()
	ny.cheese = ny.ingredientFactory.createCheese()
}

func NewNYStyleCheesePizza(ingredientFactory IPizzaIngredientFactory) *NYStyleCheesePizza {
	pizza := &NYStyleCheesePizza{
		Pizza:             NewPizza(),
		ingredientFactory: ingredientFactory,
	}
	return pizza
}
