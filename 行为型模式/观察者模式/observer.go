package 观察者模式

import "fmt"

/*
当对象间存在一对多关系时，则使用观察者模式（Observer Pattern）。比如，当一个对象被修改时，则会自动通知依赖它的对象。

意图：定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。
主要解决：一个对象状态改变给其他对象通知的问题，而且要考虑到易用和低耦合，保证高度的协作。

何时使用：一个对象（目标对象）的状态发生改变，所有的依赖对象（观察者对象）都将得到通知，进行广播通知。

关键代码：在相关类里有一个 ArrayList 存放观察者们。

应用实例： 1、拍卖的时候，拍卖师观察最高标价，然后通知给其他竞价者竞价。 2、西游记里面悟空请求菩萨降服红孩儿，菩萨洒了一地水招来一个老乌龟，这个乌龟就是观察者，他观察菩萨洒水这个动作。

优点：
	1、观察者和被观察者是抽象耦合的。
	2、建立一套触发机制。

缺点：
	1、如果一个被观察者对象有很多的直接和间接的观察者的话，将所有的观察者都通知到会花费很多时间。
	2、如果在观察者和观察目标之间有循环依赖的话，观察目标会触发它们之间进行循环调用，可能导致系统崩溃。
	3、观察者模式没有相应的机制让观察者知道所观察的目标对象是怎么发生变化的，而仅仅只是知道观察目标发生了变化。

使用场景：
	一个对象的改变将导致其他一个或多个对象也发生改变，而不知道具体有多少对象将发生改变，可以降低对象之间的耦合度。
	一个对象必须通知其他对象，而并不知道这些对象是谁。
	需要在系统中创建一个触发链，A对象的行为将影响B对象，B对象的行为将影响C对象……，可以使用观察者模式创建一种链式触发机制。
*/

/*第一步 Subject 类*/
//Subject 对象带有绑定观察者到 Client 对象和从 Client 对象解绑观察者的方法
type Subject struct {
	state    string
	observer []observer
}

func NewSubject() *Subject {
	return &Subject{
		observer: make([]observer, 0),
	}
}

//将所有依赖它的对象关联起来
func (s *Subject) Attach(o observer) {
	s.observer = append(s.observer, o)
}

//更新所有依赖它的对象的相关信息
func (s *Subject) notifyAllClass() {
	for _, o := range s.observer {
		o.Update(s)
	}
}

//用于设置相关修改参数并更新所有依赖它的对象的相关信息
func (s *Subject) setState(state string) {
	s.state = state
	s.notifyAllClass()
}

/*第二步 创建 observer 类*/
type observer interface {
	Update(*Subject)
}

/*第三步 创建 observer 实现类*/
type observerOne struct {
	name string
	s    *Subject
}

func NewObserverOne(name string, s *Subject) *observerOne {
	reader := &observerOne{name: name, s: s}
	s.observer = append(s.observer, reader)
	return reader
}

//用于更新相关信息
func (r *observerOne) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.state)
}

type observerTwo struct {
	name string
	s    *Subject
}

func NewObserverTwo(name string, s *Subject) *observerTwo {
	reader := &observerTwo{name: name, s: s}
	s.observer = append(s.observer, reader)
	return reader
}

//用于更新相关信息
func (r *observerTwo) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.state)
}

type observerThree struct {
	name string
	s    *Subject
}

func NewObserverThree(name string, s *Subject) *observerThree {
	reader := &observerThree{name: name, s: s}
	s.observer = append(s.observer, reader)
	return reader
}

//用于更新相关信息
func (r *observerThree) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.state)
}
