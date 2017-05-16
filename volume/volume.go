package volume

import (
    "log"
    "os/exec"
    "strings"
)

var vollevel = []string{"0", "20", "40", "55"}
var volclass = []string{"class0", "class1", "class2", "class3"}

func GetVolumeStatus() string {
    log.Println("Get the volume status ... ")
    cmd := exec.Command("/bin/sh", "-c", "sysint getVolume")
    bytes, err := cmd.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return err.Error()
    }
    // log.Println(string(bytes))
    return string(bytes)
}

func SetVolume(value string) {
    log.Println("Set the volume ...")
    cmdset := log.Sprintf("sysint setVolume %s", value)
    cmd := exec.Command("/bin/sh", "-c", cmdset)
    bytes, err := cmd.Output()
    if err != nil {
        log.Println("outputerr: " + err.Error())
        return
    }
    log.Println(string(bytes))
}

func GetVolumeClass() string {
    currvol := GetVolumeStatus()
    currvol = strings.TrimSpace(currvol)
    log.Println(currvol)
    switch currvol {
    case vollevel[0]:
        log.Println(0)
        return volclass[0]
    case vollevel[1]:
        log.Println(1)
        return volclass[1]
    case vollevel[2]:
        log.Println(2)
        return volclass[2]
    case vollevel[3]:
        log.Println(3)
        return volclass[3]
    default:
        log.Println(vollevel[3])
        return volclass[0]
    }
}

func SetVolumeClass(vol string) {
    switch vol {
    case volclass[0]:
        SetVolume(vollevel[0])
    case volclass[1]:
        SetVolume(vollevel[1])
    case volclass[2]:
        SetVolume(vollevel[2])
    case volclass[3]:
        SetVolume(vollevel[3])
    default:
        SetVolume(vollevel[0])
    }
}
