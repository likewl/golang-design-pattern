package 访问者模式

import "testing"

/*
@Author LiKe
@Description //TODO
@Time 2020/12/7 19:42
*/

func TestVisitor(t *testing.T) {
	computer := NewComputer()
	computer.accept(&ComputerPartDisplayVisitor{})
}
