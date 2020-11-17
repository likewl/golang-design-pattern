package 适配器模式

import "fmt"

/*
适配器模式用于转换一种接口适配另一种接口。

这种模式涉及到一个单一的类(adapter结构体)，该类负责加入独立的或不兼容的接口功能。

音频播放器设备只能播放 mp3 文件，通过使用一个更高级的音频播放器来播放 mp3 和 mp4 文件。

意图：将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

主要解决：主要解决在软件系统中，常常要将一些"现存的对象"放到新的环境中，而新环境要求的接口是现对象不能满足的。

何时使用：
	1、系统需要使用现有的类，而此类的接口不符合系统的需要。
	2、想要建立一个可以重复使用的类，用于与一些彼此之间没有太大关联的一些类，包括一些可能在将来引进的类一起工作，这些源类不一定有一致的接口。
	3、通过接口转换，将一个类插入另一个类系中。（比如老虎和飞禽，现在多了一个飞虎，在不增加实体的需求下，增加一个适配器，在里面包容一个虎对象，实现飞的接口。）

如何解决：继承或依赖（推荐）。

关键代码：适配器继承或依赖已有的对象，实现想要的目标接口。

使用场景：有动机地修改一个正常运行的系统的接口，这时应该考虑使用适配器模式。

注意事项：适配器不是在详细设计时添加的，而是解决正在服役的项目的问题。
*/
/*第一步 创建MP3播放器和MP4播放器接口*/
type player interface {
	playMp3()
}
type advancePlayer interface {
	playMp4()
}

/*第二步 创建MP3和MP4播放器实现类*/
type mp3Player struct {
}

func (m mp3Player) playMp3() {
	fmt.Println("playing:mp3")
}

type mp4Player struct {
}

func (m mp4Player) playMp4() {
	fmt.Println("playing:mp4")
}

/*第三步 创建adapter结构体 */
//匿名组合 player advancePlayer 接口，又因为Go语言中非入侵式接口特征
//所以 adapter 类也拥有 player advancePlayer 实例方法
type adapter struct {
	player
	advancePlayer
}

func NewMp3player() player {
	return &mp3Player{}
}
func NewMp4player() advancePlayer {
	return &mp4Player{}
}
func NewAdapter(mp3 player, mp4 advancePlayer) *adapter {
	return &adapter{
		mp3,
		mp4,
	}
}
