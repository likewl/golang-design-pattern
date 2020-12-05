package 中介模式

import (
	"testing"
)

func TestMeditor(t *testing.T) {
	robert:=&user{name: "Robert"}
	john :=&user{name: "John"}


	robert.sendMessage("Hi! John!")
	john.sendMessage("Hello! Robert!")
}