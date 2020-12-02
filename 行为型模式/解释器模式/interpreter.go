package 解释器模式

import "strings"
/*
提供了评估语言的语法或表达式的方式

这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。

如何解决：构建语法树，定义终结符与非终结符。

关键代码：构建环境类，包含解释器之外的一些全局信息，一般是 HashMap。

优点：
	1、可扩展性比较好，灵活。
	2、增加了新的解释表达式的方式。
	3、易于实现简单文法。

缺点：
	1、可利用场景比较少。
	2、对于复杂的文法比较难维护。
	3、解释器模式会引起类膨胀。
	4、解释器模式采用递归调用方法。

使用场景：
	1、可以将一个需要解释执行的语言中的句子表示为一个抽象语法树。
	2、一些重复出现的问题可以用一种简单的语言来进行表达。
	3、一个简单语法需要解释的场景。


*/

/*第一步 创建表达式接口*/
type expression interface {
	interpret(s string) bool
}
/*第二步 创建 expression 接口实现类*/
//① terminalExpression 作为上下文中主要解释器
type terminalExpression struct {
	data string
}
func (t *terminalExpression) interpret(s string) bool {
	return strings.Contains(s,t.data)
}
//② 创建组合式表达式
type andExpression struct {
	expr1 expression
	expr2 expression
}

func newAndExpression(expr1 expression, expr2 expression) *andExpression {
	return &andExpression{expr1: expr1, expr2: expr2}
}

func (o *andExpression) interpret(s string) bool {
	return o.expr1.interpret(s)&&o.expr2.interpret(s)
}

type orExpression struct {
	expr1 expression
	expr2 expression
}

func newOrExpression(expr1 expression, expr2 expression) *orExpression {
	return &orExpression{expr1: expr1, expr2: expr2}
}

func (o *orExpression) interpret(s string) bool {
	return o.expr1.interpret(s)&&o.expr2.interpret(s)
}