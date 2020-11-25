package 装饰器模式

import "fmt"

/*
装饰模式使用对象组合的方式动态改变或增加对象行为。

Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。

使用匿名组合，在装饰器中不必显式定义转调原对象方法。

装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。

优点：装饰类和被装饰类可以独立发展，不会相互耦合，装饰模式是继承的一个替代模式，装饰模式可以动态扩展一个实现类的功能。

缺点：多层装饰比较复杂。
*/
/*第一步 创建形状接口*/
type shape interface {
	draw()
}

/*第二步 创建接口实现类*/
type rectangle struct{}

func (r *rectangle) draw() {
	fmt.Println("draw a rectangle")
}

type circle struct{}

func (c *circle) draw() {
	fmt.Println("draw a circle")
}

/*第三步 创建匿名内嵌 shape 接口的结构体装饰器*/
type redDecorator struct {
	shape
}

//重写 shape 接口方法，实现装饰器功能
func (d *redDecorator) draw() {
	d.shape.draw()
	fmt.Println("Border Color: Red")
}

func NewRectangle() *rectangle {
	return new(rectangle)
}

func NewCircle() *circle {
	return new(circle)
}

/*第四步 构造函数*/
func NewRedDecorator(shape shape) *redDecorator {
	return &redDecorator{
		shape,
	}
}
