package 解释器模式

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInterpreter(t *testing.T) {
	t1 := &terminalExpression{
		data: "Robert",
	}
	t2 := &terminalExpression{
		data: "John",
	}
	t3 := &terminalExpression{
		data: "Julie",
	}
	t4 := &terminalExpression{
		data: "Married",
	}

	//制定规则
	isMale := newOrExpression(t1, t2)
	isMarriedWoman  := newAndExpression(t3, t4)
	//解析
	fmt.Println("John is male? " + strconv.FormatBool(isMale.interpret("John")))
	fmt.Println("Julie is a married women? "+
		strconv.FormatBool(isMarriedWoman.interpret("Married Julie")))
}
