package networking

import (
	"errors"
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

func SelectPeer() (string, error) {
	c := LookupPeers()
	peers := <-c
	// fmt.Println(len(peers))
	//Regresa el segundo elemento de la lista de peers
	if len(peers) > 0 {
		peer_selected := peers[0]
		return peer_selected, nil
	}
	return "", errors.New("you are alone")
}

func SendMessage(msg string) {
	peer, _ := SelectPeer()
	// fmt.Println(peer)
	conn, err := net.Dial("tcp", peer)
	if err != nil {
		fmt.Println("cannot connect to host")
	}
	c:= make(chan string)
	go handleConnection(conn, c)
	c <- msg
}


// func handleRequest(conn net.Conn, c chan string) {
// 	c := make(chan string)
// 	go handleConnection(conn, c)
//
// }

// func ConnectToPeer(addres string) {
// 	conn, err := net.Dial("tcp", addres)
// 	if err != nil {
// 		fmt.Println("cannot connect to host")
// 	}
//
// 	c:= make(chan string)
//
// 	go handleConnection(conn, c)
// }

func handleConnection(conn net.Conn, c chan string) {

// 	switch v := <- c ; v {
// 	case "w" :
// 		log.Println("")
// 		conn.Write([]byte("Ejecuta mi codigo"))
// 	}
	// 	switch v := <- c ; v {
	// 	case "w" :
	// 		log.Println("")
	// 		conn.Write([]byte("Ejecuta mi codigo"))
	// 	}
	// msg := <- c
	conn.Write([]byte(<- c))
	conn.Close()
}
