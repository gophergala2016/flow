package main

import (
	"fmt"
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
		}
	}
}

func netEvent(event networking.Event) {
	switch event.Type {
	case networking.PeersFound:
		log.Println("\n\n\tpeers encontrados:\n\n")
		peers, ok := event.Data.([]string)
		if !ok {
			log.Fatalf("datos incorrectos para evento 'peers-found'")
		}
		for i := range peers {
			fmt.Printf("\t%s\n", peers[i])
		}
	}
}
