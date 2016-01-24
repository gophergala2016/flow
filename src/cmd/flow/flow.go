package main

import (
	"common"
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
		ui.In() <- common.Command{
			Cmd:  "print",
			Args: map[string]string{"msg": peerMsg},
		}
	case networking.PeerSelected:
		peer, ok := event.Data.(string)
		peerMsg := ""
		if !ok {
			log.Fatalf("datos incorrectos para evento 'peers-found'")
		} else {
			peerMsg = "selected peer " + peer
		}
		ui.In() <- common.Command{
			Cmd:  "print",
			Args: map[string]string{"msg": peerMsg},
		}
	}
}

func uiEvent(event ui.Event) {
	switch event.Type {
	case ui.PeerLookupRequested:
		networking.In() <- common.Command{
			Cmd:  "lookup-peers",
			Args: map[string]string{},
		}
	case ui.PeerSelectRequested:
<<<<<<< HEAD
		networking.In() <- common.Command{
			Cmd: "select-peer",
=======
		networking.In() <- networking.Command{
			Cmd:  "select-peer",
>>>>>>> 624098b1d5252b1eeb803a49695c10eff8c1b69b
			Args: map[string]string{},
		}
	case ui.MessageSendRequested:
		networking.In() <- networking.Command{
			Cmd:  "send-message",
			Args: map[string]string{"msg": event.Data.(string)},
		}
	case ui.UserExit:
		os.Exit(0)
	default:
	}
}
