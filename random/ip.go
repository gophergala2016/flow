package main

import (
    "fmt"
    "net"
    //"os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        panic(err)
    }
    for i, addr := range addrs {
        fmt.Printf("%d %v\n", i, addr)
    }
}
