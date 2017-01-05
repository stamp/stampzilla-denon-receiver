package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	d := &denon{}
	s := NewState()
	d.setState(s)

	d.decode(
		"PWON\r" +
			"MV80\r" +
			"SIPHONO\r")
	assert.Equal(t, true, s.MainZone.Power)
	assert.Equal(t, 0.0, s.MainZone.Volume)
	assert.Equal(t, "Phono", s.MainZone.Source)

	d.decode(
		"PWSTANDBY\r" +
			"MV995\r" +
			"SICD\r")
	assert.Equal(t, false, s.MainZone.Power)
	assert.Equal(t, 19.5, s.MainZone.Volume)
	assert.Equal(t, "CD", s.MainZone.Source)

	d.decode("SITUNER\r" +
		"MV00")
	assert.Equal(t, -80.0, s.MainZone.Volume)
	assert.Equal(t, "Tuner", s.MainZone.Source)

	d.decode("SIDVD\r")
	assert.Equal(t, "Dvd", s.MainZone.Source)

	d.decode("SIBD\r")
	assert.Equal(t, "BD", s.MainZone.Source)

	d.decode("SITV\r")
	assert.Equal(t, "Tv", s.MainZone.Source)

	d.decode("SISAT/CBL\r")
	assert.Equal(t, "Sat/Cable", s.MainZone.Source)

	d.decode("SIDVR\r")
	assert.Equal(t, "Dvr", s.MainZone.Source)

	d.decode("SIGAME\r")
	assert.Equal(t, "Game", s.MainZone.Source)

	d.decode("SIGAME2\r")
	assert.Equal(t, "Game2", s.MainZone.Source)

	d.decode("SIV.AUX\r")
	assert.Equal(t, "V.aux", s.MainZone.Source)

	d.decode("SIDOCK\r")
	assert.Equal(t, "Dock", s.MainZone.Source)

	d.decode("SIHDRADIO\r")
	assert.Equal(t, "HD radio", s.MainZone.Source)

	d.decode("SIIPOD\r")
	assert.Equal(t, "iPod", s.MainZone.Source)

	d.decode("SINET/USB\r")
	assert.Equal(t, "Network/Usb", s.MainZone.Source)

	d.decode("SIRHAPSODY\r")
	assert.Equal(t, "Rhapsody", s.MainZone.Source)

	d.decode("SINAPSTER\r")
	assert.Equal(t, "Napster", s.MainZone.Source)

	d.decode("SIPANDORA\r")
	assert.Equal(t, "Pandora", s.MainZone.Source)

	d.decode("SILASTFM\r")
	assert.Equal(t, "Last FM", s.MainZone.Source)

	d.decode("SIFLICKR\r")
	assert.Equal(t, "Flickr", s.MainZone.Source)

	d.decode("SIFAVORITES\r")
	assert.Equal(t, "Favorites", s.MainZone.Source)

	d.decode("SIIRADIO\r")
	assert.Equal(t, "Internet radio", s.MainZone.Source)

	d.decode("SISERVER\r")
	assert.Equal(t, "Server", s.MainZone.Source)

	d.decode("SIUSB/IPOD\r")
	assert.Equal(t, "Usb/iPod", s.MainZone.Source)

	d.decode("SIUSB\r")
	assert.Equal(t, "Usb - Start Playback", s.MainZone.Source)

	d.decode("SIIPD\r")
	assert.Equal(t, "iPod - Start Playback", s.MainZone.Source)

	d.decode("SIIRP\r")
	assert.Equal(t, "Internet radio - Start Playback", s.MainZone.Source)

	d.decode("SIFVP\r")
	assert.Equal(t, "Favorites - Start Playback", s.MainZone.Source)

}
