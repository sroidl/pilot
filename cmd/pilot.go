package main

import (
	"github.com/gordonklaus/portaudio"
	"github.com/sroidl/pilot/screen"
	"os"
	"os/signal"
	"fmt"
	"github.com/sroidl/pilot/echo"
)

type pilot struct {
	activePilots []*echo.Pilot

}

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	pilot := &pilot{}

	menu := &screen.Menu{
		Discoverer: pilot,
		Connector:  pilot,
	}

	menu.MainMenu()

	fmt.Println("bye!")
	for _, p := range pilot.activePilots {
		p.Abort()
	}
}

func (p *pilot) DiscoverDevices() []*portaudio.DeviceInfo {
	devices, _ := portaudio.Devices()

	return devices
}

func (p *pilot) ConnectDevices(input, output *portaudio.DeviceInfo) error {
	newPilot, err := echo.StartPilot(input, output)
	if err != nil {
		return err
	}

	p.activePilots = append(p.activePilots, newPilot)
	newPilot.Start()
	return nil
}