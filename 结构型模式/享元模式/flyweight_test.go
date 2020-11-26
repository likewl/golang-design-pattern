package 享元模式

import "testing"

func TestNewFlyweightFactory(t *testing.T) {
	flyweight1 := NewFlyweightFactory("test")
	flyweight2 := NewFlyweightFactory("test")
	if flyweight1 != flyweight2 {
		t.Errorf("error")
	}
}
