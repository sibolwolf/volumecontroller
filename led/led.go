package led

import (
    "fmt"
    "os/exec"
)

/*  Tina
1. The led name for RA in tina is below:
fun_led0/  fun_led1/  net_led/   vol_led0/  vol_led1/  vol_led2/

2. The attribution for led are blow:
brightness max_brightness
*/

type LedAtt struct {
    brightness      int
    max_brightness  int
}

/*
LedName := []string{
    "fun_led0",
    "fun_led1",
    "net_led",
    "vol_led0",
    "vol_led1",
    "vol_led2",
*/
func GenFilename(led string) {
    fmt.Println("Gen filename ...")
}

func GetLedStatus(filename string) {
    fmt.Println("Get the vol_ledx status ...")
    cmdGet := "cat" + " " + filename
    cmd := exec.Command("/bin/sh", "-c", cmdGet)
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println("outputerr: " + err.Error())
        return
    }

    fmt.Println(filename +" status is: " + string(bytes))
}

func SetLed(filename string, value int) {
    fmt.Println("Set ledx with new value ...")

}
