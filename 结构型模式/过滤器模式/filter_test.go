package 过滤器模式

import (
	"fmt"
	"testing"
)

func TestCriteria(t *testing.T) {
	var p []*Person
	p = append(p, NewPerson("Robert", "Male", "Single"))
	p = append(p, NewPerson("John", "Male", "Married"))
	p = append(p, NewPerson("Laura", "Female", "Married"))
	p = append(p, NewPerson("Diana", "Female", "Single"))
	p = append(p, NewPerson("Mike", "Male", "Single"))
	p = append(p, NewPerson("Bobby", "Male", "Single"))
	criteriaMale := NewCriteriaMale()
	male := criteriaMale.meetCriteria(p)
	criteriaFemale := NewCriteriaFemale()
	female := criteriaFemale.meetCriteria(p)
	criteriaSingle := NewCriteriaSingle()
	single := criteriaSingle.meetCriteria(p)
	andCriteria := NewAndCriteria(criteriaSingle, criteriaFemale)
	and := andCriteria.meetCriteria(p)
	fmt.Printf("Male:\n")
	printPerson(male)
	fmt.Printf("Female:\n")
	printPerson(female)
	fmt.Printf("Single:\n")
	printPerson(single)
	fmt.Printf("Sing And Female:\n")
	printPerson(and)
}

func printPerson(p []*Person) {
	for _, person := range p {
		fmt.Println("Person : [ Name : " + person.Name +
			", Gender : " + person.Gender +
			", Marital Status : " + person.MaritalStatus +
			" ]")
	}
}
