package main

// package name is smartconn.cc/sibolwolf/volumecontroller
import (
    "fmt"
    "smartconn.cc/sibolwolf/volumecontroller/led"
    "smartconn.cc/sibolwolf/volumecontroller/volume"
)

func main() {
    fmt.Println("Hello, VolumeController")
    led.GetLedStatus("/sys/class/leds/vol_led0/brightness")
    volume.GetVolumeStatus()
}
