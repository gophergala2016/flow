package networking

import (
    "fmt"
    "net"
    "bytes"
    "os"
)

const (
    CONN_HOST = "0.0.0.0"
    CONN_PORT = "3569"
    CONN_TYPE = "tcp"
)

var ipTable = map[string]chan Event {}

func setServer() {
    l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        fmt.Println("connection error:", err.Error())
        os.Exit(1)
    }

    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }

        c := make(chan Event)
        ipTable[conn.RemoteAddr().String()] = c

        // Handle connections in a new goroutine.
        go handleRequest(conn, c)
    }
}

func handleRequest(conn net.Conn, can_server chan Event) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)
    // Read the incoming connection into the buffer.
    _, err := conn.Read(buf)
    if err != nil {
    fmt.Println("Error reading:", err.Error())
    }

    // Index returns the index of the first instance of sep
    n := bytes.Index(buf, []byte{0})
    request := string(buf[:n])

    //c := make(chan Event)

    switch request {
    case "usage":
        out <- Event{
          Type: UsageRequested,
          Data: conn.RemoteAddr().String(),
        }
    //case "exec":  go Exec(conn, c)
    }


    // Send a response back to person contacting us.
    conn.Write([]byte("Message received.\n"))
    conn.Write(buf)
    conn.Write([]byte("\n"))



    // Close the connection when you're done with it.
    conn.Close()
}



func AskUsage(conn net.Conn,c chan Event) {

}

// func Exec(conn net.Conn, c chan Event) {
//     com := <- c
//     switch com.Cmd {
//     case "print" : fmt.Println(com.Args["msg"])
//     default :
//     }
//}
