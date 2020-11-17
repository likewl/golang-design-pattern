package 原型模式

import (
	"testing"
)

func TestNewPrototypeManager(t *testing.T) {
	a:=type1{
		name: "haha",
	}
	prototypeManager:= NewPrototypeManager()
	prototype := prototypeManager.get("1")
	//fmt.Println(prototype)
	prototypeManager.set("1",&a)
	prototype = prototypeManager.get("1")
	//fmt.Println(prototype)
	clone := prototype.clone()
	if clone==prototype{
		t.Errorf("error! get clone not working")
	}
}
