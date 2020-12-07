package 模板模式

import "testing"

/*
@Author LiKe
@Description //TODO
@Time 2020/12/7 20:13
*/

func TestTemplate(t *testing.T) {
	cricket := NewCricket()
	football := NewFootball()
	football.play()
	cricket.play()
}
