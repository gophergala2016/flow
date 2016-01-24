package main

import (
	"common"
	"fmt"
	"usage"
)

func usageEvent(event usage.Event) {
	switch event.Type {
	case usage.SelfUsageReport:
		uiPrint(fmt.Sprintf("%f", event.Data.(float64)))
	case usage.UsageReport:
		args := map[string]string{
			"usage": event.Data.(map[string]string)["usage"],
			"peer":  event.Data.(map[string]string)["peer"],
		}
		sendNetworkingCommand("send-usage", args)
	case usage.Error:
		uiPrint(fmt.Sprintf("error: %s", event.Data.(string)))
	}
}

func sendUsageCommand(cmd string, args map[string]string) {
	usage.In() <- common.Command{
		Cmd:  cmd,
		Args: args,
	}
}
