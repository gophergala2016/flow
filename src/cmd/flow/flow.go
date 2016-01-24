package main

import (
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
