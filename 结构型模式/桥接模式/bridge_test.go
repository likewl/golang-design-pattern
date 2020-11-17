package 桥接模式

import "testing"

func TestNewBridge(t *testing.T) {
	newBridge := NewBridge(ViaSMS())
	newBridge.send("have lunch", "like")
	newBridge = NewBridge(ViaEmail())
	newBridge.send("have breakfast", "xiaoyao")
}
