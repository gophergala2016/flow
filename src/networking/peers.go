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
		var entries []string
		for entry := range entriesCh {
			entries = append(entries, net.JoinHostPort(entry.AddrV4.String(),
				fmt.Sprintf("%d", entry.Port)))
		}
		resChan <- entries
	}(resChan)
	mdns.Lookup("_flow._tcp", entriesCh)
	close(entriesCh)
	return resChan
}
