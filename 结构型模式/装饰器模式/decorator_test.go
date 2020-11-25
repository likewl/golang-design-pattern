package 装饰器模式

import "testing"

func TestNewRedDecorator(t *testing.T) {

	rectangle := NewRedDecorator(NewRectangle())
	circle := NewRedDecorator(NewCircle())
	rectangle.draw()
	circle.draw()
}
