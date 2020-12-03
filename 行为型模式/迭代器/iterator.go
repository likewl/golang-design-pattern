package 迭代器
/*
这种模式用于顺序访问集合对象的元素，不需要知道集合对象的底层表示。

如何解决：把在元素之间游走的责任交给迭代器，而不是聚合对象。

关键代码：定义接口：hasNext, next。

优点：
	1、它支持以不同的方式遍历一个聚合对象。
	2、迭代器简化了聚合类。
	3、在同一个聚合上可以有多个遍历。
	4、在迭代器模式中，增加新的聚合类和迭代器类都很方便，无须修改原有代码。

缺点：由于迭代器模式将存储数据和遍历数据的职责分离，增加新的聚合类需要对应增加新的迭代器类，类的个数成对增加，这在一定程度上增加了系统的复杂性。

使用场景：
	1、访问一个聚合对象的内容而无须暴露它的内部表示。
	2、需要为聚合对象提供多种遍历方式。
	3、为遍历不同的聚合结构提供一个统一的接口。
*/

/*第一步 创建迭代器接口*/
type Iterator interface {
	hasNext() bool
	next() interface{}
}
/*第二步 创建一个容器接口*/
type Container interface {
	getIterator() Iterator
}
/*第三步 创建容器和迭代器接口实现类*/
type NameRepository struct {
	name []string
}

func (n *NameRepository) getIterator() Iterator {
	return &NameIterator{
		index: 0,
		container: n,
	}
}
//Iterator 实现类内嵌容器接口实现类
type NameIterator struct {
	index int
	container *NameRepository
}

func (n *NameIterator) hasNext() bool {
	if len(n.container.name)<=n.index{
		return false
	}
	return true
}

func (n *NameIterator) next() interface{} {
	if n.hasNext(){
		s := n.container.name[n.index]
		n.index++
		return s
	}
	return nil
}



