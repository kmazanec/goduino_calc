package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")

	led1 := gpio.NewLedDriver(firmataAdaptor, "led1", "13")
	led2 := gpio.NewLedDriver(firmataAdaptor, "led2", "12")
	led3 := gpio.NewLedDriver(firmataAdaptor, "led3", "11")
	led4 := gpio.NewLedDriver(firmataAdaptor, "led4", "10")

	work := func() {
		led1.Toggle()

		c := 1
		gobot.Every(500*time.Millisecond, func() {
			if c == 1 {
				led1.Toggle()
				led2.Toggle()
				c = 2
			} else if c == 2 {
				led2.Toggle()
				led3.Toggle()
				c = 3
			} else if c == 3 {
				led3.Toggle()
				led4.Toggle()
				c = 4
			} else {
				led4.Toggle()
				led1.Toggle()
				c = 1
			}
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led1, led2, led3, led4},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
