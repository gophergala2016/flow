package networking

import (
    "fmt"
    "net"
    "common"
    "bytes"
)

const (
    CONN_HOST = "0.0.0.0"
    CONN_PORT = "3569"
    CONN_TYPE = "tcp"
    // Get user processing usage
    AskUsage EventType = iota
    // Ask user to exec something
    ExecuteCmd

)



func SetServer() {
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

        //logs an incoming message
        fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)
    // Read the incoming connection into the buffer.
    _, err := conn.Read(buf)
    if err != nil {
    fmt.Println("Error reading:", err.Error())
    }

    // Index returns the index of the first instance of sep
    n := bytes.Index(buf, []byte{0})
    request := string(buf[:n]))

    c := make(chan Event)

    switch request {
    case "usage": go SendUsage(conn, c)
    case "exec":  go Exec(conn, c)
    }


    // Send a response back to person contacting us.
    conn.Write([]byte("Message received.\n"))
    conn.Write(buf)
    conn.Write([]byte("\n"))



    // Close the connection when you're done with it.
    conn.Close()
}



func SendUsage(conn net.Conn,c chan Event) {

}

func Exec(conn net.Conn, c chan common.Command) {
    com := <- c 
    switch com.Cmd {
    case "print" : fmt.Println(com.Args["msg"])
    default :
    }

}
