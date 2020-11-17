package 建造者模式

import (
	"fmt"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	b := NewBuilder()
	vegMeal := b.prepareVegMeal()
	fmt.Printf("------------vegMeal------------\n")
	vegMeal.showItems()
	nonVegMeal := b.prepareNonVegMeal()
	fmt.Printf("------------nonVegMeal------------\n")
	nonVegMeal.showItems()
}
