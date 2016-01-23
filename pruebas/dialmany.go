package main

import (
    "fmt"
    "net"
    "bytes"
    // "os"
    "strconv"
)

const (
    // CONN_HOST = "10.6.0.41"
    // CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    for i := 40; i <= 57; i++ {
        CONN_HOST := "10.6.0."+strconv.Itoa(i)+":"+CONN_PORT
        conn, err := net.Dial(CONN_TYPE, CONN_HOST)
        if err != nil {
            fmt.Printf("ip: %v Connection refused\n", i)
        } else {
            fmt.Printf("ip: %v Connection accepted\n", i)
            conn.Write([]byte("hOlA."))

            buf := make([]byte, 1024)
            _, err := conn.Read(buf)
            if err != nil {
              fmt.Println("Error reading:", err.Error())
            }
            n := bytes.Index(buf, []byte{0})
            fmt.Println(string(buf[:n]))

            conn.Close()
        }
    }
}
