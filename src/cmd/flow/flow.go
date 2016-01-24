package main

import (
	"log"
	"networking"
)

func main() {
	net := networking.Start()
	cmd := networking.Command{
		Cmd:  "print",
		Args: map[string]string{"msg": "\nhola putos\n"},
	}

	for {
		select {
		case e := <-net:
			netEvent(e)
		case networking.In() <- cmd:
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
