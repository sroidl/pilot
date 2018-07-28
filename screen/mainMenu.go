package screen

import (
	"github.com/gordonklaus/portaudio"
	"github.com/manifoldco/promptui"
	"fmt"
)

var noFreeDevices = fmt.Errorf("No more devices")

type DeviceDiscoverer interface {
	DiscoverDevices() []*portaudio.DeviceInfo
}

type DeviceConnector interface {
	ConnectDevices(input, output *portaudio.DeviceInfo) error
}

type Menu struct {
	Discoverer DeviceDiscoverer
	Connector  DeviceConnector

	activeConnections []string
	activeDevices []*portaudio.DeviceInfo
}

type option struct {
	Name string
	Run  func() error
}

func (m *Menu) MainMenu() {
	printActiveConnections(m.activeConnections)

	exit := &option{Name: "Exit"}
	options := []*option{
		{Name: "Start Pilot", Run: m.startPilot},
		exit,
	}

	prompt := promptui.Select{
		Label: "Pilot",
		Items: options,
		Templates: &promptui.SelectTemplates{
			Active:   "{{ .Name | green }}",
			Inactive: "{{ .Name }}",
			Label:    "{{ .Name }}",
			Selected: "{{ .Name | green }}",
		},
	}

	i, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	opt := options[i]
	if opt == exit {
		return
	}
	err = opt.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	m.MainMenu()
}

func printActiveConnections(connections []string) {
	if len(connections) <= 0 {
		return
	}

	fmt.Printf("\nActive Pilots: \n")
	for _, conn := range connections {
		fmt.Println(conn)
	}
	fmt.Println()
}


func (m *Menu) startPilot() error {
	input, err := m.selectDevice("Select Input", m.activeDevices)
	if err != nil  {
		if err == noFreeDevices {
			return nil
		}
		return err
	}

	output, err := m.selectDevice("Select Output", append(m.activeDevices, input))
	if err != nil {
		if err == noFreeDevices {
			return nil
		}
		return err
	}

	m.activeDevices = append(m.activeDevices, input)
	m.activeDevices = append(m.activeDevices, output)

	err = m.Connector.ConnectDevices(input, output)
	if err == nil {
		m.activeConnections = append(m.activeConnections, fmt.Sprintf("%s => %s", input.Name, output.Name))
	}
	return err
}

func (m *Menu) selectDevice(label string, exclusions []*portaudio.DeviceInfo) (*portaudio.DeviceInfo, error) {
	devices := filterDevices(m.Discoverer.DiscoverDevices(), exclusions)
	abort := &portaudio.DeviceInfo{Name: "Abort"}
	devices = append(devices, abort)

	prompt := promptui.Select{
		Label: label,
		Items: devices,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Name }}",
			Selected: " ",
			Active:   "{{ .Name | cyan | bold }}",
			Inactive: "{{ .Name }}",
		},
	}
	index, _, err := prompt.Run()

	if devices[index] == abort {
		return nil, noFreeDevices
	}

	return devices[index], err
}


func filterDevices(input []*portaudio.DeviceInfo, exclude []*portaudio.DeviceInfo) []*portaudio.DeviceInfo {
	var result []*portaudio.DeviceInfo

	for _, device := range input {
		if !contains(exclude, device) {
			result = append(result, device)
		}
	}

	return result
}

func contains(list []*portaudio.DeviceInfo, value *portaudio.DeviceInfo) bool {
	for _, d := range list {
		if d == value {
			return true
		}
	}
	return false
}
