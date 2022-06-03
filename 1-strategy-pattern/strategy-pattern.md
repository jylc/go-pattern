### 策略模式

定义算法族，分别封装起来，让它们之间可以相互替换，此模式让算法的变化独立于使用算法的客户

```go 
type MallardDuck struct{
}
func (d MallardDuck) quack(){}
func (d MallardDuck) fly(){}
func (d MallardDuck) display(){}

type ModelDuck struct{
}
func (d ModelDuck) quack(){}
func (d ModelDuck) fly(){}
func (d ModelDuck) display(){}
```

如果定义一个鸭子类，需要为其定义叫声与飞行方式，不同种类的鸭子的叫声与飞行方式不同要分别定义
；但如果一个又一个定义，代码就不能复用，因此将鸭子的共同行为抽取放入接口中，把不同的行为定义为
类，以方便复用。

#### 设计原则
* > 针对接口编程，而不是针对实现编程

    接口编程，可以在“运行时”动态地“改变”对象的行为
* > 多用组合，少用继承
  
    “有一个”可能比“是一个”更好；鸭子“有一个”飞行的行为，将鸭子与飞行的类组合使用而不是鸭子继承飞行的类
  使系统具有很大的弹性，不仅可以将算法族封装成类，更可以“在运行时动态地改变行为”，只要组合的行为对象符合
 正确的接口标准。

