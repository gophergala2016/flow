package main

import (
	"log"
	"networking"
	"os"
	"ui"
)

func main() {
	f, err := os.OpenFile("flow.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("imposible crear archivo de log: %s", err.Error())
	}
	log.SetOutput(f)

	net := networking.Start()
	ui := ui.Start()

	for {
		select {
		case e := <-net:
			netEvent(e)
		case e := <-ui:
			uiEvent(e)
		}
	}
}

func netEvent(event networking.Event) {
	switch event.Type {
	case networking.PeersFound:
		peers, ok := event.Data.([]string)
		if !ok {
			log.Fatalf("datos incorrectos para evento 'peers-found'")
		}
		peerMsg := ""
		if len(peers) == 0 {
			peerMsg = "No se encontraron peers"
		} else {
			peerMsg += "[networking-module]:\n\n"
			for i := range peers {
				peerMsg += "\t" + peers[i] + "\n"
			}
		}
		ui.In() <- ui.Command{
			Cmd:  "print",
			Args: map[string]string{"msg": peerMsg},
		}
	}
}

func uiEvent(event ui.Event) {
	switch event.Type {
	case ui.PeerLookupRequested:
		networking.In() <- networking.Command{
			Cmd:  "lookup-peers",
			Args: map[string]string{},
		}
	default:
	}
}
