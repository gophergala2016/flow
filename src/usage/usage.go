package usage

import (
	"common"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// EventType define la clase de eventos que se pueden emitir
type EventType int

const (
	// UsageReport significa que el módulo está enviando el uso de cpu
	SelfUsageReport EventType = iota
	UsageReport
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

func init() {
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
		case "get-usage-self":
			usage, err := cpu.CPUPercent(time.Duration(1)*time.Second, false)
			fmt.Println(len(usage))
			if err != nil {
				out <- Event{
					Type: Error,
					Data: fmt.Sprintf("imposible obtener uso de CPU: %s", err.Error()),
				}
			} else {
				out <- Event{
					Type: SelfUsageReport,
					Data: usage[0],
				}
			}
		case "get-usage":
			usage, err := cpu.CPUPercent(time.Duration(1)*time.Second, false)
			fmt.Println(len(usage))
			if err != nil {
				out <- Event{
					Type: Error,
					Data: fmt.Sprintf("imposible obtener uso de CPU: %s", err.Error()),
				}
			} else {
				out <- Event{
					Type: UsageReport,
					Data: map[string]string{
						"peer": c.Args["peer"],
						"usage": fmt.Sprintf("%f", usage[0]),
					},
				}
			}
		default:
		}
	}
}
