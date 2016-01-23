package main

import (
    "fmt"
    "net"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    addrs, _:= net.LookupIP("hibiscus.local")
    // addrs, _:= net.LookupHost("hibiscus.local")

    for _, addr := range addrs {
        fmt.Printf("%v\n", addr)
    }
}
