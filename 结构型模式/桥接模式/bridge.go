package 桥接模式

import "fmt"

/*
桥接模式分离抽象部分和实现部分。使得两部分独立扩展。

桥接模式和适配器模式的区别
	共同点:
		桥接和适配器都是让两个东西配合工作
	不同点:
		出发点不同。
         	1）适配器：改变已有的两个接口，让他们相容。
         	2）桥接模式：分离抽象化和实现，使两者的接口可以不同，目的是分离。

	所以说，如果你拿到两个已有模块，想让他们同时工作，那么你使用的适配器。
	如果你还什么都没有，但是想分开实现，那么桥接是一个选择。

	桥接是先有桥，才有两端的东西
	适配是先有两边的东西，才有适配器

	桥接是在桥好了之后，两边的东西还可以变化。
*/

/*第一步 创建桥接实现接口*/
type abstractMessage interface {
	send(text, to string)
}

/*第二步 创建 abstractMessage 接口的实体桥接实现类*/
type messageImp interface {
	sendMsg(text, to string)
}

/*第三步 创建 messageImp 接口实现类*/
type messageSMS struct{}
type messageEmail struct{}

func (v messageSMS) sendMsg(text, to string) {
	fmt.Printf("send msg:%s to %s,by SMS\n", text, to)
}

func (v messageEmail) sendMsg(text, to string) {
	fmt.Printf("send msg:%s to %s,by Email\n", text, to)
}
func ViaSMS() *messageSMS {
	return &messageSMS{}
}
func ViaEmail() *messageEmail {
	return &messageEmail{}
}

/*第四步 创建 abstractMessage 接口实现类*/
type bridge struct {
	method messageImp
}

func (b bridge) send(text, to string) {
	b.method.sendMsg(text, to)
}

func NewBridge(msg messageImp) *bridge {
	return &bridge{
		method: msg,
	}
}
