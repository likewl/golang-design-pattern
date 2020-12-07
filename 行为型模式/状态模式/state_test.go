package 状态模式

import (
	"fmt"
	"testing"
)

/*
@Author LiKe
@Description //TODO
@Time 2020/12/7 15:50
*/

func TestState(t *testing.T) {
	c := &class{}
	var start StartState
	start.DoAction(c)
	fmt.Println(c.state)
	var stop StopState
	stop.DoAction(c)
	fmt.Println(c.state)
}
