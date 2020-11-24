package 组合模式

import (
	"fmt"
	"testing"
)

func TestNewEmployee(t *testing.T) {
	CEO := NewEmployee("John", "CEO", 30000)

	headSales := NewEmployee("Robert", "Head Sales", 20000)

	headMarketing := NewEmployee("Michel", "Head Marketing", 20000)

	clerk1 := NewEmployee("Laura", "Marketing", 10000)
	clerk2 := NewEmployee("Bob", "Marketing", 10000)

	salesExecutive1 := NewEmployee("Richard", "Sales", 10000)
	salesExecutive2 := NewEmployee("Rob", "Sales", 10000)

	CEO.add(headSales)
	CEO.add(headMarketing)

	headSales.add(salesExecutive1)
	headSales.add(salesExecutive2)

	headMarketing.add(clerk1)
	headMarketing.add(clerk2)

	//打印该组织的所有员工
	fmt.Println(CEO)
	for _, headEmployee := range CEO.getSubordinates() {
		fmt.Println(headEmployee)
		for _, em := range headEmployee.getSubordinates() {
			fmt.Println(em)
		}
	}
}
