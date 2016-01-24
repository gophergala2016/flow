package ui

import (
	"common"
	"fmt"
	"log"

	"github.com/peterh/liner"
)

// EventType define la clase de eventos que se pueden emitir
type EventType int

const (
	// PeerLookupRequested significa que el usuario ha pedido un lookup de peers
	PeerLookupRequested EventType = iota
	// UsageRequested significa que el usuario quere el uso de su CPU
	UsageRequested
	// UserExit significa que el usuario quiere salir del programa
	UserExit
)

// Event se utiliza para representar un evento emitido
type Event struct {
	Type EventType
	Data interface{}
}

var in chan common.Command
var out chan Event

func init() {
	in = make(chan common.Command)
	out = make(chan Event)
}

// Start inicia el módulo
func Start() <-chan Event {
	go uiLoop()
	go moduleLoop(in)
	return out
}

// In regresa el channel para mandar comandos al módulo
func In() chan<- common.Command {
	return in
}

func uiLoop() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)

	fmt.Print("\n\nFlow v0.1.0\n\nPresiona Ctrl+C para salir\n\n")

	for {
		if cmd, err := line.Prompt("flow> "); err == nil {
			checkCmd(cmd)
		} else if err == liner.ErrPromptAborted {
			break
		} else {
			log.Printf("error de terminal: %s\n", err.Error())
		}
	}
	line.Close()
	out <- Event{
		Type: UserExit,
		Data: nil,
	}
}

func checkCmd(cmd string) {
	switch cmd {
	case "lookup":
		fmt.Println("haciendo lookup, por favor espera...")
		out <- Event{
			Type: PeerLookupRequested,
		}
	case "usage":
		fmt.Println("obteniendo uso, por favor espera...")
		out <- Event{
			Type: UsageRequested,
		}
	case "":
	default:
		fmt.Println("comando desconocido")
	}
}

func moduleLoop(input <-chan common.Command) {
	for c := range input {
		switch c.Cmd {
		case "print":
			fmt.Printf("\n\n%s\nPresiona Enter para continuar...", c.Args["msg"])
		default:
		}
	}
}
