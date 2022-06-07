package iterator_pattern

import "fmt"

type Iterator interface {
	hasNext() bool
	next() interface{}
}

type DinerMenuIterator struct {
	items    []*MenuItem
	position int
}

func (d *DinerMenuIterator) hasNext() bool {
	if d.position >= len(d.items) || d.items[d.position] == nil {
		return false
	} else {
		return true
	}
}

func (d *DinerMenuIterator) next() interface{} {
	item := d.items[d.position]
	d.position += 1
	return item
}

func NewDinerMenuIterator(items []*MenuItem) *DinerMenuIterator {
	return &DinerMenuIterator{
		items:    items,
		position: 0,
	}
}

type DinerMenu struct {
	maxItems      int
	numberOfItems int
	menuItems     []*MenuItem
}

func (d *DinerMenu) addItem(name, description string, vegetarian bool, price float64) {
	menu := NewMenuItem(name, description, vegetarian, price)
	if d.numberOfItems >= d.maxItems {
		fmt.Println("Sorry, menu is full!!!")
	} else {
		d.menuItems[d.numberOfItems] = menu
		d.numberOfItems++
	}
}

func (d *DinerMenu) createIterator() Iterator {
	return NewDinerMenuIterator(d.menuItems)
}

func NewDinerMenu() *DinerMenu {
	maxItems := 6
	return &DinerMenu{
		maxItems:      maxItems,
		numberOfItems: 0,
		menuItems:     make([]*MenuItem, maxItems),
	}
}
