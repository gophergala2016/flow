package main

import (
	"common"
	"fmt"
	"log"
	"networking"
	"os"
	"ui"
	"usage"
)

func main() {
	f, err := os.OpenFile("flow.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("imposible crear archivo de log: %s", err.Error())
	}
	log.SetOutput(f)

	net := networking.Start()
	usage := usage.Start()
	ui := ui.Start()

	for {
		select {
		case e := <-net:
			netEvent(e)
		case e := <-ui:
			uiEvent(e)
		case e := <-usage:
			usageEvent(e)
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
	}
}

func uiEvent(event ui.Event) {
	switch event.Type {
	case ui.PeerLookupRequested:
		networking.In() <- common.Command{
			Cmd:  "lookup-peers",
			Args: map[string]string{},
		}
	case ui.UsageRequested:
		usage.In() <- common.Command{
			Cmd:  "get-usage",
			Args: map[string]string{},
		}
	case ui.UserExit:
		os.Exit(0)
	default:
	}
}

func usageEvent(event usage.Event) {
	switch event.Type {
	case usage.UsageReport:
		ui.In() <- common.Command{
			Cmd: "print",
			Args: map[string]string{
				"msg": fmt.Sprintf("%f", event.Data.(float64)),
			},
		}
	case usage.Error:
		ui.In() <- common.Command{
			Cmd: "print",
			Args: map[string]string{
				"msg": fmt.Sprintf("error: %s", event.Data.(string)),
			},
		}
	}
}
