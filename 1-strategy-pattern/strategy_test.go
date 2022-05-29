package strategy

import "testing"

func TestStrategyPattern(t *testing.T) {
	d := NewMallardDuck()
	d.setFly(&FlyWithWings{})
	d.setQuack(&Quack{})
	d.performFly()
	d.performQuack()
}
