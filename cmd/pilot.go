package main

import "github.com/gordonklaus/portaudio"

func main () {
	portaudio.Initialize()
	defer portaudio.Terminate()

}