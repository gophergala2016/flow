package networking

import (
	"log"
	"os"
	"fmt"
	"github.com/hashicorp/mdns"
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
	// A peer has been chosen
	PeerSelected
	//ERROR
	Error
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
	host, err := os.Hostname()
	if err != nil {
		log.Fatal("imposible obtener hostname para publicar servicio mDNS")
	}
	info := []string{"Flow distributed computing peer"}
	service, err := mdns.NewMDNSService(host, "_flow._tcp", "", "", 3569, nil, info)
	if err != nil {
		log.Fatal("imposible crear servicio de mDNS")
	}
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		log.Fatal("imposible iniciar servicio de mDNS")
	}
	defer server.Shutdown()

	for c := range input {
		switch c.Cmd {
		case "lookup-peers":
			p := LookupPeers()
			peerTable := <-p
			out <- Event{
				Type: PeersFound,
				Data: peerTable,
			}
		case "select-peer":
			p,err := SelectPeer()
			if err != nil {
				out <- Event{
					Type: Error,
					Data: fmt.Sprintf("error selecting peer: %s",err),
				}
			} else {
				out <- Event{
					Type: PeerSelected,
					Data: p,
				}
			}
		case "send-message":
			SendMessage(c.Args["msg"])

		default:
		}
	}
}
