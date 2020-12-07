package 策略模式

import (
	"fmt"
	"strconv"
	"testing"
)

/*
@Author LiKe
@Description //TODO
@Time 2020/12/7 17:11
*/

func TestStrategy(t *testing.T) {
	exe := NewClass(&Add{})
	fmt.Println("10 + 5 = " + strconv.Itoa(exe.executeStrategy(10, 5)))
	exe = NewClass(&Subtract{})
	fmt.Println("10 - 5 = " + strconv.Itoa(exe.executeStrategy(10, 5)))
	exe = NewClass(&Multiply{})
	fmt.Println("10 * 5 = " + strconv.Itoa(exe.executeStrategy(10, 5)))
}
