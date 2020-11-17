package 抽象工厂模式

import "testing"

func TestNewFactory(t *testing.T) {
	factory := NewFactory("haier")
	airConditioner := factory.createAirConditioner()
	airConditioner.airConditionerInfo()
	fan := factory.creatFan()
	fan.FanInfo()
	factory = NewFactory("media")
	airConditioner = factory.createAirConditioner()
	airConditioner.airConditionerInfo()
	fan = factory.creatFan()
	fan.FanInfo()
}


