package GoDS18B20

import (
  "os/exec"
  )

const (
  BASE_DIR = "/sys/bus/w1/devices/"
  PROBE_ID = "28"
  SLAVE_ID = "w1_slave"

  STATUS_REG = "crc=.* (.*)"
  TEMP_REG = "t=(\\d+)"
  )

func Initialize() {
  _, err := exec.Command("modprobe", "w1-gpio").Output()
  if err != nil {
    panic(err)
  }

  _, err = exec.Command("modprobe", "w1-therm").Output()
  if err != nil {
    panic(err)
  }
}
