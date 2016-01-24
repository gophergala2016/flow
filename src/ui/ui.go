package ui

import (
	"fmt"
	"log"
	"strings"
	"github.com/peterh/liner"
)

type EventType int

const (
	PeerLookupRequested EventType = iota

	PeerSelectRequested

	MessageSendRequested
)

// Event se utiliza para representar un evento emitido
type Event struct {
	Type EventType
	Data interface{}
}

// Command se utiliza para mandar comandos a este módulo
type Command struct {
	Cmd  string
	Args map[string]string
}

var in chan Command
var out chan Event

func init() {
	in = make(chan Command)
	out = make(chan Event)
}

// Start inicia el módulo
func Start() <-chan Event {
	go uiLoop()
	go moduleLoop(in)
	return out
}

// In regresa el channel para mandar comandos al módulo
func In() chan<- Command {
	return in
}

func uiLoop() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	fmt.Println("\nFlow v0.1.0; Presiona Ctrl+C dos veces para salir.\n")

	for {
		if input, err := line.Prompt("flow> "); err == nil {
			inputs := strings.SplitN(input," ",2)
			if len(inputs) >= 2 {
				cmd := inputs[0]
				args := inputs[1]
				checkCmd(cmd,args)
			} else {
				checkCmd(input,"")
			}
		} else if err == liner.ErrPromptAborted {
			break
		} else {
			log.Printf("error de terminal: %s\n", err.Error())
		}
	}
}

func checkCmd(cmd string, args string) {
	switch cmd {
	case "lookup":
		fmt.Println("haciendo lookup, por favor espera...")
		out <- Event{
			Type: PeerLookupRequested,
		}
	case "select-one":
		fmt.Println("Selecting a peer")
		out <- Event{
			Type: PeerSelectRequested,
		}
	case "send":
		fmt.Println("Sending message")
		out <- Event{
			Type: MessageSendRequested,
			Data: args,
		}
	case "":
	default:
		fmt.Println("comando desconocido")
	}
}

func moduleLoop(input <-chan Command) {
	for c := range input {
		switch c.Cmd {
		case "print":
			fmt.Printf("\n\n%s\nPresiona Enter para continuar...", c.Args["msg"])
		default:
		}
	}
}
