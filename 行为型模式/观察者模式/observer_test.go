package 观察者模式

import "testing"

func TestObserver(t *testing.T) {
	subject := NewSubject()
	NewObserverOne("reader1", subject)
	NewObserverTwo("reader2", subject)
	NewObserverThree("reader3", subject)

	subject.setState("observer mode1")
	subject.setState("observer mode2")
}
