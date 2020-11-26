package 原型模式

/*关键代码： 用于隔离类对象的使用者和具体类型（易变类）之间的耦合关系*/

type cloneable interface {
	clone() cloneable
}

type type1 struct {
	name string
}

func (t *type1) clone() cloneable {
	a := *t
	return &a
}

type prototypeManager struct {
	m map[string]cloneable
}

func (p *prototypeManager) get(s string) cloneable {
	c, ok := p.m[s]
	if !ok {
		return nil
	}
	return c
}
func (p *prototypeManager) set(s string, c cloneable) {
	p.m[s] = c
}
func NewPrototypeManager() *prototypeManager {
	m := make(map[string]cloneable)
	return &prototypeManager{
		m: m,
	}
}
