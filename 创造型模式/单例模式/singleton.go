package 单例模式

import (
	"sync"
)

/*
关键代码：构造函数是私有的。
*/
var once = sync.Once{}

type singleton struct{}

var instance *singleton

//获取单例模式对象
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
