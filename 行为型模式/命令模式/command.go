package 命令模式

import "fmt"

/*
请求以命令的形式包裹在对象中，并传给调用对象。调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象，该对象执行命令。

如何解决：通过调用者调用接受者执行命令，顺序：调用者→命令→接受者。

关键代码：
	定义三个角色：
		1、received 真正的命令执行对象
		2、Command
		3、invoker 使用命令对象的入口

优点：
	1、降低了系统耦合度。
	2、新的命令可以很容易添加到系统中去。

缺点：使用命令模式可能会导致某些系统有过多的具体命令类。

使用场景：认为是命令的地方都可以使用命令模式，比如： 1、GUI 中每一个按钮都是一条命令。 2、模拟 CMD。
*/

/*第一步 创建一个命令接口*/
type command interface {
	Execute()
}

/*第二步 创建请求类*/
//并给请求体实现相应的方法
type MotherBoard struct{}

func NewMotherBoard() *MotherBoard {
	return &MotherBoard{}
}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

/*第三步 创建 command 接口实现类*/
//并为实现类内嵌请求体
type StartCommand struct {
	m *MotherBoard
}

func NewStartCommand(m *MotherBoard) *StartCommand {
	return &StartCommand{m: m}
}

func (s StartCommand) Execute() {
	s.m.Start()
}

type RebootCommand struct {
	m *MotherBoard
}

func NewRebootCommand(m *MotherBoard) *RebootCommand {
	return &RebootCommand{m: m}
}

func (s RebootCommand) Execute() {
	s.m.Reboot()
}

/*第四步 创建命令调用类*/
type Box struct {
	button1 command
	button2 command
}

func NewBox(button1 command, button2 command) *Box {
	return &Box{button1: button1, button2: button2}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
