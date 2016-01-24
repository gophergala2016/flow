package main

import (
	"log"
	"networking"
)

func main() {
	net := networking.Start()
	/*cmd := networking.Command{
		Cmd:  "print",
		Args: map[string]string{"msg": "\nhola putos\n"},
	}*/

	cmd := networking.Command{
		Cmd:  "communicateToPeer",
		Args: map[string]string{
			"ip": "...",
			"port" "..."
		}

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
