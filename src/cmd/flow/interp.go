package main

import (
	"common"
	"fmt"
	"interp"
)

func interpEvent(event interp.Event) {
	switch event.Type {
	case interp.InterpDone:
		uiPrint(event.Data.(string))
	case interp.Error:
		uiPrint(fmt.Sprintf("error: %s", event.Data.(string)))
	}
}

func sendInterpCommand(cmd string, args map[string]string) {
	interp.In() <- common.Command{
		Cmd:  cmd,
		Args: args,
	}
}
