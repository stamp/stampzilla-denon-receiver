package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type denon struct {
	state *State
}

func (d *denon) setState(s *State) {
	d.state = s
}

func (d *denon) read(r io.Reader) {
	scanner := bufio.NewScanner(r)

	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	}
	scanner.Split(split)

	// Validate the input
	for scanner.Scan() {
		d.process(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

func (d *denon) decode(input string) {
	d.read(strings.NewReader(input))
}

func (d *denon) process(line string) {
	switch line {
	case "PWON":
		d.state.MainZone.Power = true
	case "PWSTANDBY":
		d.state.MainZone.Power = false
	case "SIPHONO":
		d.state.MainZone.Source = "Phono"
	case "SICD":
		d.state.MainZone.Source = "CD"
	case "SITUNER":
		d.state.MainZone.Source = "Tuner"
	case "SIDVD":
		d.state.MainZone.Source = "Dvd"
	case "SIBD":
		d.state.MainZone.Source = "BD"
	case "SITV":
		d.state.MainZone.Source = "Tv"
	case "SISAT/CBL":
		d.state.MainZone.Source = "Sat/Cable"
	case "SIDVR":
		d.state.MainZone.Source = "Dvr"
	case "SIGAME":
		d.state.MainZone.Source = "Game"
	case "SIGAME2":
		d.state.MainZone.Source = "Game2"
	case "SIV.AUX":
		d.state.MainZone.Source = "V.aux"
	case "SIDOCK":
		d.state.MainZone.Source = "Dock"
	case "SIHDRADIO":
		d.state.MainZone.Source = "HD radio"
	case "SIIPOD":
		d.state.MainZone.Source = "iPod"
	case "SINET/USB":
		d.state.MainZone.Source = "Network/Usb"
	case "SIRHAPSODY":
		d.state.MainZone.Source = "Rhapsody"
	case "SINAPSTER":
		d.state.MainZone.Source = "Napster"
	case "SIPANDORA":
		d.state.MainZone.Source = "Pandora"
	case "SILASTFM":
		d.state.MainZone.Source = "Last FM"
	case "SIFLICKR":
		d.state.MainZone.Source = "Flickr"
	case "SIFAVORITES":
		d.state.MainZone.Source = "Favorites"
	case "SIIRADIO":
		d.state.MainZone.Source = "Internet radio"
	case "SISERVER":
		d.state.MainZone.Source = "Server"
	case "SIUSB/IPOD":
		d.state.MainZone.Source = "Usb/iPod"
	case "SIUSB":
		d.state.MainZone.Source = "Usb - Start Playback"
	case "SIIPD":
		d.state.MainZone.Source = "iPod - Start Playback"
	case "SIIRP":
		d.state.MainZone.Source = "Internet radio - Start Playback"
	case "SIFVP":
		d.state.MainZone.Source = "Favorites - Start Playback"
	default:
		switch {
		case len(line) >= 4 && line[0:2] == "MV":
			v, err := strconv.Atoi(line[2:])
			if err != nil {
				return
			}
			volume := float64(v)

			if len(line[2:]) > 2 {
				volume /= 10
			}

			volume = volume - 80

			d.state.MainZone.Volume = volume
		default:
			log.Fatal(line[2:])
		}
	}
}
