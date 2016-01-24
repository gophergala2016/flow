package main

import (
	"common"
	"fmt"
	"interp"
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

	network := networking.Start()
	usage := usage.Start()
	ui := ui.Start()
	interp := interp.Start()

	for {
		select {
		case e := <-network:
			netEvent(e)
		case e := <-ui:
			uiEvent(e)
		case e := <-usage:
			usageEvent(e)
		case e := <-interp:
			interpEvent(e)
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
	case networking.UsageRequested:
		usage.In() <- common.Command{
			Cmd:  "get-usage",
			Args: map[string]string{"peer": event.Data.(string)},
		}
	case networking.Error:
		ui.In() <- common.Command{
			Cmd:  "print",
			Args: map[string]string{"msg": event.Data.(string)},
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
		networking.In() <- common.Command{
			Cmd:  "select-peer",
			Args: map[string]string{},
		}
	case ui.MessageSendRequested:
		networking.In() <- common.Command{
			Cmd:  "send-message",
			Args: map[string]string{"msg": event.Data.(string)},
		}
	case ui.UsageRequested:
		usage.In() <- common.Command{
			Cmd:  "get-usage-self",
			Args: map[string]string{},
		}
	case ui.InterpRequested: // temp; luego se hará la selección de peers
		interp.In() <- common.Command{
			Cmd:  "interp",
			Args: map[string]string{"code": event.Data.(string)},
		}
	case ui.UserExit:
		os.Exit(0)
	default:
	}
}

func usageEvent(event usage.Event) {
	switch event.Type {
	case usage.SelfUsageReport:
		ui.In() <- common.Command{
			Cmd: "print",
			Args: map[string]string{
				"msg": fmt.Sprintf("%f", event.Data.(float64)),
			},
		}
	case usage.UsageReport:
		networking.In() <- common.Command{
			Cmd: "send-usage",
			Args: map[string]string{
				"usage": event.Data.(map[string]string)["usage"],
				"peer":  event.Data.(map[string]string)["peer"],
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

func interpEvent(event interp.Event) {
	switch event.Type {
	case interp.InterpDone:
		ui.In() <- common.Command{
			Cmd: "print",
			Args: map[string]string{
				"msg": event.Data.(string),
			},
		}
	case interp.Error:
		ui.In() <- common.Command{
			Cmd: "print",
			Args: map[string]string{
				"msg": fmt.Sprintf("error: %s", event.Data.(string)),
			},
		}
	}
}
