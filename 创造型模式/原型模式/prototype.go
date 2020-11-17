package 原型模式

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
	c,ok := p.m[s]
	if !ok{
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