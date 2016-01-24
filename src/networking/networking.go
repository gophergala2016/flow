package networking

import (
	"common"
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
			p, err := SelectPeer()
			if err != nil {
				out <- Event{
					Type: Error,
					Data: fmt.Sprintf("error selecting peer: %s",err),
				}
				fmt.Printf("\nerror selecting peer:\n %s\n", err)
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
