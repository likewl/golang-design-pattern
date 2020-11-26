package 代理模式

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	image := &Proxy{
		fileName: "image.jpg",
	}
	// 图像将从磁盘加载
	image.display()
	fmt.Println("")
	// 图像不需要从磁盘加载
	image.display()
}
