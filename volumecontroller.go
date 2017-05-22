package volumecontroller

// package name is smartconn.cc/sibolwolf/volumecontroller
import (
    "log"
    LED "smartconn.cc/sibolwolf/volumecontroller/led"
    VOL "smartconn.cc/sibolwolf/volumecontroller/volume"
    KEY "smartconn.cc/liugen/input"
)

var curvolclass string

func Init() {
    curvolclass = VOL.GetVolumeClass()
    log.Println("Current volume class is", curvolclass)
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

    KEY.Connect("readingangel")
    KEY.GetButton("volume").OnPress(func() {
            log.Println("RA got a short key press event for volume")
            UpdateVol()
        })

}

func UpdateVol() {
    if curvolclass == "class0" {
        curvolclass = "class1"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(1)
    } else if curvolclass == "class1" {
        curvolclass = "class2"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(2)
    } else if curvolclass == "class2" {
        curvolclass = "class3"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(3)
    } else if curvolclass == "class3" {
        curvolclass = "class0"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(0)
    } else {
        curvolclass = "class0"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(0)
    }
}
