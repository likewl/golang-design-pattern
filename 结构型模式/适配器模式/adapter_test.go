package 适配器模式

import (
	"fmt"
	"testing"
)

func TestNewAdapter(t *testing.T) {
	mp3player := NewMp3player()
	mp4player := NewMp4player()
	mp3player.playMp3()
	mp4player.playMp4()
	fmt.Println("-------adapter---------")
	adapter := NewAdapter(mp3player, mp4player)
	adapter.playMp3()
	adapter.playMp4()
}
