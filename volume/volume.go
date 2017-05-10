package volume

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)


func GetVolumeStatus() {
    fmt.Println("Get the status ")
    cmd := exec.Command("/bin/ash", "-c", "amixer sget 'speaker volume control'")
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("StdoutPipe: " + err.Error())
        return
    }
    bytes, err := ioutil.ReadAll(stdout)
    fmt.Println("stdout: %s", bytes)
}

func SetVolume(value string) {
    fmt.Println("Get the status ")
    cmdSet := fmt.Sprintf("amixer sset 'speaker volume control' %s", value)
    cmd := exec.Command("/bin/ash", "-c", cmdSet)
    _, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("StdoutPipe: " + err.Error())
        return
    }
}
