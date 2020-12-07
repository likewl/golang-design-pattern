package 模板模式

import "fmt"

/*
@Author LiKe
@Description //TODO
@Time 2020/12/7 20:02
*/

/*因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有子类的引用。
例如：
	func NewCricket() *Cricket {
		cricket := &Cricket{}
		gameTemplate := NewGameTemplate(cricket)
		cricket.GameTemplate = gameTemplate
		return cricket
	}
*/

/*
一个抽象类公开定义了执行它的方法的方式/模板。它的子类可以按需要重写方法实现，但调用将以抽象类中定义的方式进行。

关键代码：在抽象类实现，其他步骤在子类实现。

优点：
	1、封装不变部分，扩展可变部分。
	2、提取公共代码，便于维护。
	3、行为由父类控制，子类实现。

缺点：每一个不同的实现都需要一个子类来实现，导致类的个数增加，使得系统更加庞大。

使用场景：
	1、有多个子类共有的方法，且逻辑相同。
	2、重要的、复杂的方法，可以考虑作为模板方法。

*/

/*第一步 创建模板接口，用于匿名组合到其他结构体中（相当于继承）*/
type template interface {
	initialize()
	startPlay()
	endPlay()
}

/*第二步 创建接口实现类，用作模板*/
type GameTemplate struct {
	template
}

func NewGameTemplate(template template) *GameTemplate {
	return &GameTemplate{template: template}
}

func (g *GameTemplate) play() {
	//初始化游戏
	g.initialize()

	//开始游戏
	g.startPlay()

	//结束游戏
	g.endPlay()
}

/*第三步 创建具体模板扩展类（继承）*/
type Cricket struct {
	*GameTemplate
}

/*因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有子类的引用。

例如：
	func NewCricket() *Cricket {
		cricket := &Cricket{}
		gameTemplate := NewGameTemplate(cricket)
		cricket.GameTemplate = gameTemplate
		return cricket
	}
*/
func NewCricket() *Cricket {
	cricket := &Cricket{}
	gameTemplate := NewGameTemplate(cricket)
	cricket.GameTemplate = gameTemplate
	return cricket
}

func (*Cricket) initialize() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}
func (*Cricket) startPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}
func (*Cricket) endPlay() {
	fmt.Println("Cricket Game Finished!")
}

type Football struct {
	*GameTemplate
}

func NewFootball() *Football {
	football := &Football{}
	gameTemplate := NewGameTemplate(football)
	football.GameTemplate = gameTemplate
	return football
}

func (*Football) initialize() {
	fmt.Println("Football Game Initialized! Start playing.")
}
func (*Football) startPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}
func (*Football) endPlay() {
	fmt.Println("Football Game Finished!")
}
