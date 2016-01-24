package networking

import "github.com/hashicorp/mdns"

// LookupPeers busca peers en la red
func LookupPeers() <-chan []string {
	ch := make(chan *mdns.ServiceEntry, 4)
	retChan := make(chan []string)
	mdns.Lookup("_flow._tcp", ch)
	go func(rCh chan<- []string) {
		var entries []string
		for e := range ch {
			entries = append(entries, string(e.AddrV4))
		}
		rCh <- entries
	}(retChan)
	return retChan
}
