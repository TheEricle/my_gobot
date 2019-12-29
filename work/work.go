package work

import (
  "time"
  "gobot.io/x/gobot"
  "gobot.io/x/gobot/drivers/gpio"
  )

func GetToWork(led *gpio.LedDriver) {
    gobot.Every(1*time.Second, func() {
      led.Toggle()
    })
  }