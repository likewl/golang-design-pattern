package 访问者模式

import "fmt"

/*
@Author LiKe
@Description //TODO
@Time 2020/12/7 18:59
*/

/*
给一系列对象透明的添加功能，并且把相关代码封装到一个类中。对象只要预留访问者接口Accept则后期为对象添加功能的时候就不需要改动对象。

主要解决：稳定的数据结构和易变的操作耦合问题。

关键代码：在数据基础类里面有一个方法接受访问者，将自身引用传入访问者。

优点： 1、符合单一职责原则。 2、优秀的扩展性。 3、灵活性。

缺点： 1、具体元素对访问者公布细节，违反了迪米特原则。 2、具体元素变更比较困难。 3、违反了依赖倒置原则，依赖了具体类，没有依赖抽象。

使用场景：
	1、对象结构中对象对应的类很少改变，但经常需要在此对象结构上定义新的操作。
	2、需要对一个对象结构中的对象进行很多不同的并且不相关的操作，而需要避免让这些操作"污染"这些对象的类，也不希望在增加新操作时修改这些类。

注意事项：访问者可以对功能进行统一，可以做报表、UI、拦截器与过滤器。
*/

/*第一步 创建一个元素接口*/
type ComputerPart interface {
	accept(computerPartVisitor ComputerPartVisitor)
}

/*第二步 创建元素 ComputerPart 实现类*/
type Keyboard struct {
}

func (k *Keyboard) accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.visit(k)
}

type Mouse struct {
}

func (m *Mouse) accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.visit(m)
}

type Monitor struct {
}

func (m *Monitor) accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.visit(m)
}

type Computer struct {
	parts []ComputerPart
}

func NewComputer() *Computer {
	part := []ComputerPart{&Keyboard{}, &Monitor{}, &Mouse{}}
	return &Computer{parts: part}
}

func (c *Computer) accept(computerPartVisitor ComputerPartVisitor) {
	for _, part := range c.parts {
		part.accept(computerPartVisitor)
	}
	computerPartVisitor.visit(c)
}

/*第三步 创建访问者接口*/
type ComputerPartVisitor interface {
	visit(ComputerPart)
}

/*第四步 创建 访问者接口实现类*/
type ComputerPartDisplayVisitor struct{}

func (c *ComputerPartDisplayVisitor) visit(part ComputerPart) {

	switch part.(type) {
	case *Keyboard:
		fmt.Println("Displaying Keyboard.")
	case *Monitor:
		fmt.Println("Displaying Monitor.")
	case *Mouse:
		fmt.Println("Displaying Mouse.")
	case *Computer:
		fmt.Println("Displaying Computer.")
	}

}
