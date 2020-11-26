package 代理模式

import "fmt"

/*
在代理模式中，我们创建具有现有对象的对象，以便向外界提供功能接口。

何时使用：想在访问一个类时做一些控制。

如何解决：增加中间层。

关键代码：实现与被代理类组合。

优点： 1、职责清晰。 2、高扩展性。 3、智能化。

缺点： 1、由于在客户端和真实主题之间增加了代理对象，因此有些类型的代理模式可能会造成请求的处理速度变慢。 2、实现代理模式需要额外的工作，有些代理模式的实现非常复杂。

使用场景：按职责来划分，通常有以下使用场景：
	1、远程代理。
	2、虚拟代理。
	3、Copy-on-Write 代理。
	4、保护（Protect or Access）代理。
	5、Cache代理。
	6、防火墙（Firewall）代理。
	7、同步化（Synchronization）代理。
	8、智能引用（Smart Reference）代理。
注意事项：
	1、和适配器模式的区别：适配器模式主要改变所考虑对象的接口，而代理模式不能改变所代理类的接口。
	2、和装饰器模式的区别：装饰器模式为了增强功能，而代理模式是为了加以控制。
*/

/*第一步 创建 image 接口*/
type image interface {
	display()
}

/*第二步 创建 image 接口实现类*/
type RealImage struct {
	fileName string
}

func (r *RealImage) display() {
	fmt.Println("Displaying " + r.fileName)
}

func (r *RealImage) loadFromDisk(fileName string) {
	fmt.Println("Loading " + fileName)
}
func NewRealImage(fileName string) *RealImage {
	a := &RealImage{
		fileName: fileName,
	}
	a.loadFromDisk(fileName)
	return a
}

/*第三步 创建代理类，嵌套 RealImage 类*/
type Proxy struct {
	real     *RealImage
	fileName string
}

func (p *Proxy) display() {
	if p.real == nil {
		p.real = NewRealImage(p.fileName)
	}
	p.real.display()
}
