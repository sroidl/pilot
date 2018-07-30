# Pilot

Pilot helps to easily to echo an audio source to an audio sink.

## Why?
When pair programming in noisy environments I often find myself distracted especially when my pair speaks in low volume. 
With **pilot** and a second soundcard (I use a usb adapter), we can now attach both or headsets into the machine and as if we had those cool helicopter pilot headsets which simply help both pilots to understand each other.
This is even more effective when using noise cancelling headsets.

## Dependencies
*Works on my machine™*

* Golang >= 1.10.3 
* [PortAudio](http://www.portaudio.com/)library (works with v19.6.0).

On MacOs, you can use `brew` to install it:

```bash
brew install portaudio
``` 

## Installation

...

## Usage

After attaching both headsets, you can use the `pilot` command to start.
You need to select which input device should be mapped to which output device.
