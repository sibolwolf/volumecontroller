package volume

import (
    "fmt"
    "os/exec"
    "strings"
)

var vollevel = []string{"0", "20", "40", "55"}
var volclass = []string{"class0", "class1", "class2", "class3"}

func GetVolumeStatus() string {
    fmt.Println("Get the volume status ... ")
    cmd := exec.Command("/bin/sh", "-c", "sysint getVolume")
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println("outputerr: " + err.Error())
        return err.Error()
    }
    // fmt.Println(string(bytes))
    return string(bytes)
}

func SetVolume(value string) {
    fmt.Println("Set the volume ...")
    cmdset := fmt.Sprintf("sysint setVolume %s", value)
    cmd := exec.Command("/bin/sh", "-c", cmdset)
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println("outputerr: " + err.Error())
        return
    }
    fmt.Println(string(bytes))
}

func GetVolumeClass() string {
    currvol := GetVolumeStatus()
    // fmt.Printf("%T\n", currvol)
    // fmt.Println(currvol)
    // currvol = strings.Replace(currvol, "\n", "", -1)
    // currvol = strings.Replace(currvol, "\n", "", -1)
    currvol = strings.TrimSpace(currvol)
    fmt.Println(currvol)
    switch currvol {
    case vollevel[0]:
        fmt.Println(0)
        return volclass[0]
    case vollevel[1]:
        fmt.Println(1)
        return volclass[1]
    case vollevel[2]:
        fmt.Println(2)
        return volclass[2]
    case vollevel[3]:
        fmt.Println(3)
        return volclass[3]
    default:
        fmt.Println(vollevel[3])
        return volclass[0]
    }
}
