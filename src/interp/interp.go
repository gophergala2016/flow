package interp

import (
	"common"
	"fmt"

	zygo "github.com/glycerine/zygomys/repl"
)

// EventType define la clase de eventos que se pueden emitir
type EventType int

const (
	// InterpDone señaliza que el interprete terminó de evaluar código
	InterpDone EventType = iota
	// Error reporta un error
	Error
)

// Event se utiliza para representar un evento emitido
type Event struct {
	Type EventType
	Data interface{}
}

var in chan common.Command
var out chan Event

var env *zygo.Glisp

func init() {
	env = zygo.NewGlisp()
	in = make(chan common.Command)
	out = make(chan Event)
}

// Start inicia el módulo
func Start() <-chan Event {
	go loop(in)
	return out
}

// In regresa el channel para mandar comandos al módulo
func In() chan<- common.Command {
	return in
}

func loop(input <-chan common.Command) {
	for c := range input {
		switch c.Cmd {
		case "interp":
			fmt.Println(c.Args["code"])
			if err := env.LoadString(c.Args["code"]); err != nil {
				env.Clear()
				out <- Event{
					Type: Error,
					Data: fmt.Sprintf("imposible cargar código: %s", err.Error()),
				}
			} else if expr, err := env.Run(); err != nil {
				env.Clear()
				out <- Event{
					Type: Error,
					Data: fmt.Sprintf("imposible cargar código: %s", err.Error()),
				}
			} else {
				env.Clear()
				out <- Event{
					Type: InterpDone,
					Data: map[string]string{
						"peer":   c.Args["peer"],
						"result": fmt.Sprintf(expr.SexpString()),
					},
				}
			}
		}
	}
}
