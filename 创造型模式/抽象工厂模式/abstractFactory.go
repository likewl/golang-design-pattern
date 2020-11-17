package 抽象工厂模式

import "fmt"

/*
抽象工厂模式（Abstract Factory Pattern）是围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂。
在抽象工厂模式中，接口是负责创建一个相关对象的工厂，不需要显式指定它们的类。每个生成的工厂都能按照工厂模式提供对象。

抽象工厂是创建型设计模式，它强调了一系列相关产品对象（属于同一个产品族）的创建过程。
它和工厂方法模式的侧重点不同:
	工厂模式更加侧重于同一产品等级;
	抽象工厂模式侧重的是同一产品族。

还是用冰箱来举例，美的生产的产品有美的冰箱、美的空调、美的电风扇。

在这个工厂里面生产的所有产品对象都不是同一个产品等级的，它们属于同一个产品族，该工厂生产的一系列对象都是相关的。
就像定义说的：
	抽象工厂模式提供一个创建一系列相关或相互依赖对象的接口。

如果现在我们希望生产海尔的一系列产品就可以扩展抽象产品接口创建一系列海尔产品类，创建海尔的工厂负责生产海尔的产品。

缺点：
	如果美的现在又搞洗衣机了，就需要修改接口和所有的实现类，因此对于扩展一个新的产品等级来说，不遵循开闭原则。
*/

/*第一步*/
//创建一个空调接口。
type airConditioner interface {
	airConditionerInfo()
}

//创建一个风扇接口。
type Fan interface {
	FanInfo()
}

/*第二步 创建接口产品实现类*/
//创建美的产品接口
type MediaAirConditioner struct{}
type MediaFan struct{}

func (m MediaAirConditioner) airConditionerInfo() {
	fmt.Println("美的空调")
}
func (m MediaFan) FanInfo() {
	fmt.Println("美的风扇")
}

//创建海尔产品接口
type HaierAirConditioner struct{}
type HaierFan struct{}

func (h HaierAirConditioner) airConditionerInfo() {
	fmt.Println("海尔空调")
}
func (h HaierFan) FanInfo() {
	fmt.Println("海尔风扇")
}

/*第三步 创建抽象工厂接口*/
type abstractFactory interface {
	createAirConditioner() airConditioner
	creatFan() Fan
}

/*第四步 创建抽象工厂实现类*/
//创建美的工厂
type mediaFactory struct{}

func (m mediaFactory) createAirConditioner() airConditioner {
	return &MediaAirConditioner{}
}
func (m mediaFactory) creatFan() Fan {
	return &MediaFan{}
}

//创建海尔工厂
type HaierFactory struct{}

func (h HaierFactory) createAirConditioner() airConditioner {
	return &HaierAirConditioner{}
}

func (h HaierFactory) creatFan() Fan {
	return &HaierFan{}
}

func NewFactory(s string) abstractFactory {
	if s == "haier" {
		return &HaierFactory{}
	}
	if s == "media" {
		return &mediaFactory{}
	}
	return nil
}
