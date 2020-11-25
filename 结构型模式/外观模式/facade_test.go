package 外观模式

import "testing"

func TestNewFacade(t *testing.T) {
	facade := NewFacade()
	facade.DrawCircle()
	facade.DrawSquare()
}
