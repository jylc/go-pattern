package iterator_pattern

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}

func (mi *MenuItem) getName() string {
	return mi.name
}

func (mi *MenuItem) getDescription() string {
	return mi.description
}

func (mi *MenuItem) getPrice() float64 {
	return mi.price
}
func (mi *MenuItem) isVegetarian() bool {
	return mi.vegetarian
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		name:        name,
		description: description,
		vegetarian:  vegetarian,
		price:       price,
	}
}
