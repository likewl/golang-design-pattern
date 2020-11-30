package 责任链模式

import "fmt"

/*
责任链模式（Chain of Responsibility Pattern）为请求创建了一个接收者对象的链。

在这种模式中，通常每个接收者都包含对另一个接收者的引用。如果一个对象不能处理该请求，那么它会把相同的请求传给下一个接收者，依此类推。

对请求的发送者和接收者进行解耦。

避免请求发送者与接收者耦合在一起，让多个对象都有可能接收请求，将这些对象连接成一条链，并且沿着这条链传递请求，直到有对象处理它为止。

职责链上的处理者负责处理请求，客户只需要将请求发送到职责链上即可，无须关心请求的处理细节和请求的传递，所以职责链将请求的发送者和请求的处理者解耦了。

关键代码：RequestChain 里面聚合它自己，在 HaveRight 里判断是否合适，如果没达到条件则向下传递，向谁传递之前 set 进去。

如何解决：拦截的类都实现统一接口。

应用实例：
	1、红楼梦中的"击鼓传花"。
	2、JS 中的事件冒泡。
	3、JAVA WEB 中 Apache Tomcat 对 Encoding 的处理，Struts2 的拦截器，jsp servlet 的 Filter。

优点：
	1、降低耦合度。它将请求的发送者和接收者解耦。
	2、简化了对象。使得对象不需要知道链的结构。
	3、增强给对象指派职责的灵活性。通过改变链内的成员或者调动它们的次序，允许动态地新增或者删除责任。
	4、增加新的请求处理类很方便。

缺点：
	1、不能保证请求一定被接收。
	2、系统性能将受到一定影响，而且在进行代码调试时不太方便，可能会造成循环调用。
	3、可能不容易观察运行时的特征，有碍于除错。

使用场景：
	1、有多个对象可以处理同一个请求，具体哪个对象处理该请求由运行时刻自动确定。
	2、在不明确指定接收者的情况下，向多个对象中的一个提交一个请求。
	3、可动态指定一组对象处理请求。
*/

/*第一步 创建用于chain的接口*/
type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

/*第二步 创建用于责任链结构体*/
//需要内嵌 Manager 接口
//同时需要实现 Manager 接口
type RequestChain struct {
	Manager
	successor *RequestChain
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (r *RequestChain) setSuccess(manager *RequestChain) {
	r.successor = manager
}

/*第三步 创建链的接收者*/
type ProjectManager struct{}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("Project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request\n", name, money)
	return false
}

type DepManager struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{},
	}
}

func (*DepManager) HaveRight(money int) bool {
	return money < 5000
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request\n", name, money)
	return false
}

type GeneralManager struct{}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{},
	}
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}
