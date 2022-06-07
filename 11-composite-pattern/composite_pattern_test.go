package composite_pattern

import "testing"

func TestCompositePattern(T *testing.T) {
	pancakeHouseMenu := NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	dinerMenu := NewMenu("DINER MENU", "Lunch")
	cafeMenu := NewMenu("CAFE MENU", "Dinner")
	dessertMenu := NewMenu("DESSERT MENU", "Dessert of course")
	allMenus := NewMenu("ALL MENUS", "All menus combined")

	_ = allMenus.add(pancakeHouseMenu)
	_ = allMenus.add(dinerMenu)
	_ = allMenus.add(cafeMenu)
	_ = dinerMenu.add(&MenuItem{
		MenuComponent: MenuComponent{},
		name:          "Pasta",
		description:   "Spaghetti with Marinara Sauce",
		vegetarian:    true,
		price:         3.89,
	})
	_ = dinerMenu.add(dessertMenu)
	_ = dessertMenu.add(&MenuItem{
		MenuComponent: MenuComponent{},
		name:          "Apple Pie",
		description:   "Apple Pie with a flakey crust",
		vegetarian:    true,
		price:         1.59,
	})
	waitress := NewWaitress(allMenus)
	waitress.printMenu()
}
