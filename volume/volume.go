package volume

import (
    "fmt"
    "os/exec"
)


func GetVolumeStatus() {
    fmt.Println("Get the volume status ... ")
    cmd := exec.Command("/bin/sh", "-c", "amixer sget 'speaker volume control'")
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println("outputerr: " + err.Error())
        return
    }
    fmt.Println(string(bytes))
}

func SetVolume(value string) {
    fmt.Println("Set the volume ...")
    cmdSet := fmt.Sprintf("amixer sset 'speaker volume control' %s", value)
    cmd := exec.Command("/bin/sh", "-c", cmdSet)
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println("outputerr: " + err.Error())
        return
    }
    fmt.Println(string(bytes))
}
