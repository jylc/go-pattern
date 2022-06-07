package decorator

import (
	"fmt"
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	espresso := NewEspresso()
	mocha := NewMocha(espresso)
	whip := NewWhip(mocha)
	fmt.Println(whip.getDescription())
}
