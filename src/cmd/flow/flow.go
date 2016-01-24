package main

import (
	"log"
	"networking"
)

func main() {
	net := networking.Start()
	cmd := networking.Command{
		Cmd:  "lookup-peers",
		Args: map[string]string{},
	}
	networking.In() <- cmd

	for {
		select {
		case e := <-net:
			netEvent(e)
		default:
		}
	}
}

func netEvent(event networking.Event) {
	switch event.Type {
	case networking.Connection:
		log.Println("cliente conectado")
	}
}
