package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	const BUTTON_A = 2
	const BUTTON_B = 3

	const NUM_LENGTH = 4
	const HIGH = 1
	const LOW = 0

	A_PINS := [NUM_LENGTH]int{10, 11, 12, 13}
	B_PINS := [NUM_LENGTH]int{6, 7, 8, 9}
	SUM_PINS := [NUM_LENGTH]int{0, 1, 4, 5}

	// a, b := 1, 1
	// sum := a + b

	// buttonStateA, buttonStateB := HIGH, HIGH

	var digitsA [NUM_LENGTH]*gpio.LedDriver
	var digitsB [NUM_LENGTH]*gpio.LedDriver
	var digitsSum [NUM_LENGTH]*gpio.LedDriver

	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")

	for i, _ := range digitsA {
		digitsA[i] = gpio.NewLedDriver(firmataAdaptor, fmt.Sprintf("led%v", i), fmt.Sprintf("%v", A_PINS[i]))
	}
	for i, _ := range digitsB {
		digitsB[i] = gpio.NewLedDriver(firmataAdaptor, fmt.Sprintf("led%v", i), fmt.Sprintf("%v", B_PINS[i]))
	}
	for i, _ := range digitsSum {
		digitsSum[i] = gpio.NewLedDriver(firmataAdaptor, fmt.Sprintf("led%v", i), fmt.Sprintf("%v", SUM_PINS[i]))
	}

	work := func() {
		gobot.Every(750*time.Millisecond, func() {
			for _, led := range digitsA {
				led.Toggle()
			}
			for _, led := range digitsB {
				led.Toggle()
			}
			for _, led := range digitsSum {
				led.Toggle()
			}
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{
			digitsA[0],
			digitsA[1],
			digitsA[2],
			digitsA[3],
			digitsB[0],
			digitsB[1],
			digitsB[2],
			digitsB[3],
			digitsSum[0],
			digitsSum[1],
			digitsSum[2],
			digitsSum[3],
		},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
