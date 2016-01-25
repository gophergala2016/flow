package ui

import (
	"common"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/peterh/liner"
)

// EventType define la clase de eventos que se pueden emitir
type EventType int

const (
	// PeerLookupRequested significa que el usuario ha pedido un lookup de peers
	PeerLookupRequested EventType = iota
	// PeerSelectRequested significa que el usuario ha pedido seleccionar un peer
	PeerSelectRequested
	// MessageSendRequested significa que el usuario ha pedido enviar un mensaje
	MessageSendRequested
	// UsageRequested significa que el usuario quere el uso de su CPU
	UsageRequested
	// InterpRequested significa que el usuario quiere un archivo interpretado
	InterpRequested
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

	fmt.Print("\n\nFlow v0.1.0\n\nPress Ctrl+C twice to exit\n\n")

	for {
		if input, err := line.Prompt("flow> "); err == nil {
			inputs := strings.SplitN(input, " ", 2)
			if len(inputs) > 1 {
				cmd := inputs[0]
				args := inputs[1]
				checkCmd(cmd, args)
			} else {
				checkCmd(input, "")
			}
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
	case "usage":
		fmt.Println("obteniendo uso, por favor espera...")
		out <- Event{
			Type: UsageRequested,
		}
	case "eval":
		code, err := ioutil.ReadFile(args)
		if err != nil {
			fmt.Printf("\nerror: %s\n", err.Error())
		}
		out <- Event{
			Type: InterpRequested,
			Data: string(code),
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
			fmt.Printf("\n%s\nPresiona Enter para continuar...", c.Args["msg"])
		default:
		}
	}
}
