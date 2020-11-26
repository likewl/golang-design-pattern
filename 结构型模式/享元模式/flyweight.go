package 享元模式

/*
享元模式（Flyweight Pattern）主要用于减少创建对象的数量，以减少内存占用和提高性能。

享元模式尝试重用现有的同类对象，如果未找到匹配的对象，则创建新对象。

主要解决：在有大量对象时，有可能会造成内存溢出，我们把其中共同的部分抽象出来，如果有相同的业务请求，直接返回在内存中已有的对象，避免重新创建。

关键代码：用 HashMap 存储这些对象。

注意事项：
	1、注意划分外部状态和内部状态，否则可能会引起线程安全问题。
	2、这些类必须有一个工厂对象加以控制。
*/

import "sync"

/*第一步 创建享元类*/
type flyweight struct {
	data string
}

/*第二步 创建相关工厂类*/
type flyweightFactory struct {
	maps map[string]*flyweight
}

func (f *flyweightFactory) get(s string) *flyweight {
	f1 := f.maps[s]
	if f1 == nil {
		f1 = &flyweight{
			data: "flyweight :" + s,
		}
		f.maps[s] = f1
	}
	return f1
}

/*第三步 实现享元工厂*/
var f *flyweightFactory
var once sync.Once

func NewFlyweightFactory(s string) *flyweight {
	once.Do(func() {
		f = &flyweightFactory{
			make(map[string]*flyweight),
		}
	})
	return f.get(s)
}
