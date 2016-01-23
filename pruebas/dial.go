package main

import (
    "fmt"
    "net"
    "bytes"
    // "os"
    // "strconv"
)

// const (
    // CONN_HOST = "10.6.0.41"
    // CONN_HOST = "localhost"
//     CONN_PORT = "3333"
//     CONN_TYPE = "tcp"
// )

func main() {
    conn, _ := net.Dial("tcp", "localhost:3333")
    conn.Write([]byte("hOlA."))

    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
      fmt.Println("Error reading:", err.Error())
    }
    n := bytes.Index(buf, []byte{0})
    fmt.Println(string(buf[:n]))

    conn.Close()

    // for i := 41; i <= 41; i++ {
    //     CONN_HOST := "10.6.0."+strconv.Itoa(i)+":"+CONN_PORT
    //     conn, err := net.Dial(CONN_TYPE, CONN_HOST)
    //     if err != nil {
    //         fmt.Printf("ip: %v Connection refused", i)
    //     }
    //
    //     conn.Close()
    // }
}
