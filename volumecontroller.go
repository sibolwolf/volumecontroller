package main

// package name is smartconn.cc/sibolwolf/volumecontroller
import (
    "fmt"
    "os"
    "os/signal"
    LED "smartconn.cc/sibolwolf/volumecontroller/led"
    VOL "smartconn.cc/sibolwolf/volumecontroller/volume"
    KEY "smartconn.cc/liugen/input"
)

var curvolclass string

func Init() {
    curvolclass := VOL.GetVolumeClass()
    fmt.Println(curvolclass)
    switch curvolclass {
    case "class0":
        LED.UpdateLed(0)
    case "class1":
        LED.UpdateLed(1)
    case "class2":
        LED.UpdateLed(2)
    case "class3":
        LED.UpdateLed(3)
    default:
        LED.UpdateLed(0)
    }

}

func main() {
    fmt.Println("Hello, VolumeController")

    Init()

    KEY.Connect("readingangel")
    KEY.GetButton("volume").OnPress(func() {
            fmt.Println("RA got a short key press event")
        })
    signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)
	for {
		select {
		case <-signalChanel:
			return
		}
	}
}
