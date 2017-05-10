package led

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

/*  Tina
1. The led name for RA in tina is below:
fun_led0/  fun_led1/  net_led/   vol_led0/  vol_led1/  vol_led2/

2. The attribution for led are blow:
brightness max_brightness
*/


func GetLedStatus(filename string) {
    fmt.Println("Get the status ")
    cmd := exec.Command("/bin/ash", "-c", "cat filename")
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("StdoutPipe: " + err.Error())
        return
    }
    bytes, err := ioutil.ReadAll(stdout)
    fmt.Println("stdout: %s", bytes)
}
