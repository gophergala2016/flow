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
			"ip": "10.6.0.57",
			"port": "8000",
			},
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
