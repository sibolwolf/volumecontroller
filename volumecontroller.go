package main

// package name is smartconn.cc/sibolwolf/volumecontroller
import (
    "fmt"
    LED "smartconn.cc/sibolwolf/volumecontroller/led"
    VOL "smartconn.cc/sibolwolf/volumecontroller/volume"
    KEY "smartconn.cc/liugen/input"
)

func main() {
    fmt.Println("Hello, VolumeController")
    LED.GetLedStatus("/sys/class/leds/vol_led0/brightness")
    VOL.GetVolumeStatus()
    VOL.SetVolume("50")
    VOL.GetVolumeStatus()
    KEY.Connect("readingangel")
    KEY.GetButton("volume").OnPress(
        func() {
            fmt.Println("RA got a short key press event")
        }
    )
}
