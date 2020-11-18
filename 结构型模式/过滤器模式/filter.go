package 过滤器模式

import "strings"

/*
过滤器模式（Filter Pattern）或标准模式（Criteria Pattern）是一种设计模式，
这种模式允许开发人员使用不同的标准来过滤一组对象，通过逻辑运算以解耦的方式把它们连接起来。
它结合多个标准来获得单一标准。

实现：
	我们将创建一个 Person 对象、Criteria 接口和实现了该接口的实体类，来过滤 Person 对象的列表。
	CriteriaPatternDemo 类使用 Criteria 对象，基于各种标准和它们的结合来过滤 Person 对象的列表。
*/

/*第一步 创建一个类，在此类上创建一个标准*/
type Person struct {
	Name          string
	Gender        string
	MaritalStatus string
}

func NewPerson(name string, gender string, maritalStatus string) *Person {
	return &Person{Name: name, Gender: gender, MaritalStatus: maritalStatus}
}

/*第二步 在 Person 类上创建一个标准*/
type criteria interface {
	meetCriteria(person []*Person) []*Person
}

/*第三步 创建标准的实现类，用于筛选*/
type criteriaMale struct{}

func NewCriteriaMale() *criteriaMale {
	return &criteriaMale{}
}

func (c criteriaMale) meetCriteria(person []*Person) []*Person {
	var p []*Person
	for _, p1 := range person {
		if strings.EqualFold(p1.Gender, "male") {
			p = append(p, p1)
		}
	}
	return p
}

type criteriaFemale struct{}

func NewCriteriaFemale() *criteriaFemale {
	return &criteriaFemale{}
}

func (c criteriaFemale) meetCriteria(person []*Person) []*Person {
	var p []*Person
	for _, p1 := range person {
		if strings.EqualFold(p1.Gender, "female") {
			p = append(p, p1)
		}
	}
	return p
}

type criteriaSingle struct{}

func NewCriteriaSingle() *criteriaSingle {
	return &criteriaSingle{}
}

func (c criteriaSingle) meetCriteria(person []*Person) []*Person {
	var p []*Person
	for _, p1 := range person {
		if strings.EqualFold(p1.MaritalStatus, "Single") {
			p = append(p, p1)
		}
	}
	return p
}

type andCriteria struct {
	firstCriteria criteria
	otherCriteria criteria
}

func NewAndCriteria(firstCriteria criteria, otherCriteria criteria) *andCriteria {
	return &andCriteria{
		firstCriteria: firstCriteria,
		otherCriteria: otherCriteria,
	}
}

func (a *andCriteria) meetCriteria(person []*Person) []*Person {
	f := a.firstCriteria.meetCriteria(person)
	return a.otherCriteria.meetCriteria(f)
}
