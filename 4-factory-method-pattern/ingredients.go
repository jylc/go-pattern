package factory_method

type Dough interface {
}

type ThinCrustDough struct {
	Dough
}

type Sauce interface {
}
type MarinaraSauce struct {
	Sauce
}

type Cheese interface {
}

type ReggianoCheese struct {
	Cheese
}

type Veggies interface {
}

type Garlic struct {
	Veggies
}

type Onion struct {
	Veggies
}
type Mushroom struct {
	Veggies
}

type RedPepper struct {
	Veggies
}

type Pepperoni interface {
}

type SlicedPepperoni struct {
	Pepperoni
}

type Clams interface {
}
type FreshClams struct {
	Clams
}

// IPizzaIngredientFactory 原料
type IPizzaIngredientFactory interface {
	createDough() Dough
	createSauce() Sauce
	createCheese() Cheese
	createVeggies() []Veggies
	createPepperoni() Pepperoni
	createClam() Clams
}

type NYPizzaIngredientFactory struct {
}

func (ny *NYPizzaIngredientFactory) createDough() Dough {
	return ThinCrustDough{}
}
func (ny *NYPizzaIngredientFactory) createSauce() Sauce {
	return MarinaraSauce{}
}
func (ny *NYPizzaIngredientFactory) createCheese() Cheese {
	return ReggianoCheese{}
}
func (ny *NYPizzaIngredientFactory) createVeggies() []Veggies {
	veggies := []Veggies{Garlic{}, Onion{}, Mushroom{}, RedPepper{}}
	return veggies
}
func (ny *NYPizzaIngredientFactory) createPepperoni() Pepperoni {
	return SlicedPepperoni{}
}
func (ny *NYPizzaIngredientFactory) createClam() Clams {
	return FreshClams{}
}

func NewNYPizzaIngredientFactory() *NYPizzaIngredientFactory {
	return &NYPizzaIngredientFactory{}
}
