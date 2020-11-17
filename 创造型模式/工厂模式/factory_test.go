package 工厂模式

import "testing"

func TestShapeFactory(t *testing.T) {
	shape := ShapeFactory("square")
	shape.draw()
	shape = ShapeFactory("circle")
	shape.draw()
}