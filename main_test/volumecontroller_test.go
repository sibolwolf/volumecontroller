package main

import (
    "fmt"
    "os"
    "os/signal"
    VolCtrller "smartconn.cc/sibolwolf/volumecontroller"
)

func main() {
    fmt.Println("Hello, VolumeController")

    VolCtrller.Init()

    signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)
	for {
		select {
		case <-signalChanel:
			return
		}
	}
}
