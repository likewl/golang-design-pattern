package 责任链模式

import "testing"

func TestChain(t *testing.T) {
	//创建责任链接收者之间的关系
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()
	c1.setSuccess(c2)
	c2.setSuccess(c3)

	var c Manager = c1
	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
}
