package networking

import (
	"fmt"
	"net"
	"errors"
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


func SelectPeer() (string,error) {
	c := LookupPeers()
	peers := <- c
	//Regresa el segundo elemento de la lista de peers
	if len(peers) >= 2 {
		peer_selected := peers[1]
		return peer_selected, nil
	}
	return "", errors.New("you are alone")
}

func ConnectToPeer(addres string) {
	conn, err := net.Dial("tcp", addres)
	if err != nil {
		fmt.Println("cannot connect to host")
	}

	c:= make(chan string)

	go handleRequest(conn, c)
}


func handleRequest(conn net.Conn, c chan string) {
// 	switch v := <- c ; v {
// 	case "w" :
// 		log.Println("")
// 		conn.Write([]byte("Ejecuta mi codigo"))
// 	}
	conn.Write([]byte("Ejecuta mi codigo"))
	conn.Close()
}
