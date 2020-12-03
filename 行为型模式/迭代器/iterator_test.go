package 迭代器

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	a := &NameRepository{
		name: []string{
			"Robert",
			"John",
			"Julie",
			"Lora",
		},
	}
	iterator := a.getIterator()
	for iterator.hasNext() {
		fmt.Println(iterator.next())
	}
}
