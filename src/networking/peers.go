package networking

import (
	"fmt"
	"net"

	"github.com/hashicorp/mdns"
)

// LookupPeers busca peers en la red
func LookupPeers() <-chan []string {
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	resChan := make(chan []string, 1)
	go func(ch chan<- []string) {
		var entries []string             // peers
		addrs, _ := net.InterfaceAddrs() // interfaces locales
		for entry := range entriesCh {
			addr := entry.AddrV4 // dirección de host encontrado
			skip := false        // saltar o no esta dirección
			// buscar si peer encontrado es máquina propia
			for i := range addrs {
				n := addrs[i].(*net.IPNet) // hacer type assertion
				if n.IP.Equal(addr) {
					skip = true
				}
			}
			if !skip {
				entries = append(entries, net.JoinHostPort(addr.String(),
					fmt.Sprintf("%d", entry.Port)))
			}
		}
		resChan <- entries
	}(resChan)
	mdns.Lookup("_flow._tcp", entriesCh)
	close(entriesCh)
	return resChan
}
