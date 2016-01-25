package main

import (
	"common"
	"os"
	"ui"
)

func uiEvent(event ui.Event) {
	switch event.Type {
	case ui.PeerLookupRequested:
		sendNetworkingCommand("lookup-peers", map[string]string{})
	case ui.PeerSelectRequested:
		sendNetworkingCommand("select-peer", map[string]string{})
	case ui.MessageSendRequested:
		args := map[string]string{"msg": event.Data.(string)}
		sendNetworkingCommand("send-message", args)
	case ui.UsageRequested:
		sendUsageCommand("get-usage-self", map[string]string{})
	case ui.InterpRequested: // temp; luego se hará la selección de peers
		args := map[string]string{"code": event.Data.(string)}
		sendInterpCommand("interp", args)
	case ui.UserExit:
		os.Exit(0)
	default:
	}
}

func sendUICommand(cmd string, args map[string]string) {
	ui.In() <- common.Command{
		Cmd:  cmd,
		Args: args,
	}
}

func uiPrint(msg string) {
	args := map[string]string{
		"msg": msg,
	}
	sendUICommand("print", args)
}
