package 命令模式

import "testing"

func TestCommand(t *testing.T) {
	mb := NewMotherBoard()
	rebootCommand := NewRebootCommand(mb)
	startCommand := NewStartCommand(mb)

	box1 := NewBox(rebootCommand, startCommand)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(startCommand, rebootCommand)
	box2.PressButton1()
	box2.PressButton2()
}
