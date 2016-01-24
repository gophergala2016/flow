package networking

import (
	"fmt"
	"net"
)

// EventType define la clase de eventos que se pueden emitir
type EventType int

const (
	// Connection representa cliente conectado
	Connection EventType = iota
	// Disconnection representa cliente desconectado
	Disconnection
	// PeerLookup significa que el nodo está buscando peers
	PeerLookup
	// PeersFound significa que se encontraron peers
	PeersFound
	// Interp representa una solicitud de interpretación
	Interp
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
	go loop(in)
	return out
}

// In regresa el channel para mandar comandos al módulo
func In() chan<- Command {
	return in
}

func loop(input <-chan Command) {
	for c := range input {
		switch c.Cmd {
		case "print":
			fmt.Println(c.Args["msg"])
		case "communicateToPeer":
			hostPort := net.JoinHostPort(c.Arg["ip"], c.Arg["port"])
			conn, err := net.Dial("tcp", hostPort)
			if err != nil {
				fmt.Println("Cannot connect to host")
			}
			c := make(chan string)
			c <- "w" // [w, r, x]
			go comp.HandleRequest(conn, c chan string)
		default:
		}
	}
}

// en paquete comp
func handleRequest(conn net.Conn, c chan Algo) {
	switch v := <- c ; v {
	case "w" :
		conn.Write([]byte("Ejecuta mi codigo"))	
	}
	conn.Close()
}
