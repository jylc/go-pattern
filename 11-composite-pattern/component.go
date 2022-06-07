package composite_pattern

import (
	"errors"
	"fmt"
)

type Iterator interface {
	hasNext() bool
	next() interface{}
}

type MenuComponentIterator struct {
	menuComponents []IMenuComponent
	position       int
}

func (m *MenuComponentIterator) hasNext() bool {
	if m.position >= len(m.menuComponents) {
		return false
	} else {
		return true
	}
}

func (m *MenuComponentIterator) next() interface{} {
	component := m.menuComponents[m.position]
	m.position += 1
	return component
}

func NewMenuComponentIterator(menuComponents []IMenuComponent) *MenuComponentIterator {
	return &MenuComponentIterator{
		menuComponents: menuComponents,
		position:       0,
	}
}

type IMenuComponent interface {
	add(menuComponent IMenuComponent) error
	remove(menuComponent IMenuComponent) error
	getChild(i int) (IMenuComponent, error)
	getName() (string, error)
	getDescription() (string, error)
	getPrice() (float64, error)
	isVegetarian() (bool, error)
	print() error
}

type MenuComponent struct {
}

func (m *MenuComponent) add(menuComponent IMenuComponent) error {
	err := errors.New("unsupported operation")
	return err
}

func (m *MenuComponent) remove(menuComponent IMenuComponent) error {
	err := errors.New("unsupported operation")
	return err
}

func (m *MenuComponent) getChild(i int) (IMenuComponent, error) {
	err := errors.New("unsupported operation")
	return nil, err
}

func (m *MenuComponent) getName() (string, error) {
	err := errors.New("unsupported operation")
	return "", err
}
func (m *MenuComponent) getDescription() (string, error) {
	err := errors.New("unsupported operation")
	return "", err
}
func (m *MenuComponent) getPrice() (float64, error) {
	err := errors.New("unsupported operation")
	return 0, err
}

func (m *MenuComponent) isVegetarian() (bool, error) {
	err := errors.New("unsupported operation")
	return false, err
}
func (m *MenuComponent) print() error {
	err := errors.New("unsupported operation")
	return err
}

type MenuItem struct {
	MenuComponent
	name        string
	description string
	vegetarian  bool
	price       float64
}

func (m *MenuItem) getName() (string, error) {
	return m.name, nil
}

func (m *MenuItem) getDescription() (string, error) {
	return m.description, nil
}

func (m *MenuItem) isVegetarian() (bool, error) {
	return m.vegetarian, nil
}

func (m *MenuItem) getPrice() (float64, error) {
	return m.price, nil
}

func (m *MenuItem) print() error {
	name, _ := m.getName()
	vegetarian, _ := m.isVegetarian()
	price, _ := m.getPrice()
	description, _ := m.getDescription()
	fmt.Println(" ", name, " ", vegetarian, " ", price, " ", description)
	return nil
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		MenuComponent: MenuComponent{},
		name:          name,
		description:   description,
		vegetarian:    vegetarian,
		price:         price,
	}
}

type Menu struct {
	MenuComponent
	menuComponents []IMenuComponent
	name           string
	description    string
}

func (m *Menu) add(menuComponent IMenuComponent) error {
	m.menuComponents = append(m.menuComponents, menuComponent)
	return nil
}

func (m *Menu) remove(menuComponent IMenuComponent) error {
	for i := 0; i < len(m.menuComponents); i++ {
		if m.menuComponents[i] == menuComponent {
			m.menuComponents = append(m.menuComponents[:i], m.menuComponents[i+1:]...)
			break
		}
	}
	return nil
}

func (m *Menu) getChild(i int) (IMenuComponent, error) {
	return m.menuComponents[i], nil
}

func (m *Menu) getName() (string, error) {
	return m.name, nil
}

func (m *Menu) getDescription() (string, error) {
	return m.description, nil
}

func (m *Menu) print() error {
	name, _ := m.getName()
	description, _ := m.getDescription()
	fmt.Println(" ", name, " ", description)
	fmt.Println("--------------------------")

	iterator := NewMenuComponentIterator(m.menuComponents)
	for iterator.hasNext() {
		menuComponent := iterator.next().(IMenuComponent)
		_ = menuComponent.print()
	}
	return nil
}

func NewMenu(name, description string) *Menu {
	return &Menu{
		menuComponents: make([]IMenuComponent, 0),
		name:           name,
		description:    description,
	}
}

type Waitress struct {
	allMenus IMenuComponent
}

func (w *Waitress) printMenu() {
	_ = w.allMenus.print()
}

func NewWaitress(allMenus IMenuComponent) *Waitress {
	return &Waitress{allMenus: allMenus}
}
