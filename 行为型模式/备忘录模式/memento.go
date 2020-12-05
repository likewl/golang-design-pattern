package 备忘录模式

import "fmt"

/*
备忘录模式（Memento Pattern）保存一个对象的某个状态，以便在适当的时候恢复对象。

意图：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。

如何解决：通过一个备忘录类专门存储对象状态。

关键代码：客户不与备忘录类耦合，与备忘录管理类耦合。

优点：
	1、给用户提供了一种可以恢复状态的机制，可以使用户能够比较方便地回到某个历史的状态。
	2、实现了信息的封装，使得用户不需要关心状态的保存细节。

缺点：消耗资源。如果类的成员变量过多，势必会占用比较大的资源，而且每一次保存都会消耗一定的内存。

使用场景：
	1、需要保存/恢复数据的相关状态场景。
	2、提供一个可回滚的操作。

注意事项：
	1、为了符合迪米特原则，还要增加一个管理备忘录的类。
	2、为了节约内存，可使用原型模式+备忘录模式。
*/

/*第一步 创建一个memento接口和类*/
//用作类型之间的转换
type memento interface{}

//用作储存原始数据
type gameMemento struct {
	hp, mp int
}

/*第二步 创建originator类*/
type Game struct {
	hp, mp int
}

//输出状态
func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}

//用于储存原始数据，并返回储存的原始数据类
func (g *Game) Save() memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

//操作
func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

//加载储存的原始数据
func (g *Game) Load(gm memento) {
	m := gm.(*gameMemento)
	g.mp = m.mp
	g.hp = m.hp
}
