package main

import (
  "gobot.io/x/gobot"
  "gobot.io/x/gobot/drivers/gpio"
  "gobot.io/x/gobot/platforms/firmata"
  "github.com/TheEricle/my_gobot/work"
)

func main() {
  firmataAdaptor := firmata.NewAdaptor("/dev/cu.usbmodem14101")
  led := gpio.NewLedDriver(firmataAdaptor, "13")

  worker := func(){ work.GetToWork(led) }

  robot := gobot.NewRobot("bot",
    []gobot.Connection{firmataAdaptor},
    []gobot.Device{led},
    worker,
  )

  robot.Start()
}