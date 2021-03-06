package main

import (
	"flag"
	"time"

	log "github.com/cihub/seelog"
	"github.com/stampzilla/stampzilla-go/nodes/basenode"
	"github.com/stampzilla/stampzilla-go/pkg/notifier"
	"github.com/stampzilla/stampzilla-go/protocol"
	"github.com/tarm/goserial"
)

// MAIN - This is run when the init function is done

var notify *notifier.Notify

func main() {
	log.Info("Starting stamp-presence node")

	// {{{
	// Parse all commandline arguments, host and port parameters are added in the basenode init function
	flag.Parse()

	//Get a config with the correct parameters
	config := basenode.NewConfig()

	//Activate the config
	basenode.SetConfig(config)

	node := protocol.NewNode("denon-receiver")

	//Start communication with the server
	connection := basenode.Connect()
	notify = notifier.New(connection)
	notify.SetSource(node)

	// Thit worker keeps track on our connection state, if we are connected or not
	go monitorState(node, connection)

	state := NewState()
	node.SetState(state)
	// }}}

	d := &denon{}
	d.setState(state)
	d.stateChangedFunc = func() {
		connection.Send(node.Node())
	}

	go serialConnector(d)

	for msg := range connection.Receive() {
		var err error

		switch msg.Cmd {
		case "on":
			err = d.on()
		case "off":
			err = d.off()
		}

		if err != nil {
			log.Errorf("Failed to run commmand '%s': %s", msg.Cmd, err.Error())
		}
	}
}

// WORKER that monitors the current connection state// {{{
func monitorState(node *protocol.Node, connection basenode.Connection) {
	for s := range connection.State() {
		switch s {
		case basenode.ConnectionStateConnected:
			connection.Send(node.Node())
		case basenode.ConnectionStateDisconnected:
		}
	}
}

// }}}

func serialConnector(d *denon) {
	for {
		<-time.After(time.Second)

		ports, err := GetMetaList()
		if err != nil {
			log.Warn(err)
			continue
		}

		var port OsSerialPort

		for _, val := range ports {
			if val.IdProduct == "2008" && val.IdVendor == "0557" {
				port = val
			}
		}

		if port.Name == "" {
			log.Info("List of available ports: ")
			for _, val := range ports {
				log.Infof("Port: %#v", val)
			}
			continue
		}

		log.Infof("Connecting to %s", port.Name)
		c := &serial.Config{Name: port.Name, Baud: 9600}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Error("Failed to open port: ", err)
			continue
		}

		d.read(s)
	}
}
