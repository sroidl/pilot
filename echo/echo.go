package echo
// Inspired by https://github.com/gordonklaus/portaudio/blob/master/examples/echo.go


import "github.com/gordonklaus/portaudio"

type Pilot struct {
	*portaudio.Stream
	buffer []float32
	i      int
}

func StartPilot(in, out *portaudio.DeviceInfo) (*Pilot, error) {
	pilot := new(Pilot)
	var err error

	streamParameters := portaudio.LowLatencyParameters(in, out)
	streamParameters.Input.Channels = 1
	streamParameters.Output.Channels = 1

	pilot.Stream, err = portaudio.OpenStream(streamParameters, loopback)
	if err != nil {
		return nil, err
	}

	return pilot, nil
}

func loopback(in, out []float32) {
	for i := range in {
		out[i] = in[i]
	}
}