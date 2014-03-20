package main

import (
	"io"
	"log"
	"net"
)

var _ = io.Copy

func main() {
	l, err := net.Listen("tcp", ":2593")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
            continue
		}

        go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
    buf := make([]byte, 4096)

    for {
        n, err := c.Read(buf)
        if err != nil || n == 0 {
            c.Close()
            break
        }

        n, err = c.Write(buf[0:n])
        if err != nil {
            c.Close()
            break
        }

        log.Println(string(buf[0:n]))
    }

    log.Printf("Connection from %v closed.", c.RemoteAddr())
}
