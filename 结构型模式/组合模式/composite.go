package 组合模式

import (
	"fmt"
	"strconv"
)

/*
组合模式（Composite Pattern），又叫部分整体模式，是用于把一组相似的对象当作一个单一的对象。组合模式依据树形结构来组合对象，用来表示部分以及整体层次。这种类型的设计模式属于结构型模式，它创建了对象组的树形结构。

这种模式创建了一个包含自己对象组的类。该类提供了修改相同对象组的方式。

我们通过下面的实例来演示组合模式的用法。实例演示了一个组织中员工的层次结构。
*/
type employee struct {
	name        string
	dept        string
	salary      int
	subordinate []*employee
}

func (e *employee) add(em *employee) {
	e.subordinate = append(e.subordinate, em)
}
func (e *employee) getSubordinates() []*employee {
	return e.subordinate
}

func (e *employee) String() string {
	return fmt.Sprintln("Employee :[ Name : " + e.name +
		", dept : " + e.dept +
		", salary :" + strconv.Itoa(e.salary) + " ]")
}
func NewEmployee(name string, dept string, salary int) *employee {
	return &employee{
		name:   name,
		dept:   dept,
		salary: salary,
	}
}
