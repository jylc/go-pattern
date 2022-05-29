package strategy

// Duck 鸭子的父类
type Duck struct {
	quack IQuack
	fly   IFly
}

func (d Duck) display() {

}
func (d *Duck) performQuack() {
	d.quack.quack()
}

func (d *Duck) performFly() {
	d.fly.fly()
}

func (d *Duck) setQuack(quack IQuack) {
	d.quack = quack
}

func (d *Duck) setFly(fly IFly) {
	d.fly = fly
}

func NewDuck() *Duck {
	return &Duck{}
}
