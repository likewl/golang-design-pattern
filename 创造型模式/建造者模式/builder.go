package 建造者模式

import "fmt"

/*
建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。

一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。

意图：将一个复杂的构建与其表示相分离，使得同样的构建过程可以创建不同的表示。

何时使用：一些基本部件不会变，而其组合经常变化的时候。

应用实例： 1、去肯德基，汉堡、可乐、薯条、炸鸡翅等是不变的，而其组合是经常变化的，生成出所谓的"套餐"。 2、JAVA 中的 StringBuilder。

注意事项：与工厂模式的区别是：建造者模式更加关注与零件装配的顺序。

*/

/*第一步 */
//创建食物条目接口
type item interface {
	name() string
	pack() packing
	price() float32
}

//创建食物包装接口
type packing interface {
	pack() string
}

/*第二步 创建食品包装实现类*/
type wrapper struct{}

func (w wrapper) pack() string {
	return "Wrapper"
}

type bottle struct{}

func (b bottle) pack() string {
	return "Bottle"
}

/*第三步 创建实现 item 接口的抽象类，该类提供了默认的功能*/
type burger struct{}

func (b *burger) pack() packing {
	return new(wrapper)
}

type coldDrink struct{}

func (c *coldDrink) pack() packing {
	return new(bottle)
}

/*第四步 创建扩展了 burger 和 coldDrink 的实体类*/
//创建 burger 的实体类
type vegBurger struct {
	burger
}

func (v vegBurger) name() string {
	return "VegBurger"
}

func (v vegBurger) price() float32 {
	return 25.0
}

type chickenBurger struct {
	burger
}

func (c chickenBurger) name() string {
	return "ChickenBurger"
}

func (c chickenBurger) price() float32 {
	return 50.5
}

//创建 coldDrink 的实体类
type coke struct {
	coldDrink
}

func (c coke) name() string {
	return "Coke"
}

func (c coke) price() float32 {
	return 10.5
}

type pepsi struct {
	coldDrink
}

func (p pepsi) name() string {
	return "pepsi"
}

func (p pepsi) price() float32 {
	return 20.0
}

/*第五步 创建一个 meal 类，带有上面定义的 item 对象*/
type meal struct {
	items []item
}

func (i *meal) add(Item item) {
	i.items = append(i.items, Item)
}
func (i *meal) showItems() {
	for _, v := range i.items {
		fmt.Printf("Item :%s\n", v.name())
		fmt.Println("Packing :" + v.pack().pack())
		fmt.Printf("Price :%f\n", v.price())
	}
}
func (i *meal) getCost() float32 {
	var cost float32
	for _, v := range i.items {
		cost += v.price()
	}
	return cost
}

/*第六步 创建建造者*/
type Builder struct{}

func (b Builder) prepareVegMeal() *meal {
	a := &meal{}
	a.add(&vegBurger{})
	a.add(&coke{})
	return a
}
func (b Builder) prepareNonVegMeal() *meal {
	a := &meal{}
	a.add(&chickenBurger{})
	a.add(&pepsi{})
	return a
}
func NewBuilder() *Builder {
	return new(Builder)
}
