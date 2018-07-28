package main

import (
	"github.com/gordonklaus/portaudio"
	"github.com/sroidl/pilot/screen"
	"os"
	"os/signal"
	"fmt"
)

type pilot struct {
}

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	menu := &screen.Menu{
		Discoverer: &pilot{},
		Connector: &pilot{},
	}

	menu.MainMenu()

	fmt.Println("bye!")
}

func (p *pilot) DiscoverDevices() []*portaudio.DeviceInfo {
	devices, _ := portaudio.Devices()

	return devices
}

func (p *pilot) ConnectDevices(input, output *portaudio.DeviceInfo) error {
	fmt.Println("Connecting ...")
	return nil
}