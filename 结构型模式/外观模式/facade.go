package 外观模式

import "fmt"

/*
外观模式（Facade Pattern）隐藏系统的复杂性，并向客户端提供了一个客户端可以访问系统的接口。这种类型的设计模式属于结构型模式，它向现有的系统添加一个接口，来隐藏系统的复杂性。

这种模式涉及到一个单一的类，该类提供了客户端请求的简化方法和对现有系统类方法的委托调用。

关键代码：在客户端和复杂系统之间再加一层，这一层将调用顺序、依赖关系等处理好。

意图：为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。

应用实例： 1、去医院看病，可能要去挂号、门诊、划价、取药，让患者或患者家属觉得很复杂，如果有提供接待人员，只让接待人员来处理，就很方便。 2、JAVA 的三层开发模式。
*/

/*第一步 创建形状 shape 接口*/
type shape interface {
	draw()
}

/*第二步 创建 shape 接口实现类*/
type circle struct{}

func (c *circle) draw() {
	fmt.Println("draw a circle")
}

type square struct{}

func (s *square) draw() {
	fmt.Println("draw a square")
}

/*第三步 创建装饰器*/

type facade struct {
	square shape
	circle shape
}

func (f *facade) DrawSquare() {
	f.square.draw()
}
func (f *facade) DrawCircle() {
	f.circle.draw()
}

func NewFacade() *facade {
	return &facade{
		&square{},
		&circle{},
	}
}
